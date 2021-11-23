package logic

import (
	"mall-go/internal/data"
	"mall-go/internal/data/model"
	"mall-go/pkg/enum"
	"mall-go/pkg/utils"
	"time"

	"github.com/xushuhui/goal/core"
)

type CouponChecker struct {
	Coupon *model.Coupon
}

func NewCouponChecker(coupon *model.Coupon) *CouponChecker {
	return &CouponChecker{
		Coupon: coupon,
	}
}

func (c *CouponChecker) IsOk() (err error) {
	in := utils.IsInTime(time.Now(), c.Coupon.StartTime, c.Coupon.EndTime)
	if !in {
		err = core.ParamsError(core.InvalidParams)
		return
	}
	return
}

func (c *CouponChecker) CanBeUsed(skuOrderList []data.SkuOrder, serverTotalPrice float64) (err error) {
	var orderCategoryPrice float64
	var cids []int64

	if c.Coupon.WholeStore == 1 {
		orderCategoryPrice = serverTotalPrice
	} else {
		for _, v := range c.Coupon.Edges.Category {
			cids = append(cids, v.ID)
		}
		orderCategoryPrice = getSumByCategoryList(skuOrderList, cids)
	}
	err = c.couponCanBeUsed(orderCategoryPrice)

	return
}

func (c *CouponChecker) couponCanBeUsed(orderCategoryPrice float64) (err error) {
	switch c.Coupon.Type {
	case enum.FULL_MINUS:
	case enum.FULL_OFF:
		if c.Coupon.FullMoney > orderCategoryPrice {
			err = core.ParamsError(core.InvalidParams)
			return
		}
	case enum.NO_THRESHOLD_MINUS:
		return
	default:
		err = core.ParamsError(core.InvalidParams)

		return
	}
	return
}

func getSumByCategoryList(skuOrderList []data.SkuOrder, cids []int64) (sum float64) {
	for _, cid := range cids {
		sum = sum + getSumByCategory(skuOrderList, cid)
	}
	return
}

func getSumByCategory(skuOrderList []data.SkuOrder, cid int64) (sum float64) {
	for _, v := range skuOrderList {
		if v.CategoryId == cid {
			sum = sum + v.GetTotalPrice()
		}
	}
	return
}

func (c *CouponChecker) FinalTotalPriceIsOk(orderFinalTotalPrice float64, serverTotalPrice float64) (err error) {
	var serverFinalTotalPrice float64
	switch c.Coupon.Type {
	case enum.FULL_MINUS:
	case enum.NO_THRESHOLD_MINUS:
		serverFinalTotalPrice = serverTotalPrice - c.Coupon.Minus
		if serverFinalTotalPrice <= 0 {
			err = core.ParamsError(core.InvalidParams)
			return
		}

	case enum.FULL_OFF:
		//todo
		serverFinalTotalPrice = serverTotalPrice * c.Coupon.Rate

	default:
		err = core.ParamsError(core.InvalidParams)

		return
	}
	if serverFinalTotalPrice != orderFinalTotalPrice {
		err = core.ParamsError(core.InvalidParams)

		return
	}
	return
}
