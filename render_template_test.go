package main

import (
	"os"
	"path"
	"strings"
	"testing"
)

func TestRenderServer(t *testing.T) {
	tmpDir := t.TempDir()
	f, err := os.Create(path.Join(tmpDir, "index.html"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	err = os.Mkdir(path.Join(tmpDir, "dir1"), 0640)
	if err != nil {
		t.Fatal(err)
	}
	err = os.Mkdir(path.Join(tmpDir, "dir2"), 0640)
	if err != nil {
		t.Fatal(err)
	}

	c := appConfig{
		listenAddr:  "8081",
		websitePath: tmpDir,
	}
	tData, err := buildTemplateData(&c)
	if err != nil {
		t.Fatal(err)
	}

	err = renderServer(&c, tData)
	if err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(path.Join(c.websitePath, "server.go"))
	if err != nil {
		t.Fatal(err)
	}

	expectedLines := []string{
		"//go:embed dir1 dir2",
		"//go:embed index.html",
	}

	dataStr := string(data)
	for _, line := range expectedLines {
		if !strings.Contains(dataStr, line) {
			t.Fatalf("Expected %s to have line:%s", dataStr, line)
		}
	}

	_, err = os.ReadFile(path.Join(c.websitePath, "go.mod"))
	if err != nil {
		t.Fatal(err)
	}
}
