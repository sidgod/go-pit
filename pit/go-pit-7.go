package pit

import "fmt"

// formatting test changes

func PitSeven() {
	for i := 0; i < 256; i++ {
		fmt.Printf("%d %b %x %#x %q\n", i, i, i, i, i)
	}
}
