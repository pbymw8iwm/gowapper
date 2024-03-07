package gowapper

import (
	"time"

	"github.com/go-redis/redis/v8"
)

// IRedisConnection 是 Redis 连接的抽象接口
type IRedisClient interface {
	Connect(mastername, password string, addrs []string, db int) error

	Set(key string, value string, duration time.Duration) (err error)
	Get(key string) (result string, err error)
	Keys(key string) (result []string, err error)
	Scan(key string)

	Incr(key string) (err error)
	Del(key string) (err error)

	// 加锁
	LockWithTimeout(key string, timeout int32) bool

	// 解锁
	UnLock(key string) (nums int64)

	// 向key的hash中添加元素field的值
	HashSet(key, field string, data string)

	//hash删除
	HashDel(key, field string)

	// 批量向key的hash添加对应元素field的值
	BatchHashSet(key string, fields map[string]interface{}) bool

	// 通过key获取hash的元素值
	HashGet(key, field string) string

	//获取
	HashGetAll(key string) map[string]string

	// 批量获取key的hash中对应多元素值
	BatchHashGet(key string, fields ...string) map[string]interface{}

	//删除list
	LTrim(key string, start, stop int64)
	RPush(key string, values ...interface{})
	LRange(key string, start, stop int64) (interface{}, error)
	LLen(key string) (int64, error)
	Exists(key string) (int64, error)

	// 批量向key的hash添加对应元素field的值 values ...interface{}
	BatchSAdd(key string, values ...interface{}) bool
	SCard(key string) int64

	SAdd(key string, values string) bool
	SRandMemberN(key string, count int64) (interface{}, error)
	SPopN(key string, count int64) (interface{}, error)
	SIsMember(key string, field string) (error, bool)

	SMembers(key string) (vals []string, err error)

	SRem(key string, value string) (err error)

	ZAdd(key string, score float64, member string) (err error)
	BatchZAdd(key string, member map[string]float64) (err error)
	BatchZAddMembers(key string, members *[]redis.Z) (err error)
	ZIncrBy(key string, increment float64, member string) (n float64, err error)
	ZCard(key string) (n int64, err error)
	ZRem(key string, member string) (vals int64, err error)
	ZScore(key string, member string) (vals float64, err error)
	ZRevRank(key, value string) (val int64, err error)

	//[]redis.Z{{Score: 1, Member: "one"}}
	ZRangeWithScores(key string, start, stop int64) (vals []redis.Z, err error)
	ZRevRangeWithScores(key string, start, stop int64) (vals []redis.Z, err error)
	ZRemRangeByRank(key string, start, stop int64) (vals int64, err error)
	Expire(key string, duration time.Duration) (expire bool, err error)

	ForEachMasterScan(key string, key_list *[]string) (err error)
	TTL(key string) (err error, ttl time.Duration)
	Close() error
	GetServerPing() error

	Subscribe(channels ...string) *redis.PubSub
}
