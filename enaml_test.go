package enaml

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
	dummyPaths := []string{"./src/dummydata/first.enaml", "./src/dummydata/second.enaml"}

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

// TestSave will test the Save and MassSave functions of osfile
func TestSave(t *testing.T) {

	// Dummy input paths
	dummyPaths := []string{"./src/dummydata/first.enaml", "./src/dummydata/second.enaml"}

	// Dummy output paths
	dummySavePaths := []string{"./src/dummydata/first.html", "./src/dummydata/second.html"}

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
