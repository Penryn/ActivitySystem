package dao

import (
	"activitySystem/internal/model"
	"context"
)

func (d *Dao) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := d.orm.WithContext(ctx).Where("username = ?", username).First(&user).Error
	return &user, err
}

func (d *Dao) CreateUser(ctx context.Context, user *model.User) error {
	err := d.orm.WithContext(ctx).Create(&user).Error
	return err
}

func (d *Dao) GetUserByID(ctx context.Context, uid uint) (*model.User, error) {
	var user model.User
	err := d.orm.WithContext(ctx).Where("id = ?", uid).First(&user).Error
	return &user, err
}

func (d *Dao) UpdateUser(ctx context.Context, uid uint, user *model.User) error {
	err := d.orm.WithContext(ctx).Model(&model.User{}).Where("id = ?", uid).Updates(&user).Error
	return err
}
