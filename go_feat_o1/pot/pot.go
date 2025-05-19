package pot

import (
	"fmt"

	"github.com/TaiChi112/Go/go_feat_o1/stove"
)

type Pot struct {
	Price       float32
	Width       int
	Height      int
	Stove       *stove.Stove
	AmountStove int
}

func (p *Pot) New_Streaming_Pot(Price float32, Width int, Height int, Stove *stove.Stove, AmountStove int) *Pot {
	return &Pot{
		Price:       Price,
		Width:       Width,
		Height:      Height,
		Stove:       Stove,
		AmountStove: AmountStove,
	}
}

func (p *Pot) Show_Pot_Info() {
	fmt.Printf("Price: %.2f\nWidth: %d\nHeight: %d\nStove Type: %s\n", p.Price, p.Width, p.Height, fmt.Sprintf("%d", p.Stove.Type))
}
func (p *Pot) SetPrice(price float32) {
	p.Price = price
}
func (p *Pot) SetWidth(width int) {
	p.Width = width
}
func (p *Pot) SetHeight(height int) {
	p.Height = height
}
func (p *Pot) SetStove(stove *stove.Stove) {
	p.Stove = stove
}
func (p *Pot) GetPrice() float32 {
	return p.Price
}
func (p *Pot) GetWidth() int {
	return p.Width
}
func (p *Pot) GetHeight() int {
	return p.Height
}
func (p *Pot) GetStove() *stove.Stove {
	return p.Stove
}
func (p *Pot) GetAmountStove() int {
	return p.AmountStove
}
