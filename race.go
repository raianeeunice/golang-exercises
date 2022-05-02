package race

// import (
// 	"fmt"
// 	"time"
// )

// //A race condition occurs when multiple threads try to access and modify the same data. 

// var count int

// func race(name string) {
  
//     for i := 1; true; i++ {
  
//         // prints string and number
//         fmt.Println(name, i)
          
//         // this makes the program sleep
//         time.Sleep(time.Second * 1) 
//     }
// }
// func main() {

// 	// func call simply creates a goroutine
// 	go race("First")
// 	// func call simply creates a goroutine
// 	go race("Second")
	
// 	// time.Sleep(1 * time.Second)
// 	fmt.Println("Program ends!")
// }