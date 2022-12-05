package main

import (
	f "exercise3/pkg"
	"flag"
)

func main() {
	k := flag.Int("k", 1, "колонка для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "обратная сортировка")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()
	fileData := f.NewFileMap()
	fileData.FillFileMap(flag.Arg(0))
	fileData.Sort(*k, *n, *r, *u)
}
