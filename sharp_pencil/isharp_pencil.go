// Package sharp_pencil はシャー芯を管理するパッケージです。
package sharp_pencil

import "fmt"

// ISharpPencil はシャー芯の管理を外部に示すインターフェースです。
type ISharpPencil interface {
	WriteLetters() (typeOutPencilNum, error) // 外に出ているシャー芯を1つ消費する
	PushButton() (typeOutPencilNum, error)   // 中のシャー芯を1つ消費して、外のシャー芯を1つ追加する
	RefillLead()                             // 中にシャー芯を5つ追加する
	PencilNum() typeOutPencilNum             // 外に出ているシャー芯の数を返す
}

// isharpPencil はシャー芯のアドレスを保持する構造体です。
type isharpPencil struct {
	sharpPencil *sharpPencil
}

// WriteLetters は外に出ているシャー芯を1つ消費させる関数です。
func (isp *isharpPencil) WriteLetters() (typeOutPencilNum, error) {
	// 外にシャー芯が出てないと文字は書けない
	if pencilNumIsSmallerThanOne(isp.sharpPencil.outPencilNum) {
		return isp.PencilNum(), fmt.Errorf("シャー芯が出ていません")
	}
	isp.sharpPencil = newSharpPencil(isp.sharpPencil.writeLetters())
	return isp.PencilNum(), nil
}

// PushButton は中のシャー芯を1つ消費して、外のシャー芯を1つ追加する関数です。
func (isp *isharpPencil) PushButton() (typeOutPencilNum, error) {
	// 中にシャー芯がないと外にシャー芯は出せない
	if pencilNumIsSmallerThanOne(isp.sharpPencil.innerPencilNum) {
		return isp.PencilNum(), fmt.Errorf("シャー芯の補充をしてください")
	}
	isp.sharpPencil = newSharpPencil(isp.sharpPencil.pushButton())
	return isp.PencilNum(), nil
}

// RefillLead は中にシャー芯を5つ追加する関数です。
func (isp *isharpPencil) RefillLead() {
	isp.sharpPencil = newSharpPencil(isp.sharpPencil.refillLead())
}

// PencilNum は外に出ているシャー芯の数を返す関数です。
func (isp *isharpPencil) PencilNum() typeOutPencilNum {
	return isp.sharpPencil.outPencilNum
}

// InitializeSharpPencil はシャー芯を初期化する関数です。
func InitializeSharpPencil() ISharpPencil {
	return &isharpPencil{
		sharpPencil: newSharpPencil(0, 0),
	}
}
