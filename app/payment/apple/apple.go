package apple

import (
	"context"
	"errors"
	"strings"
	"ucenter/app/logs"

	"github.com/go-pay/gopay/apple"
	"github.com/go-pay/gopay/pkg/xlog"
)

const (
	VeryUrl        = "https://buy.itunes.apple.com/verifyReceipt"
	VeryUrlSendBox = "https://sandbox.itunes.apple.com/verifyReceipt"
	PASSWORD       = "e26f52cada2641fd88e2f53673e266ba"
)

func VeryOrder(receipt string) (string, error) {
	ctx := context.Background()
	rsp, err := apple.VerifyReceipt(ctx, VeryUrl, PASSWORD, receipt)
	if err != nil {
		xlog.Error(err)
		return "", err
	}
	/**
	  response body:
	  {"receipt":{"original_purchase_date_pst":"2021-08-14 05:28:17 America/Los_Angeles", "purchase_date_ms":"1628944097586", "unique_identifier":"13f339a765b706f8775f729723e9b889b0cbb64e", "original_transaction_id":"1000000859439868", "bvrs":"10", "transaction_id":"1000000859439868", "quantity":"1", "in_app_ownership_type":"PURCHASED", "unique_vendor_identifier":"6DFDEA8B-38CE-4710-A1E1-BAEB8B66FEBD", "item_id":"1581250870", "version_external_identifier":"0", "bid":"com.huochai.main", "is_in_intro_offer_period":"false", "product_id":"10002", "purchase_date":"2021-08-14 12:28:17 Etc/GMT", "is_trial_period":"false", "purchase_date_pst":"2021-08-14 05:28:17 America/Los_Angeles", "original_purchase_date":"2021-08-14 12:28:17 Etc/GMT", "original_purchase_date_ms":"1628944097586"}, "status":0}
	*/

	if rsp.Status == 21007 {
		rsp, err = apple.VerifyReceipt(ctx, VeryUrlSendBox, PASSWORD, receipt)
		if err != nil {
			xlog.Error(err)
			return "", err
		}
	}

	if rsp.Receipt != nil {
		// b, _ := json.Marshal(rsp)
		// fmt.Println(string(b))
		// xlog.Infof("receipt:%+v", rsp.Receipt)
		if rsp.Status == 0 {
			var str []string
			for _, v := range rsp.Receipt.InApp {
				str = append(str, v.TransactionId)
			}
			return strings.Join(str, ","), nil
		} else {
			return "", errors.New("Status code is not 0")
		}
	} else {
		// b, _ := json.Marshal(rsp)
		// fmt.Println(string(b))
		logs.Logger.Error("支付错误: ", receipt)
		return "", errors.New("支付错误")
	}
	return "", errors.New("Error")
}
