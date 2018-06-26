package main
import ("fmt"
	"bufio"
	"os"
	"strings"
	"math/rand"
	"time"
	"strconv"
	)
func main () {
	args := os.Args
	reader := bufio.NewReader(os.Stdin)
	username := ""
	if len(args) < 2 {
		fmt.Println("Hello, what is your name?")
		username, _ = reader.ReadString('\n')
	} else {
		fmt.Println("hello",args[1])
		username = args[1]
	}
	username = strings.Replace(username,"\n","",-1)
	fmt.Println("hello",username,"do you know the way?[yes/no]")
	option, _   := reader.ReadString('\n')
	if strings.Replace(option,"\n","",-1)  ==  "yes" {
	fmt.Println("Oh you do know the way. Tell me the way!")
	}else{
	fmt.Println("but you do know the way!")
	}
	timeSource := rand.NewSource(time.Now().UnixNano())
	random1 := rand.New(timeSource)
	RandomNumber := random1.Intn(10)
	for {
		fmt.Println("Please,",username,"tell me the way [0-9]")
		Snumber, _ := reader.ReadString('\n')
		number,_ := strconv.Atoi(strings.Replace(Snumber,"\n","",-1))
		if RandomNumber == number {
		fmt.Println("congrats you know the way!")
		break
		}else{
		fmt.Println("you do not know the way, try again!")	
		}
	}	

}
