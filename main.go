package main

import (
	"lgwt/clockface/svg"
	"os"
	"time"
)

func main() {

	t := time.Now()
	svg.Write(os.Stdout, t)
}
