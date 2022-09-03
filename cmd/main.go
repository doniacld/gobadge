package main

import (
	"flag"
	"fmt"

	"github.com/doniacld/gobadge/cmd/logos"
)

const (
	tinygoLogo          = "./cmd/assets/tinygo.jpg"
	gopherconEU22Logo   = "./cmd/assets/gopherconeu-2022.jpg"
	gopherconUK22Logo   = "./cmd/assets/gopherconuk-2022.jpg"
	gopherconUS22Logo   = "./cmd/assets/gopherconus-2022.jpg"
	containerDays22Logo = "./cmd/assets/container-days-2022.jpg"

	defaultVarName    = "logoRGBA"
	defaultOutputFile = "logo.go"
)

func main() {
	conf := flag.String("conf", "tinygo", "Choose the conference logo you want to (e.g.: tinygo, gceu22, gcuk22, gcus22)")
	imgPath := flag.String("path", "", "Path to the image to convert")
	varName := flag.String("varName", defaultVarName, "Variable name of the RGBA array values")
	outputFile := flag.String("output", defaultOutputFile, "Name of the file where the variable will be stored")
	flag.Parse()

	c := confs()
	logo, ok := c[*conf]
	if !ok {
		fmt.Println("I do not have yet this conf in my catalog.")
		return
	}

	img := logos.ImageRGBA{
		VarName:    *varName,
		OutputFile: *outputFile,
	}

	if len(*imgPath) == 0 {
		img.Filepath = logo
	} else {
		img.Filepath = *imgPath
	}

	img.GenerateImageRGBAFile()
}

func confs() map[string]string {
	return map[string]string{
		"tinygo": tinygoLogo,
		"gceu22": gopherconEU22Logo,
		"gcuk22": gopherconUK22Logo,
		"gcus22": gopherconUS22Logo,
		"cds22":  containerDays22Logo,
	}
}
