package models

import (
	"time"
)

type AdminLogin struct {
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password" validate:"min=8,max=20"`
}

type AdminDetailsResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name" `
	Email string `json:"email" `
}
type TokenAdmin struct {
	Username     string
	RefreshToken string
	AccessToken  string
}

type UserDetailsAtAdmin struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Permission bool   `json:"permission"`
}

type CustomDates struct {
	StartingDate time.Time `json:"startingDate"`
	EndDate      time.Time `json:"endDate"`
}

type UserStats struct {
	TotalUsers            int
	AverageRevenuePerUser float32
	AverageOrderPerUSer   float32
}

type InventoryStats struct {
	MostSoldProduct string
	TrendingProduct string
	TotalProducts   int
	TotalStock      int
	AveragePrice    float32
}

type OrderStats struct {
	TotalOrders       int
	TotalRevenue      float32
	TodaysRevenue     float32
	AverageOrderValue float32
}
