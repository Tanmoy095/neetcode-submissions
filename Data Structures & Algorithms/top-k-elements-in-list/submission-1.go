func topKFrequent(nums []int, k int) []int {


	count:= make(map[int]int) //[number] -> count
	for _,num := range nums {
		count[num]++	
		
	}
	//suppose [100, 100, 100, 200, 200, 300] 
	//so the map become map[100:3 200:2 300:1]

	buckets := make([][]int , len(nums)+1)


	for num,frequency := range count{
		buckets[frequency] = append (buckets[frequency],num)
	}
	//Index 1 of buckets holds [300],means 300 appears 1 time.Index 2 of buckets holds [200]means 200 appears 2 times. Index 3 of buckets holds [100] holds 100, because 100 appears 3 times

	result := make ([]int,0,k)

	for i:= len(buckets)-1;i>=0;i--{
		 for _,num := range buckets[i]{
			result = append(result,num)
			if len (result) ==k{
				return result 
			}
		 }
	}


	return result 
}

