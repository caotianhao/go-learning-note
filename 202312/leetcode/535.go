package main

import "fmt"

type Codec struct {
	url string
}

func Constructor535() Codec {
	return Codec{url: ""}
}

// Encodes a URL to a shortened URL.
func (code *Codec) encode(longUrl string) string {
	return longUrl + "a"
}

// Decodes a shortened URL to its original URL.
func (code *Codec) decode(shortUrl string) string {
	return shortUrl[:len(shortUrl)-1]
}

func main() {
	str := Constructor535()
	s := str.encode("https://www.baidu.com")
	fmt.Println(s)
	s1 := str.decode(s)
	fmt.Println(s1)
}
