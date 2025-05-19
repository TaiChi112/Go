package streaming_pot

import (
	"fmt"

	"github.com/TaiChi112/Go/go_feat_o1/pot"
)

type Streaming_Pot struct {
	Pot   pot.Pot
	Layer int
}

func (sp *Streaming_Pot) GetLayer() int {
	return sp.Layer
}
func (sp *Streaming_Pot) GetPot() pot.Pot {
	return sp.Pot
}
func (sp *Streaming_Pot) Show_Streaming_Pot_Info() {
	fmt.Println("Streaming Pot Info:")
	sp.Pot.Show_Pot_Info()
	fmt.Println("Layer:", sp.Layer)
}
