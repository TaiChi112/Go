package noodle_pot

import (
	"fmt"

	"github.com/TaiChi112/Go/go_feat_o1/pot"
)

type Noodle_Pot struct {
	Pot         pot.Pot
	Compartment int
}

func (np *Noodle_Pot) GetCompartment() int {
	return np.Compartment
}
func (np *Noodle_Pot) GetPot() pot.Pot {
	return np.Pot
}
func (np *Noodle_Pot) Show_Noodle_Pot_Info() {
	fmt.Println("Noodle Pot Info:")
	np.Pot.Show_Pot_Info()
	fmt.Println("Compartment:", np.Compartment)
}
