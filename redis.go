package go_util

import (
	"github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
)

type RedisUtil struct {
	redisPool *redis.Pool
}

func NewRedisUtil(rp *redis.Pool) *RedisUtil{
	return &RedisUtil{redisPool:rp}
}

func (r *RedisUtil)Set(k, v string) {
	c := r.redisPool.Get()
	defer c.Close()
	_, err := c.Do("SET", k, v)
	if err != nil {
		Error("set error", err.Error())
	}
}

func (r *RedisUtil) SetInt(k string, v int){
	c := r.redisPool.Get()
	defer c.Close()
	_, err := c.Do("SET", k, v)
	if err != nil {
		Error("set error", err.Error())
	}
}

func (r *RedisUtil)SetEx(k, v string, ex int64) {
	c := r.redisPool.Get()
	defer c.Close()
	_, err := c.Do("SETEX", k,ex, v)
	if err != nil {
		Error("setex err: ", err.Error())
	}
}

func (r *RedisUtil)GetStringValue(k string) string {
	c := r.redisPool.Get()
	defer c.Close()
	username, err := redis.String(c.Do("GET", k))
	if err != nil {
		Error("GetStringValue err: ", err.Error())
		return ""
	}
	return username
}

func (r *RedisUtil)GetIntValue(k string) int{
	c := r.redisPool.Get()
	defer c.Close()
	username, err := redis.Int(c.Do("GET",k))
	if err != nil {
		Error("GetStringValue err: ", err.Error())
		return 0
	}
	return username
}

//ex seconds
func (r *RedisUtil)SetKeyExpire(k string, ex int) {
	c := r.redisPool.Get()
	defer c.Close()
	_, err := c.Do("EXPIRE", k, ex)
	if err != nil {
		Error("set key expire err: ", err.Error())
	}
}

func (r *RedisUtil)CheckKey(k string) bool {
	c := r.redisPool.Get()
	defer c.Close()
	exist, err := redis.Bool(c.Do("EXISTS", k))
	if err != nil {
		Error("CheckKey err: ",err)
		return false
	} else {
		return exist
	}
}

func (r *RedisUtil)DelKey(k string) error {
	c := r.redisPool.Get()
	defer c.Close()
	_, err := c.Do("DEL", k)
	if err != nil {
		Error("DelKey err: ",err)
		return err
	}
	return nil
}

func (r *RedisUtil)SetJson(k string, data interface{}) error {
	c := r.redisPool.Get()
	defer c.Close()
	value, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	n, err := c.Do("SET", k, value)
	if n != "OK" {
		Errorf("SET key:%v err: %v",k, err)
		return err
	}
	return nil
}

func (r *RedisUtil)SetNxJson(k string, data interface{}) error {
	c := r.redisPool.Get()
	defer c.Close()
	value, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	n, err := c.Do("SETNX", k, value)
	if n != int64(1) {
		Error("SETNX err: ",err)
		return err
	}
	return nil
}

func (r *RedisUtil)GetJsonByte(key string) ([]byte, error) {
	c := r.redisPool.Get()
	defer c.Close()
	jsonGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		Error("GetJsonByte err: ",err)
		return nil, err
	}
	return jsonGet, nil
}

func(r *RedisUtil)LoadKeys(pattern string)([]string,error){
	c := r.redisPool.Get()
	defer c.Close()
	vals, err := redis.Strings(c.Do("KEYS", pattern))
	if err != nil {
		Error("LoadKeys err: ",err)
		return nil, err
	}
	return vals,nil
}