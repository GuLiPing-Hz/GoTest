package main

import (
	"bytes"
	"container/list"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type tweet struct {
	tId int //推文ID
	seq int //全局的序列ID
}

type Node struct {
	data *tweet
	pre  *Node
	next *Node
}

type ListTweet struct {
	head *Node
	tail *Node
	size int
}

func NewList() *ListTweet {
	head := &Node{nil, nil, nil}
	tail := &Node{nil, head, nil}
	head.next = tail
	return &ListTweet{head, tail, 0}
}

func (l *ListTweet) pushBack(d *tweet) {
	node := &Node{data: d}
	node.pre, node.next = l.tail.pre, l.tail
	l.tail.pre.next, l.tail.pre = node, node
	l.size++
}

func (l *ListTweet) popFront() {
	if l.size <= 0 {
		return
	}

	node := l.head.next
	node.pre.next, node.next.pre = node.next, node.pre
	l.size--
}

func (l *ListTweet) insert(pos *Node, d *tweet) {
	node := &Node{data: d}
	node.pre, node.next = pos, pos.next
	pos.next, pos.next.pre = node, node
	l.size++
}

type user struct {
	uid    int          //用户ID
	fuids  map[int]bool //关注列表
	tweets *ListTweet   //推文列表 只有自己的推文
}

type Twitter struct {
	seq      int           //全局唯一ID
	userDict map[int]*user //map[uid]用户的推文
}

/** Initialize your data structure here. */
func ConstructorT() Twitter {
	return Twitter{seq: 1, userDict: make(map[int]*user)}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	u, ok := this.userDict[userId]
	if !ok {
		u = &user{userId, make(map[int]bool), NewList()}
	}

	d := &tweet{tweetId, this.seq}
	u.tweets.pushBack(d)
	if u.tweets.size > 10 {
		u.tweets.popFront()
	}
	this.userDict[userId] = u
	this.seq++
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	u, ok := this.userDict[userId]
	if !ok {
		return nil
	}

	arra := make([]int, 0)
	ps := make([]*Node, 0)
	for k := range u.fuids {
		if v, ok := this.userDict[k]; ok {
			ps = append(ps, v.tweets.tail.pre)
		}
	}
	ps = append(ps, u.tweets.tail.pre)

	for {
		var p *Node
		j := 0
		for i := 0; i < len(ps); i++ {
			if ps[i].data != nil {
				if p == nil {
					p = ps[i]
					j = i
				} else if ps[i].data.seq > p.data.seq {
					p = ps[i]
					j = i
				}
			}
		}

		if p == nil {
			break
		}
		arra = append(arra, p.data.tId)
		if len(arra) >= 10 {
			break
		}
		ps[j] = p.pre
	}
	return arra
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}

	u, ok := this.userDict[followerId]
	if !ok {
		u = &user{followerId, make(map[int]bool), NewList()}
	}
	u.fuids[followeeId] = true
	this.userDict[followerId] = u
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	u, ok := this.userDict[followerId]
	if !ok {
		return
	}
	delete(u.fuids, followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func testArra(arra []int) {
	arra[0] = 1
}

//aaca aaca aaca
//acaa
func count(s1, s2 string, start, limit int) (int, int, int) {
	dictPos := make(map[int]bool)
	cnt1, cnt2, pos := 0, 0, 0
	for j := 0; j < len(s2); {
		pos = 0
		for i := start; i < len(s1); i++ {
			if s2[j] == s1[i] {
				j++
				if j == len(s2) {
					cnt2++
					pos = i + 1
					if pos == len(s1) {
						pos = 0
						break
					}
					j = 0
				}
			}
		}
		cnt1++

		if start == 0 {
			if limit > 0 {
				limit--
			}
			if limit == 0 || j == 0 {
				break
			}

			if _, ok := dictPos[pos]; ok {
				break
			}
			if pos != 0 {
				dictPos[pos] = true
			}
		} else {
			start = 0
		}
	}
	return cnt1, cnt2, pos
}

// aaa aaa aaa aaa
// aa
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	cnt1, cnt2, start := count(s1, s2, 0, -1)
	if cnt1 == 0 {
		return 0
	}

	remain := n1 % cnt1
	ret := n1 / cnt1 * cnt2 / n2
	if remain > 0 {
		cnt1, cnt2, _ = count(s1, s2, start, remain)
		ret += cnt2 / n2
	}

	return ret
}

type Value struct {
	left int
	flag bool
}

func numberOfSubarrays(nums []int, k int) int {
	arra := make([]Value, 0)
	for i := 0; i < len(nums); {
		cnt := 0
		for i < len(nums) && nums[i]%2 == 0 {
			i++
			cnt++
		}

		if i >= len(nums) {
			if cnt > 0 {
				arra = append(arra, Value{cnt, false})
			}
			break
		}
		arra = append(arra, Value{cnt, true})
		i++
	}

	//3 | 2 | 3-
	ret := 0
	for i := 0; i+k <= len(arra); {
		if arra[i].flag {
			left := arra[i].left + 1
			right := 1
			if i+k < len(arra) {
				right += arra[i+k].left
			} else if !arra[i+k-1].flag {
				break
			}
			ret += left * right
		}
		i++
	}
	return ret
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	arra := make([]int, 0)
	deal(&arra, root, 0)
	return arra
}

func deal(arra *[]int, root *TreeNode, i int) {
	if root == nil {
		return
	}

	if len(*arra) <= i {
		*arra = append(*arra, root.Val)
	}

	deal(arra, root.Right, i+1) //优先遍历右节点
	deal(arra, root.Left, i+1)  //然后查询左节点
}

func waysToChange(n int) int {
	n5 := n / 5
	ret := n5 + 1 //只有5
	//只有10和5
	ret += ((n5 % 2) + 1 + n5 - 2 + 1) * (n5 / 2) / 2
	//只有25和5
	ret += ((n5 % 5) + 1 + n5 - 5 + 1) * (n / 5) / 2
	//25和10
	if n >= 35 {

	}
	return ret % 1000000007
}

func reversePairs(nums []int) int {
	n := len(nums)
	tmp := make([]int, n)
	copy(tmp, nums)
	sort.Ints(tmp)

	for i := 0; i < n; i++ {
		nums[i] = sort.SearchInts(tmp, nums[i]) + 1
	}

	bit := BIT{
		n:    n,
		tree: make([]int, n+1),
	}

	ans := 0
	for i := n - 1; i >= 0; i-- {
		ans += bit.query(nums[i] - 1)
		bit.update(nums[i])
	}
	return ans
}

type BIT struct {
	n    int
	tree []int
}

func (b BIT) lowbit(x int) int { return x & (-x) }

func (b BIT) query(x int) int {
	ret := 0
	for x > 0 {
		ret += b.tree[x]
		x -= b.lowbit(x)
	}
	return ret
}

func (b BIT) update(x int) {
	for x <= b.n {
		b.tree[x]++
		x += b.lowbit(x)
	}
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	temp2 := make([][]int, 0)
	temp2 = append(temp2, []int{nums[0]})
	for i := 1; i < len(nums); i++ {
		temp := temp2
		temp2 = make([][]int, 0)
		for j := 0; j <= i; j++ {
			for k := 0; k < len(temp); k++ {
				temp3 := make([]int, 0)
				temp3 = append(temp3, temp[k][0:j]...)
				temp3 = append(temp3, nums[i])
				temp3 = append(temp3, temp[k][j:]...)
				temp2 = append(temp2, temp3)
			}
		}
	}
	return temp2
}

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return nil
	}

	ret := make([]int, 0)
	size := len(words[0])
	num := len(words)

	dict := make(map[string]int)
	for i := range words {
		dict[words[i]]++
	}

	allsize := num * size
	for i := 0; i+allsize <= len(s); i++ {
		j := i
		cur := 0
		tempDict := make(map[string]int)

		for cur < num && j+allsize-cur*size <= len(s) {
			word := s[j : j+size]
			if _, ok := dict[word]; !ok {
				break
			}
			tempDict[word]++
			if tempDict[word] > dict[word] {
				break
			}

			cur++
			j += size
		}
		if cur == num {
			ret = append(ret, i)
		}
	}
	return ret
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := nums[0], nums[len(nums)-1]
	if target == l {
		return 0
	} else if target == r {
		return len(nums) - 1
	}

	if len(nums) == 1 {
		return -1
	}

	mid := len(nums) / 2
	v := nums[mid]
	if target == v {
		return mid
	} else if target > l && target < v {
		return search(nums[:mid], target)
	} else if target > v && target < r {
		ret := search(nums[mid:], target)
		if ret == -1 {
			return -1
		}
		return mid + ret
	} else if target > l && target > v {
		if v > l {
			ret := search(nums[mid+1:], target)
			if ret == -1 {
				return -1
			}
			return mid + 1 + ret
		} else {
			return search(nums[:mid], target)
		}
	} else if target < v && target < r {
		if v > l {
			ret := search(nums[mid+1:], target)
			if ret == -1 {
				return -1
			}
			return mid + 1 + ret
		} else {
			return search(nums[:mid], target)
		}
	}
	return -1
}

//完美诠释位运算的速度。比hashmap都快！
func singleNumbers(nums []int) []int {
	if len(nums) == 2 {
		return nums
	}

	ret := 0
	for i := range nums {
		ret ^= nums[i]
	}

	lowbit := ret & (-ret) //取到a和b在这一位bit不同的点
	a, b := 0, 0
	for i := range nums {
		if (lowbit & nums[i]) == 0 {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}
	return []int{a, b}
}

// This is the MountainArray's API interface.
// You should not implement it, or speculate about its implementation
type MountainArray struct {
	arra []int
}

func (this *MountainArray) get(index int) int { return this.arra[index] }
func (this *MountainArray) length() int       { return len(this.arra) }

func findInMountainArray(target int, mountainArr *MountainArray) int {
	i, j := 0, mountainArr.length()-1
	for i < j {
		mid := (i + j) / 2
		if mountainArr.get(mid) < mountainArr.get(mid+1) {
			i = mid + 1
		} else {
			j = mid
		}
	}

	ret := dealMoutain(target, 0, i, mountainArr, 1)
	if ret != -1 {
		return ret
	}
	return dealMoutain(-target, i, mountainArr.length()-1, mountainArr, -1)
}

func dealMoutain(target, i, j int, mountainArr *MountainArray, flag int) int {
	for i <= j {
		mid := (i + j) / 2
		midV := mountainArr.get(mid) * flag

		if midV > target {
			j = mid - 1
		} else if midV < target {
			i = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

var dict = make(map[int]bool)

func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	if _, ok := dict[n]; ok {
		return false
	}

	dict[n] = true
	ret := 0
	for n > 0 {
		v := n % 10
		ret += v * v
		n = n / 10
	}

	return isHappy(ret)
}

func solveSudoku(board [][]byte) {
	dict := make([][][]bool, 3)
	dict[0] = make([][]bool, 9)
	dict[1] = make([][]bool, 9)
	dict[2] = make([][]bool, 9)
	for i := 0; i < 9; i++ {
		dict[0][i] = make([]bool, 9)
		dict[1][i] = make([]bool, 9)
		dict[2][i] = make([]bool, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				k := i/3*3 + j/3
				index := int(board[i][j] - '1')
				dict[0][i][index] = true
				dict[1][j][index] = true
				dict[2][k][index] = true
			}
		}
	}

	dealSudo(board, 0, 0, dict)
}

func dealSudo(board [][]byte, row, col int, dict [][][]bool) bool {
	if col == 9 {
		col = 0
		row++
	}
	if row == 9 {
		return true
	}

	if board[row][col] != '.' {
		return dealSudo(board, row, col+1, dict)
	} else {
		for i := 0; i < 9; i++ {
			boxIndex := row/3*3 + col/3
			if dict[0][row][i] || dict[1][col][i] || dict[2][boxIndex][i] {
				continue
			}

			board[row][col] = byte('1' + i)
			dict[0][row][i] = true
			dict[1][col][i] = true
			dict[2][boxIndex][i] = true

			if dealSudo(board, row, col+1, dict) {
				return true
			}

			board[row][col] = '.'
			dict[0][row][i] = false
			dict[1][col][i] = false
			dict[2][boxIndex][i] = false
		}
	}
	return false
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	a := countAndSay(n - 1)
	num, cnt := a[0], 1
	ret := bytes.Buffer{}
	for i := 1; i < len(a); i++ {
		if a[i] == num {
			cnt++
		} else {
			ret.WriteString(strconv.Itoa(cnt))
			ret.WriteByte(num)
			num, cnt = a[i], 1
		}
	}
	ret.WriteString(strconv.Itoa(cnt))
	ret.WriteByte(num)
	return ret.String()
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	ret := make([][]int, 0)
	dealCombinationSum(&ret, candidates, target, make([]int, 0))
	return ret
}

func dealCombinationSum(r *[][]int, candidates []int, target int, temp []int) {
	if target == 0 {
		*r = append(*r, temp)
		return
	} else if len(candidates) == 0 {
		return
	} else if candidates[0] > target {
		return
	}

	dealCombinationSum(r, candidates[1:], target, temp)
	temp2 := make([]int, len(temp))
	copy(temp2, temp)
	temp2 = append(temp2, candidates[0])
	dealCombinationSum(r, candidates, target-candidates[0], temp2)
}

//困难-给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。  在不用排序的情况下，使用负号来标记某个索引代表的正整数已被使用，n使用下标0代替。
func firstMissingPositive(nums []int) int {
	oneExist, n := false, len(nums)
	for i := range nums {
		if nums[i] == 1 {
			oneExist = true
		} else if nums[i] <= 0 || nums[i] > n {
			nums[i] = 1
		}
	}

	if !oneExist {
		return 1
	}

	for i := range nums {
		v := abs(nums[i])
		if v == n {
			if nums[0] > 0 {
				nums[0] = -nums[0]
			}
		} else if nums[v] > 0 {
			nums[v] = -nums[v]
		}
	}

	for i := 1; i < n; i++ {
		if nums[i] > 0 {
			return i
		}
	}
	if nums[0] > 0 {
		return n
	}
	return n + 1

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dealMultiply(ret *[]int, v, pos int) {
	more := v / 10
	remain := v % 10
	if len(*ret) <= pos {
		*ret = append(*ret, remain)
		if more > 0 {
			*ret = append(*ret, more)
		}
	} else {
		remain1 := (*ret)[pos] + remain
		(*ret)[pos] = remain1 % 10

		more1 := more + (remain1 / 10)
		if more1 > 0 {
			dealMultiply(ret, more1, pos+1)
		}
	}
}

func multiply(num1 string, num2 string) string {
	arra := make([]int, 0)
	for i, p := len(num1)-1, 0; i >= 0; i-- {
		for j, q := len(num2)-1, 0; j >= 0; j-- {
			temp := int(num1[i]-'0') * int(num2[j]-'0')
			dealMultiply(&arra, temp, p+q)
			q++
		}
		p++
	}

	ret := ""
	for i := len(arra) - 1; i >= 0; i-- {
		ret += strconv.Itoa(arra[i])
	}
	return ret
}

func mincostTickets(days []int, costs []int) int {
	dict := make([]int, len(days))
	cfg := []int{1, 7, 30}
	return dpMincostTickets(0, days, costs, dict, cfg)
}

func dpMincostTickets(i int, days []int, costs, dict, cfg []int) int {
	if i >= len(days) {
		return 0 //花费0元
	}

	if dict[i] != 0 {
		return dict[i]
	}

	p := -1
	for j := range cfg {
		k := i
		for k < len(days) && days[k]-days[i] < cfg[j] {
			k++
		}
		temp := dpMincostTickets(k, days, costs, dict, cfg) + costs[j]
		if p == -1 || p > temp {
			p = temp
		}
	}

	dict[i] = p
	return p
}

//13*2+3*6=26+18
//1,4,6, 9,10,11,12,13,14,15, 16,17,18,20,21,22 ,23,27,28
func mySqrt(x int) int {
	//牛顿迭代法
	//f(x) = x^2-C
	//f'(x) = 2x
	//切线与 x轴   0-[(x1)^2-c]=2(x1)(x-x1) =>
	//               -(x1)+c/(x1) =2x-2(x1)
	//                x1-c/x1 = 2x
	//                0.5(x1+c/x1)=x

	c := float64(x)
	xi := float64(x)
	var xj float64
	for {
		xj = 0.5 * (xi + c/xi)
		if math.Abs(xi-xj) < 1e-7 {
			break
		}
		xi = xj
	}
	return int(xj)
}

type MinStack struct {
	stack []int
	val   []int
}

/** initialize your data structure here. */
func Constructor1() MinStack {
	ms := MinStack{
		stack: make([]int, 0),
		val:   make([]int, 0),
	}
	return ms
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)

	i, j := 0, len(this.val)-1
	for i <= j {
		mid := (i + j) / 2
		if this.val[mid] > x {
			j = mid - 1
		} else if this.val[mid] < x {
			i = mid + 1
		} else {
			i = mid
			break
		}
	}

	temp := make([]int, i)
	copy(temp, this.val[:i])
	temp = append(temp, x)
	this.val = append(temp, this.val[i:]...)
}

func (this *MinStack) Pop() {
	x := this.Top()
	this.stack = this.stack[:len(this.stack)-1]

	mid := 0
	i, j := 0, len(this.val)-1
	for i <= j {
		mid = (i + j) / 2
		if this.val[mid] > x {
			j = mid - 1
		} else if this.val[mid] < x {
			i = mid + 1
		} else {
			break
		}
	}

	temp := this.val[:mid]
	this.val = append(temp, this.val[mid+1:]...)
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.val[0]
}

type BIT2 struct {
	arra []int
}

func lowbit(x int) int {
	return x & -x
}

func (b *BIT2) update(v, i int) {
	for i < len(b.arra) {
		b.arra[i] += v
		i += lowbit(i)
	}
}
func (b *BIT2) get(i int) int {
	var ret = 0
	for i > 0 {
		ret += b.arra[i]
		i -= lowbit(i)
	}
	return ret
}

func subarraySum(nums []int, k int) int {
	bit := BIT2{
		arra: make([]int, len(nums)+1),
	}
	for i := range nums {
		bit.update(nums[i], i+1)
	}

	ret := 0
	for i := 1; i <= len(nums); i++ {
		if bit.get(i) == k {
			ret++
		}
		for j := i + 1; j <= len(nums); j++ {
			if bit.get(j)-bit.get(i) == k {
				ret++
			}
		}
	}
	return ret
}

type Course struct {
	pres []int
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	dict := make(map[int]*Course)

	for i := range prerequisites {
		c := prerequisites[i]
		if v, ok := dict[c[0]]; ok {
			v.pres = append(v.pres, c[1])
		} else {
			dict[c[0]] = &Course{
				pres: []int{c[1]},
			}
		}
	}

	ret := make([]int, 0)
	learned := make(map[int]bool)
	for i := 0; i < numCourses; i++ {
		if _, ok := dict[i]; !ok {
			ret = append(ret, i)
			learned[i] = true
		}
	}

	for k := range dict {
		if learned[k] {
			continue
		}

		course := dict[k]
		for i := range course.pres {
			circle := make(map[int]bool)
			circle[k] = true
			if !dealFind(course.pres[i], dict, circle, learned, &ret) {
				return []int{}
			}
		}
		if !learned[k] {
			ret = append(ret, k)
			learned[k] = true
		}
		if len(ret) == numCourses {
			break
		}
	}
	return ret
}

func dealFind(c int, dict map[int]*Course, circle, learned map[int]bool, ret *[]int) bool {
	if learned[c] {
		return true
	}
	if circle[c] {
		return false
	}

	course := dict[c]
	circle[c] = true
	for i := range course.pres {
		if !dealFind(course.pres[i], dict, circle, learned, ret) {
			return false
		}
	}
	learned[c] = true
	*ret = append(*ret, c)
	return true
}

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			if dealValidPalindrome(s[i:j]) || dealValidPalindrome(s[i+1:j+1]) {
				return true
			} else {
				return false
			}

		}
	}

	return true
}

func dealValidPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	} else if len(s) == 2 {
		return s[0] == s[1]
	}
	i, j := 0, len(s)-1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := TreeNode{
		Val: preorder[0],
	}
	if len(preorder) == 1 {
		return &root
	}

	for i := range inorder {
		if inorder[i] == preorder[0] {
			var temp = preorder[1:]
			dealBuild(&temp, inorder[0:i], inorder[i+1:], &root)
			break
		}
	}

	return &root
}

func dealBuild(preorder *[]int, left, right []int, root *TreeNode) {
	if len(*preorder) == 0 {
		return
	}
	if len(left) == 0 && len(right) == 0 {
		return
	}

	for i := range left {
		if left[i] == (*preorder)[0] {
			root.Left = &TreeNode{
				Val: (*preorder)[0],
			}
			*preorder = (*preorder)[1:]
			dealBuild(preorder, left[0:i], left[i+1:], root.Left)
			break
		}
	}

	for i := range right {
		if right[i] == (*preorder)[0] {
			root.Right = &TreeNode{
				Val: (*preorder)[0],
			}
			*preorder = (*preorder)[1:]
			dealBuild(preorder, right[0:i], right[i+1:], root.Right)
			break
		}
	}
}

type Flag struct {
	cnt int
	cur int
}

type Char struct {
	c     uint8
	index int
}

func minWindow(s string, t string) string {
	if len(t) == 0 {
		return ""
	}

	dict := make(map[uint8]*Flag)
	for i := range t {
		c := t[i]
		if v, ok := dict[c]; ok {
			v.cnt++
		} else {
			dict[t[i]] = &Flag{
				cnt: 1,
				cur: 0,
			}
		}
	}

	l, r := 0, len(s)-1
	stack := make([]Char, 0)
	findCnt := 0
	for i := range s {
		if v, ok := dict[s[i]]; ok {
			if v.cnt > v.cur {
				findCnt++
			}

			v.cur++
			stack = append(stack, Char{s[i], i})

			for len(stack) > 0 {
				temp := dict[stack[0].c]
				if temp.cur > temp.cnt {
					stack = stack[1:]
					temp.cur--
				} else {
					break
				}
			}

			if findCnt == len(t) {
				if r-l > i-stack[0].index {
					l, r = stack[0].index, i
				}
				if len(t) == 1 {
					break
				}
			}
		}
	}
	return s[l : r+1]
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	mn := m + n
	if n > m { //把空的那个放在第二位,确保m>0
		m, n = n, m
		nums1, nums2 = nums2, nums1
	}

	if n == 0 { //为空的情况
		m_2 := m / 2
		if m%2 == 0 {
			return float64(nums1[m_2]+nums1[m_2-1]) / 2.0
		} else {
			return float64(nums1[m_2])
		}
	}
	// nums1 = a1 ... ai-1 | ai ... am
	// nums2 = b1 ... bj-1 | bj ... bn
	// i+j == m-i+n-j //和为偶数
	// j = (m+n-2i)/2
	// i+j+1 == m-i+n-j //和为奇数
	// j = (m+n-1-2i)/2
	// m+n-2i 和为奇数-偶数(2i)还是奇数 => j = (m+n-2i)/2

	l, r := 0, m
	mid, j := 0, 0
	for l <= r {
		mid = (l + r) / 2
		j = (mn - 2*mid) / 2

		//if mid == 0 || mid == m || j == 0 || j == n {
		//	break
		//}

		if mid-1 >= 0 && j < n && nums1[mid-1] > nums2[j] {
			r = mid
		} else if j-1 >= 0 && mid < m && nums2[j-1] > nums1[mid] {
			l = mid
		} else {
			break
		}
	}

	if mn%2 == 0 {
		return float64(max(nums1[mid-1], nums2[j-1])+min(nums1[mid], nums2[j])) / 2.0
	} else {
		if j >= n {
			return float64(nums1[mid])
		}
		return float64(min(nums1[mid], nums2[j]))
	}
}

//func min(a,b int)int{
//	if a < b {
//		return a
//	}
//	return b
//}
//func max(a,b int)int{
//	if a > b {
//		return a
//	}
//	return b
//}

type ValCnt struct {
	k int
	v int
}

type LRUCache struct {
	//dictCnt map[int]*list.List    //map[次数] 数组序列key
	recent *list.List
	dict   map[int]*list.Element //map[key]ValCnt
	size   int
	minCnt int
}

func Constructor(capacity int) LRUCache {
	lst := list.New()
	lst.Init()
	return LRUCache{
		lst,
		make(map[int]*list.Element),
		capacity,
		0,
	}
}

func (this *LRUCache) setBack(v *list.Element, newVal *int) int {
	this.recent.Remove(v)
	vc := v.Value.(*ValCnt)
	if newVal != nil {
		vc.v = *newVal
	}
	this.dict[vc.k] = this.recent.PushBack(vc)
	return vc.v
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.dict[key]; ok {
		return this.setBack(v, nil)
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.dict[key]; ok {
		this.setBack(v, &value)
		return
	}

	//新增
	if this.size == len(this.dict) {
		//需要移除一个
		v := this.recent.Front()
		this.recent.Remove(v)
		vc := v.Value.(*ValCnt)
		delete(this.dict, vc.k)
	}

	vc := &ValCnt{key, value}
	ele := this.recent.PushBack(vc)
	this.dict[key] = ele
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func findDuplicate(nums []int) int {
	pre := 0
	slow, fast := 0, 0
	for {
		pre = slow
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			return nums[pre]
		}
	}
}

func decodeString(s string) string {
	ret, _ := dealDecode(s)
	return ret
}

func dealDecode(s string) (string, int) {
	ret := ""
	i, j := 0, -1
	for i < len(s) {
		if s[i] >= '0' && s[i] <= '9' {
			//碰到数字了解析数字
			p := i
			q := i
			for q < len(s) {
				q++
				if s[q] == '[' {
					break
				}
			}
			num, _ := strconv.Atoi(s[p:q])
			s1, pos := dealDecode(s[q+1:])
			var pre string
			if j != -1 {
				pre = s[j:i]
			}
			ret += pre + strings.Repeat(s1, num)
			i, j = q+1+pos, -1
		} else if s[i] == ']' {
			var pre string
			if j != -1 {
				pre = s[j:i]
			}
			return ret + pre, i + 1
		} else {
			if j == -1 {
				j = i
			}
			i++
		}
	}

	var pre string
	if j != -1 {
		pre = s[j:i]
	}
	return ret + pre, i
}

func isMatch(s string, p string) bool {
	p1 := strings.Builder{}
	flag := false
	for i := range p {
		if p[i] == '*' {
			if !flag {
				p1.WriteByte(p[i])
				flag = true
			}
		} else {
			p1.WriteByte(p[i])
			flag = false
		}
	}

	dict := make(map[string]bool)
	return dealMatch(s, p1.String(), 0, 0, dict)
}

func dealMatch(s string, p string, i, j int, dict map[string]bool) bool {
	key := strconv.Itoa(i) + "+" + strconv.Itoa(j)
	if v, ok := dict[key]; ok {
		return v
	}

	if len(s) == 0 {
		dict[key] = len(p) == 0 || p[0] == '*' && dealMatch(s, p[1:], i, j+1, dict)
	} else if len(p) == 0 {
		dict[key] = false
	} else if len(s) == 1 && len(p) == 1 {
		dict[key] = s[0] == p[0] || p[0] == '?' || p[0] == '*'
	} else if s[0] == p[0] || p[0] == '?' {
		dict[key] = dealMatch(s[1:], p[1:], i+1, j+1, dict)
	} else if p[0] == '*' {
		dict[key] = dealMatch(s[1:], p[1:], i+1, j+1, dict) || dealMatch(s[1:], p, i+1, j, dict) || dealMatch(s, p[1:], i, j+1, dict)
	} else {
		dict[key] = false
	}
	return dict[key]
}

func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	left, right := make([]int, len(heights)), make([]int, len(heights))
	for i := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			left[i] = stack[len(stack)-1]
		} else {
			left[i] = -1
		}
		stack = append(stack, i)
	}

	for len(stack) > 0 {
		right[stack[len(stack)-1]] = len(heights)
		stack = stack[:len(stack)-1]
	}
	// -1 -1 -1 2 2 4 4 6
	// 1  2   8 4 8 6 8 8

	ret := 0
	for i := range right {
		v := (right[i] - left[i] - 1) * heights[i]
		if v > ret {
			ret = v
		}
	}
	return ret
}

func permuteUnique(nums []int) [][]int {
	dict := make(map[int]int)
	for i := range nums {
		dict[nums[i]] ++
	}

	ret := make([][]int, 0)
	temp := make([]int, 0)
	dealP(dict, temp, len(nums), &ret)
	return ret
}

func dealP(dict map[int]int, arra []int, l int, ret *[][]int) {
	if len(arra) == l {
		temp := make([]int, l)
		copy(temp, arra)
		*ret = append(*ret, temp)
		return
	}

	for k, v := range dict {
		if v == 0 {
			continue
		}

		arra = append(arra, k)
		dict[k]--
		dealP(dict, arra, l, ret)

		arra = arra[:len(arra)-1]
		dict[k]++
	}
}

func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	if row == 0 {
		return nil
	}
	col := len(matrix[0])
	size := row * col
	ret := make([]int, size)
	i, j, k := 0, 0, 0               //行，列，计步
	p, q := 0, 1                     //一开始先向右走
	l, r, t, b := 0, col-1, 0, row-1 //左右上下边界
	for k < size {
		ret[k] = matrix[i][j]
		if p == 0 && q == 1 && j == r {
			t++
			p, q = 1, 0
		} else if p == 1 && q == 0 && i == b {
			r--
			p, q = 0, -1
		} else if p == 0 && q == -1 && j == l {
			b--
			p, q = -1, 0
		} else if p == -1 && q == 0 && i == t {
			l++
			p, q = 0, 1
		}

		i += p
		j += q
		k++
	}
	return ret
}

func equationsPossible(equations []string) bool {
	root := make([]int, 26)
	for i := range root {
		root[i] = i
	}

	var find func(a int) int
	find = func(a int) int {
		if root[a] != a {
			root[a] = root[root[a]]
			a = root[a]
		}
		return a
	}

	union := func(a, b int) {
		root[find(b)] = find(a)
	}

	for i := range equations {
		if equations[i][1] == '=' {
			union(int(equations[i][0]-'a'), int(equations[i][3]-'a'))
		}
	}

	for i := range equations {
		if equations[i][1] == '!' {
			if find(int(equations[i][0]-'a')) == find(int(equations[i][3]-'a')) {
				return false
			}
		}
	}
	return true
}

func main() {
	//twitter := Constructor()
	//// 用户1发送了一条新推文 (用户id = 1, 推文id = 5).
	//twitter.PostTweet(1, 5)
	//// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
	//twitter.GetNewsFeed(1)
	//// 用户1关注了用户2.
	//twitter.Follow(1, 2)
	//// 用户2发送了一个新推文 (推文id = 6).
	//twitter.PostTweet(2, 6)
	//// 用户1的获取推文应当返回一个列表，其中包含两个推文，id分别为 -> [6, 5].
	//// 推文id6应当在推文id5之前，因为它是在5之后发送的.
	//twitter.GetNewsFeed(1)
	//// 用户1取消关注了用户2.
	//twitter.Unfollow(1, 2)
	//// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
	//// 因为用户1已经不再关注用户2.
	//twitter.GetNewsFeed(1)
	//
	//a := make([]int, 0)
	//a = append(a, 0)
	//testArra(a)
	//fmt.Println(a)

	//nlllnl|l nl|llnl|l nl|llnll nlllnll
	//lnl
	//fmt.Println(getMaxRepetitions("nlllnll", 4, "lnl", 1))

	//fmt.Println(numberOfSubarrays([]int{7, 1, 4, 6, 7, 4, 9, 7, 9, 8, 3, 8, 0, 2, 9, 9}, 1))

	//[1,2,3,null,5,null,4]
	//   1
	//  2 3
	//   5  4
	//	root := new(TreeNode)
	//	root.Val = 1
	//	root.Left = new(TreeNode)
	//	root.Left.Val = 2
	//	root.Left.Right = new(TreeNode)
	//	root.Left.Right.Val = 5
	//
	//	root.Right = new(TreeNode)
	//	root.Right.Val = 3
	//	root.Right.Right = new(TreeNode)
	//	root.Right.Right.Val = 4
	//	fmt.Println(rightSideView(root))
	//fmt.Println(waysToChange(34))
	//fmt.Println(reversePairs([]int{0, 21, 15, 25, 24, 15, 24, 5, 7, 0}))
	//fmt.Println()
	//bit := datastruct.NewBIT(10)
	//for i := 1; i <= 10; i++ {
	//	v := int(rand.Int31n(10))
	//	fmt.Printf("%d,", v)
	//	bit.Update(i, v)
	//}
	//fmt.Println()
	//
	//for i := 1; i <= 10; i++ {
	//	fmt.Println(i, bit.Query(i))
	//}
	//fmt.Println(permute([]int{1, 2, 3}))
	//findSubstring("barfoothefoobarman", []string{"foo", "bar"})

	//fmt.Println(search([]int{8,9,2,3,4}, 9))
	//singleNumbers([]int{4, 1, 4, 6})

	//fmt.Println(search([]int{8, 9, 2, 3, 4}, 9))
	//m := MountainArray{[]int{1,2,3,4,5,3,1}}
	//findInMountainArray(2, &m)
	//isHappy(10)

	//var a = [][]byte{
	//	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	//	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	//	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	//	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	//	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	//	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	//	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	//	{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	//solveSudoku(a)
	//combinationSum([]int{7, 3, 2}, 18)
	//multiply("123", "456")
	//mincostTickets([]int{1, 4, 6, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22, 23, 27, 28}, []int{3, 13, 45})
	//mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15})
	//mySqrt(8)

	/*
		["push","push","getMin","push","pop","top","getMin","pop"]
		[[10],[-7],[],[-7],[],[],[],[]]
	*/
	//ms := Constructor1()
	//ms.Push(-10)
	//ms.Push(14)
	//ms.Push(-20)
	//ms.Pop()
	//ms.Push(10)
	//ms.Push(-7)
	//ms.Push(-7)
	//ms.Pop()
	//ms.Top()
	//ms.GetMin()

	//subarraySum([]int{1, 1, 1}, 2)

	//3
	//[[1,0],[2,1]]
	//fmt.Println(findOrder(3, [][]int{{1, 0}, {2, 1}}))
	//fmt.Println(findOrder(5, [][]int{{1, 0}, {2, 0}, {3, 4}, {3, 2}, {4, 2}}))
	//fmt.Println(findOrder(7, [][]int{{1,0},{0,3},{0,2},{3,2},{2,5},{4,5},{5,6},{2,4}}))
	//validPalindrome("abc")
	//findTheLongestSubstring("ioaio")
	//findTheLongestSubstring("leetcodeisgreat")
	//buildTree([]int{3, 1, 2, 4}, []int{1, 2, 3, 4})

	//findMedianSortedArrays([]int{1, 3, 4, 5}, []int{2})

	/*
		["LRUCache","put","put","put","get","put","put","put","put","put","put","put","put","put","put","put","get"]
		[[10],[9,12],[4,30],[9,3],[9],[6,14],[10,11],[11,4],[12,24],[5,18],[7,23],[3,27],[2,9],[13,4],[8,18],[1,7],[6]]
	*/
	//cache := Constructor(2 /* 缓存容量 */)
	//cache.Put(9, 12)
	//cache.Put(9, 3)  // 返回  1
	//cache.Get(9)     // 该操作会使得密钥 2 作废
	//cache.Put(6, 14) // 返回 -1 (未找到)
	//
	//cache.Put(10, 11)         // 该操作会使得密钥 1 作废
	//cache.Put(11, 4)          // 返回 -1 (未找到)
	//fmt.Println(cache.Get(6)) // 返回  4
	//findDuplicate([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	//decodeString("3[a2[c]]")

	//isMatch("abba", "**")
	//largestRectangleArea([]int{4, 2, 0, 3, 2, 4, 3, 4})
	//permuteUnique([]int{1, 1, 2, 2})
	//spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	equationsPossible([]string{"a==b", "b!=c", "c==a"})
	fmt.Println("end")
}
