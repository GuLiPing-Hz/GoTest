package net

import "encoding/binary"

type DataDecode struct {
}

func (this *DataDecode) GetPackageHeadLen() int {
	//单一包的头部长度 默认为2字节
	return 2
}
func (this *DataDecode) GetPackageLen(bytes []byte) int {
	//大端的形式获取包长
	return int(binary.BigEndian.Uint16(bytes))
}
