package leetcode

import (
	"fmt"
	"strings"
	"unicode"
)

type goods struct {
	weight int
	value  int
}

var allGoods = []goods{{4, 3000}, {3, 2000,}, {1, 1500}}
var best = 0
var capacity = 4

// i:第i件商品
// w:当前总重量
// v:当前总价值
func DfsBag(i int, w int, v int, chose []int) {
	// 所有商品选择完毕
	if i >= len(allGoods) {
		if v > best {
			best = v
			fmt.Println(best, chose)
		}
	} else {
		// 每件商品可以选择拿或不拿，0-不拿，1-拿
		g := allGoods[i]
		for j := 0; j <= 1; j++ {
			if w+j*g.weight <= capacity {
				chose[i] = j
				w += j * g.weight
				v += j * g.value
				DfsBag(i+1, w, v, chose)
				w -= j * g.weight
				v -= j * g.value
			}
		}
	}
}

// 求子集
func DfsCombination(arr []int, i int, tmp []int) {
	// 所有商品选择完毕
	if i >= len(arr) {
		fmt.Println(tmp)
	} else {
		// 每件商品可以选择拿或不拿，0-不拿，1-拿
		for j := 0; j <= 1; j++ {
			if j == 1 {
				tmp = append(tmp, arr[i])
			}
			DfsCombination(arr, i+1, tmp)
			if j == 1 {
				tmp = tmp[0 : len(tmp)-1]
			}
		}
	}
}

var c = 0

// row:行索引
// queen：每行摆放皇后的位置
func DfsQueen(row int, queen []int) {
	if row >= len(queen) {
		c = c + 1
		fmt.Println(queen, c)
	} else {
		for col := 0; col < len(queen); col++ {
			if isOk(queen, row, col) {
				queen[row] = col
				DfsQueen(row+1, queen)
			}
		}

	}

}
func isOk(queen []int, row int, col int) bool {
	for i := 0; i < row; i++ {
		if queen[i] == col || row-col == i-queen[i] || row+col == i+queen[i] {
			return false
		}
	}
	return true
}

func isPalindrome(s string) bool {
	if len(s)==0{
		return true
	}
	s=strings.ToLower(s)
	for i,j:=0,len(s)-1;i<j;{
		si := rune(s[i])
		if !(unicode.IsLetter(si)||unicode.IsDigit(si)){
			i++
			continue
		}
		sj := rune(s[j])
		if !(unicode.IsLetter(sj)||unicode.IsDigit(sj)){
			j++
			continue
		}
		if si!=sj{
			return false
		}
	}
	return true
}

func solveSudoku(board [][]byte) {
	black:=make([][]int,0)
	for i:=0;i<9;i++{
		for j:=0;j<9;j++{
			if board[i][j]=='.'{
				black=append(black,[]int{i,j})
			}
		}
	}
	dfs811(board,black,0)
}

func dfs811(board [][]byte,black [][]int, n int)bool{
	if n>=len(black){
		return true
	}
	i,j:=black[n][0],black[n][1]
	for num:='1';num<='9';num++{
		if isValid(board,i,j,num){
			board[i][j]=num
			if dfs811(board,black,n+1){
				return true
			}
			board[i][j]='.'
		}
	}
	return false
}

func isValid(board [][]byte,i,j int,num byte)bool{
	for n:=0;n<9;n++{
		// 遍历
		if board[i][n]==num ||board[n][j]==num{
			return false
		}
		key,value:=(i/3)*3+n/3,(j/3)*3+n%3
		if board[key][value]==num{
			return false
		}
	}
	return true
}