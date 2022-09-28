package user

import (
	"context"
	"github.com/Yideg/admybrand_challenge/internal/adapter/storage/persistence"
	"github.com/Yideg/admybrand_challenge/internal/constant/model"
	"github.com/Yideg/admybrand_challenge/internal/module"
	"time"
)

//Service defines all necessary service for the domain User
type service struct {
	userPersist    persistence.UserPersistence
	contextTimeout time.Duration
}

func (s service) GetUserByUsername(c context.Context, usr model.User) (*model.User, error) {
	//TODO implement me
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.userPersist.GetUserByUsername(ctx, usr)
}

//Initialize creates a new object with User UserUseCase type
func Initialize(userPerst persistence.UserPersistence, ut model.Utils) module.UserUseCase {
	return &service{
		userPersist:    userPerst,
		contextTimeout: ut.Timeout,
	}
}

//User gets a specific user by id number
func (s service) GetUserByID(c context.Context, user model.User) (*model.User, error) {
	//TODO implement me
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.userPersist.GetUserByID(ctx, user)
}

//Users fetches all recorded  users
func (s service) Users(c context.Context) ([]model.User, error) {
	//TODO implement me
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.userPersist.Users(ctx)
}

//UpdateUser updates Users
func (s service) UpdateUser(c context.Context, user model.User) (*model.User, error) {
	//TODO implement me
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.userPersist.UpdateUser(ctx, user)
}

//DeleteUser removes a specific User
func (s service) DeleteUser(c context.Context, user model.User) error {
	//TODO implement me
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.userPersist.DeleteUser(ctx, user)
}

// StoreUser creates new user
func (s service) StoreUser(c context.Context, user model.User) (*model.User, error) {
	//TODO implement me
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.userPersist.StoreUser(ctx, user)
}

//AssignRoleToUser assigns role
func (s service) AssignRoleToUser(c context.Context, usr model.User) (*model.User, error) {
	//TODO implement me
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.userPersist.AssignRoleToUser(ctx, usr)
}
