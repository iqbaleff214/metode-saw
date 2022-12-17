# metode-saw
Penerapan metode SAW (Simple Additive Weighting) dalam berbagai bahasa pemrograman

## Studi Kasus

Studi kasus yang digunakan adalah untuk menentukan mahasiswa yang berhak mendapatkan beasiswa Bidik Misi di Politeknik Negeri Banjarmasin. Data yang didapat berasal dari jurnal terlampir di bagian referensi.

### Kriteria Acuan

| Keterangan | Kriteria | Kecocokan | Bobot |
|------------|----------|-----------|-------|
| Penghasilan orang tua | C1 | _Cost_ | 5 |
| Jumlah tanggungan | C2 | _Benefit_ | 4 |
| Rata-rata nilai raport semester 4-5 | C3 | _Benefit_ | 4 |
| Bukti rekening listrik | C4 | _Cost_ | 3 |
| Bukti pembayaran PBB | C5 | _Cost_ | 2 |

### Alternatif Penerima Beasiswa

| Alternatif | C1 | C2 | C3 | C4 | C5 |
|------------|----|----|----|----|----|
| ke 1    | 1.500.000 | 4 | 83,541666666667 | 140.000 | 18.000 |
| ke 2    | 1.250.000 | 4 | 87,821969696970 | 150.000 | 20.000 |
| ke 3    | 1.250.000 | 4 | 92,291666666667 | 140.000 | 18.000 |
| ke 4    |   750.000 | 3 | 89,858333333333 | 150.000 | 20.000 |
| ke 5    | 1.250.000 | 3 | 88,058300000000 | 140.000 | 18.000 |
| ke 6    | 1.500.000 | 2 | 85,954700000000 | 150.000 | 20.000 |
| ke 7    | 2.500.000 | 2 | 84,345200000000 | 140.000 | 18.000 |
| ke 8    | 1.750.000 | 3 | 85,117600000000 | 150.000 | 20.000 |
| ke 9    | 1.000.000 | 6 | 83,735200000000 | 140.000 | 18.000 |
| ke 2154 | 1.250.000 | 3 | 92,291666666667 | 140.000 | 18.000 |

## Cara Menjalankan

Untuk menjalankan program, silakan gunakan CLI yang biasa Anda gunakan. Berikut cara menjalankan untuk masing-masing bahasa pemrograman:

-   [Go](https://go.dev/)
    ```sh
    go run ./go/main.go
    ```
    atau bisa juga dengan menggunakan

    ```sh
    make go-run
    ```
    *) rekomendasi: `go1.19`

-   [PHP](https://www.php.net/)
    ```sh
    php ./php/main.php
    ```
    atau bisa juga dengan menggunakan

    ```sh
    make php-run
    ```
    *) rekomendasi: `PHP 8.1.13`

-   [JavaScript](https://nodejs.org/en/)
    ```sh
    node ./js/main.js
    ```
    atau bisa juga dengan menggunakan

    ```sh
    make node-run
    ```
    *) rekomendasi: `node v18.12.1`

-   [Ruby](https://www.ruby-lang.org/en/)
    ```sh
    ruby ./ruby/main.rb
    ```
    atau bisa juga dengan menggunakan

    ```sh
    make ruby-run
    ```
    *) rekomendasi: `ruby 3.0.2p107`


## Referensi
- [Sistem Pendukung Keputusan Penerimaan Beasiswa Bidik Misi di POLIBAN Dengan Metode SAW Berbasis Web](http://join.if.uinsgd.ac.id/index.php/join/article/view/101)
- [METODE SAW (Simple Additive Weighting) | Konsep & Contoh Kasus](https://www.youtube.com/watch?v=_7-catHioro)
