package contract

import (
	"generate_stream_currency/pkg/model"
)

type ApiService interface {
	Run() error
	GetCurrency() model.Currency
}
