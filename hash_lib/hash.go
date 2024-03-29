package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"

	"github.com/spaolacci/murmur3"
)

var str = "hello world"

func md5Hash() [16]byte {
	return md5.Sum([]byte(str))
}

func sha1Hash() [20]byte {
	return sha1.Sum([]byte(str))
}

func murmur32() uint32 {
	return murmur3.Sum32([]byte(str))
}

func murmur64() uint64 {
	return murmur3.Sum64([]byte(str))
}

func main() {
	fmt.Println(
		"md5Hash: ", md5Hash(),
		"\nsha1Hash: ", sha1Hash(),
		"\nmurmur32: ", murmur32(),
		"\nmurmur64: ", murmur64(),
	)
}
