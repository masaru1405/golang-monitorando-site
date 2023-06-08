package main
import "fmt"
import "reflect"

func main(){
	frutas := []string{"banana", "limão", "laranja"}
	fmt.Println("Length:", len(frutas)) // 3
	fmt.Println("Capacity:", cap(frutas))// 3
	frutas = append(frutas, "maçã")
	fmt.Println(frutas)
	fmt.Println(reflect.TypeOf(frutas)) // []string
	fmt.Println("Length:", len(frutas)) // 4
	fmt.Println("Capacity:", cap(frutas)) // dobra a capacidade: 6
}