package main

import (
	"reflect"
	"testing"
)

// TestChangeExtension will test the ChangeExtension and MassChangeExtension
// functions of osfile
func TestChangeExtension(t *testing.T) {

	extensionTables := []struct {
		infile    string
		extension string
		outfile   string
	}{
		{"file.yml", ".html", "file.html"},
		{"file", ".html", "file.html"},
		{"file", ".o", "file.o"},
		{"file.exension with spaces", ".html", "file.html"},
		{"file.exension with spaces", ".tar.gz", "file.tar.gz"},
		{"file.tar.gz", ".html", "file.tar.html"},
	}

	// Since we know ChangeExtension will fail without strings as input,
	// we just need to test string inputs

	for _, table := range extensionTables {
		outfile := ChangeExtension(table.infile, table.extension)
		if outfile != table.outfile {
			t.Errorf("File %s could not have it's extension changed to %s, expected %s", table.infile, table.extension, table.outfile)
		}
	}

	massExtensionTables := []struct {
		infiles   []string
		extension string
		outfiles  []string
	}{
		{[]string{"file.txt", "file.png", "file.jpeg", "file.tar.gz", "file"}, ".md", []string{"file.md", "file.md", "file.md", "file.tar.md", "file.md"}},
		{[]string{"file.txt", "file.png", "file.jpeg", "file.tar.gz", "file"}, ".html", []string{"file.html", "file.html", "file.html", "file.tar.html", "file.html"}},
		{[]string{"file.txt", "file.png", "file.jpeg", "file.tar.gz", "file"}, ".yml", []string{"file.yml", "file.yml", "file.yml", "file.tar.yml", "file.yml"}},
		{[]string{"file.txt", "file.png", "file.jpeg", "file.tar.gz", "file"}, ".o", []string{"file.o", "file.o", "file.o", "file.tar.o", "file.o"}},
	}

	for _, table := range massExtensionTables {
		outfiles := MassChangeExtension(table.infiles, table.extension)
		if !reflect.DeepEqual(outfiles, table.outfiles) {
			t.Errorf("Files %v could not have it's extension changed to %s, expected %v", table.infiles, table.extension, table.outfiles)
		}
	}

}

// TestLoad will test the Load and MassLoad functions of osfile
func TestLoad(t *testing.T) {

	// Paths and expected contents we will be using throughout the tests
	dummyPaths := []string{"./test/first.enaml", "./test/second.enaml"}

	// Expected contents of first.enaml
	firstContents := []string{"[First]", "", "_@%data%@_"}

	// Expected contents of second.enaml
	secondContents := []string{"[Another Enaml File]", "", "# Hello world", "", "> welcome to enaml"}

	// Combined contents of first.enaml and second.enaml
	combinedContents := [][]string{firstContents, secondContents}

	// Test loading first file
	dummyContents, err := Load(dummyPaths[0])
	if err != nil {
		t.Errorf("Error while loading single file: %s", err)
	}

	for i, line := range dummyContents {
		if line != firstContents[i] {
			t.Errorf("Loaded files are not the same")
		}
	}

	// Test mass loading
	dummyFiles := MassLoad(dummyPaths)

	for i, file := range dummyFiles {
		for j, line := range file {
			if line != combinedContents[i][j] {
				t.Errorf("Error on massload")
			}
		}
	}

}

func TestLoadPanics(t *testing.T) {

	// Test loading bad files
	_, err := Load("/bad/file/extension/should/not/exist")
	if err == nil {
		t.Errorf("Load should have resulted in an error")
	}

	defer func() {
		{
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}
	}()

	MassLoad([]string{"/bad/file", "other/bad/file"})

}
func TestSavePanics(t *testing.T) {

	defer func() {
		{
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}
	}()

	MassSave([][]string{{"/bad/file"}}, []string{".dummy", ".will cause error on compare"})
	MassSave([][]string{{"/bad/file"}}, []string{".will cause error on save"})
	MassSave([][]string{{"/bad/file"}}, []string{})
	MassSave([][]string{{}}, []string{})

}

// TestSave will test the Save and MassSave functions of osfile
func TestSave(t *testing.T) {

	// Dummy input paths
	dummyPaths := []string{"./test/first.enaml", "./test/second.enaml"}

	// Dummy output paths
	dummySavePaths := []string{"./test/first.html", "./test/second.html"}

	// Test saving one file
	dummyContent, _ := Load(dummyPaths[0])
	err := Save(dummyContent, dummySavePaths[0])
	if err != nil {
		t.Errorf("Error while saving single file: %s", err)
	}

	// Test saving many files
	dummyContents := MassLoad(dummyPaths)
	MassSave(dummyContents, dummySavePaths)

}

// TestTranslate will test both `Translate()` and `MassTranslate()` as well as `Populate()`
func TestTranslate(t *testing.T) {
	// Averages at about 160k~ ns. Could be better
	files := [][]string{
		{"[Test Enaml 1]", "`_@%Hello World %@_`", "# H", "## e", "### l", "#### l", "##### o", "###### ."},
		{"[Test Enaml 2]", "> Blockquote", "- Bullet", "Normal Text"},
		{"[Test Enaml 3]", "[div]", "[div]", "[div]", "[div]", "[div]", "[div]", "[div]"},
		{"[Test Enaml 4]", "# Bad metadata", "[link bad line]", "[img bad img]"},
		{"Test Enaml 5, No Title", "Unclosed Tests @_`%"},
		{"Test Enaml 6, No Title, Blank lines", ""},
		{"Test Enaml 7, No Title", "[link Scala http://www.scala-lang.org/]", "[img https://assets.toptal.io/uploads/blog/category/logo/55/scala.png]"},
		{"Test Enaml 8, No Title", "[comment]"},
	}

	// Test MassTranslate
	MassTranslate(files)

}

// BENCHMARKS
// --------------------------------------------------------------

func BenchmarkFull(b *testing.B) {

	for i := 0; i < b.N; i++ {
		TranslateArgv([]string{"./test/first.enaml", "./test/second.enaml"})
	}

}

func TestFull(t *testing.T) {
	TranslateArgv([]string{"./test/first.enaml", "./test/second.enaml"})
}

func TestFullPanic(t *testing.T) {

	defer func() {
		{
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}
	}()

	TranslateArgv([]string{})

}
