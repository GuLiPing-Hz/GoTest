package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

type Env map[Var]float64

type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
}

type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

type literal float64

func (l literal) Eval(env Env) float64 {
	return float64(l)
}
func (l literal) Check(vars map[Var]bool) error {
	return nil
}

type unary struct {
	op rune //+,-
	x  Expr
}

func (l unary) Eval(env Env) float64 {
	switch l.op {
	case '+':
		return +l.x.Eval(env)
	case '-':
		return -l.x.Eval(env)
	default:
		log.Fatalf("unsupported unary operator: %q", l.op)
		return 0
	}
}
func (l unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", l.op) {
		return fmt.Errorf("unexpected unary op %q", l.op)
	}
	return l.x.Check(vars)
}

type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

func (l binary) Eval(env Env) float64 {
	switch l.op {
	case '+':
		return l.x.Eval(env) + l.y.Eval(env)
	case '-':
		return l.x.Eval(env) - l.y.Eval(env)
	case '*':
		return l.x.Eval(env) * l.y.Eval(env)
	case '/':
		return l.x.Eval(env) / l.y.Eval(env)
	default:
		log.Fatalf("unsupported binary operator: %q", l.op)
		return 0
	}
}

func (l binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", l.op) {
		return fmt.Errorf("unexpected binary op %q", l.op)
	}
	if err := l.x.Check(vars); err != nil {
		return err
	}
	return l.y.Check(vars)
}

type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

func (l call) Eval(env Env) float64 {
	switch l.fn {
	case "pow":
		return math.Pow(l.args[0].Eval(env), l.args[1].Eval(env))
	case "sin":
		return math.Sin(l.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(l.args[0].Eval(env))
	default:
		log.Fatalf("unsupported call : %q", l.fn)
		return 0
	}
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

func (l call) Check(vars map[Var]bool) error {
	num, ok := numParams[l.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", l.fn)
	}

	if num != len(l.args) {
		return fmt.Errorf("call to %s has %d args, want %d", l.fn, len(l.args), num)
	}

	for i := range l.args {
		if err := l.args[i].Check(vars); err != nil {
			return err
		}
	}

	return nil
}

type ExprParse struct {
	Type int
	S    string
	exp  Expr
}

func isFunction(s string, j int) (string, int) {
	if len(s)-j > 4 && s[j:j+4] == "sqrt" {
		s2 := s[j+4:]
		pos1 := strings.Index(s2, "(")
		if pos1 != -1 && strings.TrimSpace(s2[:pos1]) == "" {
			return "sqrt", j + 4 + pos1
		}
	} else if len(s)-j > 3 && (s[j:j+3] == "pow" || s[j:j+3] == "sin") {
		s2 := s[j+3:]
		pos1 := strings.Index(s2, "(")
		if pos1 != -1 && strings.TrimSpace(s2[:pos1]) == "" {
			return s[j : j+3], j + 3 + pos1
		}
	}

	return "", j
}

func Parse(s string) (Expr, error) {
	var eps []ExprParse
	for i := 0; i < len(s); {
		j := i

		var theType = Unknow
		for j < len(s) {
			if s[j] >= '0' && s[j] <= '9' {
				if theType == Unknow || theType == Number {
					theType = Number
					j++
				} else {
					eps = append(eps, ExprParse{theType, strings.TrimSpace(s[i:j]), nil})
					i = j
					break
				}
			} else if s[j] >= 'a' && s[j] <= 'z' || s[j] >= 'A' && s[j] <= 'Z' {
				if theType == Unknow || theType == Name {
					if theType == Unknow {
						if funcName, k := isFunction(s, j); funcName != "" {
							eps = append(eps, ExprParse{Func, funcName, nil})
							i = k
							break
						}
					}
					theType = Name
					j++
				} else {
					eps = append(eps, ExprParse{theType, strings.TrimSpace(s[i:j]), nil})
					i = j
					break
				}
			} else if strings.ContainsRune("+-*/", rune(s[j])) {
				if theType != Unknow {
					eps = append(eps, ExprParse{theType, strings.TrimSpace(s[i:j]), nil})
				}

				isJJ := strings.ContainsRune("+-", rune(s[j]))
				op := Operator1
				if isJJ {
					op = Operator2
				}
				eps = append(eps, ExprParse{op, strings.TrimSpace(s[j : j+1]), nil})
				i = j + 1
				break
			} else if s[j] == '(' || s[j] == ')' {
				if theType != Unknow {
					eps = append(eps, ExprParse{theType, strings.TrimSpace(s[i:j]), nil})
				}
				op := BracketL
				if s[j] == ')' {
					op = BracketR
				}
				eps = append(eps, ExprParse{op, strings.TrimSpace(s[j : j+1]), nil})
				i = j + 1
				break
			} else if unicode.IsSpace(rune(s[j])) {
				if theType != Unknow {
					eps = append(eps, ExprParse{theType, strings.TrimSpace(s[i:j]), nil})
					i = j + 1
					break
				}
				j++
			} else if s[j] == ',' {
				if theType != Unknow {
					eps = append(eps, ExprParse{theType, strings.TrimSpace(s[i:j]), nil})
				}
				eps = append(eps, ExprParse{Comma, strings.TrimSpace(s[j : j+1]), nil})
				i = j + 1
				break
			}
		}

		if j >= len(s) {
			if theType != Unknow {
				eps = append(eps, ExprParse{theType, strings.TrimSpace(s[i:j]), nil})
			}
			break
		}
	}

	var exprs = make([]ExprParse, 0)
	for i := 0; i < len(eps); i++ {
		data := eps[i]
		if len(exprs) == 0 {
			exp, err := Parse2(data)
			if err != nil {
				return nil, err
			}
			exprs = append(exprs, ExprParse{data.Type, data.S, exp})
		} else if data.Type == Name || data.Type == Number {
			exp, err := Parse2(data)
			if err != nil {
				return nil, err
			}

			top := exprs[len(exprs)-1]
			if top.exp == nil {
				if top.Type == Operator2 {
					top.exp = unary{rune(top.S[0]), exp}
					continue
				}

				exprs = append(exprs, ExprParse{data.Type, data.S, exp})
			} else if top.Type == Operator2 || top.Type == Operator1 {
				var nextData *ExprParse
				if i+1 < len(eps) {
					nextData = &eps[i+1]
				}

				if nextData == nil || nextData.Type <= top.Type {
					operation, ok := top.exp.(binary)
					if !ok {
						return nil, io.EOF
					}
					operation.y = exp
					exprs[len(exprs)-1].exp = operation
				} else {
					exprs = append(exprs, ExprParse{data.Type, data.S, exp})
				}
			}
		} else if data.Type == BracketL || data.Type == Func || data.Type == Comma {
			exprs = append(exprs, ExprParse{data.Type, data.S, nil})
		} else if data.Type == Operator2 { //+-
			top := exprs[len(exprs)-1]
			if top.exp == nil {
				if top.Type == BracketL || top.Type == Func {
					exprs = append(exprs, ExprParse{data.Type, data.S, nil})
				} else {
					return nil, io.EOF
				}
			} else {
				exp := binary{rune(data.S[0]), top.exp, nil}
				//top = ExprParse{data.Type, data.S, exp} //这里只是更改了临时变量的值
				exprs[len(exprs)-1] = ExprParse{data.Type, data.S, exp}
			}
		} else if data.Type == Operator1 { //*/
			top := exprs[len(exprs)-1]
			if top.exp == nil {
				return nil, io.EOF
			} else {
				exprs[len(exprs)-1] = ExprParse{data.Type, data.S, binary{rune(data.S[0]), top.exp, nil}}
			}
		} else if data.Type == BracketR {
			for {
				top := exprs[len(exprs)-1]
				if top.Type == Func {
					break
				} else if top.Type == BracketL {
					return nil, io.EOF
				}

				top2 := exprs[len(exprs)-2] // ,
				if top2.Type == Comma {
					top3 := exprs[len(exprs)-3] // exp
					top4 := exprs[len(exprs)-4] // func

					var args []Expr
					args = append(args, top3.exp, top.exp)
					exprs[len(exprs)-4].exp = call{top4.S, args}
					exprs = exprs[:len(exprs)-4+1]
				} else if top2.Type == Func {
					var args []Expr
					args = append(args, top.exp)
					exprs[len(exprs)-2].exp = call{top2.S, args}
					exprs = exprs[:len(exprs)-1]
				} else if top2.Type == BracketL {
					exprs[len(exprs)-2].exp = top.exp
					exprs = exprs[:len(exprs)-1]
				}
			}
		}
	}

	return exprs[0].exp, nil
}

const (
	Unknow = iota
	Number
	Name
	Func

	BracketR
	Comma      //小出栈
	Operator2  //+ -
	Operator1  //* / 大入栈
	BracketL
)

func Parse2(data ExprParse) (Expr, error) {
	if data.Type == Number {
		v, err := strconv.ParseFloat(data.S, 64)
		if err != nil {
			return nil, err
		}
		return literal(v), nil
	} else if data.Type == Name {
		return Var(data.S), nil
	}
	return nil, nil
}

func parseAndCheck(s string) (Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	}
	var prevExpr string
	for _, test := range tests { // Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n", test.expr, test.env, got, test.want)
		}
	}
}

func main() {
	//Parse("sqrt(A / pi)")
	//Parse("pow(x, 3) + pow(y, 3)")
	//Parse("5 / 9 * (F - 32)")
	exp, err := Parse("x+3*4")
	if err != nil {
		fmt.Printf("err=%s\n", err.Error())
		return
	}

	result := exp.Eval(Env{"x": 10})
	fmt.Printf("result=%f\n", result)
}
