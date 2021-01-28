# 7. Reverse Integer

## Problem Statement

Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).


Example 1:
```
Input: x = 123
Output: 321
```

Example 2:
```
Input: x = -123
Output: -321
```

Example 3:
```
Input: x = 120
Output: 21
```

Example 4:
```
Input: x = 0
Output: 0
```

## My thought Process

Convert to string, reverse, convert to int32

If new number is opposite sign, it has gone outside of 32 bit, right? yes.

```golang
var x int32 = -2147483648
fmt.Println(x-1) // 2147483647
x = 2147483647
fmt.Println(x+1) // -2147483648
```

## My score

Speed: faster than 100% of solutions

Memory: lower memory usage 13.53% of solutions

My memory usage was dogsh** lol.

I need to either switch to a less memory management/more friendly language like JS/TS or spend more time learning about Golang memory management.

I had to go to google and copy/pasted the String function as opposed to writing it myself.

## Other Solutions and my thoughts on them

```golang
func reverse(x int) int {
    
    res := 0 
    
  
    for x != 0 {
        
        res = res * 10 + x % 10
        x = x / 10
    }
    
    if res > math.MaxInt32 || res < math.MinInt32 {
        return 0
    }
    return res
}
```

wow I am not a smart one...why am i always going for more complicated solutions?
