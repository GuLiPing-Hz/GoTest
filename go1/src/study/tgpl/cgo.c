#include<"stdio.h">
#include<"string.h">

int cgoAdd(const int a,const int b){
    return a+b;
}

char hi[] = "Hello World";
void cgoDeal(char* buf,int* len){
    printf("sizeof=%d,strlen=%d\n",sizeof(hi),strlen(hi))
    if(!buf){
        *len = sizeof(hi);
        return;
    }

    strcpy(buf,hi)
}