package quiz

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"
)

func Quiz(fname string, wait time.Duration) error {
	file, err := os.OpenFile(fname, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	reader := csv.NewReader(file)
	correct := 0
	total := 0
	ansC := make(chan string)

	tc := time.NewTicker(wait)
	defer tc.Stop()
	defer close(ansC)

	for {
		q, err := reader.Read()
		if err == io.EOF {
			fmt.Printf("You answered %d correctly of total %d\n", correct, total)
			return nil
		}
		if err != nil {
			return err
		}
		total++

		fmt.Printf("Q: %s? \n", q[0])
		fmt.Print("Ans:")

		go func() {
			var uans string
			fmt.Scanf("%s\n", &uans)
			ansC <- uans
		}()

		select {
		case ans := <-ansC:
			if ans == q[1] {
				correct++
			}
			tc.Reset(wait)
		case <-tc.C:
			fmt.Println("Timeout")
			fmt.Printf("You answered %d correctly of total %d\n", correct, total)
			return nil
		}
	}
}
