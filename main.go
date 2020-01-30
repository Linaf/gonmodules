package main
import (
	"fmt"

	"github.com/Linaf/gonmodules/router"
)

func main() {

	fmt.Println("Server started ")

	e := router.New()

	e.Start(":1323")
}
