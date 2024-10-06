package mq

import (
	"github.com/listenGrey/lucianagRpcPKG/user"
	"testing"
)

func TestRegisterQueue(t *testing.T) {
	re := &user.RegisterFrom{
		Id:       123456,
		Email:    "test@911.com",
		Name:     "test register",
		Password: "test password",
	}
	err := RegisterQueue(re)
	if err != nil {
		t.Error(err)
	}
}
