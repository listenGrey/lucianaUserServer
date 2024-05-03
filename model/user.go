package model

import "github.com/listenGrey/lucianagRpcPKG/user"

type User struct {
	UserID   int64 `gorm:"primaryKey"`
	UserName string
	Email    string
	Password string
}

func UserUnmarshal(u *user.RegisterForm) User {
	var newUser User

	newUser.Email = u.GetEmail()
	newUser.UserID = u.GetUserId()
	newUser.UserName = u.GetUserName()
	newUser.Password = u.GetPassword()

	return newUser
}

type LogInfo struct {
	UserId   int64
	UserName string
	Exist    bool
	Success  bool
}

func LogInfoMarshal(l *LogInfo, serverErr bool) *user.LogInfo {
	var logInfo user.LogInfo

	logInfo.UserId = l.UserId
	logInfo.UserName = l.UserName
	logInfo.Exist = l.Exist
	logInfo.Success = l.Success
	logInfo.ServerError = serverErr

	return &logInfo
}
