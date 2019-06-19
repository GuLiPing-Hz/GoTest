package net

import (
	"bytes"
	"encoding/binary"
	"io"
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

func (this *Packager) ReadChar() (int8, error) {
	var val int8 = 0
	err := binary.Read(this, binary.BigEndian, &val)
	return val, err
}
func (this *Packager) ReadShort() (int16, error) {
	var val int16 = 0
	err := binary.Read(this, binary.BigEndian, &val)
	return val, err
}
func (this *Packager) ReadUShort() (uint16, error) {
	var val uint16 = 0
	err := binary.Read(this, binary.BigEndian, &val)
	return val, err
}
func (this *Packager) ReadInt() (int32, error) {
	var val int32 = 0
	err := binary.Read(this, binary.BigEndian, &val)
	return val, err
}
func (this *Packager) ReadInt64() (int64, error) {
	var val int64 = 0
	err := binary.Read(this, binary.BigEndian, &val)
	return val, err
}
func (this *Packager) ReadFloat() (float32, error) {
	var val float32 = 0
	err := binary.Read(this, binary.BigEndian, &val)
	return val, err
}
func (this *Packager) ReadSimple(flag int8, data []interface{}) []interface{} {
	switch flag {
	case CHAR:
		val, _ := this.ReadChar()
		data = append(data, val)
	case SHORT:
		val, _ := this.ReadShort()
		data = append(data, val)
	case INT:
		val, _ := this.ReadInt()
		data = append(data, val)
	case INT64:
		val, _ := this.ReadInt64()
		data = append(data, val)
	case FLOAT:
		val, _ := this.ReadFloat()
		data = append(data, val)
	case STRING:
		lenArra, _ := this.ReadUShort()
		buf := make([]byte, lenArra)
		_, _ = this.Read(buf)
		data = append(data, string(buf))
	case ARRAY + SHORT:
		lenArra, _ := this.ReadUShort()
		arra := make([]int16, 0)
		for i := uint16(0); i < lenArra; i++ {
			val, _ := this.ReadShort()
			arra = append(arra, val)
		}
		data = append(data, arra)
	case ARRAY + INT:
		lenArra, _ := this.ReadUShort()
		arra := make([]int32, 0)
		for i := uint16(0); i < lenArra; i++ {
			val, _ := this.ReadInt()
			arra = append(arra, val)
		}
		data = append(data, arra)
	case ARRAY + INT64:
		lenArra, _ := this.ReadUShort()
		arra := make([]int64, 0)
		for i := uint16(0); i < lenArra; i++ {
			val, _ := this.ReadInt64()
			arra = append(arra, val)
		}
		data = append(data, arra)
	case ARRAY + FLOAT:
		lenArra, _ := this.ReadUShort()
		arra := make([]float32, 0)
		for i := uint16(0); i < lenArra; i++ {
			val, _ := this.ReadFloat()
			arra = append(arra, val)
		}
		data = append(data, arra)
	}

	//util.D("ReadSimple data=%v", data)
	return data
}
func (this *Packager) ReadData() ([]interface{}, error) {
	var flag int8
	var err error
	var data = make([]interface{}, 0)
	for {
		//util.D("ReadData data=%v", data)
		flag, err = this.ReadChar()
		if err != nil {
			if err == io.EOF {
				return data, nil
			}
			return data, err
		}
		if flag < CHAR || flag > ARRAY+OTHER {
			return data, ErrBuffer
		}

		switch flag {
		case ARRAY + OTHER:
			lenArra, _ := this.ReadUShort()
			data1 := make([]interface{}, 0)
			for i := uint16(0); i < lenArra; i++ {
				lenStruct, _ := this.ReadUShort()
				buf := make([]byte, lenStruct)
				_, err := this.Read(buf)
				if err != nil {
					return data, ErrParam
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
			data = this.ReadSimple(flag, data)
		}
	}
}

func (this *Packager) ReadPackage() (*Head, []byte, error) {
	lenBuf := this.Len()
	lenPackage, err := this.ReadUShort()
	if err != nil {
		return nil, nil, err
	}

	if int(lenPackage+2) != lenBuf {
		return nil, nil, ErrBuffer
	}

	var head Head
	head.Cmd, _ = this.ReadUShort()
	head.Seq, _ = this.ReadUShort()
	head.Ret, _ = this.ReadShort()

	dataBuf := make([]byte, this.Len())
	_, _ = this.Read(dataBuf)
	return &head, dataBuf, nil
}

func (this *Packager) WriteChar(val int8) {
	_ = binary.Write(this, binary.BigEndian, val)
}
func (this *Packager) WriteShort(val int16) {
	_ = binary.Write(this, binary.BigEndian, val)
}
func (this *Packager) WriteUShort(val uint16) {
	_ = binary.Write(this, binary.BigEndian, val)
}
func (this *Packager) WriteInt(val int32) {
	_ = binary.Write(this, binary.BigEndian, val)
}
func (this *Packager) WriteInt64(val int64) {
	_ = binary.Write(this, binary.BigEndian, val)
}
func (this *Packager) WriteFloat(val float32) {
	_ = binary.Write(this, binary.BigEndian, val)
}
func (this *Packager) WriteSimple(data interface{}) error {
	dataType := reflect.TypeOf(data) //获取数据类型
	dataVal := reflect.ValueOf(data) //获取到数据指针的值
	if dataType.Kind() == reflect.Int8 {
		this.WriteChar(CHAR)
		this.WriteChar(data.(int8))
	} else if dataType.Kind() == reflect.Int16 {
		this.WriteChar(SHORT)
		this.WriteShort(data.(int16))
	} else if dataType.Kind() == reflect.Int || dataType.Kind() == reflect.Int32 {
		this.WriteChar(INT)
		this.WriteInt(data.(int32))
	} else if dataType.Kind() == reflect.Int64 {
		this.WriteChar(INT64)
		this.WriteInt64(data.(int64))
	} else if dataType.Kind() == reflect.Float32 {
		this.WriteChar(FLOAT)
		this.WriteFloat(data.(float32))
	} else if dataType.Kind() == reflect.String {
		val := data.(string)
		this.WriteChar(ARRAY + CHAR)
		this.WriteUShort(uint16(len(val)))
		_, _ = this.WriteString(val)
	} else if dataType.Kind() == reflect.Array || dataType.Kind() == reflect.Slice {
		elemType := dataType.Elem()
		if elemType.Kind() == reflect.Int8 || elemType.Kind() == reflect.Uint8 { //请用string
			val := data.([]byte)
			_, _ = this.Write(val)
		} else if elemType.Kind() == reflect.Int16 {
			val := data.([]int16)
			lenVal := len(val)
			this.WriteChar(ARRAY + SHORT)
			this.WriteUShort(uint16(lenVal))
			for i := 0; i < lenVal; i++ {
				this.WriteShort(val[i])
			}
		} else if elemType.Kind() == reflect.Int || dataType.Kind() == reflect.Int32 {
			val := data.([]int32)
			lenVal := len(val)
			this.WriteChar(ARRAY + INT)
			this.WriteUShort(uint16(lenVal))
			for i := 0; i < lenVal; i++ {
				this.WriteInt(val[i])
			}
		} else if elemType.Kind() == reflect.Int64 {
			val := data.([]int64)
			lenVal := len(val)
			this.WriteChar(ARRAY + INT64)
			this.WriteUShort(uint16(lenVal))
			for i := 0; i < lenVal; i++ {
				this.WriteInt64(val[i])
			}
		} else if elemType.Kind() == reflect.Float32 {
			val := data.([]float32)
			lenVal := len(val)
			this.WriteChar(ARRAY + FLOAT)
			this.WriteUShort(uint16(lenVal))
			for i := 0; i < lenVal; i++ {
				this.WriteFloat(val[i])
			}
		} else if elemType.Kind() == reflect.String {
			val := data.([]string)
			lenVal := len(val)
			this.WriteChar(ARRAY + OTHER)
			this.WriteUShort(uint16(lenVal))
			for i := 0; i < lenVal; i++ {
				val1 := val[i]
				lenVal1 := len(val1)
				this.WriteUShort(uint16(lenVal1))
				_, _ = this.WriteString(val1)
			}
		} else if elemType.Kind() == reflect.Struct {
			elemVal := dataVal.Elem()
			val := data.([]interface{})
			lenVal := len(val)
			this.WriteChar(ARRAY + OTHER)
			this.WriteUShort(uint16(lenVal))
			for i := 0; i < lenVal; i++ {
				packager := new(Packager)
				for i := 0; i < elemVal.NumField(); i++ {
					var df = elemVal.Field(i)
					_ = packager.WriteSimple(df.Interface())
				}
				this.WriteUShort(uint16(packager.Len()))
				_, _ = this.Write(packager.Bytes())
			}
		}
	} else if dataType.Kind() == reflect.Struct { //不允许Struct里面再有Struct
		for i := 0; i < dataVal.NumField(); i++ {
			var df = dataVal.Field(i)
			_ = this.WriteSimple(df.Interface())
		}
	} else {
		return ErrParam
	}

	return nil
}
func (this *Packager) WriteSimpleEx(datas []interface{}) error {
	for i := range datas {
		err := this.WriteSimple(datas[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *Packager) WritePackage(head *Head, datas []interface{}) error {
	packager := new(Packager)
	packager.WriteUShort(head.Cmd)
	packager.WriteUShort(head.Seq)
	packager.WriteShort(head.Ret)
	err := packager.WriteSimpleEx(datas)
	if err != nil {
		return err
	}

	this.WriteUShort(uint16(packager.Len() - 2)) //写入包长
	_, err = this.Write(packager.Bytes())        //写入内容
	return err
}
