// Package sharp_pencil はシャー芯を管理するパッケージです。
package sharp_pencil

import "fmt"

type typePencilNum int
type typeOutPencilNum typePencilNum
type typeInnerPencilNum typePencilNum
type comparablePencilNum interface {
	typeOutPencilNum | typeInnerPencilNum
}

type ISharpPencil interface {
	WriteLetters() (typeOutPencilNum, error) // 外に出ているシャー芯を1つ消費する
	PushButton() (typeOutPencilNum, error)   // 中のシャー芯を1つ消費して、外のシャー芯を1つ追加する
	RefillLead()                             // 中にシャー芯を5つ追加する
	PencilNum() typeOutPencilNum             // 外に出ているシャー芯の数を返す
}

type sharpPencil struct {
	outPencilNum   typeOutPencilNum
	innerPencilNum typeInnerPencilNum
}

type isharpPencil struct {
	sharpPencil *sharpPencil
}

func pencilNumIsSmallerThanOne[T comparablePencilNum](pencilNum T) bool {
	return pencilNum < 1
}

func consumeOutPencilNumByWriteLetters(outPencilNum typeOutPencilNum) typeOutPencilNum {
	return outPencilNum - 1
}

func (sp *sharpPencil) writeLetters() (typeOutPencilNum, typeInnerPencilNum) {
	return consumeOutPencilNumByWriteLetters(sp.outPencilNum), sp.innerPencilNum
}

func (isp *isharpPencil) WriteLetters() (typeOutPencilNum, error) {
	if pencilNumIsSmallerThanOne(isp.sharpPencil.outPencilNum) {
		return isp.PencilNum(), fmt.Errorf("シャー芯が出ていません")
	}
	isp.sharpPencil = newSharpPencil(isp.sharpPencil.writeLetters())
	return isp.PencilNum(), nil
}

func addOutPencilNumByPushButton(outPencilNum typeOutPencilNum) typeOutPencilNum {
	return outPencilNum + 1
}

func consumeInnerPencilNumByPushButton(innerPencilNum typeInnerPencilNum) typeInnerPencilNum {
	return innerPencilNum - 1
}

func (sp *sharpPencil) pushButton() (typeOutPencilNum, typeInnerPencilNum) {
	return addOutPencilNumByPushButton(sp.outPencilNum), consumeInnerPencilNumByPushButton(sp.innerPencilNum)
}

func (isp *isharpPencil) PushButton() (typeOutPencilNum, error) {
	if pencilNumIsSmallerThanOne(isp.sharpPencil.innerPencilNum) {
		return isp.PencilNum(), fmt.Errorf("シャー芯の補充をしてください")
	}
	isp.sharpPencil = newSharpPencil(isp.sharpPencil.pushButton())
	return isp.PencilNum(), nil
}

func addInnerPencilNumByRefillLead(innerPencilNum typeInnerPencilNum) typeInnerPencilNum {
	return innerPencilNum + 5
}

func (sp *sharpPencil) refillLead() (typeOutPencilNum, typeInnerPencilNum) {
	return sp.outPencilNum, addInnerPencilNumByRefillLead(sp.innerPencilNum)
}

func (isp *isharpPencil) RefillLead() {
	isp.sharpPencil = newSharpPencil(isp.sharpPencil.refillLead())
}

func (isp *isharpPencil) PencilNum() typeOutPencilNum {
	return isp.sharpPencil.outPencilNum
}

func (isp *isharpPencil) innerPencilNum() typeInnerPencilNum {
	return isp.sharpPencil.innerPencilNum
}

func InitializeSharpPencil() ISharpPencil {
	return &isharpPencil{
		sharpPencil: newSharpPencil(0, 0),
	}
}

func newSharpPencil(outPencilNum typeOutPencilNum, innerPencilNum typeInnerPencilNum) *sharpPencil {
	return &sharpPencil{
		outPencilNum:   outPencilNum,
		innerPencilNum: innerPencilNum,
	}
}
