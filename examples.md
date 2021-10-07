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
    if reversed > int(math.Pow(2, 31)) {
        return 0
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

## is palindrome

```go
func isPalindrome(x int) bool {
    var ret bool
    
    strX := strconv.Itoa(x)
    // fmt.Printf("%T", reverse)
    reverse := ""
    
    for i, _ := range strX {
        reverse += string(strX[len(strX) - 1 - i])
    }
    
    if strX == reverse {
        ret = true
    }
    return ret
}
```

## remove element

```go
func removeElement(nums []int, val int) int {
    for i := 0; i < len(nums); i++ {
        if nums[i] == val {
            nums = append(nums[:i], nums[i+1:]...)
            i--
        }
    }
    fmt.Println(nums)
    return len(nums)
}
```

## move zeros

```go
func moveZeroes(nums []int)  {
    l := len(nums)
    for i := 0; i < l; i++ {
        if nums[i] == 0 {
            nums = append(nums[:i], nums[i+1:]...)
            nums = append(nums, 0)
            i--
            l--
        }
    }
}
```

## search insert

```go
func searchInsert(nums []int, target int) int {
    for i := 0; i < len(nums); i++ {
        if nums[i] == target {
            return i
        }
        if nums[i] > target {
            return i
        }
    }
    return len(nums)
}
```

## merge / sort arrays

```go
func merge(nums1 []int, m int, nums2 []int, n int)  {
    nums1 = append(nums1[:m], nums2[:]...)
    var item int
    var flag int
    for i := 0; i < len(nums1); i++ {
        item = nums1[i]
		flag = 0
        
		for j := i - 1; j >= 0 && flag != 1; {
			if item < nums1[j] {
				nums1[j+1] = nums1[j]
				j = j - 1
				nums1[j+1] = item
			} else {
				flag = 1
			}
		}
    }
}
```
