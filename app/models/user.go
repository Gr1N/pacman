package models

type User struct {
	Model

	Services []Service
}

type Service struct {
	Model

	UserId int64 `sql:"not null;unique_index:idx_userid_userserviceid"`

	Name string `sql:"not null;index"`

	UserServiceId    int64 `sql:"not null;unique_index:idx_userid_userserviceid"`
	UserServiceName  string
	UserServiceEmail string
}
