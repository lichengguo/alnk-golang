package taillog

import (
	"03-logagent-v2.0/etcd"
	"fmt"
	"time"
)

// tailTask 管理者 管理具体的收集日志的任务
type tailLogMgr struct {
	logEntry    []*etcd.LogEntry      //需要收集的日志的配置信息 旧的配置信息 在运行中的配置信息
	tskMap      map[string]*TailTask  //日志收集的任务
	newConfChan chan []*etcd.LogEntry //用来保存最新从etcd拉取的配置信息 新的配置信息
}

var tskMgr *tailLogMgr

// Init 初始化taillog连接
// logEntryConf 需要运行的taillog包的配置信息
func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry:    logEntryConf, // 把当前的日志收集项配置信息保存起来
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry), // 无缓冲区的通道
	}

	//初始化的时候起了多少个tailtask 都要记下来，为了后续判断方便
	for _, logEntry := range logEntryConf {
		//logEntry.Path： 要收集的日志文件的路径
		tailObj := NewTailTask(logEntry.Path, logEntry.Topic)

		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		tskMgr.tskMap[mk] = tailObj
	}

	go tskMgr.run() //执行收集日志的任务
}

// 监听自己的newConfChan，有了新的配置过来之后就做对应的处理
func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			fmt.Println("新的配置来了！", newConf)
			//新增加的配置
			//假如是修改的配置呢？
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.tskMap[mk]
				if ok {
					// 原来就有，不需要操作
					continue
				} else {
					// 新增的
					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.tskMap[mk] = tailObj
					fmt.Println("t.tskMap[mk]: ", t.tskMap[mk])
				}
			}
			//删除的配置
			// 找出原来t.logEntry有，但是newConf中没有的，要删掉
			for _, c1 := range t.logEntry { // 从原来的配置中依次拿出配置项
				isDelete := true
				for _, c2 := range newConf { // 去新的配置中逐一进行比较
					if c2.Path == c1.Path && c2.Topic == c1.Topic {
						isDelete = false
						continue
					}
				}
				if isDelete {
					// 把c1对应的这个tailObj给停掉
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					fmt.Println("mk:", mk)
					//t.tskMap[mk] ==> tailObj
					t.tskMap[mk].cancelFunc()
				}
			}
		default:
			time.Sleep(time.Second)
		}
	}
}

// NewConfChan 向外暴露tskMgr的newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
