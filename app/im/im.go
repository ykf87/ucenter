package im

import (
	"errors"
	"ucenter/app/im/tencent"
)

type Im interface {
	GenUserSig(userid string, expire int) (string, error)
}

func Get(s string) (Im, error) {
	switch s {
	case "tencent":
		return tencent.GetClient()
	default:
		return nil, errors.New("No Im client found")
	}
}
