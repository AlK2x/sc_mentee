package main

import (
	"crypto/rand"
	"fmt"
)

type LargeObject struct {
	Data []byte
	Size int
}

func NewLargeObject() *LargeObject {
	size := 10000000
	data := make([]byte, size)
	rand.Read(data)

	return &LargeObject{
		Data: data,
		Size: len(data),
	}
}

var largeObjects []*LargeObject
var bigMap map[int]string

func PrintLargeObject(lb LargeObject) {
	fmt.Println(lb.Size)
}

func fillGlobalObjects() {
	for i := range 10000 {
		obj := NewLargeObject()
		largeObjects = append(largeObjects, obj)
		bigMap[i] = string(obj.Data)
	}
}

func main() {
	lb := NewLargeObject()
	PrintLargeObject(*lb)
}
