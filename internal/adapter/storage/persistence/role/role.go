package user

import (
	"fmt"
	"github.com/Yideg/admybrand_challenge/internal/adapter/storage/persistence"
	CustomError "github.com/Yideg/admybrand_challenge/internal/constant/errors"
	"github.com/Yideg/admybrand_challenge/internal/constant/model"

	"context"
	"gorm.io/gorm"
)

type rolePersistence struct {
	conn *gorm.DB
}

func RoleInit(conn *gorm.DB) persistence.RolePersistence {
	return &rolePersistence{
		conn: conn,
	}
}
func (r rolePersistence) GetRoleByName(c context.Context, role model.Role) (*model.Role, error) {
	//TODO implement me
	conn := r.conn.WithContext(c)
	rl := &model.Role{}
	err := conn.Model(&model.Role{}).Where("name = ?", role.Name).First(&rl).Error
	fmt.Println("user ", r)
	if err != nil {
		return nil, CustomError.ErrRecordNotFound
	}
	return rl, nil
}

func (r rolePersistence) Roles(c context.Context) ([]model.Role, error) {
	//TODO implement me
	conn := r.conn.WithContext(c)
	roles := []model.Role{}
	err := conn.Model(&model.Role{}).Find(&roles).Error
	if err != nil {
		return nil, CustomError.ErrRecordNotFound
	}
	return roles, nil
}

func (r rolePersistence) DeleteRole(c context.Context, role model.Role) error {
	//TODO implement me
	conn := r.conn.WithContext(c)
	rl := new(model.Role)
	err := conn.Model(&model.Role{}).Where("name=?", role.Name).First(&rl).Error
	if err != nil {
		return CustomError.ErrRecordNotFound
	}
	err = conn.Model(&model.Role{}).Where("name=?", rl.Name).Delete(&rl).Error
	if err != nil {
		return CustomError.ErrUnableToDelete
	}
	return nil
}

func (r rolePersistence) StoreRole(c context.Context, role model.Role) (*model.Role, error) {
	//TODO implement me
	conn := r.conn.WithContext(c)
	err := conn.Model(&model.Role{}).Create(&role).Error
	fmt.Println("---->>>>>>>>...>error ", err)
	if err != nil {
		return nil, CustomError.ErrorUnableToCreate
	}
	return &role, nil
}
