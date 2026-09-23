package main
import "fmt" 

func main() {
	var panjang_kotak int
	fmt.Println("---------------------------------------")
	fmt.Println("MENGHITUNG PANJANG PITA YANG DIBUTUHKAN")
	fmt.Println("---------------------------------------")
	fmt.Print("Masukan panjang kotak (cm) : ")
	fmt.Scanln(&panjang_kotak)
	pita_keliling := panjang_kotak*4
	pita_total := (panjang_kotak*4)+20
	fmt.Println("pita keliling 	=", pita_keliling, "cm")
	fmt.Println("pita total 	=", pita_total, "cm")
}