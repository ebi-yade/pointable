package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ebi-yade/pointable"
)

var Version = "current"

func init() {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Usage of %s:\n", os.Args[0])
		fmt.Fprintln(w, "pointable [options]")
		fmt.Fprintln(w, "")
		flag.PrintDefaults()
	}
}

func main() {
	showVersion := flag.Bool("version", false, "show the version")
	pkgPath := flag.String("path", "pkg", "the path to package. default: 'pkg'")
	pkgName := flag.String("name", "", "the name of package. default: automatically specified via -path")
	flag.Parse()

	if *showVersion {
		fmt.Println("pointable ", Version)
		return
	}

	dist := pointable.NewDist(*pkgPath)
	if *pkgName != "" {
		dist.SetName(*pkgName)
	}

	if _, err := dist.Validate(); err != nil {
		log.Println("[ERROR] pointable.Config is invalid:", err)
		os.Exit(1)
	}

	templates := []*pointable.Template{
		pointable.NewDefaultTemplate(dist, "builtin", pointable.CommonTemplate, pointable.BuiltinData()),
	}
	for _, tmpl := range templates {
		if err := tmpl.Do(); err != nil {
			log.Println("[ERROR]", err)
			os.Exit(1)
		}
	}
}
