package main

import (
	"crypto/rand"
)

func NewLargeObject(size int) *LargeObject {
	lo := &LargeObject{
		Size: size,
	}
	lo.Data = lo.generateRundomData()

	return lo
}

type LargeObject struct {
	Data []byte
	Size int
}

func (lo *LargeObject) Renew() {
	newData := lo.generateRundomData()
	copy(lo.Data, newData)
}

func (lo *LargeObject) DataUpdaterFunc(val byte) func() {
	resetFn := func() {
		for i := range lo.Data {
			lo.Data[i] ^= val
		}
	}
	return resetFn
}

func (lo *LargeObject) generateRundomData() []byte {
	data := make([]byte, lo.Size)
	rand.Read(data)
	return data
}

var largeObjects []*LargeObject
var bigMap map[int]*LargeObject

func fillGlobalObjects() {
	for i := range 10000 {
		obj := NewLargeObject(10)
		largeObjects = append(largeObjects, obj)
		bigMap[i] = obj
	}
}

func main() {
	bigMap = make(map[int]*LargeObject)
	largeObjects = make([]*LargeObject, 0)
	fillGlobalObjects()
	obj := NewLargeObject(1000000)
	for range 1000 {
		obj.Renew()
		updData := obj.DataUpdaterFunc(byte(0))
		updData()
	}
}
