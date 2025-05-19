package stove

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

func (s *Stove) Set_Stove_Type(stoveType Stove_Type) {
	s.Type = stoveType
}
func (s *Stove) Get_Stove_Type() Stove_Type {
	return s.Type
}
func (s *Stove) Show_Stove_Info() {
	fmt.Println("Stove Type:", s.Type.to_String())
}
