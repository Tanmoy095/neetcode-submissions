import "slices"
func threeSum(nums []int) [][]int {
    n := len(nums) 
    
    // Declare a set with fixed key size [3]int array key
    set := make(map[[3]int]struct{})
//Deep combinatorics via a triple nested loop
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            for k := j + 1; k < n; k++ {
                
                // for mathematical comparison
                if nums[i]+nums[j]+nums[k] == 0 {
                    
                    triplets := []int{nums[i], nums[j], nums[k]}
                    slices.Sort(triplets)
                    
                    // Convert our non-comparable slice into a comparable fixed array
                    key := [3]int{
                        triplets[0],
                        triplets[1],
                        triplets[2],
                    }
                    
                    // Insert key into set. Map structures guarantee unique keys.
                    set[key] = struct{}{}
                }
            }
        }
    }
    
    // Step 2: Convert the unique set keys back to standard output format ([][]int)
    answer := make([][]int, 0, len(set))
    for key := range set {
        answer = append(answer, []int{key[0], key[1], key[2]})
    }

    return answer
} 