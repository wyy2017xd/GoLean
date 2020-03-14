package main

import "fmt"

func main() {
	a := make([]int, 5)
	b := a[0:5]
	a = append(a, 1)//创建了新数组,cap*2
	a[1] = 5//对有影响
	fmt.Println(b, cap(b))
	// [0 0 0 0]
	fmt.Println(a, cap(a))
	// [0 5 0 0 0 1]</pre>

	s := []int{}
	for i := 0; i < 16; i++ {
		s = append(s, i)
		fmt.Println(cap(s))
	}

	/*当需要的容量超过原切片容量的两倍时，会使用需要的容量作为新容量。
	当原切片长度小于1024时，新切片的容量会直接翻倍。而当原切片的容量大于等于1024时，会反复地增加25%，直到新容量超过所需要的容量*/
}