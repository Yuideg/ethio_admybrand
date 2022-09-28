package persistence

import (
	"context"
	"github.com/Yideg/admybrand_challenge/internal/constant/model"
)

// RolePersistence interface
type RolePersistence interface {
	GetRoleByName(c context.Context, role model.Role) (*model.Role, error)
	Roles(c context.Context) ([]model.Role, error)
	DeleteRole(c context.Context, role model.Role) error
	StoreRole(c context.Context, role model.Role) (*model.Role, error)
}

//UserPersistence interface
type UserPersistence interface {
	GetUserByID(c context.Context, user model.User) (*model.User, error)
	Users(c context.Context) ([]model.User, error)
	UpdateUser(c context.Context, user model.User) (*model.User, error)
	DeleteUser(c context.Context, user model.User) error
	StoreUser(c context.Context, user model.User) (*model.User, error)
	GetUserByUsername(c context.Context, usr model.User) (*model.User, error)
	AssignRoleToUser(c context.Context, usr model.User) (*model.User, error)
}
