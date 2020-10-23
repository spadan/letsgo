package leetcode

// 1. 循环退出条件，注意是 low <= high，⽽不是 low < high。
// 2. mid 的取值，mid := low + (high-low)>>1
// 3. low 和 high 的更新。low = mid + 1，high = mid - 1。
func binarySearchMatrix(nums []int, target int) int{
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 2)
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 查找第一个目标数
func binarySearchMatrixFirst(nums []int, target int) int{
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 2)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid]<target{
			low = mid + 1
		}else{
			if mid==0 || nums[mid-1]!=target{
				return mid
			}
			high=mid-1
		}
	}
	return -1
}

// 查找最后一个目标数
func binarySearchMatrixLast(nums []int, target int) int{
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 2)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid]<target{
			low = mid + 1
		}else{
			if mid==len(nums)-1 || nums[mid+1]!=target{
				return mid
			}
			low=mid+1
		}
	}
	return -1
}

func MaximalRectangle(matrix [][]byte) int {
	m:=len(matrix)
	if m==0{
		return 0
	}
	n:=len(matrix[0])
	if n==0{
		return 0
	}
	// dp[i][j]定义,第i行第j列，在本行中连续为1的个数
	dp:=make([][]int,m)
	for i:=0;i<m;i++{
		dp[i]=make([]int,n)
	}
	res:=0
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if matrix[i][j]==0{
				dp[i][j]=0
			}else{
				dp[i][j]=1
				if j-1>=0{
					dp[i][j]+=dp[i][j-1]
				}
			}
			if dp[i][j]>=1{
				for row:=i-1;row>=0&&dp[row][j]>=1;row--{
					height:=i-row+1
					width:=min(dp[i][j],dp[row][j])
					res=max(res,height*width)
				}
			}
		}
	}
	return res
}

func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}