func productExceptSelf(nums []int) []int {
    n := len(nums)
    
    
    output := make([]int, n)
    
    //left products
    leftProduct := 1
    for i := 0; i < n; i++ {
        output[i] = leftProduct
        leftProduct *= nums[i] 
    }
    
    //  Multiply by the right products
    rightProduct := 1
    // loop backwards here: from n-1 down to 0
    for i := n - 1; i >= 0; i-- {
        output[i] *= rightProduct
        rightProduct *= nums[i] 
    }
    
    return output
}