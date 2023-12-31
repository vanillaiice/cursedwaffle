// ----------------------------------------------- //
// Shenanihacks Hackaton Submission - 31 Dec 2023 //
// ----------------------------------------------//
package main

import (
	"cursedwaffle/waffles"
	"fmt"
	"os"
	"time"

	"github.com/inancgumus/screen"
)

func readFiles(files []string) ([][]byte, error) {
	b := [][]byte{}
	for _, f := range files {
		btemp, err := os.ReadFile(f)
		if err != nil {
			return b, err
		}
		b = append(b, btemp)
	}
	return b, nil
}

func printFrame(frame string, d time.Duration) {
	fmt.Print(frame)
	time.Sleep(d)
	screen.Clear()
	screen.MoveTopLeft()
}

func main() {
	screen.Clear()
	screen.MoveTopLeft()

	w := waffles.All()

	i := 0
	for i < 5 {
		printFrame(w[0], 300*time.Millisecond)
		printFrame(w[1], 100*time.Millisecond)
		printFrame(w[2], 100*time.Millisecond)
		printFrame(w[3], 300*time.Millisecond)
		printFrame(w[2], 100*time.Millisecond)
		printFrame(w[1], 100*time.Millisecond)
		i++
	}

	printFrame(w[4], 3*time.Second)

	printFrame(w[5], 1*time.Second)
	printFrame(w[6], 1*time.Second)
	printFrame(w[7], 1*time.Second)
	printFrame(w[8], 1*time.Second)
	printFrame(w[9], 1*time.Second)
	printFrame(w[10], 1*time.Second)
	printFrame(w[11], 1*time.Second)

	printFrame(w[4], 3*time.Second)
	printFrame(w[0], 2*time.Second)

	fmt.Println("Congrats, you wasted around 20 seconds of your life :/")
	time.Sleep(1 * time.Second)
	fmt.Println("...")
	time.Sleep(1 * time.Second)
	fmt.Println("So now, GO outside and touch some grass.")
}
