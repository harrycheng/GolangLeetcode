package algorithms

func quickSort(values *[]int, left, right int) {
	temp := (*values)[left]    //第一个元素为枢轴
	p := left
	i, j := left, right

	for i <= j {
		for j > p && (*values)[j] >= temp {
			j--    //找到比枢轴小的，直到枢轴位置
		}
		
		if j >= p {
			(*values)[p] = (*values)[j]
			p = j
		}		

		if (*values)[i] < temp && i <= p {
			i++
		}

		if i <= p {
			(*values)[p] = (*values)[i]
			p = i
		}
	}

	(*values)[p] = temp

	if p-left > 1 {
		quickSort(values, left, p-1)
	}

	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

func QuickSort(values *[]int) {
	quickSort(values, 0, len(*values)-1)
}
