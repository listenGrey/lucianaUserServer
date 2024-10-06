package dao

import (
	"github.com/listenGrey/lucianagRpcPKG/user"
	"gorm.io/gorm"
	"lucianaUserServer/model"
)

// CheckEmail 检查用户是否存在
func CheckEmail(email string) (bool, error) {
	// 连接DB
	client, err := GormDB()
	if err != nil {
		return false, err
	}

	// 查找用户
	var u model.User
	info := client.Where("email = ?", email).First(&u)
	if info.Error != nil {
		if info.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, info.Error
	}

	return true, nil
}

// Login 用户登录
func Login(l *user.LoginForm) (*model.LogInfo, error) {
	// 连接DB
	client, err := GormDB()
	if err != nil {
		return nil, err
	}

	var logInfo model.LogInfo
	var u model.User
	info := client.Where("email = ?", l.Email).First(&u)

	// 检查email是否存在
	if info.Error != nil {
		if info.Error == gorm.ErrRecordNotFound {
			logInfo.Exist = false
			return &logInfo, nil
		}
		return nil, info.Error
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
