package model

type User struct {
	uid        string
	name       string
	password   string
	mail       string
	url        string
	screenName string
	created    int32
	activated  int32
	logged     int32
	group      string
	authCode   string
}

func (*User) TableName() string {
	return "typecho_users"
}
