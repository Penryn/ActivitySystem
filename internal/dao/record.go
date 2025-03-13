package dao

import (
	"activitySystem/internal/model"
	"context"
)

func (d *Dao) CreateRecord(ctx context.Context, record *model.Record) (err error) {
	err = d.orm.WithContext(ctx).Create(&record).Error
	return err
}

func (d *Dao) GetRecordList(ctx context.Context, uid uint, num, size int) (recordList []model.Record, n int64, err error) {
	query := d.orm.WithContext(ctx).Model(&model.Record{}).Where("user_id = ?", uid)

	// 统计总数
	err = query.Count(&n).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	err = query.Order("created_at DESC").Limit(size).Offset((num - 1) * size).Find(&recordList).Error
	return recordList, n, err
}

func (d *Dao) GetRecordByActivityID(ctx context.Context, aid int) (record *model.Record, err error) {
	err = d.orm.WithContext(ctx).Where("activity_id = ?", aid).First(&record).Error
	return record, err
}

func (d *Dao) DeleteRecord(ctx context.Context, uid, aid uint) error {
	return d.orm.WithContext(ctx).Where("user_id = ? AND activity_id = ?", uid, aid).Delete(&model.Record{}).Error
}

func (d *Dao) DeleteActivityAndRecordByActivityID(ctx context.Context, aid int) error {
	// 开启事务
	tx := d.orm.WithContext(ctx).Begin()

	// 删除 activity 记录
	if err := tx.Where("id = ?", aid).Delete(&model.Activity{}).Error; err != nil {
		tx.Rollback() // 事务回滚
		return err
	}

	// 删除 record 记录
	if err := tx.Where("activity_id = ?", aid).Delete(&model.Record{}).Error; err != nil {
		tx.Rollback() // 事务回滚
		return err
	}

	// 提交事务
	return tx.Commit().Error
}

func (d *Dao) GetRecordByActivityIDAndUserID(ctx context.Context, uid, aid uint) (record *model.Record, err error) {
	err = d.orm.WithContext(ctx).Where("user_id = ? AND activity_id = ?", uid, aid).First(&record).Error
	return record, err
}
