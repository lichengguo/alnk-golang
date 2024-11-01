package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
	//TaillogConf `ini:"taillog"`
}

//kafka配置
type KafkaConf struct {
	Address string `ini:"address"`
	//Topic   string `ini:"topic"`
}

//etcd配置
type EtcdConf struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
}

//------- unused ↓ -------------------
//日志文件配置
type TaillogConf struct {
	FileName string `ini:"filename"`
}
