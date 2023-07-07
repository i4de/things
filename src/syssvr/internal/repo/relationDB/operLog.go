package relationDB

import (
	"context"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/store"
	"gorm.io/gorm"
)

type OperLogRepo struct {
	db *gorm.DB
}

func NewOperLogRepo(in any) *OperLogRepo {
	return &OperLogRepo{db: store.GetCommonConn(in)}
}

type OperLogFilter struct {
	OperName     string
	OperUserName string
	BusinessType int64
}

func (p OperLogRepo) fmtFilter(ctx context.Context, f OperLogFilter) *gorm.DB {
	db := p.db.WithContext(ctx)
	if f.OperName != "" {
		db = db.Where("`operName` like ?", "%"+f.OperName+"%")
	}
	if f.OperUserName != "" {
		db = db.Where("`operUserName` like ?", "%"+f.OperUserName+"%")
	}
	if f.BusinessType > 0 {
		db = db.Where("`businessType`= ?", f.BusinessType)
	}
	return db
}

func (g OperLogRepo) Insert(ctx context.Context, data *SysOperLog) error {
	result := g.db.WithContext(ctx).Create(data)
	return store.ErrFmt(result.Error)
}

func (g OperLogRepo) FindOneByFilter(ctx context.Context, f OperLogFilter) (*SysOperLog, error) {
	var result SysOperLog
	db := g.fmtFilter(ctx, f)
	err := db.First(&result).Error
	if err != nil {
		return nil, store.ErrFmt(err)
	}
	return &result, nil
}
func (p OperLogRepo) FindByFilter(ctx context.Context, f OperLogFilter, page *def.PageInfo) ([]*SysOperLog, error) {
	var results []*SysOperLog
	db := p.fmtFilter(ctx, f).Model(&SysOperLog{})
	db = page.ToGorm(db)
	err := db.Find(&results).Error
	if err != nil {
		return nil, store.ErrFmt(err)
	}
	return results, nil
}

func (p OperLogRepo) CountByFilter(ctx context.Context, f OperLogFilter) (size int64, err error) {
	db := p.fmtFilter(ctx, f).Model(&SysOperLog{})
	err = db.Count(&size).Error
	return size, store.ErrFmt(err)
}

func (g OperLogRepo) Update(ctx context.Context, data *SysOperLog) error {
	err := g.db.WithContext(ctx).Where("`id` = ?", data.ID).Save(data).Error
	return store.ErrFmt(err)
}

func (g OperLogRepo) DeleteByFilter(ctx context.Context, f OperLogFilter) error {
	db := g.fmtFilter(ctx, f)
	err := db.Delete(&SysOperLog{}).Error
	return store.ErrFmt(err)
}

func (g OperLogRepo) Delete(ctx context.Context, id int64) error {
	err := g.db.WithContext(ctx).Where("`id` = ?", id).Delete(&SysOperLog{}).Error
	return store.ErrFmt(err)
}
func (g OperLogRepo) FindOne(ctx context.Context, id int64) (*SysOperLog, error) {
	var result SysOperLog
	err := g.db.WithContext(ctx).Where("`id` = ?", id).First(&result).Error
	if err != nil {
		return nil, store.ErrFmt(err)
	}
	return &result, nil
}
