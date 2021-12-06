package main

import (
	"os"
	"path"
	"reflect"
	"testing"
)

func TestReadDir(t *testing.T) {
	tmpDir := t.TempDir()
	f, err := os.Create(path.Join(tmpDir, "test-file"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	err = os.Mkdir(path.Join(tmpDir, "dir1"), 0640)
	if err != nil {
		t.Fatal(err)
	}

	c := appConfig{
		websitePath: tmpDir,
		listenAddr:  ":8081",
	}

	tData, err := buildTemplateData(&c)
	if err != nil {
		t.Fatal(err)
	}

	if tData.ListenAddr != c.listenAddr {
		t.Fatalf("Expected listenAddress to be:%s, got: %s", c.listenAddr, tData.ListenAddr)
	}

	expectedFilenames := []string{"test-file"}
	expectedDirnames := []string{"dir1"}

	if !reflect.DeepEqual(tData.Filenames, expectedFilenames) {
		t.Fatalf("Expected list of filenames: %v, Got: %v\n", expectedFilenames, tData.Filenames)
	}

	if !reflect.DeepEqual(tData.Dirnames, expectedDirnames) {
		t.Fatalf("Expected list of directories: %v, Got: %v\n", expectedDirnames, tData.Dirnames)
	}
}
