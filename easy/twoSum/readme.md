# TwoSum

First leetcode problem :)

## Problem Statement

Given an array/slice of integers and a target integer, provide the indices of the two integers in the array/slice that add up to equal the target integer. 

Example:
```golang
input := []int{2, 7, 15, 21}
target := 9

return []int{0,1}
```

## My thought Process

I can't think of any way to rearrange or change the input array to make it faster to search. 

If the array were given in an order like `[]int{7,2,5,9,15,3}` I wouldn't want to order from smallest to largest because I need to maintain the initial locations of the numbers. I could work around that but idk that going through that process would make it any faster.

I imagine a loop inside a loop of the same array is not the fastest method but I'm not sure what else to do at the moment.

## My score

Speed: faster than 92% of solutions

Memory: lower memory usage 99.88% of solutions

## Other Solutions and my thoughts on them

https://github.com/haoel/leetcode/blob/master/algorithms/golang/twoSum/twoSum.go

 - i don't really understand maps. and I'm not sure whether this is faster or not.

