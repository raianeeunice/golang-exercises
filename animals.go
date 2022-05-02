package animal

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type Animal struct {
// 	name string
// 	eat   string
// 	move  string
// 	speak string
// }

// func (animal Animal) Eat() {
// 	fmt.Println(animal.eat)
// }

// func (animal Animal) Move() {
// 	fmt.Println(animal.move)
// }

// func (animal Animal) Speak() {
// 	fmt.Println(animal.speak)
// }

// func main() {

// 	var animal Animal

// 	for {
// 		fmt.Printf(">")

// 		scanner := bufio.NewScanner(os.Stdin)
// 		scanner.Scan()
// 		request := scanner.Text()

// 		comand := strings.Split(request, " ")[0]
// 		name := strings.Split(request, " ")[1]
// 		information := strings.Split(request, " ")[2]

// 		switch comand {
// 		case "newanimal":
// 			animal = Animal{}

// 		}

// 		switch name {
// 		case "cow":
// 			animal = Animal{eat: "grass", move: "walk", speak: "moo"}
// 		case "bird":
// 			animal = Animal{eat: "worms", move: "fly", speak: "peep"}
// 		case "snake":
// 			animal = Animal{eat: "mice", move: "slither", speak: "hsss"}
// 		default:
// 			fmt.Printf("%s: name not found on animal type ", name)
// 			return
// 		}

// 		switch information {
// 		case "eat":
// 			animal.Eat()
// 		case "move":
// 			animal.Move()
// 		case "speak":
// 			animal.Speak()
// 		default:
// 			fmt.Printf("%s: information not found on animal type", information)
// 			return
// 		}
// 	}
// }
