package main

import (
	"os"
	"time"

	"lgwt/maths/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
