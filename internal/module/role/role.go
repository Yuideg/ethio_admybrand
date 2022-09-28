package user

import (
	"context"
	"github.com/Yideg/admybrand_challenge/internal/adapter/storage/persistence"
	"github.com/Yideg/admybrand_challenge/internal/constant/model"
	"github.com/Yideg/admybrand_challenge/internal/module"
	"time"
)

//Service defines all necessary service for the domain Role
type service struct {
	rolePersist    persistence.RolePersistence
	contextTimeout time.Duration
}

//Initialize creates a new object with Role RoleUseCase type
func Initialize(rolePerst persistence.RolePersistence, ut model.Utils) module.RoleUseCase {
	return &service{
		rolePersist:    rolePerst,
		contextTimeout: ut.Timeout,
	}
}

func (s service) GetRoleByName(c context.Context, role model.Role) (*model.Role, error) {
	//TODO implement me
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.rolePersist.GetRoleByName(ctx, role)
}

func (s service) Roles(c context.Context) ([]model.Role, error) {
	//TODO implement me
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.rolePersist.Roles(ctx)
}

func (s service) DeleteRole(c context.Context, role model.Role) error {
	//TODO implement me
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.rolePersist.DeleteRole(ctx, role)
}

func (s service) StoreRole(c context.Context, role model.Role) (*model.Role, error) {
	//TODO implement me
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.rolePersist.StoreRole(ctx, role)
}
