博客地址:https://www.liwenzhou.com/posts/Go/go_nsq/

github下载地址: https://github.com/nsqio/nsq/releases/tag/v1.2.0

简介
NSQ是目前比较流行的一个分布式的消息队列，本文主要介绍了NSQ及Go语言如何操作NSQ。

NSQ是Go语言编写的一个开源的实时分布式内存消息队列，其性能十分优异。 NSQ的优势：
    NSQ提倡分布式和分散的拓扑，没有单点故障，支持容错和高可用性，并提供可靠的消息交付保证
    NSQ支持横向扩展，没有任何集中式代理。
    NSQ易于配置和部署，并且内置了管理界面
    
NSQ组件
1.nsqd: nsqd是一个守护进程，它接收、排队并向客户端发送消息
启动nsqd，指定-broadcast-address=127.0.0.1来配置广播地址
./nsqd -broadcast-address=127.0.0.1

如果是在搭配nsqlookupd使用的模式下需要还指定nsqlookupd地址:
./nsqd -broadcast-address=127.0.0.1 -lookupd-tcp-address=127.0.0.1:4160
如果是部署了多个nsqlookupd节点的集群，那还可以指定多个-lookupd-tcp-address

2.nsqlookupd
nsqlookupd是维护所有nsqd状态、提供服务发现的守护进程。它能为消费者查找特定topic下的nsqd提供了运行时的自动发现服务。 
它不维持持久状态，也不需要与任何其他nsqlookupd实例协调以满足查询。
因此根据你系统的冗余要求尽可能多地部署nsqlookupd节点。
它们小豪的资源很少，可以与其他服务共存。我们的建议是为每个数据中心运行至少3个集群

3.nsqadmin
一个实时监控集群状态、执行各种管理任务的Web管理平台。 启动nsqadmin，指定nsqlookupd地址
./nsqadmin -lookupd-http-address=127.0.0.1:4161
我们可以使用浏览器打开http://127.0.0.1:4171/访问如下管理界面