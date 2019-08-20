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
