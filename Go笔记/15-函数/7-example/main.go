package main

import (
	"fmt"
	"strings"
)

// 你有5000枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
// 分配规则如下：
// a. 名字中每包含1个'e'或'E'分1枚金币
// b. 名字中每包含1个'i'或'I'分2枚金币
// c. 名字中每包含1个'o'或'O'分3枚金币
// d: 名字中每包含1个'u'或'U'分4枚金币

// 写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
// 程序结构如下，请实现 ‘dispatchCoin’ 函数

var (
	coins        = 5000
	users        = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	distribution = make(map[string]int, len(users))
	rule         = map[string]int{
		"e": 1,
		"i": 2,
		"o": 3,
		"u": 4,
	}
)

func main() {
	left := dispatchCoin()
	fmt.Println("剩下的金币个数：", left)
	// 输出结果循环打印
	for name, value := range distribution {
		fmt.Printf("姓名:%s \t\t金币:%d\n", name, value)
	}
}

func dispatchCoin() (left int) {
	// 1. 依次拿到每个人的名字
	// 2. 拿到一个人名根据分金币的规则去分金币,
	// 2.1 每个人分的金币数应该保存到 distribution 中
	// 2.2 还要记录下剩余的金币数
	// 3. 整个第2步执行完就能得到最终每个人分的金币数和剩余金币数

	// 循环用户名
	for _, name := range users {
		// 对每个用户进行规则循环，并且统计每个用户金币数量
		for ruleKey, ruleValue := range rule {
			count := strings.Count(strings.ToLower(name), ruleKey)
			if count > 0 {
				distribution[name] += count * ruleValue
				coins -= count * ruleValue
			}
		}
	}
	return coins
}
