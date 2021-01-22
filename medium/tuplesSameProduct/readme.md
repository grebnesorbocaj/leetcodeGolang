# 1726. Tuple with Same Product

## Problem Statement

Given an array nums of distinct positive integers, return the number of tuples (a, b, c, d) such that a * b = c * d where a, b, c, and d are elements of nums, and a != b != c != d.

Example 1:
```
Input: nums = [2,3,4,6]
Output: 8
Explanation: There are 8 valid tuples:
(2,6,3,4) , (2,6,4,3) , (6,2,3,4) , (6,2,4,3)
(3,4,2,6) , (4,3,2,6) , (3,4,6,2) , (4,3,6,2)
```

Example 2:
```
Input: nums = [1,2,4,5,10]
Output: 16
Explanation: There are 16 valids tuples:
(1,10,2,5) , (1,10,5,2) , (10,1,2,5) , (10,1,5,2)
(2,5,1,10) , (2,5,10,1) , (5,2,1,10) , (5,2,10,1)
(2,10,4,5) , (2,10,5,4) , (10,2,4,5) , (10,2,4,5)
(4,5,2,10) , (4,5,10,2) , (5,4,2,10) , (5,4,10,2)
```

Example 3:
```
Input: nums = [2,3,4,6,8,12]
Output: 40
```

Example 4:
```
Input: nums = [2,3,5,7]
Output: 0
```

## My thought Process

each set (ab and cd) will really be 8 combinations -> `abcd` `abdc` `bacd` `badc` `adcb` `adbc` `dacb` `dabc`

If the length of the array is less than 3, return 0 automatically.

For those with length of array greater than 4

`[a, b, c, d, e....]`

ab ac ad ae bc bd be cd ce de <- all possible combos

check if these match eachother

make a map of ints that is map[int][]int (is that possible?)

starting with first item:
- map[a*b]=[a,b]
- map[a*c]=[a,c]
- map[a*d]=[a,d]
- map[a*e]=[a,e]
- etc etc etc

before each is created, see if that key exists. if it does-> that will be 1 set...
how do you account for more than 1 set? like if ab bc cd de all equal same value

`ab = bc` `ab = cd` `ab = de` `bc = cd` `bc = de` `cd = de`

why don't i remember any math. isn't this algebra? or is it precal or some shit


## My score

Speed: faster than ~~70.97%~~ 91.94% of solutions after removing if statement around comboMap

Memory: lower memory usage 72.58% of solutions (resubmit same code gave 91.94%)

## Other Solutions and my thoughts on them

https://leetcode.com/problems/tuple-with-same-product/discuss/1020704/Go-Beats-100-count-pairs-O(n2)

ok this is using some math or some ish

the below is essentially the same as mine but more elegant

```golang
ans, m, n := 0, map[int]int{}, len(nums)
    
for i := 0; i < n - 1; i++ {
    for j := i + 1; j < n; j++ {
        m[nums[i] * nums[j]]++
    }
}
```

I am not sure why the below is correct

```golang
for _, v := range m {
    ans += (v - 1) * v * 4
}
```
