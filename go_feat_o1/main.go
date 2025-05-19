package main

import "fmt"

type Stove_Type int

const (
	Stove_Gas Stove_Type = iota
	Stove_Electric
	Stove_Induction
	Stove_Infrared
)

func (s Stove_Type) to_String() string {
	switch s {
	case Stove_Gas:
		return "Gas"
	case Stove_Electric:
		return "Electric"
	case Stove_Induction:
		return "Induction"
	case Stove_Infrared:
		return "Infrared"
	default:
		return "Unknown"
	}
}

type Stove struct {
	Type Stove_Type
}

type Pot struct {
	price        float32
	width        int
	height       int
	stove        *Stove
	amount_stove int
}

func (p *Pot) Show_Pot_Info() {
	if p.stove == nil {
		fmt.Println("Pot Info: No stove attached.")
		return
	}
	fmt.Printf("Price: %.2f\nWidth: %d\nHeight: %d\nStove Type: %s\n", p.price, p.width, p.height, p.stove.Type.to_String())
}
func (p *Pot) GetPrice() float32 {
	return p.price
}
func (p *Pot) GetWidth() int {
	return p.width
}
func (p *Pot) GetHeight() int {
	return p.height
}
func (p *Pot) GetStove() *Stove {
	return p.stove
}
func (p *Pot) GetAmountStove() int {
	return p.amount_stove
}

type Streaming_Pot struct {
	pot   Pot
	layer int
}

func (sp *Streaming_Pot) GetLayer() int {
	return sp.layer
}
func (sp *Streaming_Pot) GetPot() Pot {
	return sp.pot
}
func (sp *Streaming_Pot) Show_Streaming_Pot_Info() {
	fmt.Println("Streaming Pot Info:")
	sp.pot.Show_Pot_Info()
	fmt.Println("Layer:", sp.layer)
}

type Noodle_Pot struct {
	pot         Pot
	compartment int
}

func (np *Noodle_Pot) GetCompartment() int {
	return np.compartment
}
func (np *Noodle_Pot) GetPot() Pot {
	return np.pot
}
func (np *Noodle_Pot) Show_Noodle_Pot_Info() {
	fmt.Println("Noodle Pot Info:")
	np.pot.Show_Pot_Info()
	fmt.Println("Compartment:", np.compartment)
}

func main() {
	stove := Stove{Type: Stove_Gas}
	streamingPot := Streaming_Pot{
		pot: Pot{
			price:        39.99,
			width:        25,
			height:       20,
			stove:        &stove,
			amount_stove: 1,
		},
		layer: 2,
	}
	streamingPot.Show_Streaming_Pot_Info()
	println()
	noodlePot := Noodle_Pot{
		pot: Pot{
			price:        49.99,
			width:        30,
			height:       25,
			stove:        &stove,
			amount_stove: 1,
		},
		compartment: 3,
	}
	noodlePot.Show_Noodle_Pot_Info()
}
