package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func usage() {
	fmt.Println("usage: vt-color [BLACK] [RED] [GREEN] [YELLOW] [BLUE] [MAGENTA] [CYAN] [WHITE] [LIGHT_BLACK] [LIGHT_RED] [LIGHT_GREEN] [LIGHT_YELLOW] [LIGHT_BLUE] [LIGHT_MAGENTA] [LIGHT_CYAN] [LIGHT_WHITE]")
}

func main() {
	if len(os.Args) != 17 {
		usage()
		os.Exit(1)
	}

	var red []string
	var green []string
	var blue []string
	for _, color := range os.Args[1:] {
		if len(color) != 6 {
			log.Fatalf("color %v is malformed: use the form RRGGBB", color)
		}

		r, _ := strconv.ParseUint(color[0:2], 16, 64)
		g, _ := strconv.ParseUint(color[2:4], 16, 64)
		b, _ := strconv.ParseUint(color[4:6], 16, 64)

		red = append(red, strconv.FormatUint(r, 10))
		green = append(green, strconv.FormatUint(g, 10))
		blue = append(blue, strconv.FormatUint(b, 10))
	}

	fmt.Println("red: ", strings.Join(red, ","))
	fmt.Println("grn: ", strings.Join(green, ","))
	fmt.Println("blu: ", strings.Join(blue, ","))
}
