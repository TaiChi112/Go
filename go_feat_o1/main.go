package main

import (
	"github.com/TaiChi112/Go/go_feat_o1/noodle_pot"
	"github.com/TaiChi112/Go/go_feat_o1/pot"
	"github.com/TaiChi112/Go/go_feat_o1/stove"
	"github.com/TaiChi112/Go/go_feat_o1/streaming_pot"
)

func main() {
	stove := []stove.Stove{
		{Type: stove.Stove_Gas},
		{Type: stove.Stove_Electric},
		{Type: stove.Stove_Induction},
		{Type: stove.Stove_Infrared},
	}
	sp := &streaming_pot.Streaming_Pot{
		Pot: pot.Pot{
			Price:       100.0,
			Width:       30,
			Height:      20,
			Stove:       &stove[0], // Using the first stove as an example
			AmountStove: 1,
		},
		Layer: 3,
	}
	sp.Show_Streaming_Pot_Info()

	np := &noodle_pot.Noodle_Pot{
		Pot: pot.Pot{
			Price:       80.0,
			Width:       25,
			Height:      15,
			Stove:       &stove[1], // Using the second stove as an example
			AmountStove: 1,
		},
		Compartment: 2,
	}
	np.Show_Noodle_Pot_Info()
}
