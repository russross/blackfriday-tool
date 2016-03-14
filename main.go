//
// Blackfriday Markdown Processor
// Available at http://github.com/russross/blackfriday
//
// Copyright © 2011 Russ Ross <russ@russross.com>.
// Distributed under the Simplified BSD License.
// See README.md for details.
//

//
//
// Example front-end for command-line use
//
//

package main

import (
	"flag"
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
	"runtime/pprof"
	"strings"
	"path/filepath"
)

const DEFAULT_TITLE = ""

type Settings struct {
	latex, smartypants bool
	css, templateText string
	repeat,extensions,htmlFlags int
}
func main() {
	var page, toc, toconly, xhtml, latexdashes, fractions bool
	var bf Settings
	var cpuprofile, templateFile string
	// parse command-line options
	flag.BoolVar(&page, "page", false,
		"Generate a standalone HTML page (implies -latex=false)")
	flag.BoolVar(&toc, "toc", false,
		"Generate a table of contents (implies -latex=false)")
	flag.BoolVar(&toconly, "toconly", false,
		"Generate a table of contents only (implies -toc)")
	flag.BoolVar(&xhtml, "xhtml", true,
		"Use XHTML-style tags in HTML output")
	flag.BoolVar(&bf.latex, "latex", false,
		"Generate LaTeX output instead of HTML")
	flag.BoolVar(&bf.smartypants, "smartypants", true,
		"Apply smartypants-style substitutions")
	flag.BoolVar(&latexdashes, "latexdashes", true,
		"Use LaTeX-style dash rules for smartypants")
	flag.BoolVar(&fractions, "fractions", true,
		"Use improved fraction rules for smartypants")
	flag.StringVar(&bf.css, "css", "",
		"Link to a CSS stylesheet (implies -page)")
	flag.StringVar(&templateFile, "template","",
		"Template file to add the content in")
	flag.StringVar(&cpuprofile, "cpuprofile", "",
		"Write cpu profile to a file")
	flag.IntVar(&bf.repeat, "repeat", 1,
		"Process the input multiple times (for benchmarking)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Blackfriday Markdown Processor v"+blackfriday.VERSION+
			"\nAvailable at http://github.com/russross/blackfriday\n\n"+
			"Copyright © 2011 Russ Ross <russ@russross.com>\n"+
			"Distributed under the Simplified BSD License\n"+
			"See website for details\n\n"+
			"Usage:\n"+
			"  %s [options] [inputfile [outputfile]]\n\n"+
			"Options:\n",
			os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	// enforce implied options
	if bf.css != "" {
		page = true
	}
	if page {
		bf.latex = false
	}
	if toconly {
		toc = true
	}
	if toc {
		bf.latex = false
	}

	// turn on profiling?
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}


	// set up options
	bf.extensions = 0
	bf.extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	bf.extensions |= blackfriday.EXTENSION_TABLES
	bf.extensions |= blackfriday.EXTENSION_FENCED_CODE
	bf.extensions |= blackfriday.EXTENSION_AUTOLINK
	bf.extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	bf.extensions |= blackfriday.EXTENSION_SPACE_HEADERS

	bf.htmlFlags = 0
	var err error
	if !bf.latex {
		// render the data into HTML
		if xhtml {
			bf.htmlFlags |= blackfriday.HTML_USE_XHTML
		}
		if bf.smartypants {
			bf.htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
		}
		if fractions {
			bf.htmlFlags |= blackfriday.HTML_SMARTYPANTS_FRACTIONS
		}
		if latexdashes {
			bf.htmlFlags |= blackfriday.HTML_SMARTYPANTS_LATEX_DASHES
		}
		if page {
			bf.htmlFlags |= blackfriday.HTML_COMPLETE_PAGE
		}
		if templateFile != "" {
			var templateString []byte
			if templateString, err = ioutil.ReadFile(templateFile); err != nil {
				fmt.Fprintln(os.Stderr, "Error reading from", templateFile, ":", err)
				os.Exit(-1)
			}
			bf.templateText = string(templateString[:])
		}
		if toconly {
			bf.htmlFlags |= blackfriday.HTML_OMIT_CONTENTS
		}
		if toc {
			bf.htmlFlags |= blackfriday.HTML_TOC
		}
	}

	// read the input
	var input []byte
	args := flag.Args()
	switch len(args) {
	case 0:
		if input, err = ioutil.ReadAll(os.Stdin); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading from Stdin:", err)
			os.Exit(-1)
		}
		renderContent(bf,input,"","")
	case 1, 2:
		var fileNames []string
		var fileName string
		fileNames, err = filepath.Glob(args[0])
		for _, fileName = range fileNames{
			//fmt.Println(fileName)
			if input, err = ioutil.ReadFile(fileName); err != nil {
				fmt.Fprintln(os.Stderr, "Error reading from", fileName, ":", err)
				os.Exit(-1)
			}
			if len(fileName) > 3 && strings.ToLower(fileName[len(fileName)-3:len(fileName)]) == ".md" {
				fileName = fileName[:len(fileName)-3]
			}
			_,fileName = filepath.Split(fileName)
			outputFileName := ""
			if (len(args) == 2){
				if bf.latex {
					outputFileName = args[1] + fileName + ".tex"
				} else {
					outputFileName = args[1] + fileName + ".html"
				}
			}
			
			renderContent(bf,input,outputFileName,fileName)
		}
	default:
		flag.Usage()
		os.Exit(-1)
	}
}
func renderContent(bf Settings, input []byte, outputFileName string, fileName string){
	var renderer blackfriday.Renderer
	var templatewithtitle string
	if bf.latex {
		// render the data into LaTeX
		renderer = blackfriday.LatexRenderer(0)
	} else {
		// render the data into HTML
		title := DEFAULT_TITLE
		if (bf.htmlFlags | blackfriday.HTML_COMPLETE_PAGE == blackfriday.HTML_COMPLETE_PAGE) {
			title = getTitle(input)
		}
		if bf.templateText != "" {
			title = getTitle(input)
			templatewithtitle = strings.Replace(bf.templateText,"{{title}}",title,-1)
			templatewithtitle = strings.Replace(templatewithtitle,"{{filename}}",fileName,-1)
		}
		renderer = blackfriday.HtmlRenderer(bf.htmlFlags, title, bf.css)
	}

	// parse and render
	var output []byte
	for i := 0; i < bf.repeat; i++ {
		output = blackfriday.Markdown(input, renderer, bf.extensions)
	}
	
	if templatewithtitle != "" {
		output = []byte(strings.Replace(templatewithtitle,"{{content}}",string(output[:]),-1))
	}

	// output the result
	var out *os.File
	var err error
	if outputFileName != "" {
		if out, err = os.Create(outputFileName); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating %s: %v", outputFileName, err)
			os.Exit(-1)
		}
		defer out.Close()
	} else {
		out = os.Stdout
	}

	if _, err = out.Write(output); err != nil {
		fmt.Fprintln(os.Stderr, "Error writing output:", err)
		os.Exit(-1)
	}
}

// try to guess the title from the input buffer
// just check if it starts with an <h1> element and use that
func getTitle(input []byte) string {
	i := 0

	// skip blank lines
	for i < len(input) && (input[i] == '\n' || input[i] == '\r') {
		i++
	}
	if i >= len(input) {
		return DEFAULT_TITLE
	}
	if input[i] == '\r' && i+1 < len(input) && input[i+1] == '\n' {
		i++
	}

	// find the first line
	start := i
	for i < len(input) && input[i] != '\n' && input[i] != '\r' {
		i++
	}
	line1 := input[start:i]
	if input[i] == '\r' && i+1 < len(input) && input[i+1] == '\n' {
		i++
	}
	i++

	// check for a prefix header
	if len(line1) >= 3 && line1[0] == '#' && (line1[1] == ' ' || line1[1] == '\t') {
		return strings.TrimSpace(string(line1[2:]))
	}

	// check for an underlined header
	if i >= len(input) || input[i] != '=' {
		return DEFAULT_TITLE
	}
	for i < len(input) && input[i] == '=' {
		i++
	}
	for i < len(input) && (input[i] == ' ' || input[i] == '\t') {
		i++
	}
	if i >= len(input) || (input[i] != '\n' && input[i] != '\r') {
		return DEFAULT_TITLE
	}

	return strings.TrimSpace(string(line1))
}
