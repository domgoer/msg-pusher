/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : putter.go
#   Created       : 2019/1/11 17:03
#   Last Modified : 2019/1/11 17:03
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"github.com/jmoiron/sqlx"
	"uuabc.com/sendmsg/api/model"
	"uuabc.com/sendmsg/api/storer"
)

// EmailCancelMsgByID 将email信息的发送状态设置为取消
func EmailCancelMsgByID(ctx context.Context, id string) (*sqlx.Tx, error) {
	tx, err := storer.DB.Beginx()
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, "UPDATE emails SET status=2 WHERE id=?")
	if err != nil {
		return tx, err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return tx, err
	}
	if i, _ := res.RowsAffected(); i == 0 {
		return tx, ErrNoRowsEffected
	}
	return tx, nil
}

// EmailDetailByID 按照id查询email所有字段信息，如果未找到返回error
func EmailDetailByID(ctx context.Context, id string) (*model.DbEmail, error) {
	res := &model.DbEmail{}
	err := storer.DB.GetContext(ctx, res, "SELECT * FROM emails WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// InsertEmails 将消息插入emails表
func InsertEmails(ctx context.Context, email *model.DbEmail) (*sqlx.Tx, error) {
	tx, err := storer.DB.Beginx()
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, `INSERT INTO emails (id,platform,platform_key,title,content,destination,type,template,arguments,server,send_time) VALUES (?,?,?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		return tx, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(
		ctx,
		email.ID,
		email.Platform,
		email.PlatformKey,
		email.Title,
		email.Content,
		email.Destination,
		email.Type,
		email.Template,
		email.Arguments,
		email.Server,
		email.SendTime,
	)
	if err != nil {
		return tx, err
	}
	return tx, nil
}
