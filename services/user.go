package services

import (
	"dans-multi-pro/helpers"
	"dans-multi-pro/models"
	"dans-multi-pro/params"
	"dans-multi-pro/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	userRepo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

func (u *UserService) CreateUser(request params.CreateUser) *params.Response {
	user := models.User{
		Username: request.Username,
		Password: request.Password,
	}

	createUserData, err := u.userRepo.CreateUser(&user)

	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: map[string]string{
				"error": err.Error(),
			},
		}
	}

	return &params.Response{
		Status: http.StatusCreated,
		Payload: gin.H{
			"id":       createUserData.ID,
			"username": createUserData.Username,
		},
	}
}

func (u *UserService) Login(request params.CreateUser) *params.Response {
	var userDB models.User

	if request.Username == "" {
		return &params.Response{
			Status: http.StatusUnauthorized,
			Payload: gin.H{
				"message": "Username cannot be null",
			},
		}
	}

	if request.Password == "" {
		return &params.Response{
			Status: http.StatusUnauthorized,
			Payload: gin.H{
				"message": "Password cannot be null",
			},
		}
	}

	err := u.userRepo.CheckUserByUsername(request.Username, &userDB)

	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": err.Error(),
			},
		}
	}

	passwordIsOK := helpers.ComparePassword([]byte(userDB.Password), []byte(request.Password))

	if !passwordIsOK {
		return &params.Response{
			Status: http.StatusUnauthorized,
			Payload: gin.H{
				"message": "Password not match",
			},
		}
	}

	token := helpers.GenerateToken(userDB.ID, userDB.Username)

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"token": token,
		},
	}
}
