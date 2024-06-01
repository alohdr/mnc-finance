# MNC Finance

## Deskripsi Proyek
Proyek ini mencakup berbagai fungsi dan endpoint REST API untuk memenuhi persyaratan yang diberikan. Fungsi-fungsi ini termasuk validasi palindrom, struktur data dalam Go, pengelolaan API server dengan beberapa endpoint, serta tambahan algoritma dan fungsi lainnya.

## Fungsi dan Lokasi

### 1. Fungsi untuk Menentukan Palindrom
- **Deskripsi**: Fungsi ini memeriksa apakah sebuah teks adalah palindrom atau bukan.
- **Lokasi**: `util/helper.go`
- **Nama Fungsi**: `IsPalindrome`

### 2. Struct di Go
- **Deskripsi**: Struct ini menyimpan informasi tentang bahasa pemrograman.
- **Lokasi**: `internal/contract/model/language.go`

### 3. REST API Server
- **Framework**: Bebas
- **Endpoint**:
    - `GET /`: Mengembalikan respons "Hello Go developers".
    - `GET /language`: Mengembalikan data JSON sesuai dengan struct pada nomor 2.
    - Method lain selain kedua di atas: Mengembalikan error "Method not allowed".

### 4. Endpoint untuk Mengecek Palindrom
- **Endpoint**: `GET /palindrome`
- **Parameter**: `text`
- **Respons**:
    - HTTP success jika teks adalah palindrom: "Palindrome"
    - Error bad request jika teks bukan palindrom: "Not palindrome"

### 5. CRUD untuk Data Language
- **Endpoint**:
    - `POST /language`: Menambahkan data bahasa pemrograman.
    - `GET /language/{id}`: Mengembalikan data bahasa pemrograman berdasarkan ID.
    - `GET /languages`: Mengembalikan semua data bahasa pemrograman yang tersimpan.
    - `PATCH /language/{id}`: Mengubah data bahasa pemrograman berdasarkan ID.
    - `DELETE /language/{id}`: Menghapus data bahasa pemrograman berdasarkan ID.

### Contoh Data untuk Pengujian
#### POST /language
```json
{
  "language": "Go",
  "appeared": 2009,
  "created": [
    "Robert Griesemer",
    "Rob Pike",
    "Ken Thompson"
  ],
  "functional": true,
  "object-oriented": false,
  "relation": {
    "influenced-by": [
      "C",
      "Python",
      "Pascal",
      "Smalltalk",
      "Modula"
    ],
    "influences": [
      "Odin",
      "Crystal",
      "Zig"
    ]
  }
}
```

### Menjalankan API Server -> ```make run```

### 1. Pencocokan String
- **Deskripsi**: Buatlah fungsi yang menerima Inputan pertama adalah N jumlah string yang akan diambil kemudian dibaca N (s1, s2, ... sN) string, kemudian diperlukan sebuah algoritma untuk mencocokan semua string satu sama lain dan mengeluarkan nomor string yang cocok satu sama lain. Jika ada beberapa set string yang cocok, tampilkan hasilnya untuk 1 set string saja yang pertama kali ditemukan kecocokan (abaikan set string yang lain). String yang dicocokkan bersifat case insensitive. Jika tidak ditemukan string yang cocok, kembalikan hasil false. Tidak boleh menggunakan fungsi-fungsi array bawaan dari bahasa pemrograman, misalnya: Array Search, Array Filter. Gunakan percabangan dan perulangan secara manual (for / if / do while / etc..).
- **Lokasi**: ```./tes_1.go```
- **Menjalankan**: ```make run-tes-1```

### 2. Menghitung Uang Kembalian
- **Deskripsi**: Buatlah function yang dapat digunakan oleh kasir untuk menghitung nilai uang kembalian beserta dengan pecahan uang yang bisa diberikan. Input berupa total belanja dan jumlah uang yang dibayarkan oleh pembeli. Output berupa kembalian (dibulatkan ke bawah Rp.100) yang harus diberikan kasir dengan detail pecahan uang yang harus diberikan. Pecahan yang tersedia adalah: 100.000, 50.000, 20.000, 10.000, 5.000, 2.000, 1.000, 500, 200, 100. Kembalikan nilai False apabila jumlah uang yang dibayarkan kurang dari total belanjanya.
- **Lokasi**: ```./tes_2.go```
- **Menjalankan**: ```make run-tes-2```

### 3. Validasi Bracket
- **Deskripsi**: Buatlah fungsi validasi untuk string yang diberikan tanpa menggunakan regex (regular expression), sesuai dengan rule di bawah ini:
  - input = string hanya mengandung <>{}[]
  - setiap ada pembuka harus ditutup dengan karakter yang sesuai
  - tidak boleh ada penutup sebelum pembuka, misal ']<>'
  - tidak boleh ada kurung yang mengurung kurungan lain, misalnya '<[>]'
  - kurung boleh didalam kurung lain secara penuh, misal '<[]{<>}>'
  - panjang string adalah 1 – 4096
- **Lokasi**: ```./tes_3.go```
- **Menjalankan**:```make run-tes-3```

### 3. Validasi Bracket
- **Deskripsi**: Buatlah function untuk membantu menentukan apakah seorang karyawan boleh mengambil cuti pribadi atau tidak.
  - ***Peraturan*** :
    - Cuti kantor = 14 hari per tahun.
    - Jumlah cuti pribadi adalah jumlah cuti kantor - cuti bersama.
    - Untuk karyawan baru tidak berhak mengambil cuti pribadi selama 180 hari pertama.
    - Total kuota cuti untuk karyawan baru di tahun pertama adalah jumlah hari / 365 x jumlah cuti
    pribadi di tahun tersebut (pembulatan kebawah). Jumlah hari dihitung mulai tanggal masuk +
    180 hari sampai dengan 31 Desember di tahun tersebut.
    - Cuti pribadi max 3 hari berturutan.
  - ***Contoh Kasus***:
    - Karyawan baru masuk pada tanggal 1 Mei 2021.
    - Jumlah cuti bersama tahun 2021 adalah 7 hari.
    - Karyawan baru hanya dapat mengambil cuti pada tanggal: 1 Mei 2021 + 180 hari = mulai 28
    Oktober 2021.
    - Jumlah hari dihitung dari 28 Oktober 2021 s/d 31 Desember 2021, sehingga = 64 hari.
    - Jumlah cuti pribadi yang bisa diambil adalah: 64/365 x 7 hari cuti pribadi = 1 hari cuti pribadi
    (pembulatan kebawah).
    - Kesimpulan: Karyawan baru yang masuk di tanggal 1 Mei 2021 boleh mengambil cuti pribadi
    sebanyak 1 hari antara rentang tanggal 28 Oktober 2021 s/d 31 Desember 2021.
  - ***Input***:
    - Jumlah Cuti Bersama
    - Tanggal join karyawan
    - Tanggal rencana cuti
    - Durasi cuti (hari)
  - ***Output***:
    - True / False
    - Alasan
- **Lokasi**: ```./tes_4.go```
- **Menjalankan**: ```make run-tes-4```