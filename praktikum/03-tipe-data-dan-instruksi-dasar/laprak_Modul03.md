# <h1 align="center">Laporan Praktikum Modul 3 - Bahasa Pemrograman Go/Variabel,Operator</h1>
<p align="center">Clara Shevy Olya Felka - 109092600017</p>

## Dasar Teori

### A. [Bahasa Pemrograman Go]
[Go adalah bahasa pemrograman serbaguna yang dirancang dengan mempertimbangkan pemrograman sistem. Bahasa ini memiliki tipe data yang kuat dan garbage collector, serta dukungan eksplisit untuk pemrograman konkuren. Program dibangun dari paket-paket , yang propertinya memungkinkan pengelolaan dependensi yang efisien.]


#### 1. [Variabel]
[Variabel adalah wadah didalam memori yang digunakan untuk menyimpan nilai dengan tipe data tertentu. Di dalam Go variabel dapat menggunakan kata kunci var atau menggunakan := yang membedakan tipe data berdasarkan nilai yang diberikan.]

#### 2. [Operator]
[Operator dalam bahasa Go merupakan simbol khusus yang digunakan untuk melakukan berbagai manipulasi dan operasi data, meliputi perhitungan aritmatika, pembandingan nilai hubungan, penggabungan kondisi logika, serta penugasaan nilai ke dalam variabel.]

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. [konversi.go]

```go
package main

import "fmt"

func main() {
	var celcius float64

	fmt.Println("Masukkan suhu:")
	fmt.Scan(&celcius)

	fmt.Println(celcius + 273)
}
```
#### Deskripsi
[Pada pratikum ini, membuat dan menguji program konversi.go, dilakukan untuk mengonversi suhu dari satuan celcius ke satuan kelvin. Variabel celcius bertipe float64, menerima masukan nilai suhu menggunakan fmt.scan, lalu menghitung dengan (celcius + 273).]

### 2. [pertukaran_nilai.go]

```go
package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	temp := x
	x = z
	z = y
	y = temp
	fmt.Println(x, y, z)
}
```
#### Deskripsi
[Pada pratikum ini, membuat program pertukaran_nilai.go yang fungsinya menukar nilai 3 variabel. program dimulai dengan menyiapkan 3 variabel bertipe angka (int), menerima input melalui fmt.Scan, lalu menukar variabel dengan perintah x=z, z=y, y=x. Hasil akhir, variabel bisa tertukar.]

### 3. [cacahuang.go]

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
#### Deskripsi
[Pada pratikum ini, membuat program untuk menghitung jumlah lembar pecahan uang 10 ribu, 5 ribu, dan 1 ribu. programmembaca jumlah uang lalu menghitung jumlah lembar menggunakan pembagian dan mengambil sisa uang dengan modulus. Hasilnya program dapat menghitung jumlah pecahan secara berurutan.]

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. [konversi_hari.go]

```go
package main

import "fmt"

func main() {
	var hari int
	fmt.Scan(&hari)

	tahun := hari / 360
	hari = hari % 360
	bulan := hari / 30
	hari = hari % 30
	minggu := hari / 7
	hari = hari % 7
	
	fmt.Println(tahun, "tahun", bulan, "bulan", minggu, "minggu", hari, "hari")
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/03-tipe-data-dan-instruksi-dasar/unguided/konversi_hari/OUTPUT%20KONVERSI%20HARI.png)


#### Deskripsi
[Pada pratikum ini, pembuatan dan pengujian konversi_hari.go dilakukan untuk mengonversi total jumlah hari ke dalam satuan tahun, bulan, minggu, sisa hari. program ini membaca masukan nilai integer hari dengan fmt.Scan, kemudian hitung jumlah tahun (hari/360), bulan (hari/30), minggu (hari/7), serta sisa harimenggunakan modulo.]

### 2. [reamur.go]

```go
package main

import "fmt"

func main() {
	var celcius, reamur float64
	fmt.Scan(&celcius)
	reamur = celcius * 4 / 5
	fmt.Println(reamur)
}
```

##### Output
![Screenshot Output Unguided](/03-tipe-data-dan-instruksi-dasar/unguided/reamur/OUTPUT%20REAMUR.png)

#### Deskripsi
[Pada program ini, Pembuatan dan pengujian reamur.go dilakukan untuk mengonversikan dari satuan celcius ke reamur bertipe float64, membaca masukan nilai suhu menggunakan fmt.Scan, lalu menghitung hasil berdasarkan (celcius * 4/5).]

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
[Dari pratikum pratikum yang sudah di lakukan, saya jadi paham dengan dasar-dasar go dan cara mengkompilasikan kodenya, serta dapat belajar membuat beberapa program sederhana.]

## Referensi
1. [Donovan, A.A., & Kernighan, B.W. (2015). The Go Programming Language. Addison-Wesley Professional.]
2. [Prasetyo, A. (2020). Dasar Pemrograman Golang. Jakarta,PT Elex Media Komputindo.]
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
