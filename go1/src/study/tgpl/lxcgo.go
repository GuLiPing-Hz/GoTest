package main

/*
//#cgo CFLAGS: -I/usr/include    //指定cgo编译的引用头目录
//#cgo LDFLAGS: -L/usr/lib -lbz2 //指定cgo编译依赖的动态库路径
//#include <bzlib.h>
//#include <stdlib.h>
//bz_stream* bz2alloc() { return calloc(1, sizeof(bz_stream)); }
//int bz2compress(bz_stream *s, int action, char *in, unsigned *inlen, char *out, unsigned *outlen);
//void bz2free(bz_stream* s) { free(s); }

#include<"stdio.h">
bool cgoInit(const char* str){
	printf("cgo init %s\n",str)
	return true;
}
int cgoAdd(const int a,const int b);
void cgoDeal(char* buf,int* len);
void cgoUninit(){
	printf("cgo uninit\n");
	return true;
}
*/
import "C" //cgo程序依赖gcc。。得。。没得玩了

func main() {
	C.cgoInit("abcd")
}
