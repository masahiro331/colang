package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint
}

const intSize = 32 << (^uint(0) >> 63)

func (s *IntSet) Has(x int) bool {
	word, bit := x/intSize, uint(x%intSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/intSize, uint(x%intSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}

	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < intSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", intSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	len := 0
	for _, word := range s.words {
		len += popcount(uint64(word))
	}
	return len
}

func popcount(x uint64) int {
	x = (x & 0x5555555555555555) + ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x & 0x0F0F0F0F0F0F0F0F) + ((x >> 4) & 0x0F0F0F0F0F0F0F0F)
	x = (x & 0x00FF00FF00FF00FF) + ((x >> 8) & 0x00FF00FF00FF00FF)
	x = (x & 0x0000FFFF0000FFFF) + ((x >> 16) & 0x0000FFFF0000FFFF)
	x = (x & 0x00000000FFFFFFFF) + ((x >> 32) & 0x00000000FFFFFFFF)
	return int(x)
}

func (s *IntSet) Remove(x int) {
	word, bit := x/intSize, uint(x%intSize)
	if word < len(s.words) {
		s.words[word] &= ^(1 << bit)
	}
}

func (s *IntSet) Clear() {
	for i, _ := range s.words {
		s.words[i] &= 0
	}
}

func (s *IntSet) Copy() *IntSet {
	var dst IntSet
	for _, tword := range s.words {
		dst.words = append(dst.words, tword)
	}
	return &dst
}

func (s *IntSet) AddAll(args ...int) {
	for _, arg := range args {
		s.Add(arg)
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, 0)
		}

	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

func (s *IntSet) SynmetricWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}

	}
}
func (s *IntSet) Elems() []uint {
	ret := []uint{}
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < intSize; j++ {
			if word&(1<<uint(j)) != 0 {
				ret = append(ret, uint(intSize*i+j))
			}
		}
	}
	return ret
}
func main() {
	var x, y IntSet
	x.AddAll([]int{1, 63, 144, 9}...)
	x.Remove(1)
	fmt.Println("remove 1 from x:{1, 63, 144, 9}")
	fmt.Println(x.String())
	fmt.Println(x.Elems())
	x.Clear()
	fmt.Println("clear x")
	fmt.Println(x.String())
	y.Add(9)
	y.Add(42)
	fmt.Println("y:{9, 42}")
	fmt.Println(y.String())
	fmt.Println("y length")
	fmt.Println(y.Len())
	x.Add(9)
	z := x.Copy()
	fmt.Println(z.String())
	fmt.Println("z:copy x:{9}")
	z = y.Copy()
	fmt.Println("z:copy y:{9, 42}")
	fmt.Println(z.String())

	x.IntersectWith(&y)
	fmt.Println("x and y intersect :{9}")
	fmt.Println(x.String())
	fmt.Println(x.Len())
	x.SynmetricWith(&y)
	fmt.Println("x and y synmetric :{42}")
	fmt.Println(x.String())

	y.SynmetricWith(&x)
	fmt.Println("y and x difference :{9}")
	fmt.Println(y.String())
	fmt.Println(x.Has(9), x.Has(123))
}
