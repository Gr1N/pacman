package models

type Service struct {
	Model

	UserID int64 `sql:"not null;unique_index:idx_userid_userserviceid"`

	Name string `sql:"not null;index"`

	UserServiceID    int64 `sql:"not null;unique_index:idx_userid_userserviceid"`
	UserServiceName  string
	UserServiceEmail string
}
