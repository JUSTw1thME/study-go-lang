package main
import "fmt"	

func main() {
	var x, y, z, angka int
	fmt.Print("Masukan angka: ")
	fmt.Scanln(&angka)
	x = angka / 100
	y = angka / 10 % 10
	z = angka % 10
	fmt.Println(x, y, z)
}