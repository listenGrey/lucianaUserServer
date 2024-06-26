package model

import "github.com/listenGrey/lucianagRpcPKG/user"

type User struct {
	UserID   int64  `json:"user_id" gorm:"primaryKey"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogInfo struct {
	UserId   int64
	UserName string
	Exist    bool
	Success  bool
}

func LogInfoMarshal(l *LogInfo) *user.LogInfo {
	var logInfo user.LogInfo

	logInfo.UserId = l.UserId
	logInfo.UserName = l.UserName
	logInfo.Exist = l.Exist
	logInfo.Success = l.Success

	return &logInfo
}
