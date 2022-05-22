package main

import (
	"fmt"
	"github.com/Hirochon/introduction-go-test/error_handling"
	"github.com/Hirochon/introduction-go-test/sharp_pencil"
)

func main() {
	println(sum(2, 2))
	if err := error_handling.RequestValidation(500); err != nil {
		println(err.Error())
	}
	sharpenKatikati()
}

func sum(a, b int) int {
	return a * b
}

func sharpenKatikati() {
	sharpPencil := sharp_pencil.InitializeSharpPencil()
	fmt.Println("出ているシャー芯", sharpPencil.PencilNum())

	func() {
		num, err := sharpPencil.PushButton()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("シャーペンカチカチ: 出ているシャー芯の長さ", num)
	}()

	func() {
		num, err := sharpPencil.PushButton()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("シャーペンカチカチ: 出ているシャー芯の長さ", num)
	}()

	func() {
		num, err := sharpPencil.WriteLetters()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("文字を書き書き: 出ているシャー芯の長さ", num)
	}()

	sharpPencil.RefillLead()
	fmt.Println("シャー芯を補充！")

	func() {
		num, err := sharpPencil.PushButton()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("シャーペンカチカチ: 出ているシャー芯の長さ", num)
	}()

	func() {
		num, err := sharpPencil.PushButton()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("シャーペンカチカチ: 出ているシャー芯の長さ", num)
	}()

	func() {
		num, err := sharpPencil.WriteLetters()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("文字を書き書き: 出ているシャー芯の長さ", num)
	}()
}
