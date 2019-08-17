/**
公斤(千克)/斤相互转换
 */
package pkg

import "fmt"

const (
	kg2J = 2
)

type Kg float32
type Jin float32

func (imp Kg) String() string {
	return fmt.Sprintf("%g(千克/公斤)", imp)
}

func (imp Jin) String() string {
	return fmt.Sprintf("%g(斤)", imp)
}

func Kg2Jin(v Kg) Jin {
	return Jin(v * 2)
}

func Jin2Kg(v Jin) Kg {
	return Kg(v / 2)
}
