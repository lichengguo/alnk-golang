package main

// 数据模型

// 图书信息
type book struct {
	ID    int64  `db:"id"` // 这里字段名要大写，因为db.Get db.Select用了反射
	Title string `db:"title"`
	Price int64  `db:"price"`
}
