package main

import (
	// Read/write files
	"bufio"

	// Format/print strings
	"fmt"

	// Log fatal errors with file operations

	// Get files and paths, get commandline arguments
	"os"

	// Get absolute path of files to trim extensions
	"path/filepath"

	// Regex
	"regexp"

	// String manipulation
	"strings"
)

// CONSTANTS
// --------------------------------------------------------------

// Define some constants we will always be using
const (
	/* HTMLHEAD is the HTML data that is inserted into the resulting HTML file when generating
	 * rendered enaml markup. It contains the GFM CSS and the start of the body.
	 */
	HTMLHEAD = `	
<html lang="en">
  <head>
  	<meta charset="utf-8">
  	<title></title>
  	<style>
			body:after,body:before,hr::after,hr::before{display:table;content:""}blockquote,h6{color:#777}@font-face{font-family:octicons-link;src:url(data:font/woff;charset=utf-8;base64,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) format('woff')}body{-ms-text-size-adjust:100%;-webkit-text-size-adjust:100%;color:#333;font-family:"Helvetica Neue",Helvetica,"Segoe UI",Arial,freesans,sans-serif,"Apple Color Emoji","Segoe UI Emoji","Segoe UI Symbol";font-size:16px;line-height:1.6;word-wrap:break-word;background-color:#eee;margin-top:0;margin-bottom:0}.enaml{width:75%;padding:15px;margin:auto;background-color:#fff}ol,p,ul{margin-top:0}a{background-color:transparent;-webkit-text-decoration-skip:objects;color:#4078c0;text-decoration:none}a:active,a:hover{outline-width:0;text-decoration:underline}strong{font-weight:bolder}h1{margin:.67em 0}img{border-style:none}hr{box-sizing:content-box}*{box-sizing:border-box}body:after,hr::after{clear:both}p{margin-bottom:5px}blockquote{margin:0}ul{list-style:circle inside}ol,ul{padding-left:2em;margin-bottom:0}body>:first-child{margin-top:0!important}body>:last-child{margin-bottom:0!important}a:not([href]){color:inherit;text-decoration:none}h1,h2,h3,h4,h5,h6{margin-top:1em;margin-bottom:16px;font-weight:700;line-height:1.4}h1,h2{line-height:1}h3,h4{line-height:1.2}h1{padding-bottom:.3em;font-size:2.25em;border-bottom:1px solid #eee}h2{font-size:1.75em}h3{font-size:1.5em}h4{font-size:1.25em}h5,h6{font-size:1em;line-height:1.1}blockquote{margin-top:0;margin-bottom:16px;padding:0 15px;border-left:4px solid #ddd;font-style:italic}hr{overflow:hidden;background:#e7e7e7;height:4px;padding:0;margin:16px 0;border:0;border-bottom-color:#eee}blockquote>:first-child{margin-top:0}blockquote>:last-child{margin-bottom:0}code{font-family:Consolas,"Liberation Mono",Menlo,Courier,monospace;padding:.2em 0;margin:0;font-size:85%;background-color:rgba(0,0,0,.04);border-radius:3px}code:after,code:before{letter-spacing:-.2em;content:"\00a0"}
  	</style>
  	<meta name="description" content="Translated Markup">
  	<meta name="author" content="Me">
  	<link rel="stylesheet" href="css/styles.css?v=1.0">
  </head>
	
  <body>
  	<div class='enaml'>`

	/* HTMLTAIL is the HTML data that is inserted after the resulting HTML file when generating
	 * rendered enaml markup. It contains closing tags for the main enaml div, body, and html.
	 */
	HTMLTAIL = `	  </div>
  </body>
</html>`
)

// Compile some regexes needed for checking markup

// Check headers, anything from 1 to 6 `#` at the start of the string
var regexHeader, _ = regexp.Compile("^#{1,6}")

// Check blockquotes, get the first `>` at the start of the string
var regexBlockquote, _ = regexp.Compile("^>{1}")

// Check bullets, get the first `-` at the start of the string
var regexBullet, _ = regexp.Compile("^-{1}")

// Check brackets, get everything between `[` and `]`
// TODO: Needs to just return stuff between brackets
var regexBrackets, _ = regexp.Compile("\\[(.*?)\\]")

// FILE UTILITIES
// --------------------------------------------------------------

/* Extension Changing */

// ChangeExtension will change the extension of one file to another
// If there is no extension, still add the new one
func ChangeExtension(path string, extension string) string {

	oldExtension := filepath.Ext(path)
	return path[0:len(path)-len(oldExtension)] + extension

}

// MassChangeExtension will change the extensions
// of a list of paths all to the same extension
func MassChangeExtension(paths []string, extension string) []string {
	newPaths := []string{}

	for _, path := range paths {
		newPaths = append(newPaths, ChangeExtension(path, extension))
	}

	return newPaths

}

/* Saving slices of strings */

// MassSave will save files read from MassLoad to a specific set of paths
func MassSave(files [][]string, paths []string) {

	// Should panic on missmatched array sizes
	if len(files) < len(paths) {
		panic("MassSave(): There are not enough paths to write to for each given file")
	}

	for i, file := range files {

		err := Save(file, paths[i])

		if err != nil {
			panic(fmt.Sprintf("MassSave() > Save(): %s", err))
		}

	}

}

// Save will save a file read from Load to a specific path
func Save(lines []string, path string) error {

	// Create file, return error if applicable
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write contents to file
	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}

	return w.Flush()

}

/* Load files */

// MassLoad will load a list of paths to a slice of string slices
func MassLoad(paths []string) [][]string {

	files := [][]string{}

	// Iterate over paths
	for _, path := range paths {

		// Load each path, log errors if applicable
		contents, err := Load(path)
		if err != nil {
			panic("Load() tried to load an invalid path!")
		}

		files = append(files, contents)

	}

	return files

}

// Load will load a file to a slice
// where each element is a line of the file
func Load(path string) ([]string, error) {

	// Open file, log errors if applicable
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a slice to store file contents
	contents := []string{}

	// Read file line by line to contents
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	return contents, scanner.Err()

}

// TRANSLATION
// --------------------------------------------------------------

/*
Populate tags in the HTML
*/
func Populate(translatedFile []string, translatedLine []string) []string {

	// Save changes to the real line
	line := strings.Join(translatedLine[:], "")

	// Copy translated file
	file := translatedFile

	// Start checking the beginning of every line

	if strings.HasPrefix(line, "#") {

		// Headers
		match := regexHeader.FindAllString(line, -1)
		header := match[0]

		file = append(file, fmt.Sprintf("      <h%d>%s</h%d>", len(header), line[len(header)+1:], len(header)))

	} else if strings.HasPrefix(line, ">") {

		// Blockquotes
		match := regexBlockquote.FindAllString(line, -1)
		blockquote := match[0]

		file = append(file, fmt.Sprintf("      <blockquote>%s</blockquote>", line[len(blockquote)+1:]))

	} else if strings.HasPrefix(line, "-") {

		// Bullets
		match := regexBullet.FindAllString(line, -1)
		bullet := match[0]

		file = append(file, fmt.Sprintf("      <ul><li>%s</li></ul>", line[len(bullet)+1:]))

	} else if strings.HasPrefix(line, "[") {

		// Bracket metadata

		trimmedLine := line[1 : len(line)-1]
		args := strings.Split(trimmedLine, " ")

		if args[0] == "img" {

			// [img ...]
			if len(args) != 2 {
				file = append(file, "      <p style='color:red'><b>Error: Image metadata has improper syntax</b></p>")
			} else {
				file = append(file, "      <img src="+args[1]+">")
			}

		} else if args[0] == "div" {

			// [div]
			file = append(file, "      <hr>")

		} else if args[0] == "link" {

			// [link ...]
			if len(args) != 3 {
				file = append(file, "      <p style='color:red'><b>Error: Link metadata has improper syntax</b></p>")
			} else {
				file = append(file, "      <a href="+args[2]+">"+args[1]+"</a>")
			}

		}

	} else {

		// Most common cases handled here

		if line == "" {

			// Empty lines become breaks in the HTML
			file = append(file, "      </br>")

		} else {

			// Otherwise always update the translated File with the line
			file = append(file, "      <p>"+line+"</p>")

		}

	}

	return file

}

/*
Translate will turn enaml markup to HTML.
The input is expected to be a string slice where each element corresponds to each line of a file,
such as that output from `Load()`.
The resulting output will be a string slice of the same length as the input, though each element
will have been rendered from enaml markup to HTML
*/
func Translate(file []string) []string {

	// Where all rendered markup will be sent to
	translatedFile := []string{}

	// Add the html head boilerplate

	// Check to see if there is a title. A tile is written in the form of `[title]`
	// at the beginning of the file so we see if our bracket regex returns any matches
	// for the first line of the file (file[0])
	match := regexBrackets.FindAllString(file[0], -1)

	// If len(match) > 0, we have a title
	if len(match) > 0 {

		// Comb through the boilerplate to find the `<title></title>` and insert our title there
		for _, line := range addHeadWithTitle(match[0][1 : len(match[0])-1]) {
			translatedFile = append(translatedFile, line)
		}

		// Add a nice little header to the start of the HTML file
		translatedFile = append(translatedFile, "      <h1>"+match[0][1:len(match[0])-1]+"</h1>")

		// Remove the first line of the file
		file = file[1:]

	} else {

		// Otherwise use the standard title
		for _, line := range addHeadWithTitle("enaml") {
			translatedFile = append(translatedFile, line)
		}

	}

	// Keep track of what tags become opened on each line
	// Keep these outside of the loop so they can persist between lines
	openBold := false
	openCode := false
	openItalics := false
	openUnderline := false

	// Iterate over each line in the file
	for _, line := range file {

		// Keep track of changes made to the line
		translatedLine := []string{}

		// First switch statement will look at the whole line and
		// add in styling tags such as bold, code, italics, and underline
		for _, c := range line {
			switch c {
			case '@':
				if openItalics {
					translatedLine = append(translatedLine, "</i>")
				} else {
					translatedLine = append(translatedLine, "<i>")
				}
				openItalics = !openItalics
			case '_':
				if openUnderline {
					translatedLine = append(translatedLine, "</u>")
				} else {
					translatedLine = append(translatedLine, "<u>")
				}
				openUnderline = !openUnderline
			case '%':
				if openBold {
					translatedLine = append(translatedLine, "</b>")
				} else {
					translatedLine = append(translatedLine, "<b>")
				}
				openBold = !openBold
			case '`':
				if openCode {
					translatedLine = append(translatedLine, "</code>")
				} else {
					translatedLine = append(translatedLine, "<code>")
				}
				openCode = !openCode
			default:
				translatedLine = append(translatedLine, fmt.Sprintf("%c", c))
			}
		}

		// Populate other tags
		translatedFile = Populate(translatedFile, translatedLine)

	}

	// Always close your tags
	if openBold {
		translatedFile = append(translatedFile, "</b>")
	}

	if openCode {
		translatedFile = append(translatedFile, "</code>")
	}

	if openItalics {
		translatedFile = append(translatedFile, "</i>")
	}

	if openUnderline {
		translatedFile = append(translatedFile, "</u>")
	}

	// Add tail boilerplate
	translatedFile = append(translatedFile, HTMLTAIL)

	return translatedFile

}

/*
MassTranslate will use `Translate()` to generate HTML data for a slice of given file contents
acquired with `Load()`
The input is expected to be a slice of string slices where the string slices are a
line by line string representation of an enaml file. The resulting output
will be a slice of string slices of rendered enaml as HTML
*/
func MassTranslate(files [][]string) [][]string {

	// Store output from each `Translate()`
	translatedFiles := [][]string{}

	// `Translate()` all files
	for _, file := range files {
		translatedFiles = append(translatedFiles, Translate(file))
	}

	return translatedFiles

}

// addHeadWithTitle will return a string slice of HTMLHEAD where
func addHeadWithTitle(title string) []string {

	// Add title
	headWithTitle := strings.Replace(HTMLHEAD, "<title></title>", "<title>"+title+"</title>", -1)

	// Read to string slice
	contents := []string{}
	for _, line := range strings.Split(headWithTitle, "\n") {
		contents = append(contents, line)
	}

	return contents

}

// TranslateArgv all from commandline
func TranslateArgv(paths []string) {

	if len(paths) < 1 {
		panic("main: Not enough paths in arguments!")
	}

	// Load files
	// files := [][]string{}
	files := MassLoad(paths)

	// Change extensions from whatever to html
	newPaths := MassChangeExtension(paths, ".html")

	// translatedFiles := [][]string{}
	translatedFiles := MassTranslate(files)

	// Save contents
	MassSave(translatedFiles, newPaths)
}

// MAIN
// --------------------------------------------------------------

func main() {
	TranslateArgv(os.Args[1:])
}
