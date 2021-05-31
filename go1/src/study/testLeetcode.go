package main

import "fmt"

//
//import (
//	"fmt"
//	"sort"
//)
//
///*
//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//
//示例:
//给定 nums = [2, 7, 11, 15], target = 9
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//
//时间复杂度：O(n)O(n)
//空间复杂度：O(n)O(n)
//*/
//func twoSum(nums []int, target int) []int {
//	vMap := make(map[int]int)
//	for i := 0; i < len(nums); i++ {
//		if j, ok := vMap[target-nums[i]]; ok {
//			return []int{i, j}
//		}
//		vMap[nums[i]] = i
//	}
//	return []int{0, 0}
//}
//
///*
//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//*/
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	ret := new(ListNode)
//	cur := ret
//	flag := 0
//	for {
//		cur.Val = flag
//		flag = 0
//		if l1 != nil {
//			cur.Val = cur.Val + l1.Val
//			l1 = l1.Next
//		}
//		if l2 != nil {
//			cur.Val = cur.Val + l2.Val
//			l2 = l2.Next
//		}
//
//		if cur.Val > 9 {
//			flag = 1
//			cur.Val %= 10
//		}
//
//		if l1 == nil && l2 == nil && flag == 0 {
//			break
//		}
//
//		cur.Next = new(ListNode)
//		cur = cur.Next
//	}
//
//	return ret
//}
//
///**
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//*/
//func max(i, j int) int {
//	if i >= j {
//		return i
//	}
//	return j
//
//}
//
///*
//给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
//
//请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
//你可以假设 nums1 和 nums2 不会同时为空。
//*/
//func lengthOfLongestSubstring(s string) int {
//	if len(s) == 0 {
//		return 0
//	}
//
//	dict := make(map[byte]int)
//	ret := 0
//	len1 := len(s)
//	for i, j := 0, 0; j < len1; j++ {
//		if k, ok := dict[s[j]]; ok {
//			i = max(k+1, i) //防止i回退("abba")，i只能单向向右移动
//		}
//		ret = max(ret, j-i+1)
//		dict[s[j]] = j
//	}
//	return ret
//}
//
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//
//	var a = nums1
//	var b = nums2
//	var m = len(a)
//	var n = len(b)
//	if m > n {
//		a = nums2
//		b = nums1
//		m = len(a)
//		n = len(b)
//	}
//
//	min := 0
//	max := m
//	half := (m + n) / 2
//	for min <= max {
//		var i = (min + max) / 2
//		var j = half - i
//
//		if i > min && a[i-1] > b[j] {
//			max = i - 1
//		} else if i < max && b[j-1] > a[i] {
//			min = i + 1
//		} else {
//			//完美的情况
//			minRight := 0
//			if i == m {
//				minRight = b[j]
//			} else if j == n {
//				minRight = a[i]
//			} else {
//				minRight = a[i]
//				if minRight > b[j] {
//					minRight = b[j]
//				}
//			}
//			if (m+n)%2 == 1 {
//				return float64(minRight)
//			}
//
//			maxLeft := 0
//			if i == 0 {
//				maxLeft = b[j-1]
//			} else if j == 0 {
//				maxLeft = a[i-1]
//			} else {
//				maxLeft = a[i-1]
//				if maxLeft < b[j-1] {
//					maxLeft = b[j-1]
//				}
//			}
//
//			return float64(maxLeft+minRight) / 2.0
//		}
//	}
//	return 0.0
//}
//
///*
//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
//示例 1：
//
//输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
//示例 2：
//
//输入: "cbbd"
//输出: "bb"
//*/
//func findS(s string, half, min, max int) (int, int) {
//	leftLen := (half + 1) * 2
//	rightLen := (len(s) - half) * 2
//	remainLen := leftLen
//	if leftLen > rightLen {
//		remainLen = rightLen
//	}
//	if remainLen < max-min {
//		return 0, 0
//	}
//
//	//奇数检查
//	jsI := 0
//	jsJ := 0
//	i, j := half-1, half+1
//	for ; i >= 0 && j < len(s); {
//		if s[i] != s[j] {
//			break
//		}
//		i--
//		j++
//	}
//	jsI = i + 1
//	jsJ = j - 1
//	len1 := jsJ - jsI + 1
//
//	//偶数检查
//	osI1 := 0
//	osJ1 := 0
//	i, j = half-1, half
//	for ; i >= 0 && j < len(s);
//	{
//		if s[i] != s[j] {
//			break
//		}
//		i--
//		j++
//	}
//	osI1 = i + 1
//	osJ1 = j - 1
//	len2 := osJ1 - osI1 + 1
//
//	osI2 := 0
//	osJ2 := 0
//	i, j = half, half+1
//	for i, j := half, half+1; i >= 0 && j < len(s);
//	{
//		if s[i] != s[j] {
//			break
//		}
//		i--
//		j++
//	}
//	osI2 = i + 1
//	osJ2 = j - 1
//	len3 := osJ2 - osI2 + 1
//
//	if len1 > len2 && len1 > len3 {
//		return jsI, jsJ
//	} else if len2 > len1 && len2 > len3 {
//		return osI1, osJ1
//	} else if len3 > len1 && len3 > len2 {
//		return osI2, osJ2
//	}
//	return 0, 0
//}
//
//func longestPalindrome(s string) string {
//	if len(s) == 0 {
//		return ""
//	}
//
//	half := len(s) / 2
//	maxI, maxJ := 0, 0
//	for step := 0; step <= half; step++ {
//		i, j := findS(s, half+step, maxI, maxJ)
//		if maxJ-maxI < j-i {
//			maxI = i
//			maxJ = j
//		}
//
//		if step != 0 {
//			p, q := findS(s, half-step, maxI, maxJ)
//			if maxJ-maxI < q-p {
//				maxI = p
//				maxJ = q
//			}
//		}
//	}
//	return s[maxI : maxJ+1]
//}
//
//func convert(s string, numRows int) string {
//	lines := make([]string, numRows)
//	lenS := len(s)
//	step := 2*numRows - 2
//	for i := 0; i < len(s); i += step {
//		getCellLine(s, numRows, i, step, lenS, lines)
//	}
//
//	var ret string = ""
//	for i := 0; i < numRows; i++ {
//		ret += lines[i]
//	}
//	return ret
//}
//
//func getCellLine(s string, numRows, start, step, lenS int, lines []string) {
//	i := start
//	end := start + numRows
//	j := 0
//	for ; i < end; i++ {
//		if i >= lenS {
//			return
//		}
//		lines[j] += s[i : i+1]
//		j++
//	}
//
//	j--
//	end = start + step
//	for ; i < end; i++ {
//		if i >= lenS {
//			return
//		}
//		j--
//		lines[j] += s[i : i+1]
//	}
//}
//
///**
//给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
//
//'.' 匹配任意单个字符
//'*' 匹配零个或多个前面的那一个元素
//所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
//
//说明:
//
//s 可能为空，且只包含从 a-z 的小写字母。
//p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
//*/
//type Reg struct {
//	c   uint8;
//	cnt int; //0表示匹配0个或多个 1表示匹配一次
//	opt int; //0 标准匹配， 1任意字符，
//}
//
//func isMatch(s string, p string) bool {
//	if p == "" {
//		return s == p
//	} else if s == "" {
//		return false
//	}
//
//	regs := make([]Reg, 0)
//
//	for i := 0; i < len(p); i++ {
//		ch1 := p[i]
//		opt := 0
//		if ch1 == '.' {
//			opt = 1
//		}
//		if i+1 < len(p) {
//			ch2 := p[i+1]
//			if ch2 == '*' {
//				regs = append(regs, Reg{c: ch1, cnt: 0, opt: opt})
//				i++
//			} else {
//				regs = append(regs, Reg{c: ch1, cnt: 1, opt: opt})
//			}
//		} else {
//			regs = append(regs, Reg{c: ch1, cnt: 1, opt: opt})
//		}
//	}
//
//	for i, j := 0, 0; i < len(s); {
//		if j > len(regs) {
//			return false
//		}
//		reg := regs[j]
//		if reg.opt == 0 { //标准匹配
//			if reg.cnt == 0 {
//				if s[i] == reg.c {
//					i++
//				} else {
//					j++
//				}
//			} else {
//				if s[i] == reg.c {
//					i++
//					j++
//				} else {
//					return false
//				}
//			}
//		} else { //任意字符匹配
//			if reg.cnt == 0 {
//				i++
//			} else {
//				i++
//				j++
//			}
//		}
//	}
//	return true
//}
//
//func romanToInt(s string) int {
//	romans := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
//	romans2 := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
//
//	ret := 0
//	j := 0
//
//	for i := 0; i < len(romans); i++ {
//		n := 0
//		for ; j < len(s); {
//			l1 := len(romans2[i])
//			if s[j:j+l1] != romans2[i] {
//				break
//			}
//			n++
//			j += l1
//		}
//		ret += romans[i] * n
//	}
//
//	return ret
//}
//
//func threeSum(nums []int) [][]int {
//	vMap := make(map[[3]int]bool)
//	ret := make([][]int, 0)
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			for k := j + 1; k < len(nums); k++ {
//				if nums[i]+nums[j]+nums[k] == 0 {
//					arra := [3]int{nums[i], nums[j], nums[k]}
//					sort.Ints(arra[:])
//					if _, ok := vMap[arra]; !ok {
//						ret = append(ret, arra[:])
//						vMap[arra] = true
//					}
//				}
//			}
//		}
//	}
//	return ret
//}
//
//func fab(a int) int {
//	if a < 0 {
//		return -a
//	}
//	return a
//}
//
//func mint(a, b, t int) int {
//	if fab(a-t) > fab(b-t) {
//		return b
//	} else {
//		return a
//	}
//}
//func threeSumClosest(nums []int, target int) int {
//	sort.Ints(nums)
//
//	theRet := nums[0] + nums[1] + nums[2]
//	if theRet >= target {
//		return theRet
//	}
//
//	for i := 0; i < len(nums); i++ {
//		left := i + 1
//		right := len(nums) - 1
//
//		for left < right {
//			sum3 := nums[i] + nums[left] + nums[right]
//			if sum3 == target {
//				return sum3
//			} else {
//				theRet = mint(theRet, sum3, target)
//				if sum3 > target {
//					right--
//				} else {
//					left++
//				}
//			}
//		}
//	}
//	return theRet
//}
//
//func fourSum(nums []int, target int) [][]int {
//	sort.Ints(nums)
//
//	ret := make([][]int, 0)
//	for i := 0; i < len(nums)-3; i++ {
//		if i > 0 && nums[i] == nums[i-1] {
//			continue
//		}
//
//		for j := i + 1; j < len(nums)-2; j++ {
//			left := j + 1
//			right := len(nums) - 1
//
//			for left < right {
//				sum4 := nums[i] + nums[j] + nums[left] + nums[right]
//				if sum4 == target {
//					ret = append(ret, []int{nums[i], nums[j], nums[left], nums[right]})
//					for left+1 < right && nums[left+1] == nums[left] {
//						left++
//					}
//					for right-1 > left && nums[right-1] == nums[right] {
//						right--
//					}
//
//					left++
//					right--
//				} else if sum4 > target {
//					right--
//				} else {
//					left++
//				}
//			}
//		}
//	}
//	return ret
//}
//
///**
// * Definition for singly-linked list.
// * type ListNode struct {
// *     Val int
// *     Next *ListNode
// * }
// */
//func mergeKLists(lists []*ListNode) *ListNode {
//	if len(lists) == 0 {
//		return nil
//	}
//
//	//时间 O(NlgK)
//	//内存 O(1)
//	cnt := len(lists)
//	for cnt > 1 {
//		i, j := 0, cnt-1
//		for i <= j {
//			lists[i] = mergeTwoList(lists[i], lists[j])
//
//			i++
//			j--
//		}
//		cnt = i
//	}
//
//	return lists[0]
//}
//
//func mergeTwoList(a, b *ListNode) *ListNode {
//	if a == nil {
//		return b
//	} else if b == nil {
//		return a
//	} else if a == b {
//		return a
//	}
//
//	var head *ListNode = nil
//	var p *ListNode = nil
//	for a != nil && b != nil {
//		if a.Val > b.Val {
//			if head == nil {
//				head = b
//				p = head
//				b = b.Next
//			} else {
//				p.Next = b
//				p = p.Next
//				b = b.Next
//			}
//		} else {
//			if head == nil {
//				head = a
//				p = head
//				a = a.Next
//			} else {
//				p.Next = a
//				p = p.Next
//				a = a.Next
//			}
//		}
//	}
//
//	if a != nil {
//		p.Next = a
//	} else {
//		p.Next = b
//	}
//	return head
//}
//
///**
// * Definition for singly-linked list.
// * type ListNode struct {
// *     Val int
// *     Next *ListNode
// * }
// */
//func swapPairs(head *ListNode) *ListNode {
//	p := head
//
//	var ret *ListNode = nil
//	var q *ListNode = nil
//	var s *ListNode = nil
//	var t *ListNode = nil
//	for p != nil {
//		l := p
//		r := p.Next
//
//		if r != nil {
//			p = r.Next
//		}
//
//		s, t = swap(q, l, r)
//		if ret == nil {
//			ret = s
//			q = t
//		} else {
//			q = t
//		}
//	}
//	return ret
//}
//
///*
//返回 头结点，尾结点
//*/
//func swap(head, left, right *ListNode) (*ListNode, *ListNode) {
//	if left != nil && right != nil {
//		if head != nil {
//			head.Next = right
//			right.Next = left
//			left.Next = nil
//		} else {
//			head = right
//			right.Next = left
//			left.Next = nil
//		}
//	}
//
//	return head, left
//}
//
//func divide(dividend int, divisor int) int {
//	flag := true
//	if dividend < 0 {
//		flag = !flag
//		dividend = - dividend
//	}
//
//	if divisor < 0 {
//		flag = !flag
//		divisor = - divisor
//	}
//
//	if divisor == 1 {
//		limit := 1 << 31
//
//		if flag {
//			if dividend > limit-1 {
//				return limit - 1
//			} else {
//				return dividend
//			}
//		}
//
//		if dividend > limit {
//			return -limit
//		}
//		return -dividend
//	}
//
//	aa := 0
//	cur := 0
//	remain := dividend
//	for {
//		remain = remain - cur
//		cur = 0
//		n := 0
//
//		if remain < divisor {
//			break
//		}
//
//		preCur := 0
//		for {
//			preCur = cur
//			if cur == 0 {
//				cur = divisor
//			} else {
//				cur += cur
//			}
//
//			if cur <= remain {
//				n++
//			} else {
//				break
//			}
//			//n   1  2 3 4 5
//			//ret 1  2 4 8 16
//		}
//		cur = preCur
//		aa += 1 << uint(n-1)
//	}
//
//	if flag {
//		return aa
//	} else {
//		return -aa
//	}
//}
//
//func countCharacters(words []string, chars string) int {
//	dict := make(map[uint8]int)
//	for i := 0; i < len(chars); i++ {
//		dict[chars[i]] ++
//	}
//
//	ret := 0
//	for i := 0; i < len(words); i++ {
//		word := words[i]
//		flag := true
//		j := 0
//
//		tempDict := make(map[uint8]int)
//		for ; j < len(word); j++ {
//			tempDict[word[j]] ++
//		}
//		for k, v := range tempDict {
//			if cnt, ok := dict[k]; !ok || cnt < v {
//				flag = false
//				break
//			}
//		}
//
//		if flag {
//			ret += j
//		}
//	}
//	return ret
//}
//
//func getLeastNumbers(arr []int, k int) []int {
//	dealQuickSelect(arr, 0, len(arr)-1, k)
//	return arr[:k]
//}
//
//func dealQuickSelect(arra []int, l, r, k int) {
//	if l >= r {
//		return
//	}
//
//	start := l
//	end := r
//	key := arra[l]
//	for l < r {
//		for l < r && arra[r] >= key {
//			r--
//		}
//		if l >= r {
//			break
//		}
//		arra[l] = arra[r]
//		for l < r && arra[l] <= key {
//			l++
//		}
//		if l >= r {
//			break
//		}
//		arra[r] = arra[l]
//	}
//	arra[l] = key
//
//	if l == k {
//		return
//	} else if l > k {
//		dealQuickSelect(arra, start, l-1, k)
//	} else {
//		dealQuickSelect(arra, l+1, end, k)
//	}
//}
//
//func canMeasureWater(x int, y int, z int) bool {
//	if x > y {
//		x, y = y, x
//	}
//
//	if x == 0 {
//		return false
//	} else if z > x+y {
//		return false
//	} else if z == x+y || z == x || z == y || z == y-x {
//		return true
//	}
//
//	//把余水倒入大瓶，然后一直从小瓶到大瓶加水
//	if z%x == y%x {
//		return true
//	}
//
//	//把余水倒入小瓶，然后一直大瓶倒水到小瓶，清空小瓶
//	cur := y - x
//	if cur == 0 {
//		return false
//	}
//	if z%cur == 0 {
//		return true
//	}
//	if x%cur == z%cur {
//		return true
//	}
//	return false
//}
//
//func maxGCD(a, b int) int {
//	if a > b {
//		a, b = b, a
//	}
//	for {
//		c := b % a
//		if c == 0 {
//			return a
//		}
//		b, a = a, c
//	}
//}
//
//func testN() {
//	for i := -1; ; i-- {
//		if (1-i*103)%107 == 0 {
//			fmt.Printf("i=%d", i)
//		}
//	}
//}
//
//func minIncrementForUnique(A []int) int {
//	dict := make(map[int]bool)
//	arra := make([]int, 0)
//	for _, v := range A {
//		if _, ok := dict[v]; ok {
//			arra = append(arra, v)
//		}
//		dict[v] = true
//	}
//
//	ret := 0
//	cur := 0
//	sort.Ints(arra)
//	for _, num := range arra {
//		k := num
//		v := 1
//
//		for v > 0 {
//			i := k
//			if i < cur {
//				i = cur
//			}
//			for {
//				i++
//				if _, ok := dict[i]; !ok {
//					dict[i] = true
//					ret += i - k
//					cur = i
//					break
//				}
//			}
//			v--
//		}
//	}
//	return ret
//}
//
//func abs(x0, x1 int) int {
//	if x0 > x1 {
//		return x0 - x1
//	}
//	return x1 - x0
//}
//func distance(x0, y0, x1, y1 int) int {
//	return abs(x0, x1) + abs(y0, y1)
//}
//
//type Point struct {
//	x int
//	y int
//}
//
//func maxDistance(grid [][]int) int {
//
//	dict := make(map[*Point]bool)
//	for i := range grid {
//		for j := range grid[i] {
//			if grid[i][j] == 1 {
//				dict[&Point{i, j}] = true
//			}
//		}
//	}
//
//	totalMax := 0
//	for i := range grid {
//		var totalMin int
//		for j := range grid[i] {
//			if grid[i][j] == 0 {
//				totalMin = len(grid) * 2
//				for point := range dict {
//					cur := distance(point.x, point.y, i, j)
//					if totalMin > cur {
//						totalMin = cur
//					}
//				}
//
//				if totalMin > totalMax {
//					totalMax = totalMin
//				}
//			}
//		}
//	}
//	return totalMax
//}
//
//func lastRemaining(n int, m int) int { //4 101   100, 10
//	arra := make([]int, 0)
//	for i := 0; i < n; i++ {
//		arra = append(arra, i)
//	}
//
//	nLen := n
//	step := m % n
//	i := step - 1
//	for n > 1 {
//		if arra[i] == -1 {
//			i++
//			if i >= nLen {
//				i = 0
//			}
//		} else {
//			arra[i] = -1
//			n--
//
//			i += step
//			if i >= nLen {
//				i -= nLen
//			}
//		}
//	}
//
//	for i := 0; i < len(arra); i++ {
//		if arra[i] != -1 {
//			return arra[i]
//		}
//	}
//
//	return 0
//
//	//0 1 2 3 4 5 6 7 8 9 10
//	//0 1   3 4   6 7   9 10
//	//  1   3     6 7     10
//	//  1         6 7
//	//            6 7
//	//              7
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
//
//func trap(height []int) int {
//	ret := 0
//	i, j := 0, len(height)-1
//	lastH := 0
//	for i < j-1 {
//		for i < j-1 && (height[i+1] > height[i] || height[i] < lastH) {
//			i++
//		}
//
//		for i < j-1 && (height[j-1] > height[j] || height[j] < lastH) {
//			j--
//		}
//
//		if i == j-1 {
//			break
//		}
//
//		h := min(height[i], height[j])
//		hRemain := h - lastH
//		water := (j - i - 1) * hRemain
//		for k := i + 1; k < j; k++ {
//			water -= min(max(height[k], lastH)-lastH, hRemain)
//		}
//		ret += water
//		lastH = h
//
//		if height[i] > height[j] { //高度矮的往中间移动一下。。
//			j--
//		} else {
//			i++
//		}
//	}
//	return ret
//}
//
//type KeyCnt struct {
//	key  int
//	freq int
//	val  int
//
//	pre  *KeyCnt
//	next *KeyCnt
//}
//
//type List struct {
//	head *KeyCnt
//	tail *KeyCnt
//	size int
//}
//
//func New() *List {
//	head := &KeyCnt{}
//	tail := &KeyCnt{pre: head}
//	head.next = tail
//	return &List{head, tail, 0}
//}
//
//func (l *List) Shift() *KeyCnt {
//	node := l.head.next
//	l.head.next = node.next
//	node.next.pre = node.pre
//	l.size--
//	return node
//}
//
//func (l *List) Push(node *KeyCnt) {
//	l.tail.pre.next = node
//	node.pre = l.tail.pre
//	l.tail.pre = node
//	node.next = l.tail
//	l.size++
//}
//
//func (l *List) Erase(node *KeyCnt) {
//	node.pre.next = node.next
//	node.next.pre = node.pre
//	l.size--
//}
//
//type LFUCache struct {
//	dictKeyVal  map[int]*KeyCnt
//	dictFreqKey map[int]*List
//	minFreq     int
//	capacity    int
//}
//
//func Constructor(capacity int) LFUCache {
//	cache := LFUCache{}
//	cache.dictKeyVal = make(map[int]*KeyCnt)
//	cache.dictFreqKey = make(map[int]*List)
//	cache.capacity = capacity
//	cache.minFreq = 0
//	return cache
//}
//
//func (this *LFUCache) addFreq(v *KeyCnt) {
//	list, _ := this.dictFreqKey[v.freq]
//	list.Erase(v)
//
//	if list.size == 0 {
//		delete(this.dictFreqKey, v.freq)
//		if this.minFreq == v.freq {
//			this.minFreq++
//		}
//	}
//
//	v.freq++
//	list2, ok := this.dictFreqKey[v.freq]
//	if !ok {
//		list2 = New()
//	}
//	list2.Push(v)
//	this.dictFreqKey[v.freq] = list2 //更新
//}
//
//func (this *LFUCache) Get(key int) int {
//	v, ok := this.dictKeyVal[key]
//	if !ok {
//		return -1
//	}
//	this.addFreq(v)
//	return v.val
//}
//
//func (this *LFUCache) Put(key int, value int) {
//	if this.capacity == 0 {
//		return
//	}
//
//	v, ok := this.dictKeyVal[key]
//	if ok { //如果存在
//		v.val = value
//		this.addFreq(v)
//		return
//	}
//
//	if len(this.dictKeyVal) == this.capacity {
//		list := this.dictFreqKey[this.minFreq]
//		vExpire := list.Shift()
//		delete(this.dictKeyVal, vExpire.key)
//		if list.size == 0 {
//			delete(this.dictFreqKey, vExpire.freq)
//		}
//	}
//
//	this.minFreq = 0
//	list, ok := this.dictFreqKey[0]
//	if !ok {
//		list = New()
//	}
//	v = &KeyCnt{key: key, val: value}
//	list.Push(v)
//	this.dictKeyVal[key] = v
//	this.dictFreqKey[v.freq] = list //更新
//}
//
///**
// * Your LFUCache object will be instantiated and called as such:
// * obj := Constructor(capacity);
// * param_1 := obj.Get(key);
// * obj.Put(key,value);
// */
//
////horse  ros
////park   spake
////intention execution
//var dictCounter = make(map[string]int)
//
//func countMin(word1 string, word2 string) int {
//	if word1 == word2 {
//		return 0
//	} else if word1 == "" {
//		return len(word2)
//	} else if word2 == "" {
//		return len(word1)
//	}
//
//	if v, ok := dictCounter[word1+"-"+word2]; ok {
//		return v
//	}
//
//	if word1[0] == word2[0] {
//		return countMin(word1[1:], word2[1:])
//	}
//
//	//intention execution
//	a := 1 + countMin(word1[1:], word2)     //ntention execution
//	b := 1 + countMin(word1[1:], word2[1:]) //ntention xecution
//	c := 1 + countMin(word1, word2[1:])     //intention xecution
//	//取最大值
//	if a < b {
//		if a < c {
//			dictCounter[word1+"-"+word2] = a
//			return a
//		} else {
//			dictCounter[word1+"-"+word2] = c
//			return c
//		}
//	} else if b < c {
//		dictCounter[word1+"-"+word2] = b
//		return b
//	} else {
//		dictCounter[word1+"-"+word2] = c
//		return c
//	}
//}
//
//func minDistance(word1 string, word2 string) int {
//	return countMin(word1, word2)
//}
//
//func rotate(matrix [][]int) {
//	l := len(matrix)
//	p, q := 0, l-1
//	for i := l; i > 1; i = i - 2 { //i表示需要change几次外围
//		for j := p; j < q; j++ { //p表示从几号索引开始，q表示到几号索引结束
//			// j=0,p=0,q=3 第一轮
//			// j=1,p=1,q=2 第二轮
//			matrix[p][j], matrix[j][q], matrix[q][l-1-j], matrix[l-1-j][p] =
//				matrix[l-1-j][p], matrix[p][j], matrix[j][q], matrix[q][l-1-j]
//		}
//		p++
//		q--
//	}
//}
//
//func count(v int) int {
//	ret := 0
//	for v > 0 {
//		ret += v % 10
//		v /= 10
//	}
//	return ret
//}
//
//func movingCount(m int, n int, k int) int {
//	arra := make([][]int, 0)
//	cur := 1 //第一排肯定能移动到
//	for i := 0; i < m; i++ {
//		row := make([]int, n)
//		if cur > 0 { //如果上面一排没有位子可以移动，那么后面就不用计算了。
//			cur = 0 //清0，重新计算这一排可移动的位子数量
//			for j := 0; j < n; j++ {
//				if count(i)+count(j) <= k {
//					row[j] = 2 //置为2，表示可以达到，但还不确定是否能达到0,0
//					cur++
//				}
//			}
//		}
//		arra = append(arra, row)
//	}
//
//	ret3 := 0
//	cur = 1
//	for i := 0; i < m; i++ {
//		if cur > 0 {
//			cur = 0
//			for j := 0; j < n; j++ {
//				if arra[i][j] == 2 && (i == 0 && j == 0 || i == 0 && arra[0][j-1] == 1 || j == 0 && arra[i-1][0] == 1 ||
//					i > 0 && j > 0 && (arra[i-1][j] == 1 || arra[i][j-1] == 1)) {
//					arra[i][j] = 1 //置为1，表示确实可以达到
//					cur++
//					ret3++
//				}
//			}
//		}
//	}
//	return ret3
//}
//
//f(x)=kx+c
//y=kx+c
//func intersection(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {
//	//求斜率
//	Line1Flag := false //是否平行Y轴
//	k1, c1 := 0.0, 0.0
//	if end1[0]-start1[0] != 0 {
//		k1 = float64(end1[1]-start1[1]) / float64(end1[0]-start1[0])
//		c1 = float64(start1[1]) - k1*float64(start1[0])
//	} else {
//		Line1Flag = true
//	}
//
//	Line2Flag := false //是否平行Y轴
//	k2, c2 := 0.0, 0.0
//	if end2[0]-start2[0] != 0 {
//		k2 = float64(end2[1]-start2[1]) / float64(end2[0]-start2[0])
//		c2 = float64(start2[1]) - k2*float64(start2[0])
//	} else {
//		Line2Flag = true
//	}
//
//	line1YMin := min(start1[1], end1[1])
//	line1YMax := max(start1[1], end1[1])
//	line2YMin := min(start2[1], end2[1])
//	line2YMax := max(start2[1], end2[1])
//
//	//都平行于y轴的线
//	if Line1Flag && Line2Flag {
//		if start1[0] != start2[0] {
//			return nil
//		}
//
//		bottomY := max(line1YMin, line2YMin)
//		topY := min(line1YMax, line2YMax)
//
//		if bottomY > topY {
//			return nil
//		}
//
//		return []float64{float64(start1[0]), float64(bottomY)}
//	} else if Line1Flag { //线段1平行Y轴
//		y2 := k2*float64(start1[0]) + c2
//		if y2 < float64(line2YMin) || y2 > float64(line2YMax) {
//			return nil
//		}
//		return []float64{float64(start1[0]), y2}
//	} else if Line2Flag { //线段2平行Y轴
//		y1 := k1*float64(start2[0]) + c1
//		if y1 < float64(line1YMin) || y1 > float64(line1YMax) {
//			return nil
//		}
//		return []float64{float64(start2[0]), y1}
//	}
//
//	line1XMin := min(start1[0], end1[0])
//	line1XMax := max(start1[0], end1[0])
//	line2XMin := min(start2[0], end2[0])
//	line2XMax := max(start2[0], end2[0])
//
//	if k1 == k2 { //可能多个焦点或者平行，c语言需要增加浮点数判断。
//		if c1 != c2 { //平行线
//			return nil
//		}
//
//		leftX := max(line1XMin, line2XMin)
//		rightX := min(line1XMax, line2XMax)
//
//		if leftX > rightX {
//			return nil
//		}
//
//		v := float64(leftX)
//		return []float64{v, k1*v + c1}
//	} else { //检查相交点是否在线段上
//		//k1*x+c1 = k2*x+c2
//		x := float64(c1-c2) / float64(k2-k1)
//
//		if x < float64(line1XMin) || x > float64(line1XMax) { //是否在第一条线段内
//			return nil
//		}
//		if x < float64(line2XMin) || x > float64(line2XMax) { //是否在第二条线段内
//			return nil
//		}
//
//		return []float64{x, k1*x + c1}
//	}
//}

func superEggDrop(K int, N int) int {
	dict := make(map[int]int)
	return throw(&dict, K, N)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func throw(dict *map[int]int, k, n int) int {
	if v, ok := (*dict)[n*100+k]; ok {
		return v
	}

	if k == 1 {
		(*dict)[n*100+k] = n
		return n
	}

	if n == 1 {
		(*dict)[n*100+k] = 1
		return 1
	} else if n == 0 {
		(*dict)[n*100+k] = 0
		return 0
	}

	i, j := 1, n
	for i+1 < j {
		f := (i + j) / 2
		a := throw(dict, k, n-f)   //鸡蛋没有坏
		b := throw(dict, k-1, f-1) //鸡蛋坏了

		//一开始f选的楼层越小，a就越大，b反而很小
		//如果一开始f选的楼层越大，a就越小，b反而很大
		//最完美的值应该是介于a和b中间， a<=x<=b a=x=b的时候最佳
		if a > b {
			i = f
		} else if b > a {
			j = f
		} else {
			i, j = f, f
		}
	}
	ret := 1 + min(max(throw(dict, k, n-i), throw(dict, k-1, i-1)), max(throw(dict, k, n-j), throw(dict, k-1, j-1)))
	(*dict)[n*100+k] = ret
	return ret
}

func main() {
	//s := "adam"
	//fmt.Printf("s = %v,%v\n", s[1:2], s[0])
	//fmt.Printf("len = %v\n", lengthOfLongestSubstring("abcabcbb"))
	//fmt.Printf("findMedianSortedArrays=%v\n", findMedianSortedArrays([]int{1, 2}, []int{3, 4}))

	//fmt.Printf("s=%v\n",longestPalindrome(s))
	//fmt.Printf("s=%v\n", convert("PAYPALISHIRING", 3))
	//fmt.Printf("s=%v\n", romanToInt("MCMXCIV"))
	//fmt.Printf("%v\n", threeSumClosest([]int{6, -18, -20, -7, -15, 9, 18, 10, 1, -20, -17, -19, -3, -5, -19, 10,
	//	6, -11, 1, -17, -15, 6, 17, -18, -3, 16, 19, -20, -3, -17, -15, -3, 12, 1, -9, 4, 1, 12, -2, 14, 4, -4, 19, -20,
	//	6, 0, -19, 18, 14, 1, -15, -5, 14, 12, -4, 0, -10, 6, 6, -6, 20, -8, -6, 5, 0, 3, 10, 7, -2, 17, 20, 12, 19, -13,
	//	-1, 10, -1, 14, 0, 7, -3, 10, 14, 14, 11, 0, -4, -15, -8, 3, 2, -5, 9, 10, 16, -4, -3, -9, -8, -14, 10, 6, 2, -12,
	//	-7, -16, -6, 10}, -52))

	/**
	[-4, -1, -1, 0, 1, 2]
	-1

	[-5,-5,-3,-1,0,2,4,5]
	-7
	*/
	//fmt.Printf("%v\n", fourSum([]int{-4, -1, -1, 0, 1, 2}, -1))
	//[[1,4,5],[1,3,4],[2,6]]
	//a := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	//b := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	//c := &ListNode{2, &ListNode{6, nil}}
	//fmt.Printf("%v\n", mergeKLists([]*ListNode{a, b, c}))
	//fmt.Printf("%v\n",swapPairs(a))
	//fmt.Printf("%v\n", divide(10, 3))
	//fmt.Printf("%v\n", countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
	//fmt.Printf("%v\n", getLeastNumbers([]int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}, 8))
	//canMeasureWater(104579, 104593, 12444)
	//maxGCD(104579, 104593)
	//testN()
	//minIncrementForUnique([]int{2,2,2,1})
	//fmt.Printf("result=%v", twoSum([]int{2, 7, 11, 15}, 9))
	//maxDistance([][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}})
	//fmt.Println(lastRemaining(5, 3))
	//trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	//minDistance("horse", "ros")
	//fmt.Println(minDistance("intention", "execution"))
	//rotate([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}})
	//fmt.Println(movingCount(38, 15, 9))
	//fmt.Println(movingCount(16, 8, 4))
	//fmt.Println(movingCount(3, 1, 11))
	fmt.Println(superEggDrop(4, 5000))

	//fmt.Println(intersection([]int{-10, 48}, []int{-43, 46}, []int{-16, 59}, []int{-1, 85}))
}
