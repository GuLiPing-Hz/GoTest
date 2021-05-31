package main

import (
	"flag"
	"fmt"
	"go1/src/pkglearn"
	"strconv"
)

func main() {
	var changeType string
	flag.StringVar(&changeType, "type", "", "kg/jin")
	flag.Parse() //解析

	fmt.Printf("type=%s\n", changeType)
	paramStr := flag.Arg(0)
	param, err := strconv.ParseFloat(paramStr, 32)
	if err != nil {
		fmt.Printf("err=%s\n", err.Error())
	}

	switch changeType {
	case "kg":
		fmt.Printf("kg:%v\n", pkglearn.Kg2Jin(pkglearn.Kg(param)))
	case "jin":
		fmt.Printf("jin:%v\n", pkglearn.Jin2Kg(pkglearn.Jin(param)))
	default:
		fmt.Printf("unknow type")
	}
}
