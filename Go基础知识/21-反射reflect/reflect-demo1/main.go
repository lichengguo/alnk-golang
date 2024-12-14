package main

import (
	"fmt"
	"reflect"
)

// reflect 反射

type Cat struct {
}

/*
在Go语言中，使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息
*/
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("type name:%v --- type kind:%v\n", v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind() //值的类型种类

	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Int()从反射中获取整型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取整型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

//通过反射设置变量的值
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本的话，那么reflect包就会会引发panic。所以只能传递指针
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	var a float32 = 3.14
	reflectType(a)
	/*
		type:float32
		type name:float32 --- type kind:float32
	*/

	var b int64 = 100
	reflectType(b)
	/*
		type:int64
		type name:int64 --- type kind:int64
	*/

	var c = Cat{}
	reflectType(c)
	/*
		type:main.Cat
		type name:Cat --- type kind:struct
		在反射中关于类型还划分为两种：类型（Type）和种类（Kind）。
		因为在Go语言中我们可以使用type关键字构造很多自定义类型，而种类（Kind）就是指底层的类型，
		但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类（Kind）
	*/

	//valueOf
	reflectValue(a) //type is float32, value is 3.140000

	//设置值
	//通过反射设置变量的值
	//需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。而反射中使用专有的Elem()方法来获取指针对应的值
	reflectSetValue1(&b)
	fmt.Println(b) //100 修改没有成功
	reflectSetValue2(&b)
	fmt.Println(b) //200 修改成功

}
