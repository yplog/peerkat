package progress

import "fmt"

func PrintProgressBar(progress, width int) {
	fmt.Print("\r[")
	for i := 0; i < width; i++ {
		if i < progress {
			fmt.Print("=")
		} else if i == progress {
			fmt.Print(">")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Print("] ", (progress*100)/width, "%")
}
