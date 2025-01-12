package dao

import (
	"database/sql"
	"fmt"
	"github.com/listenGrey/lucianagRpcPKG/user"
	"lucianaUserServer/model"
)

// CheckEmail 检查用户是否存在
func CheckEmail(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)"
	err := PostgreConn.QueryRow(query, email).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("error querying the database: %w", err)
	}

	// 返回是否存在
	return exists, nil
}

// Login 用户登录
func Login(l *user.LoginForm) (*model.LogInfo, error) {
	// 返回的结构
	var logInfo model.LogInfo
	// 查询出来的结果
	var u model.User

	// 执行查询
	query := "SELECT user_id, user_name, password FROM users WHERE email=$1"
	row := PostgreConn.QueryRow(query, l.Email)

	err := row.Scan(&u.UserID, &u.UserName, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			logInfo.Exist = false
			return &logInfo, nil
		}
		return nil, fmt.Errorf("error querying the database: %w", err)
	}

	// 检查密码是否匹配
	if u.Password != l.Password {
		logInfo.Exist = true
		logInfo.Success = false
		return &logInfo, nil
	}

	logInfo.UserId = u.UserID
	logInfo.UserName = u.UserName
	logInfo.Exist = true
	logInfo.Success = true
	return &logInfo, nil
}
