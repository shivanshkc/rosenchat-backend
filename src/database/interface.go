package database

import (
	"sync"
)

var userInfoOnce = &sync.Once{}
var userInfoSingleton IUserInfoDB

// IUserInfoDB represents the database layer for User Info.
type IUserInfoDB interface {
	// GetUserInfo fetches the user info for the provided userID.
	GetUserInfo(userID string) (*UserInfoDTO, error)

	// PutUserInfo puts the user's info in the database.
	PutUserInfo(info *UserInfoDTO) error

	// init can be used to initialize the implementation.
	init()
}

// Get provides the IUserInfoDB singleton.
func Get() IUserInfoDB {
	userInfoOnce.Do(func() {
		userInfoSingleton = &implMongoUserInfo{}
		userInfoSingleton.init()
	})

	return userInfoSingleton
}
