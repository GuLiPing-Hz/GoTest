@echo off	

echo ��ǰ�̷���%~d0
echo ��ǰ·����%cd%
echo ��ǰִ�������У�%0
echo ��ǰbat�ļ�·����%~dp0
echo ��ǰbat�ļ���·����%~sdp0

:: ��ǰ�̷�
%~d0
:: ��ǰĿ¼
cd %~dp0
set GOARCH=amd64
set GOOS=linux
set GOTRACEBACK=all 
set GOPATH=%GOPATH%;D:\glp\Github\GoTest\go1\
go build -o ../bin/go1/go1 ./src/study/testServer.go