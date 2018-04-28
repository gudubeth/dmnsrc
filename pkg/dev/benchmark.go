package dev

import (
	"fmt"
	"time"
)

func PrintElapsedTime(name string, start time.Time) {
	elapsed := float32(time.Since(start) / time.Millisecond)
	fmt.Printf("%s: %.3f sec\n", name, elapsed/1000)
}
