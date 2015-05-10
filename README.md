# HostsSitter
HostsSitter - host更新管理工具
```
	neeke@php.net
```

### 使用

```shell
	go build
	
	查看帮助
	./HostsSitter -h
	
	查看版本
	./HostsSitter -v
	
	更新
	./HostsSitter update
	
	多源更新
	./HostsSitter update google 360kb
	
	
	启动Debug模式
	vim conf/app.conf
	
	{
	    "RunMode" : "dev"
	}
	
```

### 
 - 完成Google Hosts更新

### TODO
 - Hosts 备份
 - 手动添加\删除\更新Hosts
 - 多源的Hosts更新策略
