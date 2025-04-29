package main

import (
	"flag"
	"fmt"
	"os"
	"snapcode/internal/render"
)

func main() {
	output := flag.String("o", "generated_code.png", "Output file name")
	theme := flag.String("theme", "dark", "Theme: dark or light")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: snapcode [-o output.png] [-theme dark|light] '<code>'")
		os.Exit(1)
	}
	code := args[0]

	img, err := render.MakePNG(code, *theme)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(*output, img, 0644); err != nil {
		panic(err)
	}
	fmt.Println("Saved screenshot to:", *output)

}
