package slices

func Unique(nums []int) []int {
	seen := make(map[int]bool)
	result := []int{}
	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}
	return result
}

func Chunk(nums []int, size int) [][]int {
	if size <= 0 {
		return nil
	}
	var chunks [][]int
	for i := 0; i < len(nums); i += size {
		end := i + size
		if end > len(nums) {
			end = len(nums)
		}
		chunks = append(chunks, nums[i:end])
	}
	return chunks
}

func Contains(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}
