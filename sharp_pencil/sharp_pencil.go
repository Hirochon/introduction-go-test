package sharp_pencil

type ISharpPencil interface {
	WriteLetters() (outPencilNum, error)
	PushButton() (outPencilNum, error)
	RefillLead()
	PencilNum() outPencilNum
}
