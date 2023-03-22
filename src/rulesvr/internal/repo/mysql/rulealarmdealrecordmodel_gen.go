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
	ruleAlarmDealRecordFieldNames          = builder.RawFieldNames(&RuleAlarmDealRecord{})
	ruleAlarmDealRecordRows                = strings.Join(ruleAlarmDealRecordFieldNames, ",")
	ruleAlarmDealRecordRowsExpectAutoSet   = strings.Join(stringx.Remove(ruleAlarmDealRecordFieldNames, "`id`", "`createdTime`", "`deletedTime`", "`updatedTime`"), ",")
	ruleAlarmDealRecordRowsWithPlaceHolder = strings.Join(stringx.Remove(ruleAlarmDealRecordFieldNames, "`id`", "`createdTime`", "`deletedTime`", "`updatedTime`"), "=?,") + "=?"
)

type (
	ruleAlarmDealRecordModel interface {
		Insert(ctx context.Context, data *RuleAlarmDealRecord) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*RuleAlarmDealRecord, error)
		Update(ctx context.Context, data *RuleAlarmDealRecord) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRuleAlarmDealRecordModel struct {
		conn  sqlx.SqlConn
		table string
	}

	RuleAlarmDealRecord struct {
		Id          int64     `db:"id"`          // 编号
		AlarmID     int64     `db:"alarmID"`     // 告警配置ID
		Result      string    `db:"result"`      // 告警处理结果
		Type        int64     `db:"type"`        // 告警处理类型（1人工 2系统）
		AlarmTime   time.Time `db:"alarmTime"`   // 告警时间
		CreatedTime time.Time `db:"createdTime"` // 告警处理时间
	}
)

func newRuleAlarmDealRecordModel(conn sqlx.SqlConn) *defaultRuleAlarmDealRecordModel {
	return &defaultRuleAlarmDealRecordModel{
		conn:  conn,
		table: "`rule_alarm_deal_record`",
	}
}

func (m *defaultRuleAlarmDealRecordModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultRuleAlarmDealRecordModel) FindOne(ctx context.Context, id int64) (*RuleAlarmDealRecord, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ruleAlarmDealRecordRows, m.table)
	var resp RuleAlarmDealRecord
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

func (m *defaultRuleAlarmDealRecordModel) Insert(ctx context.Context, data *RuleAlarmDealRecord) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, ruleAlarmDealRecordRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.AlarmID, data.Result, data.Type, data.AlarmTime)
	return ret, err
}

func (m *defaultRuleAlarmDealRecordModel) Update(ctx context.Context, data *RuleAlarmDealRecord) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ruleAlarmDealRecordRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.AlarmID, data.Result, data.Type, data.AlarmTime, data.Id)
	return err
}

func (m *defaultRuleAlarmDealRecordModel) tableName() string {
	return m.table
}