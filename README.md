# FlashSale
使用Go语言实现的一个简易版秒杀系统，为了学习经典的秒杀架构及技术栈


###### 技术栈
- Go
- MySQL
- Redis
- Gin
- etcd
- xorm
- protobuf


### 特性

1、协程池、连接池
2、基于redis和内存缓存的多级缓存
3、基于etcd的分布式服务发现及配置服务（etcd也可以代替redis实现分布式锁，且比redis更稳定）
