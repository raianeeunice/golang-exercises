package Bubblesort

// import (
// 	"bufio"
// 	"errors"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func Swap(items []int, i int) {
// 	items[i+1], items[i] = items[i], items[i+1]
// }

// func Bubblesort(items []int) {
// 	n := len(items)
// 	sorted := false
// 	for !sorted {
// 		swapped := false
// 		for i := 0; i < n-1; i++ {
// 			if items[i] > items[i+1] {
// 				Swap(items, i)
// 				swapped = true
// 			}
// 		}
// 		if !swapped {
// 			sorted = true
// 		}
// 		n = n - 1
// 	}
// }

// func ReadNumbers() ([]int, error) {
// 	mySlice := make([]int, 0)
// 	fmt.Println("Enter a 10 number:")
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	strNumbers := scanner.Text()

// 	numbers := strings.Fields(strNumbers)

// 	if len(numbers) != 10{
// 		return mySlice, errors.New("enter exactly ten numbers!")
// 	}

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

// func main() {
// 	arrayNumbers, err := ReadNumbers()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		Bubblesort(arrayNumbers)
// 		fmt.Println("")
// 		fmt.Println("--- Ordered ---")
// 		fmt.Println(arrayNumbers)
// 	}
// }