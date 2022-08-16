package payment

import (
	"errors"
	"fmt"
	"ucenter/app/config"
	"ucenter/app/payment/paypal"

	P "github.com/go-pay/gopay/paypal"
)

type Payment interface {
	Pay(string, float64) (string, string, error)
	GetOrderDetail(string) (*P.OrderDetail, error)
}

func Get(lang, pm string) (p Payment, e error) {
	if pm == "" {
		pm = config.Config.Payment
	}

	_, ok := config.Config.Payments[pm]
	if !ok {
		e = errors.New(fmt.Sprintf("支付方式 [%s] 不存在!", pm))
		return
	}
	switch pm {
	case "paypal":
		return paypal.Client(lang)
	}
	return nil, errors.New("Initialization failure")
}
