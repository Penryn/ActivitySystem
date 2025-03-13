package service

import (
	"activitySystem/internal/model"
	"context"
)

func GetUserByUsername(ctx context.Context, username string) (user *model.User, err error) {
	user, err = d.GetUserByUsername(ctx, username)
	return user, err
}

func CreateUser(ctx context.Context, username string, password string, stuID string, email string) (err error) {
	err = d.CreateUser(ctx, &model.User{
		Username: username,
		Password: password,
		StuID:    stuID,
		Email:    email,
		Avatar:   "https://qiuniu.phlin.cn/bucket/icon.png",
	})
	return err
}

func GetUserByID(ctx context.Context, id uint) (user *model.User, err error) {
	user, err = d.GetUserByID(ctx, id)
	return user, err
}
