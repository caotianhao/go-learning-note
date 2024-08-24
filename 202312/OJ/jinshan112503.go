package main

import (
	"bufio"
	"container/heap"
	"container/list"
	"crypto"
	"crypto/aes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/ascii85"
	"encoding/hex"
	"fmt"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"math"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

type j11 struct {
	sort.IntSlice
}

func (g *j11) Push(x any) {
	g.IntSlice = append(g.IntSlice, x.(int))
}

func (g *j11) Pop() any {
	qu := g.IntSlice
	n := len(qu)
	o := qu[n-1]
	qu = qu[:n-1]
	return o
}

func main() {
	a := http.StatusOK
	b := math.MaxInt16
	c := ""
	d := bufio.NewScanner(os.Stdin)
	d.Scan()
	sp := strings.Split(d.Text(), "&")
	e := make([]int, 0)
	for _, v := range sp {
		t, _ := strconv.Atoi(v)
		e = append(e, t)
	}
	fr := list.List{}
	g := &j11{}
	heap.Init(g)
	h := sha256.BlockSize
	h2 := sha512.BlockSize
	h3 := sha1.Size
	h4 := md5.Size
	fmt.Println(a+b, c, e, fr, h, h2, h3, h4)
	h5 := syscall.CRYPT_NEWKEYSET
	h6 := crypto.SHA3_224
	h7 := aes.BlockSize
	h8 := reflect.DeepEqual(h5, h6)
	h9 := rand.Float32()
	h10 := ascii85.Encode(nil, nil)
	h11 := hex.EncodeToString(nil)
	h12 := utf16.Encode(nil)
	fmt.Println(h7, h8, h9, h10, h11, h12)
	fmt.Println(unsafe.Sizeof(12))
	h13 := crc32.Size + crc64.ISO + adler32.Size
	fmt.Println(h13)
}
