package main
import
( 
"fmt"
"os"
"strconv"
)

func fPrime(n int) {
first := true 

for d :=2; n > 1; d++{
	for  n % d == 0 {
		if !first{
			fmt.Print("*")
		}
	
	fmt.Print(d)
	first = false
	n /= d 

}

}
fmt.Println()
}

func main(){
	if len(os.Args) != 2 {
		return 
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return 
	}
	if n <= 1 {
		return 
	}
	fPrime(n)
}