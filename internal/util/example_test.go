package util

import "fmt"

func ExampleShortenURL() {
	short := ShortenURL("http://localhost:8080/tests")
	fmt.Println(short)
}
