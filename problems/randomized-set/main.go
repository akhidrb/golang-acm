package main

import "math/rand"

type RandomizedSet struct {
	set []int
}

func Constructor() RandomizedSet {
	set := make([]int, 0)
	return RandomizedSet{set: set}
}

func (this *RandomizedSet) Insert(val int) bool {
	for _, v := range this.set {
		if v == val {
			return false
		}
	}
	this.set = append(this.set, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	for i, v := range this.set {
		if v == val {
			this.set = append(this.set[:i], this.set[i+1:]...)
			return true
		}
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	length := len(this.set)
	randNum := rand.Intn(length)
	return this.set[randNum]
}
