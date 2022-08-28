package redis

import (
	"context"
	"errors"
	"time"

	rds "github.com/go-redis/redis/v9"
)

var RDB *rds.Client
var ctx = context.Background()

func Init(addr, pwd string, db int) error {
	rdb := rds.NewClient(&rds.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
	// key := "_tmp___"
	// err := rdb.Set(ctx, key, "a", time.Second*30).Err()
	// if err != nil {
	// 	return err
	// }
	// str, err := rdb.Get(ctx, key).Result()
	// if err != nil {
	// 	return err
	// }
	// if str != "a" {
	// 	return errors.New("获取内容失败，redis 可能链接不正常!")
	// }
	// rdb.Del(ctx, key)
	RDB = rdb
	return nil
}

func Get(key string) (string, error) {
	str, err := RDB.Get(ctx, key).Result()
	if err == rds.Nil {
		return "", errors.New("Empty")
	} else if err != nil {
		return "", errors.New("Empty")
	} else if str == "" {
		return "", errors.New("Empty")
	}
	return str, nil
}

func Set(k string, v interface{}, expir int) error {
	return RDB.Set(ctx, k, v, time.Second*time.Duration(expir)).Err()
}

func Del(keys ...string) error {
	return RDB.Del(ctx, keys...).Err()
}

func Ttl(key string) (int, error) {
	ttl, err := RDB.TTL(ctx, key).Result()
	if err == rds.Nil {
		return 0, errors.New("error")
	} else if err != nil {
		return 0, errors.New("error")
	}
	return int(ttl / time.Second), nil
}
