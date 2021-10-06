# leet code solutions

## climb stairs

```go
func climbStairs(n int) int {
    if n < 2 {
        return 1
    }
	ways := make([]int, n)
	ways[0] = 1
	ways[1] = 2
	for stairs := 3; stairs <= n; stairs++ {
		ways[stairs-1] = ways[stairs-2] + ways[stairs-3]
	}
	return ways[n-1]
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	x, y := 1, 2
	for n > 2 {
		x, y = y, x+y
		n -= 1
	}
	return y
}
```

## reverse int

```go
func reverse(x int) int {
    var neg bool 
    if x < 0 {
        neg = true
        x = int(math.Abs(float64(x)))
    }
    var reversed int
    for x > 0 {
        last := x % 10
        reversed = reversed * 10 + last
        x = x / 10 
    }
    if neg {
        reversed = reversed * -1
    }
    return reversed
}
```

## reverse array

```go
func main() {
	a := []int{1, 2, 3, 4, 5}

	ret := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		ret[i] = a[len(a)-1-i]
	}
	fmt.Println(ret)
}
```

## two sums

```go
func twoSum(nums []int, target int) []int {
    ret := make([]int, 2)
    for i := 0; i < len(nums) - 1; i++ {
        for j := i+1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                ret[0], ret[1] = i, j
            }
        }
    }
    return ret
}
```

