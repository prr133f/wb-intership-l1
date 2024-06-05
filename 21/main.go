package main

import "fmt"

type USBReader interface {
	ReadUSB()
}

type MemoryCard struct {
	Name     string
	Capasity int
}

func (m MemoryCard) ReadCard() {
	fmt.Printf("Reading %s\n", m.Name)
}

type CardAdapter struct {
	MemoryCard
}

func (ca CardAdapter) ReadUSB() {
	fmt.Print("Reading via adapter ")
	ca.ReadCard()
}

func main() {
	mc := MemoryCard{Name: "card"}
	mc.ReadCard()

	ca := CardAdapter{MemoryCard: mc}
	ca.ReadUSB()
}
