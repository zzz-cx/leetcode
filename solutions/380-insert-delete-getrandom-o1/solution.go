package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	valList []int
	valMap  map[int]int
}

func NewRandomizedSet() RandomizedSet {
	valList := make([]int, 0)
	valMap := make(map[int]int)
	return RandomizedSet{valList: valList, valMap: valMap}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valMap[val]; ok {
		return false
	}
	this.valMap[val] = len(this.valList)
	this.valList = append(this.valList, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.valMap[val]
	if !ok {
		return false
	}
	last := this.valList[len(this.valList)-1]
	this.valList[idx] = last
	this.valMap[last] = idx
	this.valList = this.valList[:len(this.valList)-1]
	delete(this.valMap, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.valList[rand.Intn(len(this.valList))]
}

func main() {
	s := NewRandomizedSet()
	r1 := s.Insert(1)
	r2 := s.Insert(2)
	r3 := s.Insert(1)
	r4 := s.Remove(2)
	r5 := s.Insert(2)
	status := "PASS"
	if !(r1 && r2 && !r3 && r4 && r5) {
		status = "FAIL"
	}
	fmt.Printf("%s | Insert(1)=%v Insert(2)=%v Insert(1)=%v Remove(2)=%v Insert(2)=%v\n", status, r1, r2, r3, r4, r5)
}
