package main

import (
	_ "embed"
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"path"
	"text/template"
)

//go:embed data/server.go.tmpl
var serverTmpl []byte

//go:embed data/go.mod.tmpl
var gomodTmpl []byte

type appConfig struct {
	websitePath string
	listenAddr  string
}

type templateData struct {
	ListenAddr string
	Dirnames   []string
	Filenames  []string
}

func renderServer(c *appConfig, tData *templateData) error {
	tmpl := template.New("website2-bin")

	t1, err := tmpl.Parse(string(serverTmpl))
	if err != nil {
		return err
	}

	f, err := os.Create(path.Join(c.websitePath, "server.go"))
	if err != nil {
		return err
	}
	defer f.Close()
	err = t1.Execute(f, tData)
	if err != nil {
		return err
	}

	t2, err := tmpl.Parse(string(gomodTmpl))
	if err != nil {
		return err
	}

	f, err = os.Create(path.Join(c.websitePath, "go.mod"))
	if err != nil {
		return err
	}
	defer f.Close()
	return t2.Execute(f, tData)
}

func setupFlags(w io.Writer, args []string) (*appConfig, error) {
	c := appConfig{}
	fs := flag.NewFlagSet("website2bin", flag.ContinueOnError)
	fs.StringVar(&c.websitePath, "website-path", "", "Directory containing the website files")
	fs.StringVar(&c.listenAddr, "listen-address", ":8080", "Address for the server to listen on")
	err := fs.Parse(args)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func validateConfig(c *appConfig) error {
	if len(c.websitePath) == 0 {
		return errors.New("website path cannot be empty")
	}
	return nil
}

func buildTemplateData(c *appConfig) (*templateData, error) {
	t := templateData{}

	dirEnts, err := os.ReadDir(c.websitePath)
	if err != nil {
		return nil, err
	}

	for _, dirEnt := range dirEnts {
		if dirEnt.IsDir() {
			t.Dirnames = append(t.Dirnames, dirEnt.Name())
			continue
		}
		t.Filenames = append(t.Filenames, dirEnt.Name())
	}
	t.ListenAddr = c.listenAddr
	return &t, nil
}

func main() {
	c, err := setupFlags(os.Stdout, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	err = validateConfig(c)
	if err != nil {
		log.Fatal(err)
	}
	tData, err := buildTemplateData(c)
	if err != nil {
		log.Fatal(err)
	}

	err = renderServer(c, tData)
	if err != nil {
		log.Fatal(err)
	}
}
