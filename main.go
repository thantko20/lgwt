package main

import (
	"lgwt/clockface"
	"os"
	"time"
)

func main() {

	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
