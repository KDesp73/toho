package main

import (
	"flag"
	"fmt"
	"toho/internal/builder"
	"toho/internal/constants"
	"toho/internal/files"
	"toho/internal/logging"
)

func main() {
	version := flag.Bool("v", false, "Prints the version of the executable")
	project := flag.String("project", "", "Specify the root path of your library")
	filename := flag.String("header", "", "Specify the generated header file name")
	define := flag.String("define", "", "Specify the library's macro prefix")
	flag.Parse()

	if(*version) {
		fmt.Printf("toho v%s\n", constants.VERSION)
		return
	}

	if(*project == ""){
		logging.Panic("Please specify a project path")
	}

	if(*filename == ""){
		logging.Panic("Please specify a header file")
	}

	if(*define == ""){
		logging.Panic("Please specify a macro prefix")
	}

	logging.Info("Project path: %s", *project)
	logging.Info("File name: %s", *filename)
	logging.Info("Library define: %s", *define)

	listConfig := files.ListConfig{
		Paths: []string {
			files.SrcDir(*project),
			files.IncludeDir(*project),
		},
		Extensions: []string {
			"c",
			"h",
		},
	}
	files, err := files.ListDirectory(listConfig)

	if(err != nil){
		logging.Panic("%v", err)
	}

	if(len(files) == 0) {
		logging.Panic("No source files found")
	}

	logging.Info("%d files found", len(files))

	builder, err := builder.Process(files)

	if err != nil {
		logging.Panic("%v", err)
	}

	err = builder.Build(*filename, *define)

	if err != nil {
		logging.Panic("%v", err)
	}

	logging.Info("%s generated successfully", *filename)
}
