package services

import (
	"net/http"
	"os"
	"strconv"

	"github.com/franzinBr/feedks-api/api/dtos"
	"github.com/franzinBr/feedks-api/api/errors"
	"github.com/franzinBr/feedks-api/constants"
	"github.com/franzinBr/feedks-api/data/db"
	"github.com/franzinBr/feedks-api/data/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	Db           *gorm.DB
	TokenService *TokenService
}

func NewUserService() *UserService {
	return &UserService{
		Db:           db.GetDB(),
		TokenService: NewTokenService(),
	}
}

func (s *UserService) CreateUser(req *dtos.CreateUserRequest) error {
	var defaultRole models.Role
	s.Db.Where("name = ?", constants.DefaultRoleName).First(&defaultRole)

	if r := s.Db.Where("email = ?", req.Email).First(&models.User{}); r.RowsAffected > 0 {
		return &errors.ApiError{Message: "user with this email already exists", StatusCode: http.StatusBadRequest}
	}

	if r := s.Db.Where("username = ?", req.Username).First(&models.User{}); r.RowsAffected > 0 {
		return &errors.ApiError{Message: "user with this username already exists", StatusCode: http.StatusBadRequest}
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return &errors.ApiError{Message: "Error on hash password", StatusCode: http.StatusInternalServerError}
	}

	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashPass),
		RoleID:    int(defaultRole.ID),
	}

	tx := s.Db.Begin()

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return &errors.ApiError{Message: "Error in create User", StatusCode: http.StatusInternalServerError}
	}

	tx.Commit()

	return nil

}

func (s *UserService) Login(req *dtos.LoginRequest) (*dtos.TokenResponse, error) {
	var user models.User

	if err := s.Db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return nil, &errors.ApiError{Message: "Username or password is invalid", StatusCode: http.StatusUnauthorized}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, &errors.ApiError{Message: "Username or password is invalid", StatusCode: http.StatusUnauthorized}
	}

	tokenClaims := map[string]any{
		"id": strconv.FormatUint(uint64(user.ID), 10),
	}

	acess_exp, err := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_EXPIRY_HOUR"))
	if err != nil {
		return nil, &errors.ApiError{Message: "Error in generate token", StatusCode: http.StatusInternalServerError}
	}

	refresh_exp, err := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_EXPIRY_HOUR"))
	if err != nil {
		return nil, &errors.ApiError{Message: "Error in generate refresh token", StatusCode: http.StatusInternalServerError}
	}

	acessToken, err := s.TokenService.CreateToken(
		os.Getenv("JWT_ACCESS_TOKEN_SECRET"),
		acess_exp,
		tokenClaims,
	)

	if err != nil {
		return nil, &errors.ApiError{Message: "Error in generate token", StatusCode: http.StatusInternalServerError, InternalError: err}
	}

	refreshToken, err := s.TokenService.CreateToken(
		os.Getenv("JWT_REFRESH_TOKEN_SECRET"),
		refresh_exp,
		tokenClaims,
	)

	if err != nil {
		return nil, &errors.ApiError{Message: "Error in generate token", StatusCode: http.StatusInternalServerError, InternalError: err}
	}

	return &dtos.TokenResponse{
		AccessToken:  acessToken,
		RefreshToken: refreshToken,
	}, nil

}
