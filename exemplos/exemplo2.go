package exemplos

import (
	"fmt"
	"os"
)

func OsView() {
	var s, step string
	fmt.Println(len(os.Args))
	for i := 0; i < len(os.Args); i++ {
		s += step + os.Args[i]
		step = " "
	}
	fmt.Println(s)
}
