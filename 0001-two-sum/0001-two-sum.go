func twoSum(nums []int, target int) []int {
    result := make([]int, 2)
    flag := false
    for idx, _ := range nums {
        for i:=idx+1; i<len(nums); i++ {
            if target == (nums[idx] + nums[i]) {
                if idx > i {
                    result[0] = i
                    result[1] = idx
                } else {
                    result[0] = idx
                    result[1] = i
                }
                break
                flag = true
            }
        }
        if flag {
            break
        }
    }
    return result
}