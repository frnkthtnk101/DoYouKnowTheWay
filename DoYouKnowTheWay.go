package main
import ("fmt"
	"bufio"
	"os"
	)
func main () {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello, what is your name?")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

}
