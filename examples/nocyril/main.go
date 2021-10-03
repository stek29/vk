package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"

	"github.com/jinzhu/configor"
	"github.com/spf13/pflag"

	"github.com/stek29/vk"
	"github.com/stek29/vk/vkapi"
	"github.com/stek29/vk/vkbot"
)

var vkClient vk.API
var cyrilRegex = regexp.MustCompile(`(?i)[а-яёй]`)

type extConfig struct {
	VKToken string
	Groups  []vkbot.CallbackGroupConfig
}

func getNameForUID(userID int) (string, error) {
	users, err := vkapi.Users{API: vkClient}.Get(vkapi.UsersGetParams{
		UserIDs: []string{strconv.Itoa(userID)},
	})

	if err != nil {
		return "", err
	}

	if len(users) < 1 || users[0].ID != userID {
		return "", errors.New("VK didn't return user we wanted")
	}

	return users[0].FirstName, nil
}

func getNameForGID(groupID int) (string, error) {
	groups, err := vkapi.Groups{API: vkClient}.GetByID(vkapi.GroupsGetByIDParams{
		GroupID: strconv.Itoa(groupID),
	})

	if err != nil {
		return "", err
	}

	if len(groups) < 1 || groups[0].ID != groupID {
		return "", errors.New("VK didn't return group we wanted")
	}

	return groups[0].Name, nil
}

func getNameForID(ID int) (string, error) {
	switch {
	case ID > 0:
		return getNameForUID(ID)
	case ID < 0:
		return getNameForGID(-ID)
	default:
		return "", errors.New("ID cant be 0!")
	}
}

type commentDeleter func(client vk.API, ownerID int, commentID int) error

func handleComment(ownerID int, comment vk.Comment, deleter commentDeleter) {
	commentText := comment.Text
	log.Printf("Handling Comment: %d_%d (%q)", ownerID, comment.ID, commentText)

	if replyTo := comment.ReplyToUser; replyTo != 0 {
		name, err := getNameForID(replyTo)

		if err != nil {
			log.Printf("Cant get first name for id%d: %v", replyTo, err)
		}

		commentText = strings.Replace(commentText, name, "", -1)
	}

	if cyrilRegex.MatchString(commentText) {
		if err := deleter(vkClient, ownerID, comment.ID); err != nil {
			log.Printf("Cant delete comment %d_%d: %v", ownerID, comment.ID, err)
		}

		log.Printf("Deleted Comment: %d_%d", ownerID, comment.ID)
	}
}

func handleDeleteComment(ok bool, err error) error {
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("deleteComment returned false")
	}
	return nil
}

func wallCommentDeleter(client vk.API, ownerID int, commentID int) error {
	return handleDeleteComment(vkapi.Wall{API: vkClient}.DeleteComment(vkapi.WallDeleteCommentParams{
		OwnerID:   ownerID,
		CommentID: commentID,
	}))
}

func photoCommentDeleter(client vk.API, ownerID int, commentID int) error {
	return handleDeleteComment(vkapi.Photos{API: vkClient}.DeleteComment(vkapi.PhotosDeleteCommentParams{
		OwnerID:   ownerID,
		CommentID: commentID,
	}))
}

func videoCommentDeleter(client vk.API, ownerID int, commentID int) error {
	return handleDeleteComment(vkapi.Video{API: vkClient}.DeleteComment(vkapi.VideoDeleteCommentParams{
		OwnerID:   ownerID,
		CommentID: commentID,
	}))
}

func marketCommentDeleter(client vk.API, ownerID int, commentID int) error {
	return handleDeleteComment(vkapi.Market{API: vkClient}.DeleteComment(vkapi.MarketDeleteCommentParams{
		OwnerID:   ownerID,
		CommentID: commentID,
	}))
}

func handleEvent(event vk.CallbackEvent) {
	log.Printf("Got event: GroupID=%v Secret=%v Event=%T", event.GroupID, event.Secret, event.Event)
	switch v := event.Event.(type) {
	case vk.WallReplyNew:
		go handleComment(v.PostOwnerID, v.Comment, wallCommentDeleter)
	case vk.WallReplyEdit:
		go handleComment(v.PostOwnerID, v.Comment, wallCommentDeleter)
	case vk.WallReplyRestore:
		go handleComment(v.PostOwnerID, v.Comment, wallCommentDeleter)
	case vk.PhotoCommentNew:
		go handleComment(v.PhotoOwnerID, v.Comment, photoCommentDeleter)
	case vk.PhotoCommentEdit:
		go handleComment(v.PhotoOwnerID, v.Comment, photoCommentDeleter)
	case vk.PhotoCommentRestore:
		go handleComment(v.PhotoOwnerID, v.Comment, photoCommentDeleter)
	case vk.VideoCommentNew:
		go handleComment(v.VideoOwnerID, v.Comment, videoCommentDeleter)
	case vk.VideoCommentEdit:
		go handleComment(v.VideoOwnerID, v.Comment, videoCommentDeleter)
	case vk.VideoCommentRestore:
		go handleComment(v.VideoOwnerID, v.Comment, videoCommentDeleter)
	case vk.MarketCommentNew:
		go handleComment(v.MarketOwnerID, v.Comment, marketCommentDeleter)
	case vk.MarketCommentEdit:
		go handleComment(v.MarketOwnerID, v.Comment, marketCommentDeleter)
	case vk.MarketCommentRestore:
		go handleComment(v.MarketOwnerID, v.Comment, marketCommentDeleter)
	}
}

func main() {
	lAddr := pflag.StringP("listen", "l", "127.0.0.1:8081", "host:port to listen")
	confPath := pflag.StringP("conf", "c", "nocyril.yaml", "Path to config file")

	pflag.Parse()

	var config extConfig

	if err := configor.Load(&config, *confPath); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	vkToken := config.VKToken
	if vkToken == "" {
		log.Fatal("vk token is required (check config file)")
	}

	if client, err := vk.NewBaseAPI(vk.BaseAPIConfig{AccessToken: vkToken}); err != nil {
		log.Fatalf("Failed to create vkClient: %v", err)
	} else {
		vkClient = client
	}

	if len(config.Groups) == 0 {
		log.Fatalf("At least one group config is required")
	}

	bot, err := vkbot.NewBot(vkClient, vkbot.BotConfig{
		GroupID: config.Groups[0].GroupID,
		Poller: &vkbot.CallbackPoller{
			Listen:       *lAddr,
			GroupConfigs: config.Groups,
		},
	})
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func(cancel context.CancelFunc) {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		cancel()
	}(cancel)

	events, err := bot.StartPolling(ctx, 0)
	if err != nil {
		log.Fatal("Cat start polling:", err)
	}

	for e := range events {
		handleEvent(e)
	}

	log.Printf("Bye!")
}
