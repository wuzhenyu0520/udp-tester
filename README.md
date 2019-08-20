# udp-tester
udp测试工具

包含内容  
服务端——server  
客户端——client  
配置文件——conf.ini  

使用ini格式配置文件  
通过修改配置文件，控制服务端监听地址及端口范围，控制客户端测试地址及端口范围  

[udp-tester]  
listenaddr=127.0.0.1  
minport=30000  
maxport=30010  
  
listenaddr：服务端监听地址，或客户端尝试连接地址  
minport：udp端口范围起始端口  
maxport：udp端口范围结束端口  

## 使用方法
1、修改配置文件进行server启动  
    修改listenaddr以配置server监听udp端口地址  
    修改minport以配置server监听udp范围端口起始端口  
    修改maxport以配置server监听udp范围端口起始端口  
    e.g：  
    listenaddr=127.0.0.1  
    minport=30000  
    maxport=30010  
    则server启动后会监听本地127.0.0.1地址的30000～30010udp端口  
   
2、修改配置文件进行client启动测试  
    修改listenaddr以配置client连接测试的udp端口地址  
    修改minport以配置client连接测试的udp范围端口起始端口  
    修改maxport以配置client连接测试的udp范围端口起始端口  

测试结果  
服务端监听  
Listen udp port 30000  
Listen udp port 30001  
Listen udp port 30002  
Listen udp port 30003  
Listen udp port 30004  
Listen udp port 30005  
Listen udp port 30006  
Listen udp port 30007  
Listen udp port 30008  
Listen udp port 30009  
Listen udp port 30010  

  
client连接测试  
[INFO]  2019-08-21 02:37:22.947435 +0800 CST m=+8.066028618  30000 Connect Sucess!!!  
[INFO]  2019-08-21 02:37:22.948025 +0800 CST m=+8.066618965  30001 Connect Sucess!!!  
[INFO]  2019-08-21 02:37:22.948312 +0800 CST m=+8.066905775  30002 Connect Sucess!!!  
[INFO]  2019-08-21 02:37:22.948526 +0800 CST m=+8.067120265  30003 Connect Sucess!!!  
[INFO]  2019-08-21 02:37:22.948734 +0800 CST m=+8.067327928  30004 Connect Sucess!!!  
[INFO]  2019-08-21 02:37:22.948945 +0800 CST m=+8.067539493  30005 Connect Sucess!!!  
[INFO]  2019-08-21 02:37:22.949703 +0800 CST m=+8.068296981  30006 Connect Sucess!!!  
[INFO]  2019-08-21 02:37:22.949923 +0800 CST m=+8.068517203  30007 Connect Sucess!!!  
[INFO]  2019-08-21 02:37:22.950129 +0800 CST m=+8.068723140  30008 Connect Sucess!!!  
[INFO]  2019-08-21 02:37:22.950443 +0800 CST m=+8.069036850  30009 Connect Sucess!!!  
[INFO]  2019-08-21 02:37:22.950678 +0800 CST m=+8.069271724  30010 Connect Sucess!!!  