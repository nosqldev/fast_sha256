package main

import (
    "fmt"
    "crypto/sha256"
)

const HashBSize = sha256.Size

func DoubleHashB(b []byte) []byte {
    first := sha256.Sum256(b)
    second := sha256.Sum256(first[:])
    return second[:]
}

func main() {
    s := "abc"
    var hash_sha_256 []byte

    for i:=0; i<10000000; i++ {
        hash_sha_256 = DoubleHashB([]byte(s))
    }

    fmt.Printf("%x\n", hash_sha_256)
}
