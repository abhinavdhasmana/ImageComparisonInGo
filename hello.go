package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math"
	"os"
	"time"
)

func main() {
	start := time.Now()
	args := os.Args[1:]

	target_file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer target_file.Close()

	source_file, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer source_file.Close()

	target, _, err := image.Decode(target_file)
	if err != nil {
		log.Fatal(err)
	}

	source, _, err := image.Decode(source_file)
	if err != nil {
		log.Fatal(err)
	}

	target_bounds := target.Bounds()
	source_bounds := source.Bounds()
	if !boundsMatch(target_bounds, source_bounds) {
		log.Fatal("Image sizes don't match!")
	}
	var diff int64
	for y := target_bounds.Min.Y; y < target_bounds.Max.Y; y++ {
		for x := target_bounds.Min.X; x < target_bounds.Max.X; x++ {
			diff += compareColor(target.At(x, y), source.At(x, y))
		}
	}
	fmt.Printf("%d\n", diff)
	elapsed := time.Since(start)
	log.Printf("image comparision took %s", elapsed)
}

func compareColor(a, b color.Color) (diff int64) {
	r1, g1, b1, a1 := a.RGBA()
	r2, g2, b2, a2 := b.RGBA()

	diff += int64(math.Abs(float64(r1 - r2)))
	diff += int64(math.Abs(float64(g1 - g2)))
	diff += int64(math.Abs(float64(b1 - b2)))
	diff += int64(math.Abs(float64(a1 - a2)))
	return diff
}

func boundsMatch(a, b image.Rectangle) bool {
	return a.Min.X == b.Min.X && a.Min.Y == b.Min.Y && a.Max.X == b.Max.X && a.Max.Y == b.Max.Y
}
