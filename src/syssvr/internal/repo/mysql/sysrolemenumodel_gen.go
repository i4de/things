// Code generated by goctl. DO NOT EDIT.

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysRoleMenuFieldNames          = builder.RawFieldNames(&SysRoleMenu{})
	sysRoleMenuRows                = strings.Join(sysRoleMenuFieldNames, ",")
	sysRoleMenuRowsExpectAutoSet   = strings.Join(stringx.Remove(sysRoleMenuFieldNames, "`id`", "`deletedTime`", "`createdTime`", "`updatedTime`"), ",")
	sysRoleMenuRowsWithPlaceHolder = strings.Join(stringx.Remove(sysRoleMenuFieldNames, "`id`", "`deletedTime`", "`createdTime`", "`updatedTime`"), "=?,") + "=?"
)

type (
	sysRoleMenuModel interface {
		Insert(ctx context.Context, data *SysRoleMenu) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysRoleMenu, error)
		FindOneByRoleIDMenuID(ctx context.Context, roleID sql.NullInt64, menuID sql.NullInt64) (*SysRoleMenu, error)
		Update(ctx context.Context, data *SysRoleMenu) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysRoleMenuModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysRoleMenu struct {
		Id          int64         `db:"id"`          // id编号
		RoleID      sql.NullInt64 `db:"roleID"`      // 角色ID
		MenuID      sql.NullInt64 `db:"menuID"`      // 菜单ID
		CreatedTime time.Time     `db:"createdTime"` // 创建时间
		UpdatedTime time.Time     `db:"updatedTime"` // 更新时间
		DeletedTime sql.NullTime  `db:"deletedTime"`
	}
)

func newSysRoleMenuModel(conn sqlx.SqlConn) *defaultSysRoleMenuModel {
	return &defaultSysRoleMenuModel{
		conn:  conn,
		table: "`sys_role_menu`",
	}
}

func (m *defaultSysRoleMenuModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysRoleMenuModel) FindOne(ctx context.Context, id int64) (*SysRoleMenu, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRoleMenuRows, m.table)
	var resp SysRoleMenu
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysRoleMenuModel) FindOneByRoleIDMenuID(ctx context.Context, roleID sql.NullInt64, menuID sql.NullInt64) (*SysRoleMenu, error) {
	var resp SysRoleMenu
	query := fmt.Sprintf("select %s from %s where `roleID` = ? and `menuID` = ? limit 1", sysRoleMenuRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, roleID, menuID)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysRoleMenuModel) Insert(ctx context.Context, data *SysRoleMenu) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, sysRoleMenuRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.RoleID, data.MenuID)
	return ret, err
}

func (m *defaultSysRoleMenuModel) Update(ctx context.Context, newData *SysRoleMenu) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysRoleMenuRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.RoleID, newData.MenuID, newData.Id)
	return err
}

func (m *defaultSysRoleMenuModel) tableName() string {
	return m.table
}
