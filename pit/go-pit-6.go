package pit

import (
	"flag"
	"fmt"
	"strings"
)

func PitSix() {
	var n = flag.Bool("n", false, "Omit trailing newline")
	var sep = flag.String("s", "", "Argument separator")

	flag.Parse()

	fmt.Printf(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}
