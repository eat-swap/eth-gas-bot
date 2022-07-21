package models

import "eth-gas-bot/modules/telegram/models/entities"

type Response[T ResponseResult] struct {
	Ok          bool   `json:"ok"`
	Result      *T     `json:"result,omitempty"`
	ErrorCode   int    `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`
}

type ResponseResult interface {
	entities.Message
}
