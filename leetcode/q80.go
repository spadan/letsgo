package leetcode

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 删除排序数组中的重复项
func RemoveDuplicate(nums *[]int) int {
	index := 0
	start, end := 0, 0
	arr := *nums
	length := len(arr)
	for end < length {
		for end < length && arr[start] == arr[end] {
			end += 1
		}
		for i := start; i < end && i < start+2; i++ {
			arr[index] = arr[i]
			index += 1
		}
		start = end
	}
	return index
}

// 迪克斯特拉算法解决有权图最短路径
func Dijkstra(graph map[string]map[string]uint, start string) {
	// 开销
	costs := make(map[string]uint)
	// 父节点
	parents := make(map[string]string)
	// 将起点到邻居节点的开销存储至开销表
	for k, v := range graph[start] {
		costs[k] = v
		parents[k] = start
	}
	// 已处理过的节点
	visited := make(map[string]struct{})
	// 取出最便宜的节点
	node, cost := findLowest(costs, visited)
	for len(node) > 0 {
		// 更新邻居节点的开销
		for k, v := range graph[node] {
			// 新开销
			newCost := cost + v
			// 旧开销
			oldCost, exist := costs[k]
			// 如果开销更低，替换
			if !exist || newCost < oldCost {
				costs[k] = newCost
				parents[k] = node
			}
		}
		visited[node] = struct{}{}
		node, cost = findLowest(costs, visited)
	}
	fmt.Println(costs)
	fmt.Println(parents)
}

// 查找开销最低的
func findLowest(costs map[string]uint, visited map[string]struct{}) (string, uint) {
	var node string
	lowestCost := uint(math.MaxUint32)
	for k, v := range costs {
		if _, exist := visited[k]; !exist && v < lowestCost {
			node = k
			lowestCost = v
		}
	}
	return node, lowestCost
}

func LongestSub(a string, b string) string {
	// 动态规划表格
	row, col := len(a), len(b)
	board := make([][]int, row)
	for i := 0; i < row; i++ {
		board[i] = make([]int, col)
	}
	index, max := -1, -1
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if a[i] != b[j] {
				board[i][j] = 0
			} else {
				if i-1 >= 0 && j-1 >= 0 {
					board[i][j] = board[i-1][j-1] + 1
				} else {
					board[i][j] = 1
				}
				if board[i][j] > max {
					max = board[i][j]
					index = i
				}
			}
		}
	}
	fmt.Println(board, max, index)
	if index >= 0 {
		return a[index-max+1 : index+1]
	}
	return ""
}

func AddStrings(num1 string, num2 string) string {
	ans := ""
	add := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add > 0; i, j = i-1, j-1 {
		tmp1, tmp2 := 0, 0
		if i >= 0 {
			tmp1 = int(num1[i] - '0')
		}
		if j >= 0 {
			tmp2 = int(num2[j] - '0')
		}
		tmp := tmp1 + tmp2 + add
		ans = strconv.Itoa(tmp%10) + ans
		add = tmp / 10
	}
	return ans
}

func LicenseKeyFormatting(S string, K int) string {
	var res, win string
	first := true
	var c byte
	for i := 0; i < len(S); i++ {
		c = S[i]
		if first && c == '-' && len(win) > 0 {
			res += win + "-"
			first = false
			win = ""
		} else if c != '-' {
			win += string(c)
		}
		if !first && len(win) == K {
			res += win + "-"
			win = ""
		}
	}
	if len(win) > 0 {
		res += win
	} else {
		res = res[:len(res)-1]
	}
	return strings.ToUpper(res)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	courses := make([][]int, 0)
	for _, v := range prerequisites {
		if courses[v[1]] == nil || len(courses[v[1]]) == 0 {
			courses[v[1]] = make([]int, 0)
		}
		courses[v[1]] = append(courses[v[1]], v[0])
	}
	valid := true
	visited := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		dfs(i, courses, visited, &valid)
		if !valid {
			return false
		}
	}
	return true
}

func dfs(i int, courses [][]int, visited []int, valid *bool) {
	visited[i] = 1
	for _, v := range courses[i] {
		if visited[v] == 0 {
			dfs(v, courses, visited, valid)
			if !*valid {
				return
			}
		} else if visited[v] == 1 {
			*valid = false
			return
		}
	}
	visited[i] = 2
}

func FrogPosition(n int, edges [][]int, t int, target int) float64 {
	adj := make([][]int, n+1)
	for i := 1; i < n+1; i++ {
		adj[i] = make([]int, 0)
	}
	for _, v := range edges {
		adj[v[0]] = append(adj[v[0]], v[1])
		adj[v[1]] = append(adj[v[1]], v[0])
	}
	visited := make([]int, n+1)
	var res float64 = 1
	find := false
	search(1, adj, visited, t, target, &res, &find)
	if find {
		return res
	}
	return 0
}

func search(i int, adj [][]int, visited []int, t int, target int, res *float64, find *bool) {
	if i == target {
		path := 0
		for _, neighbor := range adj[i] {
			if visited[neighbor] == 0 {
				path++
			}
		}
		if t == 0 || path == 0 {
			*find = true
			return
		}
	}
	if t > 0 {
		visited[i] = 1
		path := 0
		for _, neighbor := range adj[i] {
			if visited[neighbor] == 0 {
				path++
			}
		}
		if path > 0 {
			for _, neighbor := range adj[i] {
				if !*find && visited[neighbor] == 0 {
					*res = *res / float64(path)
					t--
					search(neighbor, adj, visited, t, target, res, find)
					if *find {
						return
					}
					*res = *res * float64(path)
					t++
				}
			}
		}
	}
}

func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	max := 0
	start := 0
	for i := 0; i < len(s); i++ {
		cnt1, s1 := centerExpand(s, i, i)
		if cnt1 > max {
			max, start = cnt1, s1
		}
		if i+1 < len(s) {
			cnt2, s2 := centerExpand(s, i, i+1)
			if cnt2 > max {
				max, start = cnt2, s2
			}
		}
	}
	return s[start : start+max]
}

func centerExpand(s string, i int, j int) (int, int) {
	cnt := 0
	for ; i >= 0 && j < len(s); i, j = i-1, j+1 {
		if s[i] != s[j] {
			return cnt, i + 1
		}
		if i == j {
			cnt++
		} else {
			cnt += 2
		}
	}
	return cnt, i + 1
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	cur, smaller, bigger := head, &ListNode{},&ListNode{}
	curs,curb:=smaller,bigger
	for cur!=nil{
		tmp:=cur.Next
		if cur.Val<x{
			curs.Next=cur
			curs=curs.Next
			curs.Next=nil
		}else{
			curb.Next=cur
			curb=curb.Next
			curb.Next=nil
		}
		cur=tmp
	}
	curs.Next=bigger.Next
	bigger.Next=nil
	return smaller.Next
}
