package vkCallbackApi

// APIWall implements VK API namespace `wall`
type APIWall struct {
	API *API
}

// WallDeleteCommentParams are params for Wall.DeleteComment
type WallDeleteCommentParams struct {
	OwnerID   int `url:"owner_id"`
	CommentID int `url:"comment_id"`
}

// DeleteComment is wall.deleteComment
func (v APIWall) DeleteComment(params WallDeleteCommentParams) (bool, error) {
	r, err := v.API.Request("wall.deleteComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
