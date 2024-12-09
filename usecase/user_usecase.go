package usecase

import (
	"backend/model"
	"backend/repository"
	"backend/validator"
	"os"
	"time"

	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"backend/auth"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SingUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
	GetGoogleAuthURL() string
	GoogleCallback(code string) (string, error)
}

type userUsecase struct {
	ur               repository.IUserRepository
	uv               validator.IUserValidator
	googleAuthConfig auth.GoogleAuthConfig
}

func NewUserUsecase(
	ur repository.IUserRepository, uv validator.IUserValidator, gac auth.GoogleAuthConfig) IUserUsecase {
	return &userUsecase{
		ur:               ur,
		uv:               uv,
		googleAuthConfig: gac,
	}
}

func (uu *userUsecase) SingUp(user model.User) (model.UserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}

	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (uu *userUsecase) GetGoogleAuthURL() string {
	return uu.googleAuthConfig.GetConfig().AuthCodeURL("state")
}

func (uu *userUsecase) GoogleCallback(code string) (string, error) {
	config := uu.googleAuthConfig.GetConfig()
	// Googleからトークンを取得
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	// ユーザー情報を取得
	res, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	userInfo := model.GoogleUserInfo{}
	if err := json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return "", err
	}

	// DBからユーザーを検索
	user := model.User{}
	exists, err := uu.ur.ExistsUserByEmail(userInfo.Email)
	if err != nil {
		return "", fmt.Errorf("failed to check user existence: %w", err)
	}

	if !exists {
		// ユーザーが存在しない場合は新規作成
		newUser := model.User{
			Email: userInfo.Email,
			Name:  userInfo.Name,
		}
		if err := uu.ur.CreateUser(&newUser); err != nil {
			return "", fmt.Errorf("failed to create new user: %w", err)
		}
		user = newUser
	} else {
		// ユーザーが存在する場合は取得
		if err := uu.ur.GetUserByEmail(&user, userInfo.Email); err != nil {
			return "", fmt.Errorf("failed to get existing user: %w", err)
		}
	}

	// JWTトークンを生成
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})

	return jwtToken.SignedString([]byte(os.Getenv("SECRET")))
}
