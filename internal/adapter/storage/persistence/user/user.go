package user

import (
	"fmt"
	"github.com/Yideg/admybrand_challenge/internal/adapter/storage/persistence"
	CustomError "github.com/Yideg/admybrand_challenge/internal/constant/errors"
	"github.com/Yideg/admybrand_challenge/internal/constant/model"
	"gorm.io/gorm/clause"

	"context"
	"gorm.io/gorm"
)

//userPersistence receiver all ports
type userPersistence struct {
	conn *gorm.DB
}

//UserInit initializes user persistence
func UserInit(conn *gorm.DB) persistence.UserPersistence {
	return &userPersistence{
		conn: conn,
	}
}

//GetUserByID fetches a specific user identified by an id number
func (u *userPersistence) GetUserByID(c context.Context, usr model.User) (*model.User, error) {
	conn := u.conn.WithContext(c)
	user := &model.User{}
	err := conn.Model(&model.User{}).Where("id = ?", usr.ID).First(&user).Error
	fmt.Println("user ", user)
	if err != nil {
		return nil, CustomError.ErrRecordNotFound
	}
	user.Sanitize()
	return user, nil
}

//Users fetches all registered accounts
func (u *userPersistence) Users(c context.Context) ([]model.User, error) {
	conn := u.conn.WithContext(c)
	user := []model.User{}

	err := conn.Model(&model.User{}).Find(&user).Error
	if err != nil {
		return nil, CustomError.ErrRecordNotFound
	}
	for i, u := range user {
		u.Sanitize()
		user[i] = u
	}
	return user, nil
}

//GetUserByUsername fetch a user that matches with a given username
func (u *userPersistence) GetUserByUsername(c context.Context, usr model.User) (*model.User, error) {
	conn := u.conn.WithContext(c)
	user := &model.User{}
	err := conn.Model(&model.User{}).Where("user_name = ?", usr.UserName).First(user).Error
	fmt.Println("by username error ", err)
	if err != nil {
		return nil, CustomError.ErrRecordNotFound
	}
	user.Sanitize()

	return user, nil
}

//UpdateUser is used to edit user account after it is created
func (u *userPersistence) UpdateUser(c context.Context, user model.User) (*model.User, error) {
	conn := u.conn.WithContext(c)
	err := conn.Model(&model.User{}).
		Where("id = ?", user.ID).
		Updates(&user).Error

	if err != nil {
		return nil, CustomError.ErrUnableToSave
	}
	user.Sanitize()

	return &user, nil
}

//DeleteUser is used to delete user account
func (u *userPersistence) DeleteUser(c context.Context, user model.User) error {
	conn := u.conn.WithContext(c)
	usr := new(model.User)
	err := conn.Model(&model.User{}).Where("id=?", user.ID).First(&usr).Error
	if err != nil {
		return CustomError.ErrRecordNotFound
	}
	err = conn.Model(&model.User{}).Where("id=?", usr.ID).Delete(&usr).Error
	if err != nil {
		return CustomError.ErrUnableToDelete
	}
	return nil
}

//StoreUser used to create new account for all type of users
func (u *userPersistence) StoreUser(c context.Context, user model.User) (*model.User, error) {
	//TODO implement me
	conn := u.conn.WithContext(c)
	user.RoleName = "anonymous"
	err := conn.Model(&model.User{}).Create(&user).Error
	if err != nil {
		return nil, CustomError.ErrorUnableToCreate
	}
	user.Sanitize()
	return &user, nil
}

// AssignRoleToUser assignes new role to a new registered  user
func (u *userPersistence) AssignRoleToUser(c context.Context, usr model.User) (*model.User, error) {
	//TODO implement me
	conn := u.conn.WithContext(c)
	err := conn.Model(&model.User{}).
		Select("role_name").
		Preload(clause.Associations).
		Where("id=?", usr.ID).
		Updates(usr).Error
	if err != nil {
		return nil, CustomError.ErrorUnableToCreate
	}
	usr.Sanitize()
	return &usr, nil
}
