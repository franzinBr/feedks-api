package dtos

type CreateFeedBackRequest struct {
	Comment string `json:"comment" binding:"required"`
}

type UserFeedBack struct {
	ID       uint   `json:"id"`
	UserName string `json:"userName"`
}

type FeedBackResponse struct {
	ID      uint   `json:"id"`
	Comment string `json:"comment"`
	User    UserFeedBack
}
