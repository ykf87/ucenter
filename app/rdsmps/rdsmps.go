package rdsmps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"ucenter/app/dbs/redis"
	"ucenter/app/logs"
)

type Mmpp struct {
	Prev    string `json:"prev"`
	Timeout int    `json:"timeout"`
}

//生成key
func (this *Mmpp) k(str string) string {
	return this.Prev + str
}

//设置内容,存在将重置
func (this *Mmpp) Set(kstr, val string, times int, reset bool) error {
	val = fmt.Sprintf("%s:%d:%d", val, times, time.Now().Unix())
	key := this.k(kstr)

	ss, err := redis.Get(key)
	if err == nil && ss != "" {
		if reset == false {
			logs.Logger.Error(fmt.Sprintf("设置的内容 %s 已经存在!", kstr))
			return errors.New("Existed")
		}
		redis.Del(key)
	}

	err = redis.Set(key, val, this.Timeout)
	if err != nil {
		logs.Logger.Error(err)
	}
	return err
}

//获取内容
func (this *Mmpp) Get(kstr string) (val string, times int, ttl int, addtime int64, err error) {
	key := this.k(kstr)
	ss, errs := redis.Get(key)
	if errs != nil {
		err = errs
		return
	}

	sps := strings.Split(ss, ":")
	if len(sps) != 3 {
		logs.Logger.Error(fmt.Sprintf("获取的内容不正经! %s - %s", key, ss))
		err = errors.New("Error")
		return
	}
	times, _ = strconv.Atoi(sps[1])
	val = sps[0]
	ttl, _ = redis.Ttl(key)
	atstr, _ := strconv.Atoi(sps[2])
	addtime = int64(atstr)
	return
}

//错误次数加1
func (this *Mmpp) Increment(kstr string) error {
	val, errtimes, ttl, addtime, err := this.Get(kstr)
	if err != nil {
		return err
	}

	key := this.k(kstr)
	val = fmt.Sprintf("%s:%d:%d", val, (errtimes + 1), addtime)
	return redis.Set(key, val, ttl)
}
