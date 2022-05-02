package anima

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type Animal interface {
// 	Eat()
// 	Move()
// 	Speak()
// }

// type Cow struct {
// 	eat string
// 	move string
// 	speak string
// }

// type Bird struct {
// 	eat string
// 	move string 
// 	speak string
// }

// type Snake struct {
// 	eat string
// 	move string
// 	speak string
// }

// func (cow Cow) Eat() {
// 	fmt.Println(cow.eat)
// }

// func (cow Cow) Move() {
// 	fmt.Println(cow.move)
// }

// func (cow Cow) Speak() {
// 	fmt.Println(cow.speak)
// }

// func (bird Bird) Eat() {
// 	fmt.Println(bird.eat)
// }

// func (bird Bird) Move() {
// 	fmt.Println(bird.move)
// }

// func (bird Bird) Speak() {
// 	fmt.Println(bird.speak)
// }

// func (snake Snake) Eat() {
// 	fmt.Println(snake.eat)
// }

// func (snake Snake) Move() {
// 	fmt.Println(snake.move)
// }

// func (snake Snake) Speak() {
// 	fmt.Println(snake.speak)
// }

// func main() {

// 	var nameType map[string]string

// 	nameType = make(map[string]string)

// 	for {
// 		fmt.Printf(">")

// 		scanner := bufio.NewScanner(os.Stdin)
// 		scanner.Scan()
// 		request := scanner.Text()

// 		if len(strings.Split(request, " ")) != 3 {
// 			fmt.Println("Command input number is not equal to 3")
// 			continue
// 		}

// 		command := strings.Split(request, " ")[0]

// 		switch command {
// 		case "newanimal":

// 			name := strings.Split(request, " ")[1]
// 			animalType := strings.Split(request, " ")[2]

// 			_, boolean := nameType[name]
// 			if boolean == false {
// 				nameType[name] = animalType
// 				fmt.Println("Created it!")
// 			} else {
// 				fmt.Println("Please try another name")
// 			}
// 		case "query":

// 			name := strings.Split(request, " ")[1]
// 			information := strings.Split(request, " ")[2]

// 			switch nameType[name] {
// 			case "cow":
// 				cow := Cow{"grass", "walk", "moo"}
// 				switch information {
// 				case "eat":
// 					cow.Eat()
// 				case "move":
// 					cow.Move()
// 				case "speak":
// 					cow.Speak()
// 				default:
// 					fmt.Printf("%s: information not found on animal type", information)
// 					continue
// 				}
// 			case "bird":
// 				bird := Bird{"worms", "fly", "peep"}
// 				switch information {
// 				case "eat":
// 					bird.Eat()
// 				case "move":
// 					bird.Move()
// 				case "speak":
// 					bird.Speak()
// 				default:
// 					fmt.Printf("%s: information not found on animal type", information)
// 					continue
// 				}
// 			case "snake":
// 				snake := Snake{"mice", "slither", "hsss"}
// 				switch information {
// 				case "eat":
// 					snake.Eat()
// 				case "move":
// 					snake.Move()
// 				case "speak":
// 					snake.Speak()
// 				default:
// 					fmt.Printf("%s: information not found on animal type", information)
// 					continue
// 				}
// 			default:
// 				fmt.Printf("%s: name not found on animal type\n", nameType[name])
// 				continue
// 			}
// 		default:
// 			fmt.Printf("%s is not in (newanimal, query)... Try again!\n", command)
// 			continue
// 		}
// 	}
// }