package builder

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"toho/internal/files"
	"toho/internal/logging"
	"toho/internal/strutils"
)

type Builder struct {
	Files []files.File
	Includes []string
	Implementations string
	Definitions string
}

func (builder *Builder) filterIncludes() {
	includes := []string{}
	for _, include := range builder.Includes {
		if(strutils.StartsWith(include, "#include <")){
			includes = append(includes, include)
		}
	}

	builder.Includes = strutils.RemoveDuplicates(includes)
}

func extractNumber(str string) (int, error) {
    pattern := `\{(\d+)\}`
    
    re, err := regexp.Compile(pattern)
    if err != nil {
        return -1, err
    }
    
    matches := re.FindStringSubmatch(str)
    if len(matches) < 2 {
        return -1, fmt.Errorf("no number found in the string")
    }
    
    num, err := strconv.Atoi(matches[1])
    if err != nil {
        return -1, err
    }
    
    return num, nil
}

func Process(filesStructs []files.File) (*Builder, error) {
    builder := &Builder{}
    builder.Files = filesStructs

    orderedFiles := make(map[int]files.File)
    maxIndex := 0

    for _, file := range filesStructs {
        if file.Extension == "h" {
            hasIndex := false
            var index int
            for lineNum, line := range strings.Split(file.Content, "\n") {
                if strutils.StartsWith(line, "// {") {
                    var err error
                    index, err = extractNumber(line)
                    if err != nil {
                        return nil, fmt.Errorf("error extracting number from file %s, line %d: %v", file.Name, lineNum+1, err)
                    }

                    logging.Info("Index found: %d", index)
                    hasIndex = true
                    break
                }
            }

            if hasIndex {
                if prevFile, ok := orderedFiles[index]; !ok {
                    logging.Info("Added file: %s with index: %d", file.Name, index)
                    orderedFiles[index] = file
                    if index > maxIndex {
                        maxIndex = index
                    }
                } else {
                    // Replace the file with the duplicate index
                    logging.Info("Replaced file: %s with index: %d", prevFile.Name, index)
                    orderedFiles[index] = file
                }
            } else {
                maxIndex++
                logging.Info("Added file: %s with index: %d", file.Name, maxIndex)
                orderedFiles[maxIndex] = file
            }
        }
    }

    keys := make([]int, 0, len(orderedFiles))
    for k := range orderedFiles {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    for _, i := range keys {
        if file, ok := orderedFiles[i]; ok {
            for _, line := range strings.Split(file.Content, "\n") {
                if strutils.IsWhitespace(line) {
                    continue
                }

                if strutils.StartsWith(line, "#include") {
                    builder.Includes = append(builder.Includes, line)
                } else if file.Extension == "h" {
                    builder.Definitions += line + "\n"
                } else if file.Extension == "c" {
                    builder.Implementations += line + "\n"
                }
            }
        }
    }

    builder.filterIncludes()

    return builder, nil
}

func (builder *Builder) Build(headerFile string, define string) error {
	outFile, err := os.Create(headerFile)
	if err != nil {
		return err
    }
    defer outFile.Close()

	writer := bufio.NewWriter(outFile)
    defer writer.Flush()

	writer.WriteString(fmt.Sprintf("#ifndef %s_H\n", define))
	writer.WriteString(fmt.Sprintf("#define %s_H\n", define))
	writer.WriteString("\n")

	writer.WriteString(fmt.Sprintf("#ifndef %sAPI\n", define))
	writer.WriteString(fmt.Sprintf("\t#define %sAPI static inline\n", define))
	writer.WriteString("#endif\n")
	
	writer.WriteString("\n")
	writer.WriteString("\n")

	for _, include := range builder.Includes {
		writer.WriteString(include + "\n")
	}
	writer.WriteString("\n")
	writer.WriteString("\n")

	writer.WriteString(builder.Definitions)

	writer.WriteString("\n")
	writer.WriteString("\n")

	writer.WriteString(fmt.Sprintf("#ifdef %s_IMPLEMENTATION\n", define))

	writer.WriteString("\n")
	writer.WriteString("\n")

	writer.WriteString(builder.Implementations)

	writer.WriteString("\n")
	writer.WriteString("\n")

	writer.WriteString("#endif\n")
	writer.WriteString("#endif\n")

	return nil
}
