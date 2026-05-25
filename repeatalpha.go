// Online Go compiler to run Golang program online
// Print "Start small. Ship something." message

package main
import "fmt"

func RepeatAlpha(s string) string {
    res:= ""
    for _, ch := range s {
        if ch >='a' && ch <= 'z' {
            count := int( ch -'a' + 1)
            for i:= 0 ; i < count ; i++{
                res +=string(ch)
            }
        } else if ch >='A' && ch <= 'Z' {
            count := int( ch -'A' + 1)
            for i:= 0 ; i < count ; i++{
                res +=string(ch)
            }
        } else {
            res+= string(ch)
        }
    }
    return res
}
 func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}