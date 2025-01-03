package main

import "testing"

// TestAdd 单元测试 用于测试Add函数的正确性
func TestAdd(t *testing.T) {
	got := Add(1, 2)
	want := 3
	if got != want {
		t.Errorf("want:%v but got:%v\n", want, got)
	}
}

// BenchmarkAdd 基准测试 用于测试Add函数的性能
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := Add(1, 2)
		if result != 3 {
			b.Errorf("want:%v but got:%v\n", 3, result)
		}
	}
}

/*
单元测试
alnk@Alnk-MacBook-Air test-01-add % go test -v
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
PASS
ok      test-01-add     0.242s

基准测试
alnk@Alnk-MacBook-Air test-01-add % go test -bench=.
goos: darwin
goarch: arm64
pkg: test-01-add
BenchmarkAdd-8          1000000000               0.2922 ns/op
PASS
ok      test-01-add     0.588s
*/
