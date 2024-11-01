package conf

type KafkaConf struct {
	Address     string `ini:"address"`       //kafka服务端地址
	ChanMaxSize int    `ini:"chan_max_size"` //通道初始化大小
}

type EtcdConf struct {
	Address string `ini:"address"`         //etcd服务端地址
	Key     string `ini:"collect_log_key"` //etcd中的key，分类用
	Timeout int    `ini:"timeout"`         //超时时间
}

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
}
