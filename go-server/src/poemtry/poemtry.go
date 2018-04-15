package poemtry

import (
       "bufio"
       "fmt"
       "log"
       "strings"
       "os"
)

type Poem struct {
     Name string
     Author string
     Content string
}

func parsePoem(s string) Poem {
     fields := strings.Split(s, "-")
     if ( len(fields) >= 3) {
     	return Poem {
	       Name: fields[0],
	       Author: fields[1],
	       Content: fields[2],
	}
     }
     return Poem {}
}

func Generator(file_path string) chan Poem {
     c := make(chan Poem, 4)

     go func() {
    	src, err := os.Open(file_path)
	if err != nil {
     	   log.Printf("cannot open file %s, error %v", file_path, err)
	}
	defer src.Close()
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
     	    line := scanner.Text()
	    p := parsePoem(line)
	    print_poem(p)
	    c <- p
     	}
	close(c)
     }()
     return c
}

func print_poem(p Poem) {
     log.Println(fmt.Sprintf("poem name %s, author %s, content %s",
		       p.Name, p.Author, p.Content))
}
