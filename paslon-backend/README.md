# Proyek CRUD sederhana dengan Go, Echo, dan PostgreSQL

Proyek ini adalah contoh aplikasi CRUD (Create, Read, Update, Delete) sederhana yang dibangun menggunakan bahasa pemrograman Go, framework Echo, dan basis data PostgreSQL. Aplikasi ini digunakan untuk mengelola data paslon pemilu.

## Instalasi dan Penggunaan

1. Pastikan Anda memiliki Go dan PostgreSQL terinstal di komputer Anda.

2. Clone repositori ini ke komputer Anda.

   ```bash
   git clone https://github.com/alifdwt/dumbways-b49Add-alif-golang.git
   cd dumbways-b49Add-alif-golang
   ```

3. Atur koneksi database Anda pada `pkg/postgres/postgres.go` (jika menggunakan PostgreSQL). Petunjuk lengkap dapat dibaca di dokumentasi [GORM](https://gorm.io/docs/)

4. Jalankan aplikasi.

   ```bash
   go run main.go
   ```

5. Aplikasi akan berjalan di `http://localhost:8000`.

## API Endpoints

### Get Paslons

- URL: `/api/v1/paslons`
- Method: GET
- Deskripsi: Mengambil daftar paslon pemilu.

### Get Paslon by ID

- URL: `/api/v1/paslon/:paslonId`
- Method: GET
- Deskripsi: Mengambil detail paslon berdasarkan ID.

### Create Paslon

- URL: `/api/v1/paslon`
- Method: POST
- Deskripsi: Membuat paslon baru.

### Update Paslon

- URL: `/api/v1/paslon/:paslonId`
- Method: PUT
- Deskripsi: Mengupdate paslon berdasarkan ID.

### Delete Paslon

- URL: `/api/v1/paslon/:paslonId`
- Method: DELETE
- Deskripsi: Menghapus paslon berdasarkan ID.

## Kontribusi

Anda dipersilakan untuk berkontribusi pada proyek ini dengan cara melakukan pull request atau melaporkan masalah yang Anda temui.
