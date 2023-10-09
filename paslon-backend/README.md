# Proyek CRUD sederhana dengan Go, Echo, dan PostgreSQL

Proyek ini adalah contoh aplikasi CRUD (Create, Read, Update, Delete) sederhana yang dibangun menggunakan bahasa pemrograman Go, framework Echo, dan basis data PostgreSQL. Aplikasi ini digunakan untuk mengelola data paslon pemilu.

## Instalasi dan Penggunaan

1. Pastikan Anda memiliki Go dan PostgreSQL terinstal di komputer Anda.

2. Clone repositori ini ke komputer Anda.

   ```bash
   git clone https://github.com/alifdwt/dumbways-b49Add-alif-golang.git
   cd dumbways-b49Add-alif-golang/paslon-backend
   ```

3. Atur koneksi database Anda pada `pkg/postgres/postgres.go` (jika menggunakan PostgreSQL). Petunjuk lengkap dapat dibaca di dokumentasi [GORM](https://gorm.io/docs/)

4. Jalankan aplikasi.

   ```bash
   go run main.go
   ```

5. Aplikasi akan berjalan di `http://localhost:8000`.

# API Endpoints

## Register

- URL: `/api/v1/auth/register`
- Method: POST
- Deskripsi: Mendaftarkan user ke database
- Membutuhkan token: `Tidak`
- Request:

| Key       | Type | Value             |
| --------- | ---- | ----------------- |
| full_name | text | Dominic Deccoco   |
| email     | text | dominic@gmail.com |
| password  | text | h1st0r10gr4f1     |

- Response:

```
{
    "code": 200,
    "data": {
        "id": 3,
        "full_name": "Dominic Deccoco",
        "email": "dominic@gmail.com"
    }
}
```

### Login

- URL: `/api/v1/auth/login`
- Method: POST
- Deskripsi: Melakukan Login
- Membutuhkan token: `Tidak`
- Request:

| Key      | Type | Value             |
| -------- | ---- | ----------------- |
| email    | text | dominic@gmail.com |
| password | text | h1st0r10gr4f1     |

- Response:

```
{
    "code": 200,
    "data": {
        "id": 3,
        "full_name": "Dominic Deccoco",
        "email": "dominic@gmail.com"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRvbWluaWNAZ21haWwuY29tIiwidXNlcl9pZCI6M30.aoPV5Ds-BCwY04ZD-7fG8xT_joFqTeudeoRFFZ7-zDw"
}
```

## Parties

### Get Parties

- URL: `/api/v1/parties`
- Method: GET
- Deskripsi: Mengambil semua daftar partai.
- Membutuhkan token: `Tidak`
- Response:

```
{
    "code": 200,
    "data": [
        {
            "id": 5,
            "name": "Partai Golongan Kartun",
            "paslon_id": 3,
            "created_at": "2023-10-06T09:08:49.298476+07:00",
            "updated_at": "2023-10-06T09:08:49.298476+07:00",
            "paslons": {
                "id": 3,
                "name": "Patrick Stars",
                "vision": "Mengembangkan potensi sumber daya alam dan mengedukasi penduduk tentang pentingnya menjaga lingkungan laut.",
                "image": "https://res.cloudinary.com/dxirtmo5t/image/upload/v1696439225/Pemilu/patrick_stars.webp",
                "posted_at": "2023-10-06T02:45:56.265315+07:00",
                "updated_at": "2023-10-06T02:45:56.265315+07:00",
                "party": null
            }
        },
        {
            "id": 4,
            "name": "Partai Bikini Bottom Sejahtera",
            "paslon_id": 1,
            "created_at": "2023-10-06T09:08:02.303082+07:00",
            "updated_at": "2023-10-06T10:09:05.476106+07:00",
            "paslons": {
                "id": 1,
                "name": "Spongebob Squarepants",
                "vision": "Membuat Bikini Bottom menjadi tempat yang lebih ceria dan ramah lingkungan",
                "image": "https://res.cloudinary.com/dxirtmo5t/image/upload/v1696438683/Pemilu/spongebob_squarepants.jpg",
                "posted_at": "2023-10-06T02:37:42.30249+07:00",
                "updated_at": "2023-10-06T02:37:42.30249+07:00",
                "party": null
            }
        },
        {
            "id": 8,
            "name": "Partai Nasional Bikini Bottom",
            "paslon_id": 2,
            "created_at": "2023-10-06T10:40:03.419812+07:00",
            "updated_at": "2023-10-06T10:40:03.419812+07:00",
            "paslons": {
                "id": 2,
                "name": "Squidward Tentacles",
                "vision": "Menghargai seni dan budaya, serta meningkatkan kualitas hidup melalui seni dan musik",
                "image": "https://res.cloudinary.com/dxirtmo5t/image/upload/v1696438791/Pemilu/squidward_tentacles.jpg",
                "posted_at": "2023-10-06T02:38:49.676298+07:00",
                "updated_at": "2023-10-06T02:38:49.676298+07:00",
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
                "id": 2,
                "name": "Squidward Tentacles",
                "vision": "Menghargai seni dan budaya, serta meningkatkan kualitas hidup melalui seni dan musik",
                "image": "https://res.cloudinary.com/dxirtmo5t/image/upload/v1696438791/Pemilu/squidward_tentacles.jpg",
                "posted_at": "2023-10-06T02:38:49.676298+07:00",
                "updated_at": "2023-10-06T02:38:49.676298+07:00",
                "party": null
            }
        }
    ]
}
```

### Get Party by ID

- URL: `/api/v1/party/:partyId` (misalnya: `4`)
- Method: GET
- Deskripsi: Mengambil data partai berdasarkan ID
- Membutuhkan token: `Tidak`
- Response:

```
{
    "code": 200,
    "data": {
        "id": 4,
        "name": "Partai Bikini Bottom Sejahtera",
        "paslon_id": 1,
        "created_at": "2023-10-06T09:08:02.303082+07:00",
        "updated_at": "2023-10-06T10:09:05.476106+07:00",
        "paslons": {
            "id": 1,
            "name": "Spongebob Squarepants",
            "vision": "Membuat Bikini Bottom menjadi tempat yang lebih ceria dan ramah lingkungan",
            "image": "https://res.cloudinary.com/dxirtmo5t/image/upload/v1696438683/Pemilu/spongebob_squarepants.jpg",
            "posted_at": "2023-10-06T02:37:42.30249+07:00",
            "updated_at": "2023-10-06T02:37:42.30249+07:00",
            "party": null
        }
    }
}
```

### Create Party

- URL: `/api/v1/party`
- Method: POST
- Deskripsi: Menambah partai ke dalam database
- Membutuhkan token: `Tidak`
- Request:

| Key       | Type | Value                                     |
| --------- | ---- | ----------------------------------------- |
| name      | text | Partai Demokrasi Bikini Bottom Perjuangan |
| paslon_id | text | 3                                         |

- Response

```
{
    "code": 200,
    "data": {
        "id": 9,
        "name": "Partai Demokrasi Bikini Bottom Perjuangan",
        "paslon_id": 3,
        "created_at": "2023-10-09T15:55:36.356665+07:00",
        "updated_at": "2023-10-09T15:55:36.356665+07:00",
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
}
```

### Update Party

- URL: `/api/v1/party/:partyId` (misalnya: `9`)
- Method: PATCH
- Deskripsi: Mengubah data partai
- Membutuhkan token: `Tidak`
- Request:

```
{
    "name": "Partai Demokrasi Bikini Bottom",
    "paslon_id": 1
}
```

- Response:

```
{
    "code": 200,
    "data": {
        "id": 9,
        "name": "Partai Demokrasi Bikini Bottom",
        "paslon_id": 1,
        "created_at": "2023-10-09T15:55:36.356665+07:00",
        "updated_at": "2023-10-09T16:01:23.3681348+07:00",
        "paslons": {
            "id": 3,
            "name": "Patrick Stars",
            "vision": "Mengembangkan potensi sumber daya alam dan mengedukasi penduduk tentang pentingnya menjaga lingkungan laut.",
            "image": "https://res.cloudinary.com/dxirtmo5t/image/upload/v1696439225/Pemilu/patrick_stars.webp",
            "posted_at": "2023-10-06T02:45:56.265315+07:00",
            "updated_at": "2023-10-06T02:45:56.265315+07:00",
            "party": null
        }
    }
}
```

### Delete Party

- URL: `/api/v1/party/:partyId` (misalnya: `9`)
- Method: DELETE
- Deskripsi: Menghapus data satu partai
- Membutuhkan token: `Tidak`
- Response:

```
{
    "code": 200,
    "data": {
        "id": 9,
        "name": "Partai Demokrasi Bikini Bottom",
        "paslon_id": 1,
        "created_at": "2023-10-09T15:55:36.356665+07:00",
        "updated_at": "2023-10-09T16:01:23.368134+07:00",
        "paslons": {
            "id": 1,
            "name": "Spongebob Squarepants",
            "vision": "Membuat Bikini Bottom menjadi tempat yang lebih ceria dan ramah lingkungan",
            "image": "https://res.cloudinary.com/dxirtmo5t/image/upload/v1696438683/Pemilu/spongebob_squarepants.jpg",
            "posted_at": "2023-10-06T02:37:42.30249+07:00",
            "updated_at": "2023-10-06T02:37:42.30249+07:00",
            "party": null
        }
    }
}
```

## Voters

### Get Voters by ID

- URL: `/api/v1/voter/voterId` (misalnya: `2`)
- Method: GET
- Deskripsi: Mendapatkan data voter berdasarkan ID
- Response:

```
{
    "code": 200,
    "data": {
        "id": 2,
        "paslon_id": 3,
        "user_id": 2,
        "voter_name": "Enzo Gorlomi",
        "created_at": "2023-10-09T13:59:52.537864+07:00",
        "updated_at": "2023-10-09T13:59:52.537864+07:00",
        "paslons": {
            "id": 3,
            "name": "Patrick Stars",
            "vision": "Mengembangkan potensi sumber daya alam dan mengedukasi penduduk tentang pentingnya menjaga lingkungan laut.",
            "image": "https://res.cloudinary.com/dxirtmo5t/image/upload/v1696439225/Pemilu/patrick_stars.webp",
            "posted_at": "2023-10-06T02:45:56.265315+07:00",
            "updated_at": "2023-10-06T02:45:56.265315+07:00",
            "party": null
        }
    }
}
```

### Create Vote

- URL: `/api/v1/voter`
- Method: POST
- Deskripsi: Menambahkan vote ke database
- Request:

| Key        | Type | Value               |
| ---------- | ---- | ------------------- |
| paslon_id  | text | 3                   |
| voter_name | text | Antonio Margharetti |
| user_id    | text | 1                   |

- Response:

```
{
    "code": 200,
    "data": {
        "id": 3,
        "voter_name": "Antonio Margharetti",
        "paslon_id": 3,
        "user_id": 0
    }
}
```

## Paslon

### Get Paslons

- URL: `/api/v1/paslons`
- Method: GET
- Deskripsi: Mengambil daftar paslon pemilu.
- Membutuhkan token: `Tidak`
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
- Membutuhkan token: `Tidak`
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
- Membutuhkan token: `Tidak`
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
- Membutuhkan token: `Tidak`
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
- Membutuhkan token: `Tidak`
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
