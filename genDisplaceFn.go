package GenDisplaceFn

// import "fmt"

// var(
// 	a float64 //acceleration
// 	vo float64 //velocity
// 	d float64 //displacement
// 	t float64 //time
// )
// func GenDisplaceFn(a, vo, so float64) func(float64) float64 {

// 	return func(t float64) float64 {
// 		return ((0.5 * a * (t * t)) + (vo * t) + so)
// 	}
// }

// func main() {
// 	fmt.Println("------ s = ½ a t2 + vot + so ------")
// 	fmt.Print("Enter acceleration: ")
// 	fmt.Scan(&a)
// 	fmt.Print("Enter initial velocity: ")
// 	fmt.Scan(&vo)
// 	fmt.Print("Enter initial displacement: ")
// 	fmt.Scan(&d)
// 	fmt.Print("Enter time (seconds): ")
// 	fmt.Scan(&t)

// 	fn := GenDisplaceFn(a, vo, d)

// 	fmt.Printf("\nDisplacement after %v seconds  --->  %v\n", t, fn(t))

// }