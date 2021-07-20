//package main

//import (
//	"fmt"
//	"math/rand"
//)

type Solution struct {
    arr []int
    sol []int
}


func Constructor(nums []int) Solution {
    solution := Solution{arr: nums}
    solution.sol = make([]int, len(nums))
    copy(solution.sol, nums)
    return solution
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    copy(this.sol, this.arr)
    return this.arr
}


/** Returns a random shuffling of the array. */
func (solution *Solution) Shuffle() []int {
    n := len(solution.arr)
    for i := 0; i<n; i++ {
        j := i + rand.Intn(n-i)
        solution.sol[i], solution.sol[j] = solution.sol[j], solution.sol[i]
    }
    return solution.sol
}


/**
 * Your Solution object wil be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */l
