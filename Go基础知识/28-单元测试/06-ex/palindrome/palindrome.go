package palindrome

/*
编写一个回文检测函数，并为其编写单元测试和基准测试，根据测试的结果逐步对其进行优化。
回文：一个字符串正序和逆序一样，如“Madam,I’mAdam”、“油灯少灯油”等
*/

func Palindrome(str string) bool {
	//1.将字符串转为[]rune类型的切片
	runeSlice := []rune(str)
	//2.拿切片的第一个元素和最后一个元素对比
	//再拿第二个元素和倒数第二个元素对比
	//对比的次数为切片的总长度 /2 取商
	for i := 0; i < len(runeSlice)/2; i++ {
		if runeSlice[i] != runeSlice[len(runeSlice)-1-i] {
			return false
		}
	}
	return true
}
