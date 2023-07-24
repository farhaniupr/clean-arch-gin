package models

import "time"

type Author struct {
	Id         int
	Name       string `form:"name" json:"name"`
	Created_at time.Time
	Updated_at time.Time
}

func (a Author) TableName() string {
	return "author"
}
