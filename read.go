package read

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type Name struct {
// 	fname string
// 	lname string
// }

// func main() {
// 	slice := make([]Name, 0, 3)
// 	fmt.Println("Enter file name:\n(Important!!! Same directory)")

// 	var fileName string
// 	fmt.Scan(&fileName)

// 	file, e := os.Open(fileName)
// 	if e != nil {
// 		fmt.Println("Error when trying to open file = ", e)
// 	}

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {

// 		s := strings.Split(scanner.Text(), " ")
// 		var names Name
// 		names.fname, names.lname = s[0], s[1]
// 		slice = append(slice, names)

// 	}

// 	file.Close()
// 	fmt.Printf("\n--------------------\n")

// 	fmt.Println("First Names:")
// 	for _, v := range slice {
// 		fmt.Println(v.fname)
// 	}
// 	fmt.Println("")
// 	fmt.Println("Last Names:")
// 	for _, v := range slice {
// 		fmt.Println(v.lname)
// 	}
