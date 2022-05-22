package sharp_pencil

type typePencilNum int
type typeOutPencilNum typePencilNum
type typeInnerPencilNum typePencilNum
type comparablePencilNum interface {
	typeOutPencilNum | typeInnerPencilNum
}

type sharpPencil struct {
	outPencilNum   typeOutPencilNum
	innerPencilNum typeInnerPencilNum
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

func addOutPencilNumByPushButton(outPencilNum typeOutPencilNum) typeOutPencilNum {
	return outPencilNum + 1
}

func consumeInnerPencilNumByPushButton(innerPencilNum typeInnerPencilNum) typeInnerPencilNum {
	return innerPencilNum - 1
}

func (sp *sharpPencil) pushButton() (typeOutPencilNum, typeInnerPencilNum) {
	return addOutPencilNumByPushButton(sp.outPencilNum), consumeInnerPencilNumByPushButton(sp.innerPencilNum)
}

func addInnerPencilNumByRefillLead(innerPencilNum typeInnerPencilNum) typeInnerPencilNum {
	return innerPencilNum + 5
}

func (sp *sharpPencil) refillLead() (typeOutPencilNum, typeInnerPencilNum) {
	return sp.outPencilNum, addInnerPencilNumByRefillLead(sp.innerPencilNum)
}

func newSharpPencil(outPencilNum typeOutPencilNum, innerPencilNum typeInnerPencilNum) *sharpPencil {
	return &sharpPencil{
		outPencilNum:   outPencilNum,
		innerPencilNum: innerPencilNum,
	}
}
