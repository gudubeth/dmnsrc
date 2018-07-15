package dev

import (
	"fmt"
	"time"
)

//PrintElapsedTime prints elapsed time in seconds based on start
func PrintElapsedTime(name string, start time.Time) {
	elapsed := float32(time.Since(start) / time.Millisecond)
	fmt.Printf("%s: %.3f sec\n", name, elapsed/1000)
}
