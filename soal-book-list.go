package main

import (
	"fmt"
)

type Book struct {
	Time  int64
	Alice bool
	Bob   bool
}

func Construct(time int64, alice bool, bob bool) Book {
	return Book{
		Time:  time,
		Alice: alice,
		Bob:   bob,
	}
}

func main() {
	totalChosen := 4
	listBook := []Book{
		Construct(7, true, true),
		Construct(2, true, true),
		Construct(4, false, true),
		Construct(8, true, true),
		Construct(1, false, true),
		Construct(1, true, true),
		Construct(1, false, true),
		Construct(3, false, false),
	}
	fmt.Println(resolve(totalChosen, listBook))
}

func resolve(totalChosen int, bookList []Book) int64 {
	aliceBook := []Book{}
	bobBook := []Book{}
	var totalAlice, totalBob int64
	for _, b := range bookList {
		if b.Alice {
			aliceBook = append(aliceBook, b)
		}

		if b.Bob {
			bobBook = append(bobBook, b)
		}
	}

	sortedAlice := BubbleSort(aliceBook)
	sortedBob := BubbleSort(bobBook)

	if len(sortedBob) < totalChosen || len(sortedAlice) < totalChosen {
		return 0
	}

	for i := 0; i < totalChosen; i++ {
		totalAlice += sortedAlice[i].Time
		totalBob += sortedBob[i].Time
	}

	return totalBob + totalAlice
}

func BubbleSort(array []Book) []Book {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j].Time > array[j+1].Time {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
