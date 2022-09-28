package auth

import (
	"context"
	"fmt"
	"github.com/Yideg/admybrand_challenge/internal/adapter/storage/persistence"
	"github.com/Yideg/admybrand_challenge/internal/constant"
	CustomError "github.com/Yideg/admybrand_challenge/internal/constant/errors"
	"github.com/Yideg/admybrand_challenge/internal/constant/model"
	"github.com/Yideg/admybrand_challenge/internal/module"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userPersist    persistence.UserPersistence
	contextTimeout time.Duration
}

func Initialize(userPerst persistence.UserPersistence, ut model.Utils) module.AuthUseCase {
	return &service{
		userPersist:    userPerst,
		contextTimeout: ut.Timeout,
	}
}
func (s service) UserLogin(c context.Context, user model.User) (*constant.Token, error) {

	userDb, errs := s.userPersist.GetUserByUsername(c, user)
	if errs != nil {
		return nil, errs
	}
	er := constant.CompareHashAndPassword([]byte(user.Password), []byte(userDb.Password))
	if er == bcrypt.ErrMismatchedHashAndPassword {
		return nil, CustomError.ErrrHashPasswordMissMatched
	}
	fmt.Println("user id when logged ", userDb.ID)
	if userDb.RoleName == "anonymous" {
		return nil, CustomError.UnverifiedAccount
	}
	token, err := CreateToken(userDb.ID, user.UserName)
	if err != nil {
		return nil, CustomError.ErrInvalidToken
	}
	data := &constant.Token{
		AccessToken: token,
		UserID:      userDb.ID,
	}
	return data, nil
}
func (s service) GetUserByID(c context.Context, user model.User) (*model.User, error) {
	//TODO implement me
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.userPersist.GetUserByID(ctx, user)
}
