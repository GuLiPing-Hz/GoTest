package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

//学习 字符串常用函数，数组，切片(动态数组),字典

//字符串操作
func testStr() {
	fmt.Println("\n\n" + strings.Repeat("*", 15) + "字符串" + strings.Repeat("*", 15))
	var s = "hello world"
	var s1 = `hello world` //go语言不能使用单引号表示字符串，单引号只能表示某个字符
	s2 := `hello
	world`
	fmt.Println(s, " ", s1, " ", s2)
	fmt.Println("字符串拼接:" + strconv.Itoa(123))
	fmt.Println(strings.Repeat("*", 30))

	//go对字符串的访问
	s3 := s[:] //字符串拷贝出来的地址 python一样，go不一样
	fmt.Println(s3, " ", &s, " ", &s3, " ", s == s3, " ", &s == &s3)
	fmt.Println(s[1:])  //索引起点都是0
	fmt.Println(s[1:4]) // 前闭后开区间
	//fmt.Println(s[3:2]) // 前大后小会报错
	//fmt.Println(s[-5:-2]) // 也不支持负数
	fmt.Println("一个英文长度为1，utf8中文是3，长度=", len(s), len("中"), len("中文")) // 查看字符串长度

	fmt.Println("字符串复制", strings.Repeat("*", 30))

	fmt.Println(fmt.Sprintf("字符串格式化 the first code is %s %d %s", s, 1, "jack son"))
	fmt.Println("字符串成员运算符in,Hello 是否存在", strings.Contains(s, "Hello"))
	fmt.Println("字符串成员运算符in,Hello 是否存在", !strings.Contains(s, "Hello"))
	fmt.Println("字符串全部转为大写字母", strings.ToUpper(s))
	fmt.Println("字符串全部转为小写字母", strings.ToLower(s))
	fmt.Println("Title", strings.Title(s))     // 把每个单词的第一个字母转化为大写，其余小写
	fmt.Println("ToTitle", strings.ToTitle(s)) //作用同ToUpper

	fmt.Println("字符串查找Index", strings.Index(s, "l")) // 4指定起始位置 返回在字符串中的开始位置 找不到返回-1
	fmt.Println("字符串查找IndexAny", strings.IndexAny(s, "l"))
	fmt.Println("字符串查找IndexAny2", strings.IndexAny(s, "le")) //any的意思就是找到指定字符串中的任意字符
	fmt.Println("字符串查找LastIndex", strings.LastIndex(s, "l"))
	fmt.Println("字符串查找IndexRune", strings.IndexRune(s, 'l')) //l字节
	fmt.Println("字符串替换", strings.Replace(s, "l", "L", 2))    // 最后一个参数是要替换的次数，-1全部替换
	fmt.Println("字符串替换", strings.Replace(s, "l", "L", -1))
	////去除字符串左空格 TrimLeft ，右空格 TrimRight,左右空格Trim
	fmt.Println("Trim=" + strings.Trim(" ss s1  ", " ")) //等价于TrimSpace
	fmt.Println("Trim=" + strings.TrimSpace(" ss s1  "))

	ss := []string{"a", "b", "c"} //申明字符串数组
	ss2 := ss[:]                  //数组拷贝出来的地址 python不一样，go不一样,似乎只有字符串的处理是不一样的
	fmt.Println(ss2, " ", &ss, " ", &ss2, " ", &ss == &ss2)
	fmt.Println("字符串以指定连接符连接", strings.Join(ss, "_")) //这个要比用for循环 [s = s+sep+s] 更高效
	fmt.Println("字符串以指定字符串分隔", strings.Split(s, "l"))
	fmt.Println("字符串检查指定字符串重复出现次数", strings.Count(s, "l"))
	fmt.Println("字符串开头检查", s, strings.HasPrefix(s, "Hell"))
	fmt.Println("字符串结束检查", strings.HasSuffix(s, "world"))
	fmt.Println("字符串比较", s, "\n", s2, "\n", strings.EqualFold(s, s2), s == s2)
	fmt.Println("字符串分割Fields", strings.Fields(s)) //以空格分割字符串返回字符串数组

	var plainText = "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(plainText[0 : len(plainText)-1])

	var sign = md5.Sum([]byte("123"))
	sign2 := fmt.Sprintf("%x", sign)
	fmt.Println(sign2)

	//字符串转byte []byte(字符串)
	//byte转字符串 string([]byte)
}

//数组操作
func testArra() {
	//数组测试
	//静态数组
	animal := []string{"cat", "dog", "fish", "bird"}
	for i, v := range animal {
		fmt.Println("animal["+strconv.Itoa(i)+"]=", v)
	}
	//go 不支持负索引
	//fmt.Println(animal[-1])

	//动态数组，使用切片
	var slice1 []int //申明切片
	if slice1 == nil { //默认是nil值
		fmt.Println("slice1 is nil")
	}
	fmt.Println("append之前 slice1=", slice1, len(slice1), cap(slice1)) //空切片
	slice2 := append(slice1, 1, 2)                                    //可同时添加多个元素
	fmt.Println("append之后 slice1=", slice1, len(slice1), cap(slice1), slice1 == nil)
	fmt.Println("slice2=", slice2, len(slice2), cap(slice2))

	//使用make创建一个切片
	slice3 := make([]int, 1, 5) //类型，大小(默认插入0，1个就插入1个0)，容量
	//
	fmt.Println("make创建切片 slice3=", slice3, len(slice3), cap(slice3))
	copy(slice3, slice2) //拷贝slice2到slice3中，如果目标长度小于源，则剪裁
	fmt.Println("copy后的切片 slice3=", slice3, len(slice3), cap(slice3))
	slice3 = append(slice3, 10)
	fmt.Println("copy后的切片 slice3=", slice3, len(slice3), cap(slice3))

	//多维数组
	//number := [2][2]int{{1, 2}, {2, 3}}
	number := [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
	fmt.Println(len(number), len(number[0]))
	for i, v1 := range number {
		fmt.Println("number=", i, v1)
		for j, v2 := range v1 {
			fmt.Println("number2=", j, v2)
		}
	}
	var arras [][]int
	var nums = 3
	for i := 0; i < nums; i++ {
		var temp []int
		for j := 0; j < nums; j++ {
			vi := i + 1
			vj := j + 1
			temp = append(temp, vi*vj)
			//fmt.Println("arras[%d][%d]=%d" % (vi,vj,vi*vj))
		}
		//把切边(数组)添加到另一个切片(数组)
		//arras = append(arras, temp...)
		arras = append(arras, temp)
	}

	for i := 0; i < nums; i++ {
		for j := 0; j < nums; j++ {
			fmt.Printf("arras[%d][%d]=%d \n", i, j, arras[i][j])
		}
	}

	fmt.Println(strings.Repeat("*", 30))
	fmt.Println(arras[0])
	fmt.Println(arras[0:2])                  // python支持列表片段访问
	fmt.Println("number[1:2]=", number[1:2]) //跟操作字符串一样，前闭后开区间

	arras[0][0] = 2
	fmt.Println("修改 arras[0][0] = ", arras[0][0])

	//删除序列中指定索引的元素
	var line = 0   //行
	var column = 0 //列
	//把一个切片增加到另一个切片中
	arras[line] = append(arras[line][:column], arras[line][column+1:]...)
	fmt.Println("删除 arras[0][0] 后 arras[0]=", arras[line])

	//不允许列表相加
	//print("列表相加", [1, 2, 3]+[4, 5, 6])
	//不允许列表乘
	//print("列表相乘", ["Hi"]*4)
}

//字典操作
func testMap() {
	//Map 是一种无序的键值对的集合.
	//Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
	//不过，Map 是无序的，我们无法决定它的返回顺序，
	//这是因为 Map 是使用 hash 表来实现的。

	var map1 map[string]int //申明一个map
	fmt.Println("map1=", map1)
	if map1 == nil {
		fmt.Println("map1 is nil")
	}
	//map1["A"] = 90//对nil赋值，崩溃
	fmt.Println("map1=", map1)
	map1 = make(map[string]int) //创建map，实例化
	map1["A"] = 90              //添加元素
	fmt.Println("插入字典 map1=", map1)
	map1["B"] = 80 //添加元素
	fmt.Println("map1=", map1)
	delete(map1, "A")
	fmt.Println("移除字典 map1=", map1)

	v, ok := map1["C"]
	if ok {
		fmt.Println("map1[C] 存在", v)
	} else {
		fmt.Println("map1[C] 不存在")
	}

	delete(map1, "C") //删除指定元素
	fmt.Println("删除C map1=", map1)
	delete(map1, "A")
	fmt.Println("删除A map1=", map1)

	//遍历字典
	for k, v := range map1 {
		fmt.Println("map1[", k, "]=", v)
	}
}

type Student struct {
	id   int32  `json:"id"`
	name string `json:"name"`
}

func testPointer1() *Student {
	student := Student{
		101, "Jack",
	}

	//在c++中，这是个局部变量，返回局部变量，该指针就会变成野指针，打印的信息会乱码
	//但是在go中却不会，应该是动态创建的变量
	return &student
}
func testPointer() {
	fmt.Printf("student=%v\n", testPointer1())
	fmt.Printf("student=%v\n", *testPointer1())
}

func main() {
	testStr()
	//testArra()
	//testMap()
	testPointer()
}
