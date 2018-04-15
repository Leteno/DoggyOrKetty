package poemtry

import (
       "testing"
)

func init() {
}

func TestPoemtry(t *testing.T) {
     wait := make(chan int)
     c := generator("poem.txt")
     go func() {
     	for poem := range c {
	    print_poem(poem)
	}
	wait <- 1
     }()
     <- wait
}