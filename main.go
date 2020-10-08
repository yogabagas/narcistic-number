package main

import (
	"fmt"
	"math/rand"
)

type sorts struct {
	nums []int64
}

func main() {
	var reqs []int64
	for {
		req := rand.Int63()
		reqs = append(reqs, req)
		if len(reqs) >= 7 {
			break
		}
	}

	tes := sorts{nums: reqs}
	res := tes.sortingNumber()

	fmt.Println("reqs", reqs)
	fmt.Println("res", res)

}

func (s sorts) sortingNumber() []int64 {

	if len(s.nums) < 1 {
		return s.nums
	}

	fmt.Println("nums", s.nums)

	// var iP int = 1

	for i := 0; i < len(s.nums); i++ {
		for j := 0; j < len(s.nums); j++ {
			if s.nums[i] < s.nums[j] {
				num := s.nums[i]
				s.nums[i] = s.nums[j]
				s.nums[j] = num
			}
		}
	}
	return s.nums
}
