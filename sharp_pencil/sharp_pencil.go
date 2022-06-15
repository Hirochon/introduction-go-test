// Package sharp_pencil はシャー芯を管理するパッケージです。
package sharp_pencil

// type ISharpPencil interface {
// 	WriteLetters() (typeOutPencilNum, error) // 外に出ているシャー芯を1つ消費する
// 	PushButton() (typeOutPencilNum, error)   // 中のシャー芯を1つ消費して、外のシャー芯を1つ追加する
// 	RefillLead()                             // 中にシャー芯を5つ追加する
// 	PencilNum() typeOutPencilNum             // 外に出ているシャー芯の数を返す
// }

type sharpPencil struct {
	innerPencilNum int
	outPencilNum   int
}

func (sp *sharpPencil) PencilNum() int {
	return sp.outPencilNum
}

func (sp *sharpPencil) PushButton() (int, error) {
	sp.innerPencilNum -= 1
	sp.outPencilNum += 1
	return sp.outPencilNum, nil
}

func (sp *sharpPencil) WriteLetters() (int, error) {
	sp.outPencilNum -= 1
	return sp.outPencilNum, nil
}

func (sp *sharpPencil) RefillLead() {
	sp.innerPencilNum += 5
}

type ISharpPencil interface {
	WriteLetters() (int, error) // 外に出ているシャー芯を1つ消費する
	PushButton() (int, error)   // 中のシャー芯を1つ消費して、外のシャー芯を1つ追加する
	RefillLead()                // 中にシャー芯を5つ追加する
	PencilNum() int             // 外に出ているシャー芯の数を返す
}

func InitializeSharpPencil() ISharpPencil {
	// a := sharpPencil{pencilNum: 3}
	// a.PencilNum()
	return &sharpPencil{
		innerPencilNum: 5,
		outPencilNum:   0,
	}
}
