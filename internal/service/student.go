package service

import (
	"activitySystem/internal/model"
	"context"
	"time"
)

func CreateActivity(ctx context.Context, uid uint, title, content, category, location, img string, deadline, startTime time.Time) (err error) {
	err = d.CreateActivity(ctx, &model.Activity{
		UserID:    uid,
		Title:     title,
		Content:   content,
		Category:  category,
		Location:  location,
		Upvote:    0,
		Img:       img,
		Deadline:  deadline,
		StartTime: startTime,
	})
	return err
}

func GetNewestActivityList(ctx context.Context, category string, num, size int) (activityList []model.Activity, n int64, err error) {
	activityList, n, err = d.GetNewestActivityList(ctx, category, num, size)
	return activityList, n, err
}

func GetLatestActivityList(ctx context.Context, num, size int) (activityList []model.Activity, n int64, err error) {
	activityList, n, err = d.GetLatestActivityList(ctx, num, size)
	return activityList, n, err
}

func GetHottestActivityList(ctx context.Context, num, size int) (activityList []model.Activity, n int64, err error) {
	activityList, n, err = d.GetHottestActivityList(ctx, num, size)
	return activityList, n, err
}

func UpdateActivity(ctx context.Context, aid int, title, content, category, location string, deadline, startTime time.Time) (err error) {
	err = d.UpdateActivity(ctx, aid, title, content, category, location, deadline, startTime)
	return err
}

func GetRecordList(ctx context.Context, uid uint, num, size int) (RecordList []model.Record, n int64, err error) {
	RecordList, n, err = d.GetRecordList(ctx, uid, num, size)
	return RecordList, n, err
}

func GetActivityByID(ctx context.Context, aid int) (activity *model.Activity, err error) {
	activity, err = d.GetActivityByID(ctx, aid)
	return activity, err
}

func DeleteActivityAndRecordByActivityID(ctx context.Context, aid int) (err error) {
	err = d.DeleteActivityAndRecordByActivityID(ctx, aid)
	return err
}

func UpvoteActivity(ctx context.Context, aid int) (err error) {
	err = d.UpvoteActivity(ctx, aid)
	return err
}

func SignUpActivity(ctx context.Context, uid, aid uint) (err error) {
	return d.CreateRecord(ctx, &model.Record{
		UserID:     uid,
		ActivityID: aid,
	})
}

func CancelSignUpActivity(ctx context.Context, uid, aid uint) (err error) {
	return d.DeleteRecord(ctx, uid, aid)
}

func GetRecordByActivityIDAndUserID(ctx context.Context, uid, aid uint) (record *model.Record, err error) {
	record, err = d.GetRecordByActivityIDAndUserID(ctx, uid, aid)
	return record, err
}

func UpdateUser(ctx context.Context, uid uint, username, password, stuID, email, avatar, profile string) (err error) {
	err = d.UpdateUser(ctx, uid, &model.User{
		Username: username,
		Password: password,
		StuID:    stuID,
		Email:    email,
		Avatar:   avatar,
		Profile:  profile,
	})
	return err
}
