package models

import "main/pkg/domain"

type Wallet struct {
	Balance int                    `json:"balance"`
	History []domain.WalletHistory `json:"history"`
}
