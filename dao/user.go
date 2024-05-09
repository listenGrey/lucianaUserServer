package dao

import (
	"errors"
	"github.com/listenGrey/lucianagRpcPKG/user"
	"lucianaUserServer/model"
	"gorm.io/gorm"
)

// CheckEmail 检查用户是否存在
func CheckEmail(email string) (bool, error) {
	// 连接DB
	client := MySQLClient()
	if client == nil {
		return false, errors.New("无法连接到MySQL")
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

// Register 用户注册
func Register(u *model.User) error {
	// 连接DB
	client := MySQLClient()
	if client == nil {
		return errors.New("无法连接到MySQL")
	}

	// 插入数据
	info := client.Create(*u)
	if info.Error != nil {
		return info.Error
	}

	return nil
}

// Login 用户登录
func Login(l *user.LoginForm) (*model.LogInfo, error) {
	// 连接DB
	client := MySQLClient()
	if client == nil {
		return nil, errors.New("无法连接到MySQL")
	}

	var logInfo model.LogInfo
	var u model.User
	info := client.Where("email = ?", l.Email).First(&u)

	if info.Error != nil {
		if info.Error == gorm.ErrRecordNotFound{
			logInfo.Exist = false
			return &logInfo, nil
		}
		return nil, info.Error
	}

	// 检查email是否存在

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
