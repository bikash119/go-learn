package main
import "fmt"

const(
	Pi = 31.4
	Trurth = false
	Big = 1 << 62
	Small = Big >> 61
)

func main(){
	const Greeting = "bikash"
	fmt.Println(Greeting)
	fmt.Println(Pi)
	fmt.Println(Trurth)
	fmt.Println(Big)
	fmt.Println(Small)
}