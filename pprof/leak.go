package main

var leakSlice []int

func memoryLeak() {
	slc := make([]int, 1000000000)
	leakSlice = slc[:100000:100000]
	leakSlice = append(leakSlice, 1)
}
