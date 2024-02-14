package services

import (
	"fmt"
	"net/http"

	"github.com/franzinBr/feedks-api/api/dtos"
	"github.com/franzinBr/feedks-api/api/errors"
	"github.com/franzinBr/feedks-api/api/helpers"
	"github.com/franzinBr/feedks-api/constants"
	"github.com/franzinBr/feedks-api/data/db"
	"github.com/franzinBr/feedks-api/data/models"
	"gorm.io/gorm"
)

type FeedBackService struct {
	Db *gorm.DB
}

func NewFeedBackService() *FeedBackService {
	return &FeedBackService{
		Db: db.GetDB(),
	}
}

func (s *FeedBackService) CreateFeedBack(req *dtos.CreateFeedBackRequest, userId string) (*dtos.FeedBackResponse, error) {
	var user models.User
	if err := s.Db.First(&user, userId).Error; err != nil {
		return nil, &errors.ApiError{Message: "Error on get user", StatusCode: http.StatusInternalServerError}
	}

	feedBack := models.Feedback{
		Comment: req.Comment,
		UserID:  int(user.ID),
	}

	tx := s.Db.Begin()

	if err := tx.Create(&feedBack).Error; err != nil {
		tx.Rollback()
		return nil, &errors.ApiError{Message: "Error in create Feedback", StatusCode: http.StatusInternalServerError}
	}

	tx.Commit()

	return &dtos.FeedBackResponse{
		ID:        feedBack.ID,
		Comment:   feedBack.Comment,
		CreatedAt: feedBack.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		User: dtos.UserFeedBack{
			ID:       user.ID,
			Username: user.Username,
		},
	}, nil
}

func (s *FeedBackService) ListFeedBacks(req *dtos.PaginationRequest, userId string) (*dtos.PaginationResponse[dtos.FeedBackResponse], error) {
	var user models.User
	if err := s.Db.Preload("Role").First(&user, userId).Error; err != nil {
		return nil, &errors.ApiError{Message: "Error on get user", StatusCode: http.StatusInternalServerError}
	}

	var feedbacks []*models.Feedback
	var feedbacksResponse []*dtos.FeedBackResponse

	paginationResponse := new(dtos.PaginationResponse[dtos.FeedBackResponse])

	query := s.Db.Scopes(helpers.Paginate[dtos.FeedBackResponse](feedbacks, req, paginationResponse, s.Db))

	switch user.Role.Name {
	case constants.AdminRole:
		query.Preload("User").Find(&feedbacks)
	default:
		query.Where("user_id = ?", user.ID).Preload("User").Find(&feedbacks)
	}

	for _, feedback := range feedbacks {

		feedbacksResponse = append(feedbacksResponse, &dtos.FeedBackResponse{
			ID:        feedback.ID,
			Comment:   feedback.Comment,
			CreatedAt: feedback.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			User: dtos.UserFeedBack{
				ID:       feedback.User.ID,
				Username: feedback.User.Username,
			},
		})
	}

	paginationResponse.Items = &feedbacksResponse

	return paginationResponse, nil
}

func (s *FeedBackService) DeleteFeedback(req *dtos.DeleteFeedbackRequest, userId string) error {
	var user models.User
	if err := s.Db.Preload("Role").First(&user, userId).Error; err != nil {
		return &errors.ApiError{Message: "Error on get user", StatusCode: http.StatusInternalServerError}
	}

	var feedBack models.Feedback
	if err := s.Db.Preload("User").First(&feedBack, req.ID).Error; err != nil {
		return &errors.ApiError{Message: fmt.Sprintf("Feedback %d Not Found", req.ID), StatusCode: http.StatusNotFound}
	}

	if feedBack.UserID != int(user.ID) && user.Role.Name != constants.AdminRole {
		return &errors.ApiError{Message: fmt.Sprintf("You cannot delete this feedback %d", req.ID), StatusCode: http.StatusForbidden}
	}

	tx := s.Db.Begin()

	if err := tx.Delete(&feedBack).Error; err != nil {
		tx.Rollback()
		return &errors.ApiError{Message: "Error on delete Feedback", StatusCode: http.StatusInternalServerError}
	}

	tx.Commit()

	return nil
}
