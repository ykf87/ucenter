package paypal

import (
	"context"
	"errors"
	"fmt"
	"ucenter/app/config"

	"github.com/go-pay/gopay"
	P "github.com/go-pay/gopay/paypal"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xlog"
)

type Pp struct {
	Clientid  string `json:"clientid"`
	Secret    string `json:"secret"`
	ReturnUrl string `json:"return_url"`
	CancelUrl string `json:"cancel_url"`
	Lang      string `json:"lang"`
	Client    *P.Client
}

func Client(lang string) (c *Pp, err error) {
	cres, ok := config.Config.Payments["paypal"]
	if !ok {
		err = errors.New("Please set Paypal config")
		return
	}
	rrs, errs := P.NewClient(cres.Appid, cres.Secret, false)
	if errs != nil {
		xlog.Error(errs)
		err = errs
		return
	}
	rrs.DebugSwitch = gopay.DebugOn

	cp := new(Pp)
	cp.Clientid = cres.Appid
	cp.Secret = cres.Secret
	cp.Client = rrs
	cp.Lang = lang
	cp.ReturnUrl = cres.ReturnUrl
	cp.CancelUrl = cres.CancelUrl

	c = cp
	return
}

func (this *Pp) Pay(currency string, price float64) (string, error) {
	var pus []*P.PurchaseUnit

	var item = &P.PurchaseUnit{
		ReferenceId: util.RandomString(16),
		Amount: &P.Amount{
			CurrencyCode: currency,
			Value:        fmt.Sprintf("%0.2f", price),
		},
	}
	pus = append(pus, item)

	bm := make(gopay.BodyMap)
	bm.Set("intent", "CAPTURE").
		Set("purchase_units", pus).
		SetBodyMap("application_context", func(b gopay.BodyMap) {
			b.Set("brand_name", config.Config.APPName).
				Set("locale", this.Lang).
				Set("return_url", this.ReturnUrl).
				Set("cancel_url", this.CancelUrl)
		})

	//var ctx context.Background()
	ppRsp, err := this.Client.CreateOrder(context.Background(), bm)
	if err != nil {
		xlog.Error(err)
		return "", err
	}
	if ppRsp.Code != P.Success {
		return "", errors.New("error")
	}
	for _, v := range ppRsp.Response.Links {
		if v.Rel == "approve" {
			return v.Href, nil
		}
	}
	return "", errors.New("error")
}
