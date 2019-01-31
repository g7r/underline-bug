package main

type C struct{}

func NewC() C { return C{} }
func (C) Do() {}

func donor() {
	zzz := NewC()
	zzz.Do()
}

func main() {
	zz := NewC()
	zz.Do()
	zz.Do()
}
