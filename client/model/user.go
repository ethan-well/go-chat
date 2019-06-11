package model

type User struct{}

type CurrentUserInfo struct {
	UserID   int
	UserName string
}

var CurrentUser CurrentUserInfo

func (this User) InitCurrentUser(userID int, userName string) (err error) {
	CurrentUser = CurrentUserInfo{
		UserID:   userID,
		UserName: userName,
	}

	return
}
