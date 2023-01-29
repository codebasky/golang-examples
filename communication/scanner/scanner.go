package scanner

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		fmt.Print(scan.Text() + "\n")
	}
	fmt.Println("End of scanner")
}
