/*
http://weibo.com/1915548291/yDjra5ajO#1348465113389
一个大小为n的数组，里面的数都属于范围[0, n-1]，有不确定的重复元素，
找到至少一个重复元素，要求O(1)空间和O(n)时间。
*/
package main

import "fmt"

func find_duplicate(ia []int) bool {
	for i := 0; i < len(ia); {
		if ia[i] < 0 {
			return true
		}
		t := ia[i]
		ia[i] = -1
		i = t
	}
	return false
}

func main() {
	ia := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 8}
	fmt.Println(find_duplicate(ia))
}
