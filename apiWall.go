package vkCallbackApi

// APIWall implements VK API namespace `wall`
type APIWall struct {
	api *APIBase
}

// WallDeleteCommentParams are params for Wall.DeleteComment
type WallDeleteCommentParams struct {
	OwnerID   int `url:"owner_id"`
	CommentID int `url:"commend_id"`
}

// DeleteComment is wall.deleteComment
func (v APIWall) DeleteComment(params WallDeleteCommentParams) (bool, error) {
	r, err := v.api.Request("wall.deleteComment", params)
	if err != nil {
		return false, err
	}

	return decodeBoolIntResponse(r)
}
