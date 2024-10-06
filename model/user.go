package model

import "github.com/listenGrey/lucianagRpcPKG/user"

type User struct {
	UserID   int64  `json:"user_id" gorm:"column:user_id; primaryKey"`
	UserName string `json:"user_name" gorm:"column:user_name"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
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

func UserUnmarshal(re *user.RegisterFrom) *User {
	var res User

	res.UserID = re.Id
	res.Email = re.Email
	res.UserName = re.Name
	res.Password = re.Password

	return &res
}
