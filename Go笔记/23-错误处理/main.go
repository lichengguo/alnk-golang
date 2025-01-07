package main

import (
	// "errors"
	"fmt"
)

// 定义一个函数，用于除法运算
func divide(x, y int) (int, error) {
	if y == 0 {
		// 返回一个错误，表示除数不能为零
		// return 0, errors.New("除数不能为零")
		return 0, &customError{y, "除数不能为零"}
	}
	// 返回除法结果和nil，表示没有错误发生
	return x / y, nil
}

// 自定义 error 类型
type customError struct {
	num     int
	message string
}

// 通过实现 Error() 方法来自定义 error 类型
func (e *customError) Error() string {
	return fmt.Sprintf("%d - %s", e.num, e.message)
}

func main() {
	// 调用 divide 函数进行除法运算
	result, err := divide(10, 2)
	if err != nil {
		// 如果发生错误，则打印错误信息
		fmt.Println("错误:", err)
	} else {
		// 如果没有发生错误，则打印计算结果
		fmt.Println("计算结果:", result)
	}

	// 再次调用 divide 函数进行除法运算
	result, err = divide(10, 0)
	if err != nil {
		// 如果发生错误，则打印错误信息
		fmt.Println("自定义错误: ", err)
	} else {
		// 如果没有发生错误，则打印计算结果
		fmt.Println("计算结果:", result)
	}
}
