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
		fmt.Fprintln(w, "Usage: pointable [FLAGS]")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "Flags:")
		flag.PrintDefaults()
	}
}

func main() {
	showVersion := flag.Bool("version", false, "show the version")
	pkgPath := flag.String("path", "ptr", "the path to package")
	pkgName := flag.String("name", "", "the name of package (default: automatically specified via -path)")
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
		pointable.NewDefaultTemplate(dist, "builtin", pointable.CommonTemplate, pointable.BuiltinTypes()),
	}
	for _, tmpl := range templates {
		if err := tmpl.Do(); err != nil {
			log.Println("[ERROR]", err)
			os.Exit(1)
		}
	}
}
