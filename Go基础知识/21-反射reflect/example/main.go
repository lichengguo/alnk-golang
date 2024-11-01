package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

//ini配置文件解析器（尽量掌握）

//MysqlConfig MySQL配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

//RedisConfig ...
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
	Test     bool   `ini:"test"`
}

//Config ...
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

//解析函数
func loadIni(fileName string, data interface{}) (err error) {
	//0.参数的校验
	//0.1 传进来的data参数必须是指针类型（因为需要在函数中对其赋值）
	t := reflect.TypeOf(data)
	//fmt.Println("type: ", t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer") //新创建一个错误
		return
	}
	//0.2 传进来的data参数必须是结构体类型指针（因为配置文件中各种键值对需要赋值给结构体的字段）
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer") //新建一个错误
		return
	}

	//1.读取文件得到字节类型的数据
	b, err := os.ReadFile(fileName)
	if err != nil {
		return
	}
	//string(b) 将字节类型的文件内容转换成字符串
	lineSlice := strings.Split(string(b), "\n")
	//fmt.Printf("--%#v\n", lineSlice)
	//[]string{"; mysql config", "[mysql]", "address=10.20.30.40", "port=3306", "username=root", "password=rootroot", "", "", "", "",
	// " # redis config", ";  [", "; [  ]", "[redis]", "; heihei", "; = hahaha", "xxx=", "host =11.22.33.44", "port=6379", "password=root123", "database=0", "test=false", ""}

	//2.一行一行的读取数据
	var structName string
	for idx, line := range lineSlice {
		//去掉字符串首尾空格
		//fmt.Println(line)
		line = strings.TrimSpace(line)
		//2.0 如果是空行就跳过
		if len(line) == 0 {
			continue
		}
		//2.1 如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		//2.2 如果开头是 [ ，就表示是节点(section)
		if strings.HasPrefix(line, "[") {
			//2.2.0 如果不是以 [ 开头 和 ] 结尾的话，就报错
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			//2.2.1 把这一行首尾 [ ] 去掉，取到中间的内容，把首尾的空格去掉，拿到内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			//fmt.Println("sectionName: ", sectionName)
			if len(sectionName) == 0 { //表示 [ ] 中没有内容，那么报错
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			//2.2.2 根据字符串sectionName去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					//说明找到了对应的嵌套结构体，把字段名字记下来
					structName = field.Name
					//fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			//2.3 如果不是 [ 开头，那么就是 = 分割的键值对
			//2.3.1 以 =  分割这一行，等号左边是key,等号右边是value
			//strings.Index(line, "=") == -1 没有等号的情况
			//strings.HasPrefix(line, "=") //以等号开头的情况
			//strings.HasSuffix(line, "=") //以等号结尾的情况
			//if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") || strings.HasSuffix(line, "=") {
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=") //找到=的索引，准备切片
			key := strings.TrimSpace(line[:index])
			//fmt.Println("key:", key)
			value := strings.TrimSpace(line[index+1:])
			//fmt.Println("value:", value)
			//2.3.2 根据structName 去 data 里面把对应的嵌套结构体给取出来
			v := reflect.ValueOf(data)
			//fmt.Printf("v: %#v\n", v)
			// v:&main.Config{MysqlConfig:main.MysqlConfig{Address:"", Port:0, Username:"", Password:""}, RedisConfig:main.RedisConfig{Host:"", Port:0, Password:"", Database:0, Test:false}}
			sValue := v.Elem().FieldByName(structName) //拿到嵌套结构体的值信息
			//fmt.Printf("structName: %#v\n", structName) //structName: "MysqlConfig"
			//fmt.Printf("sValue: %#v\n", sValue)         //sValue: main.MysqlConfig{Address:"", Port:0, Username:"", Password:""}
			sType := sValue.Type() //拿到嵌套结构体的类型信息
			//fmt.Printf("sType: %#v\n", sType.Kind())
			//fmt.Println("reflect.Struct: ", reflect.Struct)
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			var fileType reflect.StructField
			//2.3.3 遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				filed := sType.Field(i) //tag信息是存储在类型信息中
				if filed.Tag.Get("ini") == key {
					//找到对应的字段
					fieldName = filed.Name
					fileType = filed
					break
				}
			}
			//2.3.4 如果 key = tag，给这个字段赋值
			//2.3.4.1 根据fieldName去取出这个字段
			if len(fieldName) == 0 {
				//在结构体中找不到对应的字段
				continue
			}
			fileObj := sValue.FieldByName(fieldName)

			//2.3.4.2 对其赋值
			//fmt.Println("===", fieldName, fileType.Type.Kind())
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetFloat(valueFloat)
			}
		}
	}
	return
}

func main() {
	var cfg Config
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}

	fmt.Printf("%#v\n", cfg)
}
