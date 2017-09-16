package count_nonatomic

var counter uint64 = 0

func Inc_count() {
	counter++
}

func Read_count() uint64 {
	return counter
}


