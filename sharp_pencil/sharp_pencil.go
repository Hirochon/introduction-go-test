// Package sharp_pencil はシャー芯を管理するパッケージです。
package sharp_pencil

// type ISharpPencil interface {
// 	WriteLetters() (typeOutPencilNum, error) // 外に出ているシャー芯を1つ消費する
// 	PushButton() (typeOutPencilNum, error)   // 中のシャー芯を1つ消費して、外のシャー芯を1つ追加する
// 	RefillLead()                             // 中にシャー芯を5つ追加する
// 	PencilNum() typeOutPencilNum             // 外に出ているシャー芯の数を返す
// }

type sharpPencil struct {
	pencilNum int
}

func (sp sharpPencil) PencilNum() int {
	return sp.pencilNum
}

type ISharpPencil interface {
// 	WriteLetters() (typeOutPencilNum, error) // 外に出ているシャー芯を1つ消費する
// 	PushButton() (typeOutPencilNum, error)   // 中のシャー芯を1つ消費して、外のシャー芯を1つ追加する
// 	RefillLead()                             // 中にシャー芯を5つ追加する
	PencilNum() int            				 // 外に出ているシャー芯の数を返す
}

func InitializeSharpPencil() ISharpPencil {
	// a := sharpPencil{pencilNum: 3}
	// a.PencilNum()
	return sharpPencil{
		pencilNum: 5,
	}
}