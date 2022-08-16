package paypal

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"ucenter/app/config"
	"ucenter/app/funcs"

	"github.com/go-pay/gopay"
	P "github.com/go-pay/gopay/paypal"
	"github.com/go-pay/gopay/pkg/util"

	// "github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/gopay/pkg/xlog"
)

const (
	Success = 0

	HeaderAuthorization       = "Authorization" // 请求头Auth
	AuthorizationPrefixBasic  = "Basic "
	AuthorizationPrefixBearer = "Bearer "

	baseUrlProd    = "https://api-m.paypal.com"         // 正式 URL
	baseUrlSandbox = "https://api-m.sandbox.paypal.com" // 沙箱 URL

	// 获取AccessToken
	getAccessToken = "/v1/oauth2/token" // 获取AccessToken POST

	// 订单相关
	orderCreate    = "/v2/checkout/orders"                           // 创建订单 POST
	orderUpdate    = "/v2/checkout/orders/%s"                        // order_id 更新订单 PATCH
	orderDetail    = "/v2/checkout/orders/%s"                        // order_id 订单详情 GET
	orderAuthorize = "/v2/checkout/orders/%s/authorize"              // order_id 订单支付授权 POST
	orderCapture   = "/v2/checkout/orders/%s/capture"                // order_id 订单支付捕获 POST
	orderConfirm   = "/v2/checkout/orders/%s/confirm-payment-source" // order_id 订单支付确认 POST

	// 支付相关
	paymentAuthorizeDetail  = "/v2/payments/authorizations/%s"             // authorization_id 支付授权详情 GET
	paymentAuthorizeCapture = "/v2/payments/authorizations/%s/capture"     // authorization_id 支付授权捕获 POST
	paymentReauthorize      = "/v2/payments/authorizations/%s/reauthorize" // authorization_id 重新授权支付授权 POST
	paymentAuthorizeVoid    = "/v2/payments/authorizations/%s/void"        // authorization_id 作废支付授权 POST
	paymentCaptureDetail    = "/v2/payments/captures/%s"                   // capture_id 支付捕获详情 GET
	paymentCaptureRefund    = "/v2/payments/captures/%s/refund"            // capture_id 支付捕获退款 POST
	paymentRefundDetail     = "/v2/payments/refunds/%s"                    // refund_id 支付退款详情 GET

	// 支出相关
	createBatchPayout         = "/v1/payments/payouts"                // 创建批量支出 POST
	showPayoutBatchDetail     = "/v1/payments/payouts/%s"             // payout_batch_id 获取批量支出详情 GET
	showPayoutItemDetail      = "/v1/payments/payouts-item/%s"        // payout_item_id 获取支出项目详情 GET
	cancelUnclaimedPayoutItem = "/v1/payments/payouts-item/%s/cancel" // payout_item_id 取消支出项目 POST

	// 订阅相关
	subscriptionCreate = "/v1/billing/plans" // 创建订阅 POST

)

type Pp struct {
	Clientid    string `json:"clientid"`
	Secret      string `json:"secret"`
	ReturnUrl   string `json:"return_url"`
	CancelUrl   string `json:"cancel_url"`
	Lang        string `json:"lang"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
	IsProd      bool   `json:"-"`
	DebugSwitch int8   `json:"-"`
}

type AccessToken struct {
	Scope       string `json:"scope"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Appid       string `json:"app_id"`
	ExpiresIn   int    `json:"expires_in"`
	Nonce       string `json:"nonce"`
}

func Client(lang string) (c *Pp, err error) {
	cres, ok := config.Config.Payments["paypal"]
	if !ok {
		err = errors.New("Please set Paypal config")
		return
	}
	// rrs, errs := P.NewClient(cres.Appid, cres.Secret, false)
	// if errs != nil {
	// 	xlog.Error(errs)
	// 	err = errs
	// 	return
	// }
	// rrs.IsProd = true
	// rrs.DebugSwitch = gopay.DebugOn

	cp := new(Pp)
	cp.Clientid = cres.Appid
	cp.Secret = cres.Secret
	// cp.Client = rrs
	cp.Lang = lang
	cp.ReturnUrl = cres.ReturnUrl
	cp.CancelUrl = cres.CancelUrl
	cp.IsProd = true
	cp.DebugSwitch = gopay.DebugOn

	_, err = cp.GetAccessToken()
	if err != nil {
		return
	}

	c = cp
	return
}

// 获取AccessToken（Get an access token）
func (c *Pp) GetAccessToken() (token *AccessToken, err error) {
	baseUrl := baseUrlProd
	// var (
	// // baseUrl = baseUrlProd
	// // url     string
	// )
	if !c.IsProd {
		baseUrl = baseUrlSandbox
	}
	url := baseUrl + getAccessToken
	// Authorization
	authHeader := AuthorizationPrefixBasic + base64.StdEncoding.EncodeToString([]byte(c.Clientid+":"+c.Secret))

	header := make(map[string]string)
	header["Authorization"] = authHeader
	header["Accept"] = "*/*"
	header["Content-Type"] = "application/json"
	// Body
	body := `grant_type=client_credentials`
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("PayPal_RequestBody: %s", body)
		xlog.Debugf("PayPal_Authorization: %s", authHeader)
	}
	resp, err := funcs.Request("POST", url, []byte(body), header, "")
	// fmt.Println(string(resp), err)

	if err != nil {
		return nil, err
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("PayPal_Response: %s", string(resp))
		// xlog.Debugf("PayPal_Headers: %#v", res.Header)
	}
	token = new(AccessToken)
	if err = json.Unmarshal(resp, token); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(resp), err)
	}
	c.AccessToken = token.AccessToken
	c.ExpiresIn = token.ExpiresIn
	// res, bs, err := httpClient.Type(xhttp.TypeForm).Post(url).SendBodyMap(bm).EndBytes(c.ctx)
	// if err != nil {
	// 	return nil, err
	// }
	// if c.DebugSwitch == gopay.DebugOn {
	// 	xlog.Debugf("PayPal_Response: %d > %s", res.StatusCode, string(bs))
	// 	xlog.Debugf("PayPal_Headers: %#v", res.Header)
	// }
	// if res.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	// }
	// token = new(AccessToken)
	// if err = json.Unmarshal(bs, token); err != nil {
	// 	return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	// }
	// c.Appid = token.Appid
	// c.AccessToken = token.AccessToken
	// c.ExpiresIn = token.ExpiresIn
	return
}

func (c *Pp) Pay(currency string, price float64) (orderid string, urls string, errs error) {
	var url = baseUrlProd + orderCreate
	if !c.IsProd {
		url = baseUrlSandbox + orderCreate
	}
	authHeader := AuthorizationPrefixBearer + c.AccessToken

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
				Set("locale", c.Lang).
				Set("return_url", c.ReturnUrl).
				Set("cancel_url", c.CancelUrl)
		})
	body := bm.JsonBody()
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("PayPal_RequestBody: %s", body)
		xlog.Debugf("PayPal_Authorization: %s", authHeader)
	}

	header := make(map[string]string)
	header["Authorization"] = authHeader
	header["Accept"] = "*/*"
	header["Content-Type"] = "application/json"

	bs, err := funcs.Request("POST", url, []byte(body), header, "")

	if err != nil {
		errs = err
		fmt.Println("---", err)
		return
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("PayPal_Response: %s", string(bs))
	}

	orderDetail := new(P.OrderDetail)
	if err = json.Unmarshal(bs, orderDetail); err != nil {
		errs = fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
		return
	}
	orderid = orderDetail.Id
	for _, v := range orderDetail.Links {
		if v.Rel == "approve" {
			urls = v.Href
			return
		}
	}
	orderid = ""
	errs = errors.New("error")
	fmt.Println(*orderDetail)
	//var ctx context.Background()
	// ppRsp, err := this.Client.CreateOrder(context.Background(), bm)
	// if err != nil {
	// 	xlog.Error(err)
	// 	errs = err
	// 	return
	// }
	// if ppRsp.Code != P.Success {
	// 	errs = errors.New("error")
	// 	return
	// }
	// orderid = ppRsp.Response.Id
	// for _, v := range ppRsp.Response.Links {
	// 	if v.Rel == "approve" {
	// 		url = v.Href
	// 		return
	// 	}
	// }
	// orderid = ""
	// errs = errors.New("error")
	return
}

//获取订单详情
func (this *Pp) GetOrderDetail(orderid string) (odt *P.OrderDetail, err error) {
	// rrs, err := P.NewClient(this.Clientid, this.Secret, false)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// resp, err := rrs.OrderDetail(context.Background(), orderid, nil)
	// fmt.Println(resp, err)
	if orderid == "" {
		err = errors.New("Order Id is empty")
		return
	}

	uri := fmt.Sprintf(orderDetail, orderid)
	od, errs := this.get(uri)
	if errs != nil {
		err = errs
		return
	}
	odt = od
	return
}

//发起请求
func (c *Pp) get(uri string) (ress *P.OrderDetail, errs error) {
	var url = baseUrlProd + uri
	if !c.IsProd {
		url = baseUrlSandbox + uri
	}
	authHeader := AuthorizationPrefixBearer + c.AccessToken
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("PayPal_Url: %s", url)
		xlog.Debugf("PayPal_Authorization: %s", authHeader)
	}
	header := map[string]string{
		HeaderAuthorization: authHeader,
		"Accept":            "*/*",
	}

	rs, err := funcs.Request("GET", url, nil, header, "")
	if err != nil {
		errs = err
		return
	}
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("PayPal_Body: %s", string(rs))
	}

	orderDetail := new(P.OrderDetail)
	if err = json.Unmarshal(rs, orderDetail); err != nil {
		errs = fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(rs))
		return
	}
	ress = orderDetail
	return
}
