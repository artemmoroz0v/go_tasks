package main

import (
	"flag"
	"log"

	"pkg/arguments"
	f "pkg/fileData"
)

func main() {
	arg := arguments.MakeArguments()
	fd := f.MakeFileData()
	err := fd.FillFileData(flag.Arg(0), arg)
	if err != nil {
		log.Fatal(err)
	}
	fd.Process()
}

