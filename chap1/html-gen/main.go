package main

import (
"errors"
"flag"
"fmt"
"html/template"
"io"
"os"
)

type config struct {
	name       string
	outputPath string
}

func main() {
	err := runCmd(os.Stdout, os.Args)
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			os.Exit(0)
			}
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
	}
}

func runCmd(w io.Writer, args []string) error {
	c, err := parseArgs(w, args[1:])
	if err != nil {
		return err
	}

	if c.outputPath != "" {
		return createHTMLFile(c)
	}

	fmt.Fprintf(w, "Hello %s\n", c.name)
	return nil
}

func parseArgs(w io.Writer, args []string) (config, error) {
	var c config
	fs := flag.NewFlagSet("greeter", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&c.name, "name", "Guest", "имя для приветствия")
	fs.StringVar(&c.outputPath, "o", "", "путь к выходному HTML-файлу")

	err := fs.Parse(args)
	if err != nil {
		return c, err
	}

	if fs.NArg() > 0 {
		return c, errors.New("недопустимые позиционные аргументы")
	}

	return c, nil
}

func createHTMLFile(c config) error {
const tpl = "<h1>Hello {{.Name}}</h1>"
t, err := template.New("webpage").Parse(tpl)
if err != nil {
return err
}

f, err := os.Create(c.outputPath)
if err != nil {
return err
}
defer f.Close()

return t.Execute(f, struct{ Name string }{Name: c.name})
}