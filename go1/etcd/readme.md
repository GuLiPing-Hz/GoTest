## etcd 一个分布式键值对服，用于服务注册、发现
[详细介绍参见](https://github.com/etcd-io/etcd)

##一.启动etcd cluster
###1. [下载etcd release版本](https://github.com/etcd-io/etcd/releases) windows需要配置Path增加访问路径;linux则直接mv 到 /usr/bin
###2. [下载goreman release版本](https://github.com/mattn/goreman/releases) windows需要配置Path增加访问路径;linux则直接mv 到 /usr/bin
###3. 编写Procfile，可以参考本目录的一个文件，编写来3个node组成的cluster,3个node是最少的数量了。还可以是5，7奇数。

下面是对ProcFile的具体说明
```markdown
# Change the path of bin/etcd if etcd is located elsewhere
etcd1: bin/etcd --name infra1 --listen-client-urls http://127.0.0.1:2379 --advertise-client-urls http://127.0.0.1:2379 --listen-peer-urls http://127.0.0.1:12380 --initial-advertise-peer-urls http://127.0.0.1:12380 --initial-cluster-token etcd-cluster-1 --initial-cluster infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380 --initial-cluster-state new --enable-pprof --logger=zap --log-outputs=stderr

etcd1: Procfile语法命令开始  模块名: 命令行   
        冒号后必须跟一个空格。

bin/etcd 启动进程的exe
--name  指定etcd名
--listen-client-urls 指定监听客户端请求地址
--advertise-client-urls 指定客户广播地址
--listen-peer-urls 指定监听node的地址
--initial-advertise-peer-urls 初始化广播地址
--initial-cluster-token 初始化cluster token，同一个token属于同一个cluster。 
--initial-cluster 初始化cluster下的所有监听地址     etcd名=etcd监听node地址,   
--initial-cluster-state 指定启动状态是 新建   
--enable-pprof 启用分析工具  这个生产环境可以关掉。 
--logger=zap 指定日志打印方式为zap。 
--log-outputs=stderr 指定日志输出IO

etcd2: bin/etcd --name infra2 --listen-client-urls http://127.0.0.1:22379 --advertise-client-urls http://127.0.0.1:22379 --listen-peer-urls http://127.0.0.1:22380 --initial-advertise-peer-urls http://127.0.0.1:22380 --initial-cluster-token etcd-cluster-1 --initial-cluster infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380 --initial-cluster-state new --enable-pprof --logger=zap --log-outputs=stderr
etcd3: bin/etcd --name infra3 --listen-client-urls http://127.0.0.1:32379 --advertise-client-urls http://127.0.0.1:32379 --listen-peer-urls http://127.0.0.1:32380 --initial-advertise-peer-urls http://127.0.0.1:32380 --initial-cluster-token etcd-cluster-1 --initial-cluster infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380 --initial-cluster-state new --enable-pprof --logger=zap --log-outputs=stderr

以上3个都是etcd的普通node，他们可以组成一个cluster。
下面语句是在本机开启一个代理。用于跟客户端交互，

proxy: bin/etcd grpc-proxy start --endpoints=127.0.0.1:2379,127.0.0.1:22379,127.0.0.1:32379 --listen-addr=127.0.0.1:23790 --advertise-client-url=127.0.0.1:23790 --enable-pprof
bin/etcd
grpc-proxy start 指定代理模式启动
--endpoints 指定接入的点，地址指向上面3个配置的IP端口。
--listen-addr 指定监听地址
--advertise-client-url 指定广播客户地址
--enable-pprof 启动分析工具


# A learner node can be started using Procfile.learner
```

###4. 启动cmd，输入CHCP 65001 防止乱码
###5. 定位到当前目录 goreman start

##二.增加etcd节点到当前cluster中。
###1.请求添加节点
etcdctl member add infra4 --peer-urls="http://127.0.0.1:42380" --learner=true

###2.goreman -f ./Procfile.learner start

###3.查看当前node列表
etcdctl --endpoints=127.0.0.1:2379,127.0.0.1:22379,127.0.0.1:32379,127.0.0.1:42379 member list

结果见下：
```
261b63fccf6619e6, started, infra4, http://127.0.0.1:42380, http://127.0.0.1:42379, true #这里的true表示新加入的
8211f1d0f64f3269, started, infra1, http://127.0.0.1:12380, http://127.0.0.1:2379, false
91bc3c398fb3c146, started, infra2, http://127.0.0.1:22380, http://127.0.0.1:22379, false
fd422379fda50e48, started, infra3, http://127.0.0.1:32380, http://127.0.0.1:32379, false
```

###4.更新learner投票
etcdctl --endpoints=127.0.0.1:2379,127.0.0.1:22379,127.0.0.1:32379,127.0.0.1:42379 member promote <memberid>
<memberid> 这里的id是指新加入的node节点成员id。可以根据member list来选择。

###5.增加节点之后，启动的配置里面需要增加新节点的配置，如下
```
# Use goreman to run `go get github.com/mattn/goreman`
# Change the path of bin/etcd if etcd is located elsewhere

etcd1: etcd --name infra1 --listen-client-urls http://127.0.0.1:2379 --advertise-client-urls http://127.0.0.1:2379 --listen-peer-urls http://127.0.0.1:12380 --initial-advertise-peer-urls http://127.0.0.1:12380 --initial-cluster-token etcd-cluster-1 --initial-cluster infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380,infra4=http://127.0.0.1:42380 --initial-cluster-state new --enable-pprof --logger=zap --log-outputs=stderr
etcd2: etcd --name infra2 --listen-client-urls http://127.0.0.1:22379 --advertise-client-urls http://127.0.0.1:22379 --listen-peer-urls http://127.0.0.1:22380 --initial-advertise-peer-urls http://127.0.0.1:22380 --initial-cluster-token etcd-cluster-1 --initial-cluster infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380,infra4=http://127.0.0.1:42380 --initial-cluster-state new --enable-pprof --logger=zap --log-outputs=stderr
etcd3: etcd --name infra3 --listen-client-urls http://127.0.0.1:32379 --advertise-client-urls http://127.0.0.1:32379 --listen-peer-urls http://127.0.0.1:32380 --initial-advertise-peer-urls http://127.0.0.1:32380 --initial-cluster-token etcd-cluster-1 --initial-cluster infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380,infra4=http://127.0.0.1:42380 --initial-cluster-state new --enable-pprof --logger=zap --log-outputs=stderr
etcd4: etcd --name infra4 --listen-client-urls http://127.0.0.1:42379 --advertise-client-urls http://127.0.0.1:42379 --listen-peer-urls http://127.0.0.1:42380 --initial-advertise-peer-urls http://127.0.0.1:42380 --initial-cluster-token etcd-cluster-1 --initial-cluster infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380,infra4=http://127.0.0.1:42380 --initial-cluster-state new --enable-pprof --logger=zap --log-outputs=stderr
proxy: etcd grpc-proxy start --endpoints=127.0.0.1:2379,127.0.0.1:22379,127.0.0.1:32379,127.0.0.1:42379 --listen-addr=127.0.0.1:23790 --advertise-client-url=127.0.0.1:23790 --enable-pprof
```


##三.etcd 运维
###1.查看节点状态
`etcdctl --endpoints=127.0.0.1:2379,127.0.0.1:22379,127.0.0.1:32379 endpoint status -w=table`

`etcdctl help endpoint status` 可以知道 -w表示打印方式。以表格打印。默认simple。
###2.查看节点ttl
`etcdctl --endpoints=127.0.0.1:2379,127.0.0.1:22379,127.0.0.1:32379 endpoint health -w=table`

###3.[更多](https://etcd.io/docs/v3.4.0/demo/)

##四.客户端服务发现请求
```
配置ngxin代理，设置body为固定的  c2VydmVy  即base64加密的 server 作为KEY，查找可用服务器列表。

配置三个可访问的节点。本来想用etcd的代理 proxy的，结果没有这个功能。。
通过设置proxy_set_body 我们可以访问到servers 前缀的服务器地址，建议使用 {ip,port,weight}的json字符串让客户端选择服务器。
location /servers {
    proxy_pass    http://127.0.0.1:2379/v3/kv/range;
    proxy_method POST;
    proxy_redirect  off;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_body  "{\"key\":\"c2VydmVy\",\"range_end\":\"c2VydmVz\"}";
}

location /servers1 {
    proxy_pass    http://127.0.0.1:22379/v3/kv/range;
    proxy_method POST;
    proxy_redirect  off;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_body  "{\"key\":\"c2VydmVy\",\"range_end\":\"c2VydmVz\"}";
}

location /servers2 {
    proxy_pass    http://127.0.0.1:32379/v3/kv/range;
    proxy_method POST;
    proxy_redirect  off;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_body  "{\"key\":\"c2VydmVy\",\"range_end\":\"c2VydmVz\"}";
}
```



