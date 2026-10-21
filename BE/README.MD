# BE Website Artikel

Cara menjalankan be

```
go mod download
go run main.go
```

Jika ingin mengubah port aplikasi sesuaikan main.go

```
http.ListenAndServe("localhost:8080", cors(router))
```

Kemudian sesuaikan konfigurasi database pada .env

```
DB_HOST=
DB_PORT=
DB_DATABASE=
DB_USERNAME=
DB_PASSWORD=
```

Jika ingin membuat tabel manual ada di folder Raw Sql
