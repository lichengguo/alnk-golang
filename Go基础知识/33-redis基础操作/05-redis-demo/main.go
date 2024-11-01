/*
安装: go get  github.com/go-redis/redis
*/

package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

//在全局中声明一个db变量，好让所有的函数都能调用
var rdb *redis.Client

//initClient 普通连接
func initClient() (err error) {
	//1.配置信息
	rdb = redis.NewClient(&redis.Options{
		Addr: "192.168.3.121:6379",
		DB:   0,
	})
	//2.尝试连接
	str, err := rdb.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println("str:", str) //str: PONG
	return
}

/*
连接Redis哨兵模式
var rdb *redis.Client

func initClient()(err error){
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
*/

/*
连接Redis集群
var rdb *redis.Client

func initClient()(err error){
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

*/

//set/get示例 redisExample
func redisExample() {
	//1.插入值
	err := rdb.Set("score", 100, 0).Err() //set(key,values,失效时间 0表示永不失效)
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}
	//2.取值
	//可以取到值的情况
	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score value failed, err:%v\n", err)
		return
	}
	fmt.Println("keys:score , val:", val)

	//键值对不存在或者值为nil的情况
	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("keys: name val: ", val2)
	}
}

//zset示例
func redisExample2() {
	zsetKey := "language_rank"
	languages := []redis.Z{
		redis.Z{Score: 90.0, Member: "Golang"}, //Socre:相当于值 Member:相当于键
		redis.Z{Score: 98.0, Member: "Java"},
		redis.Z{Score: 95.0, Member: "Python"},
		redis.Z{Score: 97.0, Member: "JavaScript"},
		redis.Z{Score: 99.0, Member: "C/C++"},
	}

	//zadd
	num, err := rdb.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d success!\n", num)

	//把Golang的分数加10
	newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	//取分数最高的3个
	ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		return
	}
	for _, z := range ret {
		fmt.Println(z)
	}

	//取95-100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

//inserRedis1W 测试redis
func inserRedis1W() {
	startTime := time.Now().Unix()
	for i := 0; i < 10000; i++ {
		keyName := fmt.Sprintf("name%00000d", i)
		err := rdb.Set(keyName, "我是很占内存的一串字符串!!!!!!", time.Second*60).Err() //设置了key过期的时间
		if err != nil {
			fmt.Printf("insert failed, err:%v\n", err)
			return
		}
	}
	fmt.Printf("耗时:%v\n", time.Now().Unix()-startTime) //耗时:41
}

/*
pipLine 测试
Pipeline 主要是一种网络优化。它本质上意味着客户端缓冲一堆命令并一次性将它们发送到服务器。
这些命令不能保证在事务中执行。这样做的好处是节省了每个命令的网络往返时间（RTT）
*/
func pipLine() {
	startTime := time.Now().Unix()

	pipe := rdb.Pipeline()

	for i := 0; i < 10000; i++ {
		keyName := fmt.Sprintf("name%00000d", i)
		pipe.Set(keyName, "我是很占内存的一串字符串!!!!!!", time.Second*60)
	}

	pipe.Exec()

	fmt.Printf("耗时:%v\n", time.Now().Unix()-startTime) //耗时:0
}

func main() {
	err := initClient()
	if err != nil {
		fmt.Printf("connect redis server failed, err:%v\n", err)
		return
	}
	fmt.Println("redis connect sucess!")

	//set/get示例
	redisExample()

	//zadd
	//redisExample2()

	//测试插入1万键值对
	//inserRedis1W()

	//pipLine 测试
	//pipLine()

}

/*
192.168.3.121:6379> zrange language_rank 0 100
1) "Golang"
2) "Python"
3) "JavaScript"
4) "Java"
5) "C/C++"

192.168.3.121:6379> zrange language_rank 0 100 withscores
 1) "Golang"
 2) "90"
 3) "Python"
 4) "95"
 5) "JavaScript"
 6) "97"
 7) "Java"
 8) "98"
 9) "C/C++"
10) "99"
*/
