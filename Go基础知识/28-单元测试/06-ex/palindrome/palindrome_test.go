package palindrome

import (
	"reflect"
	"testing"
)

//TestPalindrome 单元测试
//func TestPalindrome(t *testing.T) {
//	got := Palindrome("油灯少灯油")
//	want := true
//	if got != want {
//		t.Errorf("want:%v but got:%v\n", want, got)
//	}
//}

//TestPalindrome 子测试
func TestPalindrome(t *testing.T) {
	//定义一个测试用的类型
	type test struct {
		input string //输入
		want  bool   //输出
	}

	//测试用例使用map存储
	tests := map[string]test{
		"t1": {input: "abcdedcba", want: true},
		"t2": {input: "油灯少灯油", want: true},
		"t3": {input: "Madam,I’mAdam", want: false},
	}

	//遍历map，逐一测试
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Palindrome(tc.input)
			want := tc.want
			if !reflect.DeepEqual(got, want) {
				t.Errorf("name:%s wang:%v got:%v\n", name, want, got)
			}
		})
	}
}

/*
lichengguo@lichengguodeMacBook-Pro palindrome % go test -v
=== RUN   TestPalindrome
=== RUN   TestPalindrome/t1
=== RUN   TestPalindrome/t2
=== RUN   TestPalindrome/t3
--- PASS: TestPalindrome (0.00s)
    --- PASS: TestPalindrome/t1 (0.00s)
    --- PASS: TestPalindrome/t2 (0.00s)
    --- PASS: TestPalindrome/t3 (0.00s)
PASS
ok      code.oldboyedu.com/gostudy/day09/99homework/palindrome  0.005s
*/

//BenchmarkPalindrome 基准测试
func BenchmarkPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Palindrome("再或者对于同一个任务究竟使用哪种算法性能最佳?")
	}
}

/*
lichengguo@lichengguodeMacBook-Pro palindrome % go test -bench=Palindrome  -benchmem
goos: darwin
goarch: amd64
pkg: code.oldboyedu.com/gostudy/day09/99homework/palindrome
BenchmarkPalindrome-4            3765415               308 ns/op               0 B/op          0 allocs/op
PASS
ok      code.oldboyedu.com/gostudy/day09/99homework/palindrome  1.490s
*/
