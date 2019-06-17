package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// hash is a perfect hash function for keywords.
// It assumes that s has at least length 2.
func hash(s []byte) uint {
	return (uint(s[0])<<4 ^ uint(s[1]) + uint(len(s))) & uint(len(keywordMap)-1)
}

var keywordMap [1 << 6]uint // size must be power of two

var wwww = []string{
	"w",
	"W",
}

var m = map[string]struct{}{}

var tok uint

func ww() string {
LOOP:
	for {
		var wwwww string
		for i := 0; i < 2; i++ {
			wwwww += wwww[rand.Int()%2]
		}
		for {
			r := rand.Int() % 2
			wwwww += wwww[r]
			h := hash([]byte(
				wwwww,
			))
			if keywordMap[h] == 0 {
				if _, ok := m[wwwww]; !ok {
					m[wwwww] = struct{}{}
					keywordMap[h] = tok
					tok++
					return wwwww
				}
				continue LOOP

			}
		}
	}
}

func main() {
	for i := 0; i < 25; i++ {
		fmt.Println(ww())
	}
}
