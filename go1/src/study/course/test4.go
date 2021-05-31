package course

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

//学习 字符串常用函数，数组，切片(动态数组),字典

/*
字符串操作
第i个字节并不一定是字符串的第i个字符，
因为对于非ASCII字符的UTF8编码会要两个或多个 字节。

go中的rune类型（int32） 对应unicode的码点，
UTF-32或UCS-4 每个字符都对应一个 int32,32位字节

UTF-8 是对UTF-32内存占用大的优化
0xxxxxxx 							runes 0-127 (ASCII)
110xxxxx 10xxxxxx 					128-2047 (values <128 unused)
1110xxxx 10xxxxxx 10xxxxxx 			2048-65535 (values <2048 unused)
11110xxx 10xxxxxx 10xxxxxx 10xxxxxx 65536-0x10ffff (other values unused)
*/
func testStr() {
	fmt.Println("\n\n" + strings.Repeat("*", 15) + "字符串" + strings.Repeat("*", 15))
	var s = "hello world"
	//反引号表示原生字符串 `
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
	fmt.Println(s[1:1]) //=》空字符串
	fmt.Println(s[1:4]) // 前闭后开区间
	//fmt.Println(s[3:2]) // 前大后小会报错
	//fmt.Println(s[-5:-2]) // 也不支持负数
	fmt.Println("一个英文长度为1，utf8中文是3,个别utf8是4。长度=", len(s), len("中"), len("中文"), len("𝌆")) // 查看字符串长度

	fmt.Println("字符串复制", strings.Repeat("*", 30))

	fmt.Println(fmt.Sprintf("字符串格式化 the first code is %s %d %s", s, 1, "jack son"))
	fmt.Println("字符串成员运算符in,Hello 是否存在", strings.Contains(s, "Hello"))
	fmt.Println("字符串成员运算符in,Hello 是否存在", !strings.Contains(s, "Hello"))
	fmt.Println("字符串全部转为大写字母", strings.ToUpper(s))
	fmt.Println("字符串全部转为小写字母", strings.ToLower(s))
	fmt.Println("Title", strings.Title(s))     // 把每个单词的第一个字母转化为大写，其余小写
	fmt.Println("ToTitle", strings.ToTitle(s)) //作用同ToUpper

	fmt.Println("字符串查找Index", strings.Index(s, "l"))        // 4指定起始位置 返回在字符串中的开始位置 找不到返回-1
	fmt.Println("字符串查找IndexAny", strings.IndexAny(s, "le")) //any的意思就是找到指定字符串中的任意字符
	fmt.Println("字符串查找LastIndex", strings.LastIndex(s, "l"))
	fmt.Println("字符串查找IndexRune", strings.IndexRune(s, '国')) //l字节
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
	fmt.Println(sign2, hex.EncodeToString(sign[:]))

	//字符串转byte []byte(字符串)
	//byte转字符串 string([]byte)

	//unicode 码点：
	//'世' '\u4e16' '\U00004e16'
	fmt.Printf("unicode码点: %c %c %c\n", '世', '\u4e16', '\U00004e16')
	str := "hello 世界"
	fmt.Printf("【hello 世界】utf8字节长度len=%d\n"+
		"utf8字符长度RuneCountInString=%d\n", len(str), utf8.RuneCountInString(str))
	for i, r := range str { //遍历出来的是utf8实际字符个数
		fmt.Printf("str[i] %d,%q\t%[2]c\t%[2]d\n", i, str[i])
		fmt.Printf("r      %d,%q\t%[2]c\t%[2]d\n", i, r)
	}

	fmt.Println(string(0x4e16))  //码点转换成utf8字符串
	fmt.Println(string(1234567)) //如果对应码点的字符是无效的，则用'\uFFFD'无效字符作为替换

	//strings strconv bytes unicode
}

//数组和切片操作
func testArra() {
	//数组和切片
	//在go语言中，传递给函数的数组对象是一个复制后的副本，不像其他语言是一个引用或是指针
	//当然我们可以直接传递 数组指针 来达到高效传递参数的目的。

	//数组
	//数组的长度是数组类型的一个组成部分，因此[3]int和[4]int是两种不同的数组类型。
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	var s [10]int = [10]int{9: 1} //用冒号可以指定位置初始化，其他填充默认零值
	fmt.Println(r, q, s)          // "0"

	//如果在数组的长度位置出现的是“...”省略号，则表示数组的长度是根据初始 化值的个数来计算
	animal := [...]string{"cat", "dog", "fish", "bird"}
	for i, v := range animal {
		fmt.Println("animal["+strconv.Itoa(i)+"]=", v)
	}
	//go 不支持负索引
	//fmt.Println(animal[-1])

	//动态数组-slice 跟数组相比，数组是可以当map的key的而slice不可以。
	/**
	var s []int // len(s) == 0, s == nil
	s = nil // len(s) == 0, s == nil
	s = []int(nil) // len(s) == 0, s == nil
	s = []int{} // len(s) == 0, s != nil

	//切片的两个创建方式。
	make([]T, len)
	make([]T, len, cap) // same as make([]T, cap)[:len]
	*/
	var slice1 []int   //申明切片
	if slice1 == nil { //默认是nil值
		fmt.Println("slice1 is nil")
	}
	fmt.Println("append之前 slice1=", slice1, len(slice1), cap(slice1)) //空切片
	slice2 := append(slice1, 1, 2)                                    //可同时添加多个元素
	fmt.Println("append之后 slice1=", slice1, len(slice1), cap(slice1), slice1 == nil)
	fmt.Println("slice2=", slice2, len(slice2), cap(slice2))

	//定义slice，
	fmt.Println("slice实验，修改前")
	var bytes []byte
	bytes = append(bytes, 1, 2, 3, 4, 5, 6, 7, 8)
	for _, b := range bytes {
		fmt.Print(b, " ")
	}
	fmt.Println()
	pos := 2
	bytes1 := bytes[pos : pos+2] //获取其中的部分切片
	bytes1[0] = 30               //修改其中的内容
	bytes1[1] = 40
	fmt.Println("slice实验，修改后")
	for _, b := range bytes {
		fmt.Print(b, " ")
	}
	fmt.Println()

	//使用make创建一个切片
	slice3 := make([]int, 1, 3) //类型，大小(默认插入0，1个就插入1个0)，容量
	//
	fmt.Println("make创建切片 slice3=", slice3, len(slice3), cap(slice3))
	copy(slice3, slice2) //拷贝slice2到slice3中，如果目标长度小于源，则剪裁min(len(dst),len(src))
	fmt.Println("copy后的切片 slice3=", slice3, len(slice3), cap(slice3))
	slice3 = append(slice3, 10, 11, 12, 13, 14, 15)
	fmt.Println("append后的切片 slice3=", slice3, len(slice3), cap(slice3))
	//指定位置copy
	slice4 := []int{101, 102, 103, 104}
	copy(slice3[3:], slice4[1:3])
	fmt.Println("指定位置copy后的切片 slice3=", slice3, len(slice3), cap(slice3))
	slice5 := append(slice3[3:3], 1000)
	fmt.Println("指定位置copy后的切片 slice3=", slice3, len(slice3), cap(slice3))
	fmt.Println("指定位置copy后的切片 slice5=", slice5, len(slice5), cap(slice5))

	//注意！！！！
	slice6 := slice4[:]                //这里无法实现python的深拷贝
	slice6[0] = 106                    //这样做会修改slice4中的值
	slice7 := make([]int, len(slice4)) //正确的深拷贝示范。。！！
	copy(slice7, slice4)
	slice7[0] = 107
	fmt.Printf("slice4=%v,slice6=%v,slice7=%v\n", slice4, slice6, slice7)

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

	var arras [][]int //数组长度不能使用变量定义

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
	_, ok := map1["A"] //但是可以对ni值查询。
	fmt.Printf("map1=%v,ok=%t", map1, ok)
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

	map1["A"] = 66
	map1["AA"] = 67
	map1["AAA"] = 68
	map1["AAAA"] = 69
	//遍历字典
	for k, v := range map1 {
		fmt.Println("map1[", k, "]=", v)
		delete(map1, k) //@注意 go map边遍历，边移除是安全的。！！！
	}

	for k, v := range map1 {
		fmt.Println("map1[", k, "]=", v)
	}
}

type Student struct {
	id    int32  `json:"id"`
	name  string `json:"name"`
	score int    `json:"score"`
}

func testPointer1() *Student {
	student := Student{
		101, "Jack", 0,
	}

	//在c++中，这是个局部变量，返回局部变量，该指针就会变成野指针，打印的信息会乱码
	//但是在go中却不会，应该是动态创建的变量
	return &student
}
func testPointer() {
	fmt.Printf("student=%v\n", testPointer1())
	fmt.Printf("student=%v\n", *testPointer1())
}

func testMapStruct() {
	dict := make(map[int32]Student)
	dict[1] = Student{1, "Jack", 59}
	dict[2] = Student{2, "Tom", 90}

	if v, ok := dict[1]; ok {
		//dict[1].score++ //这里语法错误，无法改变结构中的数据
		v.score++ //这里也是错误的用法，，v只是个深拷贝临时变量，不影响原值
	}

	fmt.Printf("修改临时深拷贝变量 name:%s,score:%d\n", dict[1].name, dict[1].score)

	//正确的方法是修改成指针map
	dict2 := make(map[int32]*Student)
	dict2[1] = &Student{1, "Jack", 59}
	if v, ok := dict2[1]; ok {
		dict2[1].score++
		v.score += 2
	}
	fmt.Printf("修改指针变量 name:%s,score:%d\n", dict[1].name, dict2[1].score)
	//@实际打印62，说明上面对原值修改和对临时指针变量的修改都生效了
}

func Course4() {
	testStr()
	testArra()
	testMap()
	testPointer()

	//@注意临时深拷贝的修改不改变原值
	testMapStruct()
}
