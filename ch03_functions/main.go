package main

import "fmt"

func concat(s1, s2) string {
	return s1 + s2
}

// don't touch below this line

func main() {
	test("Lane,", " happy birthday!")
	test("Zuck,", " hope that Metaverse thing works out")
	test("Go", " is fantastic")
}

func test(s1 string, s3 string) {
	fmt.Println(concat(s1, s2))
}
