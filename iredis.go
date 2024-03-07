package databasewapper

import (
	"time"

	"github.com/go-redis/redis/v8"
)

// IRedisConnection 是 Redis 连接的抽象接口
type IRedisConnection interface {
	// Connect 连接到 Redis 实例
	Connect(mastername, password string, addrs []string, db int) error

	// Close 关闭 Redis 连接
	Close() error

	// Set 设置指定 key 的值为指定字符串，并设置过期时间
	Set(key string, value string, duration time.Duration) (err error)

	// Get 获取指定 key 的值
	Get(key string) (result string, err error)

	// Keys 获取匹配指定模式的所有 key
	Keys(key string) (result []string, err error)

	// Scan 扫描指定 key
	Scan(key string)

	// Incr 将 key 中储存的数字值增一
	Incr(key string) (err error)

	// Del 删除指定 key
	Del(key string) (err error)

	// LockWithTimeout 尝试获取锁
	LockWithTimeout(key string, timeout int32) bool

	// UnLock 释放锁
	UnLock(key string) (nums int64)

	// HashSet 向 key 的 hash 中添加元素 field 的值
	HashSet(key, field string, data string)

	// HashDel 删除 key 的 hash 中指定 field
	HashDel(key, field string)

	// BatchHashSet 批量向 key 的 hash添加对应元素 field 的值
	BatchHashSet(key string, fields map[string]interface{}) bool

	// HashGet 通过 key 获取 hash 的元素值
	HashGet(key, field string) string

	// HashGetAll 获取 key 的 hash 中所有元素值
	HashGetAll(key string) map[string]string

	// BatchHashGet 批量获取 key 的 hash 中对应多元素值
	BatchHashGet(key string, fields ...string) map[string]interface{}

	// LTrim 对列表进行修剪
	LTrim(key string, start, stop int64)

	// LRange 获取列表指定范围的元素
	LRange(key string, start, stop int64) (interface{}, error)

	// Exists 检查 key 是否存在
	Exists(key string) (int64, error)

	// BatchSAdd 批量向 key 的集合添加元素
	BatchSAdd(key string, values ...interface{}) bool

	// SCard 获取集合的基数
	SCard(key string) int64

	// SAdd 向集合添加一个或多个成员
	SAdd(key string, values string) bool

	// SRandMemberN 从集合中随机获取指定数量的元素
	SRandMemberN(key string, count int64) (interface{}, error)

	// SPopN 从集合中移除并返回指定数量的随机元素
	SPopN(key string, count int64) (interface{}, error)

	// SIsMember 判断 member 元素是否是集合 key 的成员
	SIsMember(key string, field string) (error, bool)

	// SMembers 返回集合中的所有成员
	SMembers(key string) (vals []string, err error)

	// SRem 移除集合中一个或多个成员
	SRem(key string, value string) (err error)

	// ZAdd 向有序集合添加一个或多个成员，或更新已存在成员的分数
	ZAdd(key string, score float64, member string) (err error)

	// BatchZAdd 批量向有序集合添加成员及分数
	BatchZAdd(key string, member map[string]float64) (err error)

	// BatchZAddMembers 批量向有序集合添加成员及分数
	BatchZAddMembers(key string, members *[]redis.Z) (err error)

	// ZIncrBy 有序集合中对指定成员的分数加上增量 increment
	ZIncrBy(key string, increment float64, member string) (n float64, err error)

	// ZCard 获取有序集合的基数
	ZCard(key string) (n int64, err error)

	// ZRem 移除有序集合中的一个或多个成员
	ZRem(key string, member string) (vals int64, err error)

	// ZRangeByScore 返回有序集合中指定分数范围内的成员
	ZRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error)

	// ZRevRangeByScore 返回有序集合中指定分数范围内的成员，按分数从高到低排序
	ZRevRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error)

	// ZRangeByLex 返回有序集合中指定字典区间内的成员
	ZRangeByLex(key string, opt *redis.ZRangeBy) ([]string, error)

	// ZRevRangeByLex 返回有序集合中指定字典区间内的成员，按成员字典顺序递减返回
	ZRevRangeByLex(key string, opt *redis.ZRangeBy) ([]string, error)

	// ZRank 返回有序集合中指定成员的排名
	ZRank(key, member string) (int64, error)

	// ZRevRank 返回有序集合中指定成员的排名，按分数从高到低排序
	ZRevRank(key, member string) (int64, error)

	// ZScore 返回有序集合中指定成员的分数
	ZScore(key, member string) (float64, error)

	// ZRange 返回有序集合中指定范围内的成员
	ZRange(key string, start, stop int64) ([]string, error)

	// ZRevRange 返回有序集合中指定范围内的成员，按分数从高到低排序
	ZRevRange(key string, start, stop int64) ([]string, error)
}
