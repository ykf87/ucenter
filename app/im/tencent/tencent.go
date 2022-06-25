package tencent

import (
	"errors"
	"strconv"
	"ucenter/app/config"

	"github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"
)

type Client struct {
	Id  int
	Key string
}

func GetClient() (cli *Client, err error) {
	c, ok := config.Config.Im["tencent"]
	if !ok {
		err = errors.New("No config for tencent im")
		return
	}
	cli = new(Client)
	cli.Id, _ = strconv.Atoi(c.Id)
	cli.Key = c.Key
	return
}

func (this *Client) GenUserSig(userid string, expire int) (string, error) {
	sig, err := tencentyun.GenUserSig(this.Id, this.Key, userid, expire)
	if err != nil {
		return "", err
	}
	return sig, nil
}
