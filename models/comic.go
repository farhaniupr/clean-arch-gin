package models

import "time"

type Comic struct {
	Id                         int
	Title                      string  `form:"title" json:"title"`
	Desc                       string  `form:"desc" json:"desc"`
	Thumbnail                  string  `form:"thumbnail" json:"thumbnail"`
	Total_like                 int     `form:"total_like" json:"total_like"`
	Total_views                int     `form:"total_views" json:"total_views"`
	Vote                       int     `form:"vote" json:"vote"`
	Member_id                  int     `form:"member_id" json:"member_id"`
	Approve                    int     `form:"approve" json:"approve"`
	Is_18                      int     `form:"is_18" json:"is_18"`
	Status_episode             int     `form:"status_episode" json:"status_episode"`
	Paid                       int     `form:"paid" json:"paid"`
	Price                      int     `form:"price" json:"price"`
	Price_discount             int     `form:"price_discount" json:"price_discount"`
	Price_discount_percent     float64 `form:"price_discount_percent" json:"price_discount_percent"`
	Price_ios                  int     `form:"price_ios" json:"price_ios"`
	Price_discount_ios         int     `form:"price_discount_ios" json:"price_discount_ios"`
	Price_discount_ios_percent float64 `form:"price_discount_ios_percent" json:"price_discount_ios_percent"`
	Created_at                 time.Time
	Created_at_wib             time.Time
	Updated_at                 time.Time
}

func (c Comic) TableName() string {
	return "Comic"
}
