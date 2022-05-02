package goroutine

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"
// 	"strconv"
// 	"strings"
// 	"sync"
// )

// var (
// 	numInts  int
// 	inputNum int
// 	wg       sync.WaitGroup
// )

// func ReadNumbers() ([]int, error) {
// 	mySlice := make([]int, 0)
// 	fmt.Println("Type whole numbers to add to an array on the same line and separated by space! Press 'Enter' when done :)")
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	strNumbers := scanner.Text()

// 	numbers := strings.Fields(strNumbers)

// 	i := 0
// 	for _, s := range numbers {
// 		num, err := strconv.Atoi(s)
// 		if err == nil {
// 			mySlice = append(mySlice, num)
// 			if i >= 9 {
// 				break
// 			}
// 			i++
// 		}
// 	}
// 	return mySlice, nil
// }

// func Partition(sliceOfInts []int) ([]int, []int, []int, []int) {
// 	var length = len(sliceOfInts)

// 	sliceSize := length / 4
// 	slice1 := sliceOfInts[:sliceSize]
// 	slice2 := sliceOfInts[sliceSize : 2*(sliceSize)]
// 	slice3 := sliceOfInts[2*(sliceSize) : 3*(sliceSize)]
// 	slice4 := sliceOfInts[3*(sliceSize):]

// 	return slice1, slice2, slice3, slice4
// }

// func sortList(unsortedSlice []int) []int {
// 	sort.Ints(unsortedSlice)
// 	return unsortedSlice

// }

// func mergeAndSort(list1 []int, list2 []int, list3 []int, list4 []int) []int {
// 	newSlice := []int{}
// 	newSlice = append(list1, list2...)
// 	newSlice = append(newSlice, list3...)
// 	newSlice = append(newSlice, list4...)
// 	sort.Ints(newSlice)
// 	return newSlice

// }

// func main() {
// 	arrayNumbers, err := ReadNumbers()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		slice1, slice2, slice3, slice4 := Partition(arrayNumbers)
// 		fmt.Println("\n--- Array ---")
// 		fmt.Println(arrayNumbers)

// 		fmt.Println("\nHere are your arrays when partitioned: ", slice1, slice2, slice3, slice4)

// 		wg.Add(4)
// 		go sortList(slice1)
// 		wg.Done()
// 		go sortList(slice2)
// 		wg.Done()
// 		go sortList(slice3)
// 		wg.Done()
// 		go sortList(slice4)
// 		wg.Done()
// 		wg.Wait()

// 		newSlice := mergeAndSort(slice1, slice2, slice3, slice4)

// 		fmt.Println("\nYour Slice merged, sorted and printed is as follows:", newSlice)
// 	}
// }