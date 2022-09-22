package main

import "fmt"

type FlyBehavior interface {
	fly()
}

type QuackBehavior interface {
	quack()
}

type Duck struct {
	display       func()
	FlyBehavior   FlyBehavior
	QuackBehavior QuackBehavior
}

func (d *Duck) PerformFly() {
	d.FlyBehavior.fly()
}

type Fly struct {
}

func (f Fly) fly() {
	fmt.Println("I'm flying")
}

type Nofly struct{}

func (nf Nofly) fly() {
	fmt.Println("Cannot fly")
}

type Squeak struct{}

func (s Squeak) quack() {
	fmt.Println("squeaking")
}

type Mute struct {
}

func (m Mute) quack() {
	fmt.Println("No sound")
}

type Quack struct {
}

func (q Quack) quack() {
	fmt.Println("Quacking")
}

func (d *Duck) PerformQuack() {
	d.QuackBehavior.quack()
}

type Mallard struct {
	*Duck
}

func NewMallard() *Mallard {
	d := &Duck{
		display: func() {
			fmt.Println("Mallard duck")
		},
		QuackBehavior: &Quack{},
		FlyBehavior:   &Fly{},
	}
	return &Mallard{d}
}

type DuckDuckGo struct {
	*Duck
}

func NewDuckDuckGo() *DuckDuckGo {
	d := &Duck{
		display: func() {
			fmt.Println("DuckDuckGo poiskovik")
		},
		QuackBehavior: &Mute{},
		FlyBehavior:   &Nofly{},
	}
	return &DuckDuckGo{d}
}

func main() {
	myMallard := NewMallard()
	myMallard.display()
	myMallard.PerformFly()
	myMallard.PerformQuack()

	myDuckDuckGo := NewDuckDuckGo()
	myDuckDuckGo.display()
	myDuckDuckGo.PerformFly()
	myDuckDuckGo.PerformQuack()
}
