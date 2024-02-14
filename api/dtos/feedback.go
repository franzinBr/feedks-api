package dtos

type CreateFeedBackRequest struct {
	Comment string `json:"comment" binding:"required"`
}

type DeleteFeedbackRequest struct {
	ID uint `uri:"id" binding:"required"`
}

type UserFeedBack struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type FeedBackResponse struct {
	ID        uint         `json:"id"`
	Comment   string       `json:"comment"`
	CreatedAt string       `json:"created_at"`
	User      UserFeedBack `json:"user,omitempty"`
}
