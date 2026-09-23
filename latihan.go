package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main() {
	akun_terdaftar := map[string]string{
		"Fazril": "rahasia",
		"asep": "123",
		"alif": "qwerty",
		"amba": "sad",
	}
	var umur int
	
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan nama anda : ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)
	
	fmt.Print("Masukan umur anda : ")
	umurStr, _ := reader.ReadString('\n')
	umur, err := strconv.Atoi(strings.TrimSpace(umurStr))
	if err != nil {
		fmt.Println("Umur harus berupa angka")
		return
	}
	
	if umur >= 17 {
		
		for {
			var punya_akun string
			fmt.Print("Apakah kamu sudah memiliki akun y/t : ")
			fmt.Scanln(&punya_akun)
			
			if punya_akun == "y" {
				var Username string
				var Password string
				fmt.Println("Anda sudah mencukupi umur", nama)
				
				fmt.Print("Masukan Username anda : ")
				fmt.Scanln(&Username)
				fmt.Print("Masukan Password anda : ")
				fmt.Scanln(&Password)		
				
				pass, ada:= akun_terdaftar[Username]
				 if ada && pass == Password {
					 fmt.Println("Login berhasil", Username)
					 break
				 } else {
					 fmt.Println("Username atau Password salah.")
					 continue
				 }
			} else if punya_akun == "t" {
				var buat_akun string
				fmt.Print("Apakah anda ingin membuat akun y/t : ")
				fmt.Scanln(&buat_akun)
				
				if buat_akun == "y" {
					var userbaru string
					var passbaru string
					fmt.Print("Masukan Username akun baru anda : ")
					fmt.Scanln(&userbaru)
					fmt.Print("Masukan Password akun baru anda : ")
					fmt.Scanln(&passbaru)
					
					akun_terdaftar[userbaru] = passbaru
					fmt.Println("akun berhasil dibuat! Silahkan coba login lagi.")
				} else {
					fmt.Println("Terima kasih sudah datang")
					break
				}
			}
		}
	} else {
		fmt.Println("Anda belum mecukupi umur", nama)
		fmt.Println("Maaf anda belum bisa menggunakan aplikasi ini")
	}
}