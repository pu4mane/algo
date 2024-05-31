package insertsort

func InsertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		curr := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > curr {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = curr
	}
	return arr
}
