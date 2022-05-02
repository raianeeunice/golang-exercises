package module

// import "fmt"
// import "bufio"
// import "os"
// import "strings"
// import "strconv"

// func GenDisplaceFn(acceleration float64, velocity float64, initial_displacement float64) func(float64) float64 {
// 	return func(time float64) float64 {
// 		return 1/2*acceleration*time*time + velocity*time + initial_displacement
// 	}
// }

// func main() {
// 	fmt.Println("Please, write the acceleration, initial velocity, and initial displacement separated by space:")
// 	reader := bufio.NewReader(os.Stdin)
// 	in, err := reader.ReadString('\n')

// 	if err != nil {
// 		panic(err)
// 	}

// 	in = strings.TrimRight(in, "\n")

// 	list_of_string_numbers := strings.Split(in, " ")

// 	if len(list_of_string_numbers) < 3 {
// 		panic("Three arguments required")
// 	}

// 	acceleration, a_err := strconv.ParseFloat(list_of_string_numbers[0], 64)
// 	velocity, v_err := strconv.ParseFloat(list_of_string_numbers[1], 64)
// 	displacement, d_err := strconv.ParseFloat(list_of_string_numbers[2], 64)

// 	if a_err != nil || v_err != nil || d_err != nil {
// 		panic("Invalid numbers")
// 	}

// 	var time float64

// 	fmt.Println("Write a value for time")
// 	fmt.Scan(&time)
// 	fn := GenDisplaceFn(acceleration, velocity, displacement)
// 	fmt.Println(fn(time))

// }
