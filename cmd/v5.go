package main

import (
	"flag"
	"fmt"
	v5 "github.com/tiyee/sqlgen/v5"
	"io"
	"os"
)

func main() {
	in := flag.String("in", "", "create table sql file")
	out := flag.String("out", "", "go struct save file")
	pkg := flag.String("pkg", "", "go package name")

	flag.Parse()
	if *in == "" || *pkg == "" {
		panic(" arg error")
	}
	f, err := os.OpenFile(*in, os.O_RDONLY, 644)
	if err != nil {
		panic(err)
	}
	bs, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var wr io.Writer
	if *out == "" {
		wr = os.Stdout
	} else {

		f, err := os.OpenFile(*out, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("openFile err: ", err.Error())
		}
		defer f.Close()
		wr = f
	}

	creator := v5.ParseSql(string(bs))

	if err := creator.Save(wr); err != nil {
		fmt.Println(err.Error())
	}

}
