package main
import (
	"bufio"
	"fmt"
	"os"
)

func add(a int, b int) int {
	return a + b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println("sum is", add(10, 19))
	fmt.Println(text)
	text, _ = reader.ReadString('\n')
	fmt.Println(text)
}
