package main

import (
	"fmt"
	"math"
	"strings"
)

type shape interface {
	luas() float32
	keliling() float32
}

type persegi struct {
	sisi float32
}

func newPersegi(sisi float32) shape {
	return &persegi{
		sisi: sisi,
	}
}

func (p *persegi) luas() float32 {
	return p.sisi * p.sisi
}

func (p *persegi) keliling() float32 {
	return 4 * p.sisi
}

type lingkaran struct {
	diameter float32
}

func newLingkaran(diameter float32) shape {
	return &lingkaran{diameter: diameter}
}

func (p *lingkaran) display() {
	fmt.Println(p.luas())
}

func (p *lingkaran) luas() float32 {
	jari := p.diameter / 2
	return math.Pi * jari * jari
}

func (p *lingkaran) keliling() float32 {
	return math.Pi * p.diameter
}

func main() {
	p := newPersegi(30)
	l := newLingkaran(21)

	display(p, "persegi")
	display(l, "lingkaran")
}

func display(obj shape, nama string) {
	fmt.Println(strings.Repeat("=", 20))
	fmt.Println("Object :", nama)
	fmt.Println("Luas :", obj.luas())
	fmt.Println("Keliling :", obj.keliling())
}
