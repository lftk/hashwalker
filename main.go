package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

var (
	dir    = flag.String("dir", "", "walks the file tree rooted at dir")
	ignore = flag.String("ignore", "", "ignore sub folder or file")
	out    = flag.String("out", "", "save result to out file")
)

func main() {
	flag.Parse()
	if *dir == "" {
		*dir = filepath.Dir(os.Args[0])
	}
	writer := os.Stdout
	if *out != "" {
		f, err := os.Create(*out)
		if err != nil {
			log.Fatal(err)
		}
		writer = f
	}
	ignores := strings.Split(*ignore, ",")

	runtime.GOMAXPROCS(runtime.NumCPU())

	in := make(chan *file, 1024)
	out := make(chan *file, 1024)
	wg := new(sync.WaitGroup)
	done := make(chan interface{})

	// write result to local file
	go func() {
		defer func() {
			if writer != os.Stdout {
				writer.Close()
			}
			close(done)
		}()

		err := write(writer, out)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// calc hash for file
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			err := hash(out, in)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	// walks file
	err := walk(*dir, ignores, in)
	if err != nil {
		log.Fatal(err)
	}
	close(in)

	wg.Wait()
	close(out)
	<-done
}
