package models

type Repo struct {
	Model

	ServiceID int64 `sql:"not null;unique_index:idx_serviceid_name"`

	Name        string `sql:"not null;unique_index:idx_serviceid_name"`
	Description string
	Private     bool
	Fork        bool
	URL         string
	Homepage    string
}
