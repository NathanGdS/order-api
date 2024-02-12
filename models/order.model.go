package models

import (
	"time"

	"github.com/nathangds/order-api/helpers"
)

type NewOrderRequest struct {
	Value         int64    `json:"value" validate:"required"`
	DiscountValue float64  `json:"discount_value"`
	DeliveryFee   float64  `json:"delivery_fee"`
	ItemsIds      []string `json:"items_ids" validate:"required"`
}

type ResponseOrder struct {
	OrderId       string    `json:"order_id"`
	OriginalValue float64   `json:"value"`
	DiscountValue float64   `json:"discount_value"`
	DeliveryFee   float64   `json:"delivery_fee"`
	TotalValue    float64   `json:"total_value"`
	PaidAt        time.Time `json:"paid_at"`
	DeliveryAt    time.Time `json:"delivery_at"`
	BaseModels
}

type Order struct {
	OrderId       string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"order_id"`
	OriginalValue int64     `json:"value" validate:"required"`
	DiscountValue int64     `json:"discount_value" validate:"optional"`
	DeliveryFee   int64     `json:"delivery_fee" validate:"optional"`
	TotalValue    int64     `json:"total_value" validate:"required"`
	PaidAt        time.Time `json:"paid_at" validate:"optional"`
	DeliveryAt    time.Time `json:"delivery_at" validate:"optional"`
	Items         []Item    `gorm:"foreignKey:OrderId;references:OrderId" json:"items"`

	BaseModels
}

func (o *Order) NewOrder(discountValue float64, deliveryFee float64, items []Item) *Order {
	var originalValue int64
	for i := range items {
		originalValue += items[i].Value
	}

	var deliveryFeeInCents = helpers.DecimalToCents(deliveryFee)
	var discountValueInCents = helpers.DecimalToCents(discountValue)
	var sum = originalValue + deliveryFeeInCents - discountValueInCents

	return &Order{
		OriginalValue: originalValue,
		DiscountValue: discountValueInCents,
		DeliveryFee:   deliveryFeeInCents,
		TotalValue:    sum,
	}
}

func (o *Order) ShowOrder() ResponseOrder {
	return ResponseOrder{
		OrderId:       o.OrderId,
		OriginalValue: helpers.CentsToDecimal(o.OriginalValue),
		DiscountValue: helpers.CentsToDecimal(o.DiscountValue),
		DeliveryFee:   helpers.CentsToDecimal(o.DeliveryFee),
		TotalValue:    helpers.CentsToDecimal(o.TotalValue),
		PaidAt:        o.PaidAt,
		DeliveryAt:    o.DeliveryAt,
		BaseModels:    o.BaseModels,
	}
}
