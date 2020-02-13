package algoritmic_on_Golang

//простая сортировка вставками, метод грубой силы
func InsertionSort(ar []int) {
	for i:=0; i<len(ar)-2; i++ {
		min := i
		for j:= i+1; j<len(ar)-1; j++ {
			if ar[min]>ar[j] {min = j}}
		ar[i], ar[min] = ar[min], ar[i]}
}
