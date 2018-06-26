package main
import ("fmt"
	"bufio"
	"os"
	"strings"
	)
func main () {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello, what is your name?")
	username, _ := reader.ReadString('\n')
	fmt.Println("hello",username,"do you know the way?[yes/no]")
	option, _   := reader.ReadString('\n')
	if strings.Replace(option,"\n","",-1)  ==  "yes" {
	fmt.Println("Oh you do know the way. Tell me the way!")
	}else{
	fmt.Println("but you do know the way!")
	}
	

}
