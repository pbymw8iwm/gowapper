package gowapper

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/astaxie/beego"
	"github.com/go-redis/redis/v8"
)

type RedisSentinelClient struct {
	RedisClient *redis.Client
}

func (r *RedisSentinelClient) Connect(mastername, password string, addrs []string, db int) error {
	sc := redis.FailoverOptions{
		MasterName:    mastername,
		SentinelAddrs: addrs,
		Password:      password,
		DB:            db,

		SlaveOnly: false,

		//每一个redis.Client的连接池容量及闲置连接数量，而不是clusterClient总体的连接池大小。
		//实际上没有总的连接池而是由各个redis.Client自行去实现和维护各自的连接池。
		PoolSize:     5 * runtime.NumCPU(), // 连接池最大socket连接数，默认为5倍CPU数， 5 * runtime.NumCPU
		MinIdleConns: 10,                   //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量。

		//命令执行失败时的重试策略
		MaxRetries:      1,                      // 命令执行失败时，最多重试多少次，默认为0即不重试
		MinRetryBackoff: 8 * time.Millisecond,   //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
		MaxRetryBackoff: 512 * time.Millisecond, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔

		//超时
		DialTimeout:  5 * time.Second, //连接建立超时时间，默认5秒。
		ReadTimeout:  3 * time.Second, //读超时，默认3秒， -1表示取消读超时
		WriteTimeout: 3 * time.Second, //写超时，默认等于读超时，-1表示取消读超时
		PoolTimeout:  4 * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。

		IdleTimeout: 5 * time.Minute, //闲置超时，默认5分钟，-1表示取消闲置超时检查
		MaxConnAge:  0 * time.Second, //连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接
	}

	r.RedisClient = redis.NewFailoverClient(&sc)
	return r.GetServerPing()
}

func (client *RedisSentinelClient) Set(key string, value string, duration time.Duration) (err error) {
	err = client.RedisClient.Set(context.Background(), key, value, duration).Err()
	return
}
func (client *RedisSentinelClient) Get(key string) (result string, err error) {
	result, err = client.RedisClient.Get(context.Background(), key).Result()
	return
}

func (client *RedisSentinelClient) Keys(key string) (result []string, err error) {
	res := client.RedisClient.Keys(context.Background(), key)
	err = res.Err()
	result = res.Val()

	return
}

func (client *RedisSentinelClient) Scan(key string) {
	iter := client.RedisClient.Scan(context.Background(), 0, key, 10).Iterator()
	for iter.Next(context.Background()) {
		beego.Error("Redis Scan Error:", iter.Val())
	}

	return
}

func (client *RedisSentinelClient) Incr(key string) (err error) {
	err = client.RedisClient.Incr(context.Background(), key).Err()
	return
}
func (client *RedisSentinelClient) Del(key string) (err error) {
	err = client.RedisClient.Del(context.Background(), key).Err()
	return
}

// 加锁
func (client *RedisSentinelClient) LockWithTimeout(key string, timeout int32) bool {
	// ex:设置默认过期时间10秒，防止死锁
	ex := time.Duration(timeout) * time.Second
	ret, _ := client.RedisClient.SetNX(context.Background(), key, `{"lock":1}`, ex).Result()
	return ret
}

// 解锁
func (client *RedisSentinelClient) UnLock(key string) (nums int64) {
	var err error
	nums, err = client.RedisClient.Del(context.Background(), key).Result()
	if err != nil {
		beego.Error("UnLock:%+v", err)
		return 0
	}
	return nums
}

// 向key的hash中添加元素field的值
func (client *RedisSentinelClient) HashSet(key, field string, data string) {
	err := client.RedisClient.HSet(context.Background(), key, field, data).Err()
	if err != nil {
		beego.Error("Redis HSet Error:", err)
	}
}

//hash删除
func (client *RedisSentinelClient) HashDel(key, field string) {
	hDel := client.RedisClient.HDel(context.Background(), key, field)
	if hDel.Err() != nil {
		beego.Error("Redis HSet Error:", hDel.Err())
	}
}

// 批量向key的hash添加对应元素field的值
func (client *RedisSentinelClient) BatchHashSet(key string, fields map[string]interface{}) bool {
	val, err := client.RedisClient.HMSet(context.Background(), key, fields).Result()
	if err != nil {
		beego.Error("Redis HMSet Error:", err)
	}
	return val
}

// 通过key获取hash的元素值
func (client *RedisSentinelClient) HashGet(key, field string) string {
	result := ""
	val, err := client.RedisClient.HGet(context.Background(), key, field).Result()
	if err == redis.Nil {
		//beego.Informational("Key Doesn't Exists:", field)
		return result
	} else if err != nil {
		beego.Informational("Redis HGet Error:", err)
		return result
	}
	return val
}

//获取
func (client *RedisSentinelClient) HashGetAll(key string) map[string]string {
	val, err := client.RedisClient.HGetAll(context.Background(), key).Result()
	if err == redis.Nil {
		beego.Informational("Key Doesn't Exists:", key)
		return val
	} else if err != nil {
		beego.Informational("Redis HGet Error:", err)
		return val
	}
	return val
}

// 批量获取key的hash中对应多元素值
func (client *RedisSentinelClient) BatchHashGet(key string, fields ...string) map[string]interface{} {
	resMap := make(map[string]interface{})
	for _, field := range fields {
		var result interface{}
		val, err := client.RedisClient.HGet(context.Background(), key, fmt.Sprintf("%s", field)).Result()
		if err == redis.Nil {
			beego.Informational("Key Doesn't Exists:", field)
			resMap[field] = result
		} else if err != nil {
			beego.Informational("Redis HMGet Error:", err)
			resMap[field] = result
		}
		if val != "" {
			resMap[field] = val
		} else {
			resMap[field] = result
		}
	}
	return resMap
}

//删除list
func (client *RedisSentinelClient) LTrim(key string, start, stop int64) {
	er, err := client.RedisClient.LTrim(context.Background(), key, start, stop).Result()
	if err != nil {
		beego.Error("Redis HSet Error:", err, er)
	}
}
func (client *RedisSentinelClient) RPush(key string, values ...interface{}) {
	err := client.RedisClient.RPush(context.Background(), key, values).Err()
	if err != nil {
		beego.Error("Redis HSet Error:", err)
	}
}
func (client *RedisSentinelClient) LRange(key string, start, stop int64) (interface{}, error) {
	strings, err := client.RedisClient.LRange(context.Background(), key, start, stop).Result()
	if err != nil {
		beego.Error("Redis HSet Error:", err)
	}
	return strings, err
}
func (client *RedisSentinelClient) LLen(key string) (int64, error) {
	cnt, err := client.RedisClient.LLen(context.Background(), key).Result()
	if err != nil {
		beego.Error("Redis HSet Error:", err)
	}
	return cnt, err
}

func (client *RedisSentinelClient) Exists(key string) (int64, error) {
	cnt, err := client.RedisClient.Exists(context.Background(), key).Result()
	if err != nil {
		beego.Error("Redis HSet Error:", err)
	}
	return cnt, err
}

// 批量向key的hash添加对应元素field的值 values ...interface{}
func (client *RedisSentinelClient) BatchSAdd(key string, values ...interface{}) bool {
	val, err := client.RedisClient.SAdd(context.Background(), key, values).Result()
	if err != nil {
		beego.Error("Redis BatchSAdd Error:", err, key, values)
	}
	return val > 0
}
func (client *RedisSentinelClient) SCard(key string) int64 {
	sCard := client.RedisClient.SCard(context.Background(), key)
	if sCard.Err() != nil {
		return 0
	}
	return sCard.Val()
}

func (client *RedisSentinelClient) SAdd(key string, values string) bool {
	val, err := client.RedisClient.SAdd(context.Background(), key, values).Result()
	if err != nil {
		beego.Error("Redis SAdd Error:", err)
	}
	return val > 0
}
func (client *RedisSentinelClient) SRandMemberN(key string, count int64) (interface{}, error) {
	members, err := client.RedisClient.SRandMemberN(context.Background(), key, count).Result()
	if err != nil {
		beego.Error("Redis SAdd Error:", err)
	}
	return members, err
}
func (client *RedisSentinelClient) SPopN(key string, count int64) (interface{}, error) {
	members, err := client.RedisClient.SPopN(context.Background(), key, count).Result()
	if err != nil {
		beego.Error("Redis SAdd Error:", err)
	}
	return members, err
}
func (client *RedisSentinelClient) SIsMember(key string, field string) (error, bool) {
	exists, err := client.RedisClient.SIsMember(context.Background(), key, field).Result()
	if err != nil {
		beego.Error("Redis SAdd Error:", err)
		return err, false
	}
	return err, exists
}

func (client *RedisSentinelClient) SMembers(key string) (vals []string, err error) {
	vals, err = client.RedisClient.SMembers(context.Background(), key).Result()
	if err != nil {
		beego.Error("Redis SMembers Error:", err)
	}
	return
}

func (client *RedisSentinelClient) SRem(key string, value string) (err error) {
	_, err = client.RedisClient.SRem(context.Background(), key, value).Result()
	if err != nil {
		beego.Error("Redis SRem Error:", err)
	}
	return
}

func (client *RedisSentinelClient) ZAdd(key string, score float64, member string) (err error) {
	err = client.RedisClient.ZAdd(context.Background(), key, &redis.Z{
		Score:  score,
		Member: member,
	}).Err()
	return
}
func (client *RedisSentinelClient) BatchZAdd(key string, member map[string]float64) (err error) {
	add_list := make([]*redis.Z, 0)
	for key, value := range member {
		add_list = append(add_list, &redis.Z{
			Score:  value,
			Member: key,
		})
	}
	err = client.RedisClient.ZAdd(context.Background(), key, add_list...).Err()
	return
}
func (client *RedisSentinelClient) BatchZAddMembers(key string, members *[]redis.Z) (err error) {
	add_list := make([]*redis.Z, 0)
	for i, _ := range *members {
		add_list = append(add_list, &(*members)[i])
	}
	err = client.RedisClient.ZAdd(context.Background(), key, add_list...).Err()
	return
}

func (client *RedisSentinelClient) ZIncrBy(key string, increment float64, member string) (n float64, err error) {
	n, err = client.RedisClient.ZIncrBy(context.Background(), key, increment, member).Result()
	return n, err
}

func (client *RedisSentinelClient) ZCard(key string) (n int64, err error) {
	n, err = client.RedisClient.ZCard(context.Background(), key).Result()
	return n, err
}
func (client *RedisSentinelClient) ZRem(key string, member string) (vals int64, err error) {
	vals, err = client.RedisClient.ZRem(context.Background(), key, member).Result()
	return vals, err
}
func (client *RedisSentinelClient) ZScore(key string, member string) (vals float64, err error) {
	vals, err = client.RedisClient.ZScore(context.Background(), key, member).Result()
	return vals, err
}

func (client *RedisSentinelClient) ZRevRank(key, value string) (val int64, err error) {
	val, err = client.RedisClient.ZRevRank(context.Background(), key, value).Result()
	if err != nil {
		beego.Informational("Redis ZRevRank Error:", err)
	}
	return
}

//[]redis.Z{{Score: 1, Member: "one"}}
func (client *RedisSentinelClient) ZRangeWithScores(key string, start, stop int64) (vals []redis.Z, err error) {
	vals, err = client.RedisClient.ZRangeWithScores(context.Background(), key, start, stop).Result()
	return vals, err
}
func (client *RedisSentinelClient) ZRevRangeWithScores(key string, start, stop int64) (vals []redis.Z, err error) {
	vals, err = client.RedisClient.ZRevRangeWithScores(context.Background(), key, start, stop).Result()
	return vals, err
}
func (client *RedisSentinelClient) ZRemRangeByRank(key string, start, stop int64) (vals int64, err error) {
	vals, err = client.RedisClient.ZRemRangeByRank(context.Background(), key, start, stop).Result()
	return vals, err
}
func (client *RedisSentinelClient) Expire(key string, duration time.Duration) (expire bool, err error) {
	expire, err = client.RedisClient.Expire(context.Background(), key, duration).Result()
	return expire, err
}

func (client *RedisSentinelClient) ForEachMasterScan(key string, key_list *[]string) (err error) {

	beego.Informational("Redis ForEachMasterScan not implement")
	return
}
func (client *RedisSentinelClient) TTL(key string) (err error, ttl time.Duration) {
	ttl, err = client.RedisClient.TTL(context.Background(), key).Result()
	if err != nil {
		beego.Error("Redis TTL Error:", err)
	}
	return
}

func (client *RedisSentinelClient) Close() {
	client.RedisClient.Close()
}

func (client *RedisSentinelClient) GetServerPing() error {
	if client.RedisClient == nil {
		return fmt.Errorf("redis Client connect is nil")
	}
	err := client.RedisClient.Ping(context.Background()).Err()
	if err != nil {
		return err
	}
	return nil
}
