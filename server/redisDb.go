package server

import (
	"errors"
	"redis/adt"
)

// todo 先实现 dict 不然不好存储　redisDb
type RedisDb struct {
	dict *adt.Dict // 数据库空间，保存所有键值对
}

func NewRedisDb() *RedisDb {
	db := &RedisDb{}
	dict := adt.NewDict()
	db.dict = dict
	return db
}

func (r *RedisDb) Set(key, value string) {
	r.dict.Hset(adt.NewRedisObject().Set(&key), adt.NewRedisObject().Set(&value))
}

func (r *RedisDb) Get(key string) string {
	tarObj := r.dict.Hget(adt.NewRedisObject().Set(&key))
	if tarObj == nil {
		return "<nil>"
	}
	return *tarObj.Sdshdr.Get()
}

func (r *RedisDb) HSet(key, filed, value string) (err error) {

	k := adt.NewRedisObject().Set(&key)

	// 先找到 redisDb 中实际存值的　hash
	existsRedisObj := r.dict.Hget(k)

	if existsRedisObj != nil {
		if existsRedisObj.GetType() != adt.REDIS_HASH {
			return errors.New("can not use another type")
		}

		existsRedisObj.Hset(&filed, &value)
	} else {
		redisObj := adt.NewRedisObject()
		redisObj.Hset(&filed, &value)

		// 再把 key => dict 存入 r.dict
		r.dict = adt.NewDict().Hset(k, redisObj)
	}

	return nil
}

func (r *RedisDb) HGet(key, filed string) string {

	k := adt.NewRedisObject().Set(&key)
	f := adt.NewRedisObject().Set(&filed)
	existsRedisObj := r.dict.Hget(k)

	if existsRedisObj != nil {
		if existsRedisObj.GetType() != adt.REDIS_HASH {
			return "can not use this get " + existsRedisObj.GetType()
		}

		targetObj := existsRedisObj.Hget(f)

		if targetObj == nil {
			return "<nil>"
		}

		return *targetObj.Sdshdr.Get()
	}

	return "<nil>"
}

func (r *RedisDb) SetList(key, value *adt.StringObject) {

}
