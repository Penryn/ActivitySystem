package dao

import (
	"activitySystem/internal/model"
	"context"
	"gorm.io/gorm"
	"time"
)

func (d *Dao) CreateActivity(ctx context.Context, Activity *model.Activity) (err error) {
	err = d.orm.WithContext(ctx).Create(&Activity).Error
	return err
}

func (d *Dao) GetNewestActivityList(ctx context.Context, category string, num, size int) (activityList []model.Activity, n int64, err error) {
	query := d.orm.WithContext(ctx).Model(&model.Activity{})

	// 处理 category 模糊查询
	if category != "" {
		query = query.Where("category = ?", category)
	}

	// 统计总数
	err = query.Count(&n).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	err = query.Order("created_at DESC").Limit(size).Offset((num - 1) * size).Find(&activityList).Error
	return activityList, n, err
}

func (d *Dao) GetLatestActivityList(ctx context.Context, num, size int) (activityList []model.Activity, n int64, err error) {
	query := d.orm.WithContext(ctx).Model(&model.Activity{}).Where("start_time > NOW()")

	// 统计未来活动的总数
	err = query.Count(&n).Error
	if err != nil {
		return nil, 0, err
	}

	// 查询最近的未来活动
	err = query.Order("start_time ASC").Limit(size).Offset((num - 1) * size).Find(&activityList).Error
	return activityList, n, err
}

func (d *Dao) GetHottestActivityList(ctx context.Context, num, size int) (activityList []model.Activity, n int64, err error) {
	query := d.orm.WithContext(ctx).Model(&model.Activity{})

	// 统计未来活动的总数
	err = query.Count(&n).Error
	if err != nil {
		return nil, 0, err
	}

	// 查询最近的未来活动
	err = query.Order("upvote DESC").Limit(size).Offset((num - 1) * size).Find(&activityList).Error
	return activityList, n, err
}

func (d *Dao) UpdateActivity(ctx context.Context, aid int, title, content, category, location string, deadline, startTime time.Time) (err error) {
	err = d.orm.WithContext(ctx).Model(&model.Activity{}).Where("id = ?", aid).Updates(&model.Activity{
		Title:     title,
		Content:   content,
		Category:  category,
		Location:  location,
		Deadline:  deadline,
		StartTime: startTime,
	}).Error
	return err
}

func (d *Dao) DeleteActivity(ctx context.Context, aid int) (err error) {
	err = d.orm.WithContext(ctx).Delete(&model.Activity{}, aid).Error
	return err
}

func (d *Dao) GetActivityByID(ctx context.Context, aid int) (Activity *model.Activity, err error) {
	err = d.orm.WithContext(ctx).Where("id = ?", aid).First(Activity).Error
	return Activity, err
}

func (d *Dao) GetAllActivityByID(ctx context.Context, aid int) (Activity *model.Activity, err error) {
	Activity = new(model.Activity)
	err = d.orm.WithContext(ctx).Unscoped().Where("id = ?", aid).First(Activity).Error
	return Activity, err
}

func (d *Dao) UpvoteActivity(ctx context.Context, aid int) (err error) {
	err = d.orm.WithContext(ctx).Model(&model.Activity{}).Where("id = ?", aid).UpdateColumn("upvote", gorm.Expr("upvote + ?", 1)).Error
	return err
}
