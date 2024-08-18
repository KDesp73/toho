package main

import (
	"fmt"
	"os"
	"strings"
	"toho/internal/builder"
	"toho/internal/files"
	"toho/internal/constants"
	"toho/internal/logging"
)

func main() {
	if(os.Args[1] == "-h" || os.Args[1] == "--help"){
		logging.Info("Usage: %s <project-path> <filename> <library-define>", os.Args[0]);
		return
	}

	if(os.Args[1] == "-v" || os.Args[1] == "--version"){
		fmt.Printf("toho v%s\n", constants.VERSION)
		return
	}

	if(len(os.Args) < 4) {
		logging.Info("Usage: %s <project-path> <filename> <library-define>", os.Args[0]);
		return
	}

	project := os.Args[1]
	filename := os.Args[2]
	define := strings.ToUpper(os.Args[3])

	logging.Info("Project path: %s", project)
	logging.Info("File name: %s", filename)
	logging.Info("Library define: %s", define)

	listConfig := files.ListConfig{
		Paths: []string {
			files.SrcDir(project),
			files.IncludeDir(project),
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

	err = builder.Build(filename, define)

	if err != nil {
		logging.Panic("%v", err)
	}

	logging.Info("%s generated successfully", filename)
}
