package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func Quiz(fname string) error {
	r, err := os.OpenFile(fname, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	reader := csv.NewReader(r)
	correct := 0
	total := 0
	ir := bufio.NewScanner(os.Stdin)
	for {
		q, err := reader.Read()
		if err == io.EOF {
			fmt.Printf("You answered %d correctly of total %d\n", correct, total)
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println(q)
		total++
		fmt.Printf("Q: %s? \n", q[0])
		fmt.Println("Ans:")
		var ans string
		if ir.Scan() {
			ans = ir.Text()
		}

		if ans == q[1] {
			correct++
		}
	}
}
