package net2

import (
	"bytes"
	"encoding/binary"
	"io"
	"mybase"
	"reflect"
)

const (
	CHAR   int8 = 1
	SHORT  int8 = 2
	INT    int8 = 3
	INT64  int8 = 4
	FLOAT  int8 = 5
	OTHER  int8 = 6
	ARRAY  int8 = 16
	STRING int8 = 17
)

type Head struct {
	Cmd uint16
	Seq uint16
	Ret int16
}

type Packager struct {
	bytes.Buffer
}

func (imp *Packager) ReadChar() (int8, error) {
	var val int8 = 0
	err := binary.Read(imp, binary.BigEndian, &val)
	return val, err
}
func (imp *Packager) ReadShort() (int16, error) {
	var val int16 = 0
	err := binary.Read(imp, binary.BigEndian, &val)
	return val, err
}
func (imp *Packager) ReadUShort() (uint16, error) {
	var val uint16 = 0
	err := binary.Read(imp, binary.BigEndian, &val)
	return val, err
}
func (imp *Packager) ReadInt() (int32, error) {
	var val int32 = 0
	err := binary.Read(imp, binary.BigEndian, &val)
	return val, err
}
func (imp *Packager) ReadInt64() (int64, error) {
	var val int64 = 0
	err := binary.Read(imp, binary.BigEndian, &val)
	return val, err
}
func (imp *Packager) ReadFloat() (float32, error) {
	var val float32 = 0
	err := binary.Read(imp, binary.BigEndian, &val)
	return val, err
}
func (imp *Packager) ReadSimple(flag int8, data []interface{}) []interface{} {
	switch flag {
	case CHAR:
		val, _ := imp.ReadChar()
		data = append(data, val)
	case SHORT:
		val, _ := imp.ReadShort()
		data = append(data, val)
	case INT:
		val, _ := imp.ReadInt()
		data = append(data, val)
	case INT64:
		val, _ := imp.ReadInt64()
		data = append(data, val)
	case FLOAT:
		val, _ := imp.ReadFloat()
		data = append(data, val)
	case STRING:
		lenArra, _ := imp.ReadUShort()
		buf := make([]byte, lenArra)
		_, _ = imp.Read(buf)
		data = append(data, string(buf))
	case ARRAY + SHORT:
		lenArra, _ := imp.ReadUShort()
		arra := make([]int16, 0)
		for i := uint16(0); i < lenArra; i++ {
			val, _ := imp.ReadShort()
			arra = append(arra, val)
		}
		data = append(data, arra)
	case ARRAY + INT:
		lenArra, _ := imp.ReadUShort()
		arra := make([]int32, 0)
		for i := uint16(0); i < lenArra; i++ {
			val, _ := imp.ReadInt()
			arra = append(arra, val)
		}
		data = append(data, arra)
	case ARRAY + INT64:
		lenArra, _ := imp.ReadUShort()
		arra := make([]int64, 0)
		for i := uint16(0); i < lenArra; i++ {
			val, _ := imp.ReadInt64()
			arra = append(arra, val)
		}
		data = append(data, arra)
	case ARRAY + FLOAT:
		lenArra, _ := imp.ReadUShort()
		arra := make([]float32, 0)
		for i := uint16(0); i < lenArra; i++ {
			val, _ := imp.ReadFloat()
			arra = append(arra, val)
		}
		data = append(data, arra)
	}

	//mybase.D("ReadSimple data=%v", data)
	return data
}
func (imp *Packager) ReadData() ([]interface{}, error) {
	var flag int8
	var err error
	var data = make([]interface{}, 0)
	for {
		//mybase.D("ReadData data=%v", data)
		flag, err = imp.ReadChar()
		if err != nil {
			if err == io.EOF {
				return data, nil
			}
			return data, err
		}
		if flag < CHAR || flag > ARRAY+OTHER {
			return data, mybase.ErrBuffer
		}

		switch flag {
		case ARRAY + OTHER:
			lenArra, _ := imp.ReadUShort()
			data1 := make([]interface{}, 0)
			for i := uint16(0); i < lenArra; i++ {
				lenStruct, _ := imp.ReadUShort()
				buf := make([]byte, lenStruct)
				_, err := imp.Read(buf)
				if err != nil {
					return data, mybase.ErrParam
				}

				packager := new(Packager)
				_, err = packager.Write(buf)
				if err != nil {
					return data, err
				}

				data2, _ := packager.ReadData()
				data1 = append(data1, data2)
			}
			data = append(data, data1)
		default:
			data = imp.ReadSimple(flag, data)
		}
	}
}

func (imp *Packager) ReadPackage() (*Head, []byte, error) {
	lenBuf := imp.Len()
	lenPackage, err := imp.ReadUShort()
	if err != nil {
		return nil, nil, err
	}

	if int(lenPackage+2) != lenBuf {
		return nil, nil, mybase.ErrBuffer
	}

	var head Head
	head.Cmd, _ = imp.ReadUShort()
	head.Seq, _ = imp.ReadUShort()
	head.Ret, _ = imp.ReadShort()

	dataBuf := make([]byte, imp.Len())
	_, _ = imp.Read(dataBuf)
	return &head, dataBuf, nil
}

func (imp *Packager) WriteChar(val int8) {
	_ = binary.Write(imp, binary.BigEndian, val)
}
func (imp *Packager) WriteShort(val int16) {
	_ = binary.Write(imp, binary.BigEndian, val)
}
func (imp *Packager) WriteUShort(val uint16) {
	_ = binary.Write(imp, binary.BigEndian, val)
}
func (imp *Packager) WriteInt(val int32) {
	_ = binary.Write(imp, binary.BigEndian, val)
}
func (imp *Packager) WriteInt64(val int64) {
	_ = binary.Write(imp, binary.BigEndian, val)
}
func (imp *Packager) WriteFloat(val float32) {
	_ = binary.Write(imp, binary.BigEndian, val)
}
func (imp *Packager) WriteSimple(data interface{}) error {
	dataType := reflect.TypeOf(data) //获取数据类型
	dataVal := reflect.ValueOf(data) //获取到数据指针的值
	if dataType.Kind() == reflect.Int8 {
		imp.WriteChar(CHAR)
		imp.WriteChar(data.(int8))
	} else if dataType.Kind() == reflect.Int16 {
		imp.WriteChar(SHORT)
		imp.WriteShort(data.(int16))
	} else if dataType.Kind() == reflect.Int || dataType.Kind() == reflect.Int32 {
		imp.WriteChar(INT)
		imp.WriteInt(data.(int32))
	} else if dataType.Kind() == reflect.Int64 {
		imp.WriteChar(INT64)
		imp.WriteInt64(data.(int64))
	} else if dataType.Kind() == reflect.Float32 {
		imp.WriteChar(FLOAT)
		imp.WriteFloat(data.(float32))
	} else if dataType.Kind() == reflect.String {
		val := data.(string)
		imp.WriteChar(ARRAY + CHAR)
		imp.WriteUShort(uint16(len(val)))
		_, _ = imp.WriteString(val)
	} else if dataType.Kind() == reflect.Array || dataType.Kind() == reflect.Slice {
		elemType := dataType.Elem()
		if elemType.Kind() == reflect.Int8 || elemType.Kind() == reflect.Uint8 { //请用string
			val := data.([]byte)
			_, _ = imp.Write(val)
		} else if elemType.Kind() == reflect.Int16 {
			val := data.([]int16)
			lenVal := len(val)
			imp.WriteChar(ARRAY + SHORT)
			imp.WriteUShort(uint16(lenVal))
			for i := 0; i < lenVal; i++ {
				imp.WriteShort(val[i])
			}
		} else if elemType.Kind() == reflect.Int || dataType.Kind() == reflect.Int32 {
			val := data.([]int32)
			lenVal := len(val)
			imp.WriteChar(ARRAY + INT)
			imp.WriteUShort(uint16(lenVal))
			for i := 0; i < lenVal; i++ {
				imp.WriteInt(val[i])
			}
		} else if elemType.Kind() == reflect.Int64 {
			val := data.([]int64)
			lenVal := len(val)
			imp.WriteChar(ARRAY + INT64)
			imp.WriteUShort(uint16(lenVal))
			for i := 0; i < lenVal; i++ {
				imp.WriteInt64(val[i])
			}
		} else if elemType.Kind() == reflect.Float32 {
			val := data.([]float32)
			lenVal := len(val)
			imp.WriteChar(ARRAY + FLOAT)
			imp.WriteUShort(uint16(lenVal))
			for i := 0; i < lenVal; i++ {
				imp.WriteFloat(val[i])
			}
		} else if elemType.Kind() == reflect.String {
			val := data.([]string)
			lenVal := len(val)
			imp.WriteChar(ARRAY + OTHER)
			imp.WriteUShort(uint16(lenVal))
			for i := 0; i < lenVal; i++ {
				val1 := val[i]
				lenVal1 := len(val1)
				imp.WriteUShort(uint16(lenVal1))
				_, _ = imp.WriteString(val1)
			}
		} else if elemType.Kind() == reflect.Struct {
			elemVal := dataVal.Elem()
			val := data.([]interface{})
			lenVal := len(val)
			imp.WriteChar(ARRAY + OTHER)
			imp.WriteUShort(uint16(lenVal))
			for i := 0; i < lenVal; i++ {
				packager := new(Packager)
				for i := 0; i < elemVal.NumField(); i++ {
					var df = elemVal.Field(i)
					_ = packager.WriteSimple(df.Interface())
				}
				imp.WriteUShort(uint16(packager.Len()))
				_, _ = imp.Write(packager.Bytes())
			}
		}
	} else if dataType.Kind() == reflect.Struct { //不允许Struct里面再有Struct
		for i := 0; i < dataVal.NumField(); i++ {
			var df = dataVal.Field(i)
			_ = imp.WriteSimple(df.Interface())
		}
	} else {
		return mybase.ErrParam
	}

	return nil
}
func (imp *Packager) WriteSimpleEx(datas []interface{}) error {
	for i := range datas {
		err := imp.WriteSimple(datas[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (imp *Packager) WritePackage(head *Head, datas []interface{}) error {
	packager := new(Packager)
	packager.WriteUShort(head.Cmd)
	packager.WriteUShort(head.Seq)
	packager.WriteShort(head.Ret)
	err := packager.WriteSimpleEx(datas)
	if err != nil {
		return err
	}

	imp.WriteUShort(uint16(packager.Len() - 2)) //写入包长
	_, err = imp.Write(packager.Bytes())        //写入内容
	return err
}
