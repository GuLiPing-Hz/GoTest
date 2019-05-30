@echo off	

echo 当前盘符：%~d0
echo 当前路径：%cd%
echo 当前执行命令行：%0
echo 当前bat文件路径：%~dp0
echo 当前bat文件短路径：%~sdp0

:: 当前盘符
%~d0
:: 当前目录
cd %~dp0
set GOARCH=amd64
set GOOS=linux
set GOTRACEBACK=all 
set GOPATH=%GOPATH%;D:\glp\Github\GoTest\go1\
go build -o ../bin/go1/go1 ./src/study/testServer.go