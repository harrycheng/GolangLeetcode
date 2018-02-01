package algorithms

func MySort(values *[]int){
	qSort(values , 0, len(*values)-1)
}

func qSort(values *[]int , left, right int){
	if left < right {
		pivotPositon := partition(values, left, right)
		qSort(values, left, pivotPositon -1)
		qSort(values, pivotPositon + 1, right)
	}
}

func partition(values *[]int , left, right int) int {	
	pivot := (*values)[left]
	
	for left < right {
		for left < right && (*values)[right] >= pivot {
			right--
		}
		(*values)[left] = (*values)[right];

		for left < right && (*values)[left] <= pivot {
			left++
		}
		(*values)[right] = (*values)[left]
	}

	(*values)[left] = pivot

	return left;
}