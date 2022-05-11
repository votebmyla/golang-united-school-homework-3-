package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var tmpSlice []int

	for k, _ := range input {
		tmpSlice = append(tmpSlice, k)
	}

	sort.Ints(tmpSlice)

	for _, v := range tmpSlice {
		result = append(result, input[v])
	}
	return
}
