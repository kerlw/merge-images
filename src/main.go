package main

import (
	"flag"
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/mergi/loader"
	"go.uber.org/zap"
	"image"
	"io/ioutil"
	"path"
	"strings"
)

var logger = createLogger()

func main()  {
	// var iFlags []string
	// var oFlags []string

	const (
		iFlagMsg = "Enter input directory path"
		oFlagMsg = "Enter output file name, default is out.png"
	)

	input := flag.String("i", ".", iFlagMsg)
	output := flag.String("o", "out.png", oFlagMsg)

	flag.Parse()

	files, err := ioutil.ReadDir(*input)
	if err != nil {
		logger.Error("failed to read files in input dir.", zap.Error(err))
	}

	var imgs []image.Image
	var template []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		importer := loader.NewFileImporter(path.Join(*input, file.Name()))
		img, err := importer.Import()
		if err != nil {
			logger.Warn(file.Name() + " is not a valid image file")
			continue
		}

		imgs = append(imgs, img)
		if len(template) == 0 {
			template = append(template, "T")
		} else {
			template = append(template, "B")
		}
	}

	if len(imgs) > 1 {
		img, err := mergi.Merge(strings.Join(template, ""), imgs)
		if err != nil {
			logger.Error("failed to merge image", zap.Error(err))
		}
		mergi.Export(loader.NewFileExporter(img, *output))
	}
}


