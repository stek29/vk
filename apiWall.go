package vkCallbackApi

import (
	"net/url"
	"strconv"
)

type WallDeleteCommentParams struct {
	OwnerID   int
	CommentID int
}

func (vk *VKApi) WallDeleteComment(params WallDeleteCommentParams) (bool, error) {
	method := "wall.deleteComment"
	query := url.Values{}

	query.Set("owner_id", strconv.Itoa(params.OwnerID))
	query.Set("comment_id", strconv.Itoa(params.CommentID))

	r, err := vk.Request(method, query)
	if err != nil {
		return false, err
	}

	var resp int

	if resp, err = strconv.Atoi(string(r)); err != nil {
		return false, err
	}

	return resp == 1, nil
}
