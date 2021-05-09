package model

import (
	"spectrum/common/pb"
	"spectrum/service/mvp/utils"
	"time"
)

type Good struct {
	ID              int64     `gorm:"id"`
	CreatedAt       time.Time `gorm:"created_at"`
	UpdatedAt       time.Time `gorm:"updated_at"`
	OrderID         int64     `gorm:"order_id"`
	MainElementID   int64     `gorm:"main_element_id"`
	Expense         float64   `gorm:"expense"`
	CheckOutAt      time.Time `gorm:"check_out_at"`
	NonFavorExpense float64   `gorm:"non_favor_expense"`
}

func (g *Good) GetID() int64 {
	return g.ID
}

func (g *Good) TableName() string {
	return g.GetChargeableObjectName()
}

func (*Good) GetChargeableObjectName() string {
	return utils.ChargeableObjectNameOfGood
}

func (g *Good) GetExpenseInfo(mainElement *pb.Element, attachElement []*pb.Element, favors []*pb.Favor) *pb.ExpenseInfo {
	if g.CheckOutAt != utils.NilTime {
		return &pb.ExpenseInfo{
			NonFavorExpense: g.NonFavorExpense,
			CheckOutAt:      g.CheckOutAt.Unix(),
			Expense:         g.Expense,
		}
	}
	nonFavorExpense := g.getNonFavorExpense(append(attachElement, mainElement))
	return &pb.ExpenseInfo{
		NonFavorExpense: nonFavorExpense,
		CheckOutAt:      utils.NilTime.Unix(),
		Expense:         GetFavorExpense(nonFavorExpense, favors),
	}
}

func (g *Good) getNonFavorExpense(elements []*pb.Element) float64 {
	expense := 0.0
	for _, element := range elements {
		priceString := GetPbElementSelectSizeInfo(element).Price
		expense += utils.GetDbPrice(priceString)
	}
	return expense
}
