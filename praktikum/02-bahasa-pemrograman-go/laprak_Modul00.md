# <h1 align="center">Laporan Praktikum Modul 2 - Algoritma dan Pemrograman</h1>
<p align="center">Clara Shevy Olya Felka - 109092600017</p>

## Dasar Teori

### A. [Bahasa Pemrograman Go]
[Go adalah bahasa pemrograman serbaguna yang dirancang dengan mempertimbangkan pemrograman sistem. Bahasa ini memiliki tipe data yang kuat dan garbage collector, serta dukungan eksplisit untuk pemrograman konkuren. Program dibangun dari paket-paket , yang propertinya memungkinkan pengelolaan dependensi yang efisien.]

### B. [Package dan Struktur Program di Go]

#### 1. [Pengertian Package main dan func main()]
[package main merupakan penanda bahwa file ini berisi program utama. func main() berisi kode utama dari sebuah program Go. ]

#### 2. [Koding, Kompilasi, dan Eksekusi Go]
[2.1 Koding: Kode program Go ditulis menggunakan penyunting teks dan disimpan dalam format teks biasa. Setiap berkas program wajib menggunakan ekstensi .go . Satu program lengkap disimpan dalam satu folder tersendiri dan dapat dipecah kedalam beberapa berkas .go di folder yang sama. 2.2 Komplikasi: Bahasa Go menggunakan kompilator yang memeriksa seluruh kode dan mengoptimalkan menjadi file eksekutabel sebelum di jalankan, dimana proses komplikasinya dilakukan dengan perintah go build melalui terminal. 2.3 Eksekusi Go: Eksekusi Go dilakukan melalui terminal atau command prompt dengan memanggil langsung file ekskutabel seperti .exe, yang dihasilkan dari proses kompilasi go build.]

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. [hello.go]

```go
package main
import "fmt"

func main(){
    fmt.Println("Hello World!")
}
```
#### Deskripsi
[Pada pratikum ini, membuat dan menguji program hello.go, dan diimplementasikan dengan main yang mengimport fmt untuk mencetak teks. Setelah itu dijalankan, program berhasil menghasilkan output "Hello World!"]

### 2. [Tukar.go]

```go
package main
import "fmt"

func main() {
	var a,b int

	//Membaca input
	fmt.Scan(&a)
	fmt.Scan(&b)

	//Menukar a dan b
	a,b = b,a 

	//Menampilkan output
	fmt.Println(a)
	fmt.Println(b)
}
```
#### Deskripsi
[Pada pratikum ini, membuat program Tukar.go yang fungsinya menukar nilai 2 variabel. program dimulai dengan menyiapkan 2 variabel bertipe angka (int), menerima input melalui fmt.Scan, lalu menukar variabel dengan perintah a,b = b,a. Hasil akhir, variabel bisa tertukar.]

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. [cacahuang]

```go
package main

import "fmt"

func main() {
	var uang int 

	// Membaca masukan nominal uang
	fmt.Scan(&uang)

	// Menghitung lembar 10.000
	sepuluhribu := uang / 10000
	sisa := uang % 10000

	// Menghitung lembar 5.000 dari sisa uang
	limaribu := sisa / 5000
	sisa = sisa % 5000

	// Menghitung lembar 1.000 dari sisa uang (sisa < 1000 akan diabaikan secara otomatis)
	seribu := sisa / 1000

	// Menampilkan jumlah lembar pecahan di pisahkan oleh spasi
	fmt.Println(sepuluhribu, limaribu, seribu)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](02-bahasa-pemrograman-go\unguided\cacahuang\OUTPUT CACAHUANG.png)


#### Deskripsi
[Pada pratikum ini, membuat program untuk menghitung jumlah lembar pecahan uang 10 ribu, 5 ribu, dan 1 ribu. programmembaca jumlah uang lalu menghitung jumlah lembar menggunakan pembagian dan mengambil sisa uang dengan modulus. Hasilnya program dapat menghitung jumlah pecahan secara berurutan.]

### 2. [kalkulator]

```go
package main

import "fmt"

func main() {
	var a,b int

	// Membaca masukan dua bilangan bulat
	fmt.Scan(&a, &b)

	// Menghitung hasil operasi aritmatika
	tambah := a + b
	kurang := a - b
	kali := a * b
	bagi := a / b
	modulo := a % b

	// Menampilkan keluaran dipisahkan oleh spasi
	fmt.Println(tambah, kurang, kali, bagi, modulo)
}
```

##### Output
![Screenshot Output Unguided](02-bahasa-pemrograman-go\unguided\kalkulator\OUTOUT KALKULATOR.png)

#### Deskripsi
[Pada program ini, membuat program Kalkulator sederhana. Program membaca 2 angka lalu melakukan operasi matematika dasar (tambah, kurang, kali, bagi, dan modulo). Hasil dari seluruh perhitungan tersebut kemudian ditampilkan berurutan dan dipisahkan oleh spasi.]

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
[Dari pratikum pratikum yang sudah di lakukan, saya jadi paham dengan dasar-dasar go dan cara mengkompilasikan kodenya, serta dapat belajar membuat beberapa program sederhana.]

## Referensi
1. [Robert Griesemer, Rob Pike, dan Ken Thompson]. ([2009]). *[The Go Programming Language Specification]*.[Google Project Go]. Diakses pada [27 Septermber] melalui [(https://go.dev/ref/spec)]
2. [Fakultas Informatika Telkom University,Modul 2 BAHASA PEMROGRAMAN GO, Telkom University.]
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
