package model

import (
	"spectrum/common/pb"
	"time"
)

type Space struct {
	ID               int64          `gorm:"id"`
	CreatedAt        time.Time      `gorm:"created_at"`
	UpdatedAt        time.Time      `gorm:"updated_at"`
	Name             string         `json:"name"`       // 桌球、麻将...
	ClassName        string         `json:"class_name"` // 1、2、3...
	Price            float64        `json:"price"`
	BillingType      pb.BillingType `json:"billing_type"` // 计场、计时
	PictureStorePath string         `json:"picture_store_path"`
}

func (*Space) TableName() string {
	return "space"
}

func (s *Space) ToPb() *pb.Space {
	return &pb.Space{
		Id:               s.ID,
		Name:             s.Name,
		ClassName:        s.ClassName,
		Price:            GetPbPrice(s.Price),
		BillingType:      s.BillingType,
		PictureStorePath: s.PictureStorePath,
	}
}
