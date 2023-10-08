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

# API Endpoints

## Paslon

### Get Paslons

- URL: `/api/v1/paslons`
- Method: GET
- Deskripsi: Mengambil daftar paslon pemilu.
- Response:

```
{
    "code": 200,
    "data": [
        {
            "id": 1,
            "name": "Spongebob Squarepants",
            "vision": "Membuat Bikini Bottom menjadi tempat yang lebih ceria dan ramah lingkungan",
            "image": "https://res.cloudinary.com/dxirtmo5t/image/upload/v1696438683/Pemilu/spongebob_squarepants.jpg",
            "posted_at": "2023-10-06T02:37:42.30249+07:00",
            "updated_at": "2023-10-06T02:37:42.30249+07:00",
            "party": [
                {
                    "id": 4,
                    "name": "Partai Bikini Bottom Sejahtera",
                    "paslon_id": 1,
                    "created_at": "2023-10-06T09:08:02.303082+07:00",
                    "updated_at": "2023-10-06T10:09:05.476106+07:00",
                    "paslons": {
                        "id": 0,
                        "name": "",
                        "vision": "",
                        "image": "",
                        "posted_at": "0001-01-01T00:00:00Z",
                        "updated_at": "0001-01-01T00:00:00Z",
                        "party": null
                    }
                }
            ]
        },
        {
            "id": 2,
            "name": "Squidward Tentacles",
            "vision": "Menghargai seni dan budaya, serta meningkatkan kualitas hidup melalui seni dan musik",
            "image": "https://res.cloudinary.com/dxirtmo5t/image/upload/v1696438791/Pemilu/squidward_tentacles.jpg",
            "posted_at": "2023-10-06T02:38:49.676298+07:00",
            "updated_at": "2023-10-06T02:38:49.676298+07:00",
            "party": [
                {
                    "id": 8,
                    "name": "Partai Nasional Bikini Bottom",
                    "paslon_id": 2,
                    "created_at": "2023-10-06T10:40:03.419812+07:00",
                    "updated_at": "2023-10-06T10:40:03.419812+07:00",
                    "paslons": {
                        "id": 0,
                        "name": "",
                        "vision": "",
                        "image": "",
                        "posted_at": "0001-01-01T00:00:00Z",
                        "updated_at": "0001-01-01T00:00:00Z",
                        "party": null
                    }
                },
                {
                    "id": 3,
                    "name": "Partai Solidaritas Bikini Bottom",
                    "paslon_id": 2,
                    "created_at": "2023-10-06T02:47:25.825489+07:00",
                    "updated_at": "2023-10-06T10:41:00.7036+07:00",
                    "paslons": {
                        "id": 0,
                        "name": "",
                        "vision": "",
                        "image": "",
                        "posted_at": "0001-01-01T00:00:00Z",
                        "updated_at": "0001-01-01T00:00:00Z",
                        "party": null
                    }
                }
            ]
        },
        {
            "id": 3,
            "name": "Patrick Stars",
            "vision": "Mengembangkan potensi sumber daya alam dan mengedukasi penduduk tentang pentingnya menjaga lingkungan laut.",
            "image": "https://res.cloudinary.com/dxirtmo5t/image/upload/v1696439225/Pemilu/patrick_stars.webp",
            "posted_at": "2023-10-06T02:45:56.265315+07:00",
            "updated_at": "2023-10-06T02:45:56.265315+07:00",
            "party": [
                {
                    "id": 5,
                    "name": "Partai Golongan Kartun",
                    "paslon_id": 3,
                    "created_at": "2023-10-06T09:08:49.298476+07:00",
                    "updated_at": "2023-10-06T09:08:49.298476+07:00",
                    "paslons": {
                        "id": 0,
                        "name": "",
                        "vision": "",
                        "image": "",
                        "posted_at": "0001-01-01T00:00:00Z",
                        "updated_at": "0001-01-01T00:00:00Z",
                        "party": null
                    }
                }
            ]
        }
    ]
}
```

### Get Paslon by ID

- URL: `/api/v1/paslon/:paslonId`
- Method: GET
- Deskripsi: Mengambil detail paslon berdasarkan ID.
- Request:

```
{
  "paslon_id": 1
}
```

- Response:

```
{
    "code": 200,
    "data": {
        "id": 1,
        "name": "Spongebob Squarepants",
        "vision": "Membuat Bikini Bottom menjadi tempat yang lebih ceria dan ramah lingkungan",
        "image": "https://res.cloudinary.com/dxirtmo5t/image/upload/v1696438683/Pemilu/spongebob_squarepants.jpg",
        "posted_at": "2023-10-06T02:37:42.30249+07:00",
        "updated_at": "2023-10-06T02:37:42.30249+07:00",
        "party": [
            {
                "id": 4,
                "name": "Partai Bikini Bottom Sejahtera",
                "paslon_id": 1,
                "created_at": "2023-10-06T09:08:02.303082+07:00",
                "updated_at": "2023-10-06T10:09:05.476106+07:00",
                "paslons": {
                    "id": 0,
                    "name": "",
                    "vision": "",
                    "image": "",
                    "posted_at": "0001-01-01T00:00:00Z",
                    "updated_at": "0001-01-01T00:00:00Z",
                    "party": null
                }
            }
        ]
    }
}
```

### Create Paslon

- URL: `/api/v1/paslon`
- Method: POST
- Deskripsi: Membuat paslon baru.
- Form data:

| Key    | Type | Value                                                         |
| ------ | ---- | ------------------------------------------------------------- |
| name   | text | Sandy Cheeks                                                  |
| vision | text | Melindungi ekosistem laut dan mempromosikan penelitian ilmiah |
| image  | file | C:\Users\aputr\Pictures\sandy.jpg                             |

- Response:

```
{
    "code": 200,
    "data": {
        "id": 0,
        "name": "Sandy Cheeks",
        "vision": "Melindungi ekosistem laut dan mempromosikan penelitian ilmiah",
        "image": "https://res.cloudinary.com/dxirtmo5t/image/upload/v1696786081/Pemilu/sandy_cheeks.jpg",
        "posted_at": "2023-10-09T00:28:00.4129385+07:00",
        "updated_at": "2023-10-09T00:28:00.4129385+07:00",
        "party": null
    }
}
```

### Update Paslon

- URL: `/api/v1/paslon/:paslonId` (misal: 4)
- Method: PUT
- Deskripsi: Mengupdate paslon berdasarkan ID.
- Form data:

| Key    | Type | Value                                                                                                   |
| ------ | ---- | ------------------------------------------------------------------------------------------------------- |
| name   | text | Plankton                                                                                                |
| vision | text | Menciptakan inovasi teknologi untuk memajukan industri makanan cepat saji dan mengurangi limbah plastik |
| image  | file | C:\Users\aputr\Pictures\plankton.webp                                                                   |

- Response

```
{
    "code": 200,
    "data": {
        "id": 4,
        "name": "Plankton",
        "vision": "Menciptakan inovasi teknologi untuk memajukan industri makanan cepat saji dan mengurangi limbah plastik",
        "image": "https://res.cloudinary.com/dxirtmo5t/image/upload/v1696786671/Pemilu/plankton.webp",
        "voters": null
    }
}
```

### Delete Paslon

- URL: `/api/v1/paslon/:paslonId` (misalnya: 4)
- Method: DELETE
- Deskripsi: Menghapus paslon berdasarkan ID.
- Response:

```
{
    "code": 200,
    "data": {
        "id": 4,
        "name": "Plankton",
        "vision": "Menciptakan inovasi teknologi untuk memajukan industri makanan cepat saji dan mengurangi limbah plastik",
        "image": "https://res.cloudinary.com/dxirtmo5t/image/upload/v1696786671/Pemilu/plankton.webp",
        "voters": null
    }
}
```

## Kontribusi

Anda dipersilakan untuk berkontribusi pada proyek ini dengan cara melakukan pull request atau melaporkan masalah yang Anda temui.
