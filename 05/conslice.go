package main
import "fmt"

func ConcatSlice(slice1, slice2 []int) []int {
result := []int{}

for i := 0 ; i < len(slice1); i++{
	result = append(result, slice1[i])
}
for i := 0 ; i < len(slice2); i++{
	result = append(result, slice2[i])
}

return result
}


func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}