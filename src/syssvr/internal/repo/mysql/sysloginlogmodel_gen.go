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
	sysLoginLogFieldNames          = builder.RawFieldNames(&SysLoginLog{})
	sysLoginLogRows                = strings.Join(sysLoginLogFieldNames, ",")
	sysLoginLogRowsExpectAutoSet   = strings.Join(stringx.Remove(sysLoginLogFieldNames, "`id`", "`deletedTime`", "`createdTime`", "`updatedTime`"), ",")
	sysLoginLogRowsWithPlaceHolder = strings.Join(stringx.Remove(sysLoginLogFieldNames, "`id`", "`deletedTime`", "`createdTime`", "`updatedTime`"), "=?,") + "=?"
)

type (
	sysLoginLogModel interface {
		Insert(ctx context.Context, data *SysLoginLog) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysLoginLog, error)
		Update(ctx context.Context, data *SysLoginLog) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysLoginLogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysLoginLog struct {
		Id            int64     `db:"id"`            // 编号
		Uid           int64     `db:"uid"`           // 用户id
		UserName      string    `db:"userName"`      // 登录账号
		IpAddr        string    `db:"ipAddr"`        // 登录IP地址
		LoginLocation string    `db:"loginLocation"` // 登录地点
		Browser       string    `db:"browser"`       // 浏览器类型
		Os            string    `db:"os"`            // 操作系统
		Code          int64     `db:"code"`          // 登录状态（200成功 其它失败）
		Msg           string    `db:"msg"`           // 提示消息
		CreatedTime   time.Time `db:"createdTime"`   // 登录时间
	}
)

func newSysLoginLogModel(conn sqlx.SqlConn) *defaultSysLoginLogModel {
	return &defaultSysLoginLogModel{
		conn:  conn,
		table: "`sys_login_log`",
	}
}

func (m *defaultSysLoginLogModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysLoginLogModel) FindOne(ctx context.Context, id int64) (*SysLoginLog, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysLoginLogRows, m.table)
	var resp SysLoginLog
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

func (m *defaultSysLoginLogModel) Insert(ctx context.Context, data *SysLoginLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysLoginLogRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Uid, data.UserName, data.IpAddr, data.LoginLocation, data.Browser, data.Os, data.Code, data.Msg)
	return ret, err
}

func (m *defaultSysLoginLogModel) Update(ctx context.Context, data *SysLoginLog) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysLoginLogRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Uid, data.UserName, data.IpAddr, data.LoginLocation, data.Browser, data.Os, data.Code, data.Msg, data.Id)
	return err
}

func (m *defaultSysLoginLogModel) tableName() string {
	return m.table
}
