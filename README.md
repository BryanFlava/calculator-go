# Membuat Kalkulator Sederhana Menggunakan Golang
- Bryan Samperura a.k.a [[BryanFlava](https://github.com/BryanFlava)]

![Golang Logo](https://savvycomsoftware.com/wp-content/uploads/2021/06/hire-golang-developer-savvycom-12.jpg)

Ini adalah README.md yang menjelaskan cara membuat kalkulator sederhana menggunakan bahasa pemrograman Golang. Kalkulator akan menerima input dari pengguna dan menghasilkan hasil perhitungan matematika sederhana. Berikut adalah langkah-langkah untuk membangun kalkulator menggunakan Golang.

## Tools yang Diperlukan
1. Golang: Pastikan Anda memiliki Golang terinstal di komputer Anda. Anda dapat mengunduhnya dari [situs web resmi Golang](https://golang.org/dl/).

## Fitur
Kalkulator yang akan kita bangun akan memiliki fitur-fitur berikut:
- Penjumlahan
- Pengurangan
- Perkalian
- Pembagian

## Langkah-langkah Pembuatan
1. Buatlah file baru dengan ekstensi `.go`, misalnya `main.go`, untuk menulis kode program Golang kita.

2. Lalu sebelum ke tahap selanjutnya buatkan terlebih dahulu modulenya di terminal dengan cara `go mod init <nama-module>`, misalnya `go mod init bryan-calculator`, dan juga setelah membuat module jangan lupa melakukan `go mod tidy` pada module yang tadi dibuat.
   
3. Impor paket yang diperlukan. Kita akan menggunakan paket `fmt` untuk input/output. Tulislah kode berikut di awal file `main.go`:
    ```go
    package main

    import (
        "fmt"
    )
    ```

4. Buat fungsi `main` sebagai fungsi utama program kita. Tulislah kode berikut di bawah impor paket:
    ```go
    func main() {
        // Kode program akan ditulis di sini
    }
    ```
5. Buat variabel untuk menyimpan input dari pengguna. Tulislah kode berikut di dalam fungsi `main`:
    ```
	var x, y int
	var operation string

	fmt.Println("Enter the first numbers:")
	fmt.Scanf("%d %d", &y)

	fmt.Println("Enter the seconds numbers:")
	fmt.Scanf("%d %d", &x)

	fmt.Println("Enter an operation (+, -, *, /):")
	fmt.Scanf("%s", &operation)
    ```

6. Buat fungsi `swicth operation` untuk melakukan perhitungan berdasarkan operator yang dimasukkan oleh pengguna dan sekaligus memprint hasil dari result yang datanya disimpan dalam variable. Tulislah kode berikut di bawah fungsi `main`:
    ```go
  	switch operation {
  	case "+":
  		result := x + y
  		fmt.Println("The result is:", result)
  	case "-":
  		result := x - y
  		fmt.Println("The result is:", result)
  	case "*":
  		result := x * y
  		fmt.Println("The result is:", result)
  	case "/":
  		result := x / y
  		fmt.Println("The result is:", result)
  	default:
  		fmt.Println("Invalid operation.")
  	 }
    }
    ```

8. Simpan dan jalankan program dengan menggunakan perintah `go run main.go` pada terminal. Anda akan diminta untuk memasukkan operator dan dua angka. Setelah memasukkan input, program akan menghitung hasilnya dan mencetaknya di layar.

## Hasil Output

![Golang Logo](https://raw.githubusercontent.com/BryanFlava/asset-gambar/main/Screenshot%202023-07-05%20191405.png)

## Instalasi Golang
Berikut adalah langkah-langkah untuk menginstal Golang:

1. Kunjungi [situs web resmi Golang](https://golang.org/dl/) untuk mengunduh paket instalasi Golang sesuai dengan sistem operasi Anda.

2. Ikuti petunjuk instalasi yang tertera pada situs web.

3. Setelah selesai menginstal, buka terminal atau command prompt dan ketikkan perintah `go version` untuk memastikan bahwa Golang terinstal dengan benar. Anda akan melihat versi Golang yang terpasang.

Selamat! Anda sekarang telah berhasil menginstal Golang.

---

Sekarang Anda telah memiliki kalkulator sederhana yang dibangun dengan menggunakan Golang. Anda dapat mengembangkan dan menyesuaikan kalkulator ini dengan fitur-fitur tambahan yang Anda inginkan. Selamat mencoba!

