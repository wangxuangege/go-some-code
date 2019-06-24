package main

import "log"

func main() {
	slice1()
	slice2()
	slice3()
}

func slice1() {
	slice := []int{1, 2, 3, 4}

	newSlice := slice[1:3]
	newSlice[0] += 10
	newSlice[1] += 10
	log.Println(slice)
	log.Println(newSlice)
}

func slice2() {
	slice := []int{1, 2, 3, 4}

	newSlice := append(slice, 5)
	log.Println(slice)
	log.Println(newSlice)

	// append超过长度后，会新分配空间，指向就不同了
	newSlice[0] += 10
	newSlice[1] += 10
	log.Println(slice)
	log.Println(newSlice)
}

func slice3() {
	slice := []int{1, 2, 3, 4}

	// 第三个限制容量，要根据分隔的数组或者切片大小和起始索引位置来看
	newSlice := slice[1:3:3]
	log.Println(slice)
	log.Println(newSlice)

	newSlice[0] += 10
	newSlice[1] += 10
	log.Println(slice)
	log.Println(newSlice)
}
