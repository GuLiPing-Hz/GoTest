package main

/**
按道理我们应该去这个网站下代码,
https://github.com/grpc/grpc-go
也就是 go get github.com/grpc/grpc-go
但是好像临近国庆节国外网站貌似都不容易访问了。

我们还是从码云的镜像直接下载吧
go get gitee.com/mirrors/grpc-go
然后去 $GoPATH 修改目录名把  gitee.com/mirrors/grpc-go 改成 google.golang.org/grpc

然后对着proto文件再次编译生成gRPC文件
protoc --go_out=plugins=grpc:. PbLobbyData.proto

*/
func main() {

}
