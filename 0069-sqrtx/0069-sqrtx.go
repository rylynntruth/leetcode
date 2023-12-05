func mySqrt(x int) int {
    if x == 0 || x == 1 {
        return x
    }

    left, right := 1, x
    var result int

    for left <= right {
        mid := left + (right - left) / 2
        if mid <= x / mid {
            left = mid + 1
            result = mid
        } else {
            right = mid - 1
        }
    }

    return result
}