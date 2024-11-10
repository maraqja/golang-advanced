package main

type User struct {
	Name string
}

//  go run -gcflags '-m -l'  main.go - для просмотра как используется память (какие переменные в heap, а какие в stack)

func main() {
	age := getAge()
	/* переменная age будет в heap:
	она ссылается на переменную, которая храниися в stack frame getAge, а не в main
	а stack pointer после вызова функции будет на main указывать
	*/
	canDrink(age) /*
		вызов canDrink полностью замещает stack frame getAge, поэтому ссылки на age не будет
		то есть из-за того что текущий stack не может больше хранить переменную из верхнего stack frame, тк ей надо вызвать новый stack frame
		GO решает хранить ее в HEAP

		все на что есть референс уходит на heap (чтобы на него можно было дальше ссылаться) при вызове нового stack frame
	*/
}

func canDrink(age *int) bool {
	return *age >= 18
}

func getAge() *int {
	age := 18
	return &age
}