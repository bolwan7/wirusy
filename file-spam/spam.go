package main

import (
	"flag"
	"math/rand"
	"os"
)

func main(){
	n := flag.Int("n", 0, "number of folders to generate")
	p := flag.String("p", "mocny przekaz", "jaki przekaz wariacie")
	flag.Parse()
	flag.Usage()

	var p2 string = *p
	for b:=0;b<1000;b++{
		p2 = p2 + *p + "\n"
	}
	f := ""
	charset := "abcdefghijklmnopqrstuvwxyz"
			for y := 0; y<*n; y++{
				for i := 0; i<5; i++{
					x := string(charset[rand.Intn(26)])
					f = f+x
				}
				f = f + ".txt"
				os.WriteFile(f, []byte(p2), 0644)
				f=""
			}
}