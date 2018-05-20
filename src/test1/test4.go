package main

import (
	"fmt"
	"strconv"
	"strings"
)

//字符串操作

func testStr() {
	fmt.Println("字符串拼接:" + strconv.Itoa(123))
	s := "Hello "
	//s1 := 'Hello World'//go语言不能使用单引号表示字符串，单引号只能表示某个字符
	fmt.Println(s)
	s += "World"
	fmt.Println(s)

	//go对字符串的访问
	fmt.Println(s[:])
	fmt.Println(s[1:])  //索引起点都是0
	fmt.Println(s[1:4]) // 前闭后开区间
	//fmt.Println(s[3:2]) // 前大后小会报错
	//fmt.Println(s[-5:-2]) // 也不支持负数
	fmt.Println("一个英文长度为1，utf8中文是3，长度=", len(s), len("中"), len("中文")) // 查看字符串长度

	//fmt.Println("字符串复制","*"*30)
	// python 3 不同于 python2的%表示法 ，都用大括号表示占位，并且支持索引和关键字，跟c的fmt.Printlnf相比只是多了{:},不用写%了
	fmt.Println(fmt.Sprintf("字符串格式化 the first code is %s %d %s", s, 1, "jack son"))
	s2 := strings.ToUpper(s)
	fmt.Println("字符串全部转为大写字母", s2)
	fmt.Println("字符串全部转为小写字母", strings.ToLower(s))
	fmt.Println("Title", strings.Title(s))           // 把每个单词的第一个字母转化为大写，其余小写
	fmt.Println("ToTitle", strings.ToTitle(s))       //作用同ToUpper
	fmt.Println("字符串查找Index", strings.Index(s, "l")) // 4指定起始位置 返回在字符串中的开始位置 找不到返回-1
	fmt.Println("字符串查找IndexAny", strings.IndexAny(s, "l"))
	fmt.Println("字符串查找IndexAny2", strings.IndexAny(s, "le")) //any的意思就是找到指定字符串中的任意字符
	fmt.Println("字符串查找LastIndex", strings.LastIndex(s, "l"))
	fmt.Println("字符串查找IndexRune", strings.IndexRune(s, 'l')) //l字节
	fmt.Println("字符串替换", strings.Replace(s, "l", "L", 2))    // 最后一个参数是要替换的次数，不填，默认全部替换
	fmt.Println("字符串替换", strings.Replace(s, "l", "L", -1))
	////去除字符串左空格 TrimLeft ，右空格 TrimRight,左右空格Trim
	fmt.Println("Trim=" + strings.Trim(" ss s1  ", " ")) //等价于TrimSpace
	fmt.Println("Trim=" + strings.TrimSpace(" ss s1  "))

	ss := []string{"a", "b", "c"} //申明字符串数组
	fmt.Println("字符串以指定连接符连接", strings.Join(ss, "_"))
	fmt.Println("字符串以指定字符串分隔", strings.Split(s, "l"))
	fmt.Println("字符串检查指定字符串重复出现次数", strings.Count(s, "l"))
	fmt.Println("字符串开头检查", s, strings.HasPrefix(s, "Hell"))
	fmt.Println("字符串结束检查", strings.HasSuffix(s, "world"))
	fmt.Println("字符串比较", s, "\n", s2, "\n", strings.EqualFold(s, s2), s == s2)
	fmt.Println("字符串分割Fields", strings.Fields(s)) //以空格分割字符串返回字符串数组

}

func testArra() {
	//数组测试
	//静态数组
	animal := []string{"cat", "dog", "fish", "bird"}
	for i, v := range animal {
		fmt.Println("animal["+strconv.Itoa(i)+"]=", v)
	}

	//number := [2][2]int{{1, 2}, {2, 3}}
	number := [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
	fmt.Println(len(number), len(number[0]))
	for i, v1 := range number {
		fmt.Println("number=", i, v1)
		for j, v2 := range v1 {
			fmt.Println("number2=", j, v2)
		}
	}

	fmt.Println("number[1:2]=", number[1:2]) //跟操作字符串一样，前闭后开区间

	//动态数组，使用切片
	var slice1 []int //申明切片
	if slice1 == nil { //默认是nil值
		fmt.Println("slice1 is nil")
	}
	fmt.Println("append之前 slice1=", slice1, len(slice1), cap(slice1)) //空切片
	slice2 := append(slice1, 1, 2)                                    //                           //可同时添加多个元素
	fmt.Println("append之后 slice1=", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2=", slice2, len(slice2), cap(slice2))

	//使用make创建一个切片
	slice3 := make([]int, 1, 5) //类型，大小(默认插入0，1个就插入1个0)，容量
	//
	fmt.Println("make创建切片 slice3=", slice3, len(slice3), cap(slice3))
	copy(slice3, slice2) //拷贝slice2到slice3中，如果目标长度小于源，则剪裁
	fmt.Println("copy后的切片 slice3=", slice3, len(slice3), cap(slice3))
	slice3 = append(slice3, 10)
	fmt.Println("copy后的切片 slice3=", slice3, len(slice3), cap(slice3))
}

func testMap() {
	//Map 是一种无序的键值对的集合.
	//Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
	//不过，Map 是无序的，我们无法决定它的返回顺序，
	//这是因为 Map 是使用 hash 表来实现的。

	var map1 map[string]string //申明一个map
	fmt.Println("map1=", map1)
	if map1 == nil {
		fmt.Println("map1 is nil")
	}
	//map1["A"] = "优秀"//对nil赋值，崩溃
	fmt.Println("map1=", map1)
	map1 = make(map[string]string) //创建map，实例化
	map1["A"] = "优秀"               //添加元素
	fmt.Println("map1=", map1)
	map1["B"] = "良好" //添加元素
	fmt.Println("map1=", map1)

	for k, v := range map1 {
		fmt.Println("map1[", k, "]=", v)
	}

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
}

func main() {
	testStr()
	testArra()
	testMap()
}
