package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/dave/jennifer/jen"
)

var packageName = flag.String("pkg", "webclient", "name of package of out file")
var inputFilePath = flag.String("in", "./examples/webclient/index.ts", "path of input file")
var outputFilePath = flag.String("out", "./pkg/factory/webclient/stringified_decls.go", "path of output file")

var literalMatcher = regexp.MustCompile(`([a-zA-Z])\w*`)

func declarationName(decl string) string {
	matches := literalMatcher.FindAllString(decl, 3)

	if len(matches) < 3 {
		panic(fmt.Sprintf("found < 3 matches for this decl:\n%s", decl))
	}

	if matches[0] == "export" {
		return fmt.Sprintf("%s_%s", matches[1], matches[2])
	}

	return fmt.Sprintf("%s_%s", matches[0], matches[1])
}

func main() {
	flag.Parse()

	content, err := os.ReadFile(*inputFilePath)
	if err != nil {
		panic(err)
	}

	decls := strings.Split(string(content), "\n\n")

	file := jen.NewFile(*packageName)
	for _, d := range decls {
		file.Const().Id(declarationName(d)).Op("=").Id("`" + escapeBackticks(strings.TrimSpace(d)) + "`")
	}

	buf := bytes.NewBuffer(nil)
	file.Render(buf)
	if err := os.WriteFile(*outputFilePath, buf.Bytes(), 0644); err != nil {
		panic(err)
	}
}

func escapeBackticks(s string) string {
	return strings.Replace(s, "`", "` + \"`\" +  `", -1)
}