package relationDB

import (
	"context"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/store"
	"gorm.io/gorm"
)

type RoleInfoRepo struct {
	db *gorm.DB
}

func NewRoleInfoRepo(in any) *RoleInfoRepo {
	return &RoleInfoRepo{db: store.GetCommonConn(in)}
}

type RoleInfoFilter struct {
	Name   string
	Status int64
	*RoleInfoWith
}
type RoleInfoWith struct {
	WithMenus bool
}

func (p RoleInfoRepo) fmtWith(db *gorm.DB, f *RoleInfoWith) *gorm.DB {
	if f == nil {
		return nil
	}
	if f.WithMenus {
		db = db.Preload("Menus")
	}
	return db
}

func (p RoleInfoRepo) fmtFilter(ctx context.Context, f RoleInfoFilter) *gorm.DB {
	db := p.db.WithContext(ctx)
	db = p.fmtWith(db, f.RoleInfoWith)
	if f.Name != "" {
		db = db.Where("`name` like ?", "%"+f.Name+"%")
	}
	if f.Status > 0 {
		db = db.Where("`status`= ?", f.Status)
	}
	return db
}

func (g RoleInfoRepo) Insert(ctx context.Context, data *SysRoleInfo) error {
	result := g.db.WithContext(ctx).Create(data)
	return store.ErrFmt(result.Error)
}

func (g RoleInfoRepo) FindOneByFilter(ctx context.Context, f RoleInfoFilter) (*SysRoleInfo, error) {
	var result SysRoleInfo
	db := g.fmtFilter(ctx, f)
	err := db.First(&result).Error
	if err != nil {
		return nil, store.ErrFmt(err)
	}
	return &result, nil
}
func (p RoleInfoRepo) FindByFilter(ctx context.Context, f RoleInfoFilter, page *def.PageInfo) ([]*SysRoleInfo, error) {
	var results []*SysRoleInfo
	db := p.fmtFilter(ctx, f).Model(&SysRoleInfo{})
	db = page.ToGorm(db)
	err := db.Find(&results).Error
	if err != nil {
		return nil, store.ErrFmt(err)
	}
	return results, nil
}

func (p RoleInfoRepo) CountByFilter(ctx context.Context, f RoleInfoFilter) (size int64, err error) {
	db := p.fmtFilter(ctx, f).Model(&SysRoleInfo{})
	err = db.Count(&size).Error
	return size, store.ErrFmt(err)
}

func (g RoleInfoRepo) Update(ctx context.Context, data *SysRoleInfo) error {
	err := g.db.WithContext(ctx).Where("`id` = ?", data.ID).Save(data).Error
	return store.ErrFmt(err)
}

func (g RoleInfoRepo) DeleteByFilter(ctx context.Context, f RoleInfoFilter) error {
	db := g.fmtFilter(ctx, f)
	err := db.Delete(&SysRoleInfo{}).Error
	return store.ErrFmt(err)
}

func (g RoleInfoRepo) Delete(ctx context.Context, id int64) error {
	err := g.db.WithContext(ctx).Where("`id` = ?", id).Delete(&SysRoleInfo{}).Error
	return store.ErrFmt(err)
}
func (g RoleInfoRepo) FindOne(ctx context.Context, id int64, with *RoleInfoWith) (*SysRoleInfo, error) {
	var result SysRoleInfo
	db := g.db.WithContext(ctx)
	db = g.fmtWith(db, with)
	err := db.Where("`id` = ?", id).First(&result).Error
	if err != nil {
		return nil, store.ErrFmt(err)
	}
	return &result, nil
}