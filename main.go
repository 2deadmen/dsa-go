package main

func main() {
	in := []int{1, 2, 3, 1}
	print(containsDuplicate(in))
}

func containsDuplicate(nums []int) bool {
	if len(nums) < 1 {
		return false
	}

	dup := map[int]string{}
	for _, i := range nums {

		if _, ok := dup[i]; ok {

			return true
		}

		dup[i] = "ha"

	}

	return false
}
