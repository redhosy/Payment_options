# Payment Options API

Aplikasi backend sederhana menggunakan bahasa Go (Golang) yang menyediakan endpoint untuk menampilkan daftar metode pembayaran seperti OVO, DANA, GoPay, ShopeePay, OneKlik, BRIDD, dan LinkAja.

## Fitur Utama

- Endpoint REST API GET `/payment/options`
- Pemrosesan paralel menggunakan goroutine
- Simulasi pemanggilan API untuk setiap metode pembayaran
- Koneksi ke database MySQL (produk dan payment)
- Thread-safe dengan penggunaan mutex

## Struktur Proyek

```
.
├── cmd/
│   └── main.go                  # Entry point aplikasi
├── internal/
│   ├── handler/
│   │   └── payment.go           # Handler untuk API payment
│   ├── model/
│   │   └── payment.go           # Model data
│   └── repository/
│       └── db.go                # Koneksi database
├── go.mod                       # Dependency management
└── go.sum                       # Checksums dependency
```

## Teknologi yang Digunakan

- Go (Golang) 1.22.2
- MySQL Driver for Go (github.com/go-sql-driver/mysql)
- HTTP Server dari standard library Go

## Cara Menjalankan

1. Pastikan Go sudah terinstal di sistem anda
2. Clone repository ini
3. Masuk ke direktori proyek

```bash
cd Payment-options
```

4. Download dependency

```bash
go mod tidy
```

5. Jalankan aplikasi

```bash
go run cmd/main.go
```

Server akan berjalan pada port 8081.

## API Endpoint

### GET /payment/options

Endpoint ini mengembalikan daftar semua metode pembayaran yang tersedia beserta informasi detailnya.

#### Response

```json
{
  "returnCode": "200",
  "returnDesc": "success",
  "data": {
    "bridd": {
      "account": "628867890123",
      "status": "Active",
      "balance": "100000",
      "icon": "https://sampleurl.com/bridd.jpg"
    },
    "dana": {
      "account": "628823456789",
      "status": "Active",
      "balance": "15000",
      "icon": "https://sampleurl.com/dana.jpg"
    },
    "gopay": {
      "account": "628834567890",
      "status": "Active",
      "balance": "25000",
      "icon": "https://sampleurl.com/gopay.jpg"
    },
    "linkaja": {
      "account": "628878901234",
      "status": "Active",
      "balance": "75000",
      "icon": "https://sampleurl.com/linkaja.jpg"
    },
    "oneklik": {
      "account": "628856789012",
      "status": "Active",
      "balance": "50000",
      "icon": "https://sampleurl.com/oneklik.jpg"
    },
    "ovo": {
      "account": "628812345678",
      "status": "Active",
      "balance": "10000",
      "icon": "https://sampleurl.com/ovo.jpg"
    },
    "shopeepay": {
      "account": "628845678901",
      "status": "Active",
      "balance": "30000",
      "icon": "https://sampleurl.com/shopeepay.jpg"
    }
  }
}
```

## Implementasi Teknis

### Penggunaan Goroutines

Aplikasi menggunakan goroutines untuk memproses setiap permintaan API secara paralel. Setiap metode pembayaran memiliki fungsi simulasi yang berjalan secara independen.

```go
// Process each payment method in parallel
for _, method := range paymentMethods {
    wg.Add(1)
    go func(method string) {
        defer wg.Done()
        
        var option *model.PaymentOption
        
        switch method {
        case "ovo":
            option = getOVOProfile()
        // ...dll
        }
        
        // Store result in map with mutex protection
        if option != nil {
            mutex.Lock()
            data[method] = option
            mutex.Unlock()
        }
    }(method)
}
```

### Sinkronisasi

Untuk memastikan thread safety, aplikasi menggunakan `sync.WaitGroup` untuk menunggu semua goroutines selesai dan `sync.Mutex` untuk melindungi akses konkuren ke map data.

### Struktur Data

```go
// PaymentOption stores payment data including account information, status, balance, and icon
type PaymentOption struct {
    Account string `json:"account"`
    Status  string `json:"status"`
    Balance string `json:"balance"`
    Icon    string `json:"icon"`
}

// Response represents the API response format
type Response struct {
    ReturnCode string                    `json:"returnCode"`
    ReturnDesc string                    `json:"returnDesc"`
    Data       map[string]*PaymentOption `json:"data"`
}
```

### Koneksi Database

Aplikasi mengkonfigurasi koneksi ke dua database MySQL (produk dan payment) melalui DBManager yang diterapkan di `internal/repository/db.go`. Saat ini koneksi database sudah diatur namun belum digunakan dalam operasi API.

## Pengujian API

Anda bisa menguji API dengan menggunakan curl:

```bash
curl -X GET http://localhost:8081/payment/options
```

Atau dengan format yang lebih mudah dibaca:

```bash
curl -X GET http://localhost:8081/payment/options | python3 -m json.tool
```

## Catatan Pengembangan

- Credentials database saat ini menggunakan placeholder dan perlu diganti dengan nilai yang sebenarnya
- Aplikasi ini masih dalam tahap pengembangan awal dan belum memiliki sistem logging yang komprehensif
- Untuk pengembangan lebih lanjut, penambahan unit test dan middleware autentikasi diperlukan
# Payment_options
