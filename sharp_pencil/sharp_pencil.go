// Package sharp_pencil はシャー芯を管理するパッケージです。
package sharp_pencil

import "fmt"

type pencilNum int
type outPencilNum pencilNum
type innerPencilNum pencilNum
type comparablePencilNum interface {
	outPencilNum | innerPencilNum
}

type ISharpPencil interface {
	WriteLetters() (outPencilNum, error) // 外に出ているシャー芯を1つ消費する
	PushButton() (outPencilNum, error)   // 中のシャー芯を1つ消費して、外のシャー芯を1つ追加する
	RefillLead()                         // 中にシャー芯を5つ追加する
	PencilNum() outPencilNum             // 外に出ているシャー芯の数を返す
}

type sharpPencil struct {
	outPencilNum   outPencilNum
	innerPencilNum innerPencilNum
}

type isharpPencil struct {
	sharpPencil *sharpPencil
}

func pencilNumIsSmallerThanOne[T comparablePencilNum](pencilNum T) bool {
	return pencilNum < 1
}

func (p *isharpPencil) WriteLetters() (outPencilNum, error) {
	if pencilNumIsSmallerThanOne(p.sharpPencil.outPencilNum) {
		return p.PencilNum(), fmt.Errorf("シャー芯が出ていません")
	}
	p.sharpPencil = newSharpPencil(p.sharpPencil.outPencilNum-1, p.sharpPencil.innerPencilNum)
	return p.PencilNum(), nil
}

func (p *isharpPencil) PushButton() (outPencilNum, error) {
	if pencilNumIsSmallerThanOne(p.sharpPencil.innerPencilNum) {
		return p.PencilNum(), fmt.Errorf("シャー芯の補充をしてください")
	}
	p.sharpPencil = newSharpPencil(p.sharpPencil.outPencilNum+1, p.sharpPencil.innerPencilNum-1)
	return p.PencilNum(), nil
}

func (p *isharpPencil) RefillLead() {
	p.sharpPencil = newSharpPencil(p.sharpPencil.outPencilNum, p.sharpPencil.innerPencilNum+5)
}

func (p *isharpPencil) PencilNum() outPencilNum {
	return p.sharpPencil.outPencilNum
}

func (p *isharpPencil) innerPencilNum() innerPencilNum {
	return p.sharpPencil.innerPencilNum
}

func InitializeSharpPencil() ISharpPencil {
	return &isharpPencil{
		sharpPencil: newSharpPencil(0, 0),
	}
}

func newSharpPencil(outPencilNum outPencilNum, innerPencilNum innerPencilNum) *sharpPencil {
	return &sharpPencil{
		outPencilNum:   outPencilNum,
		innerPencilNum: innerPencilNum,
	}
}
