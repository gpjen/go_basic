# Menggunakan golang alpine image
FROM golang:1.23-alpine

# Set working directory
WORKDIR /app

# Install air untuk hot reload
RUN go install github.com/air-verse/air@latest

# Salin file go.mod dan go.sum untuk mengunduh dependensi
COPY go.mod go.sum ./
RUN go mod tidy && go mod download

# Salin kode aplikasi ke dalam container
COPY . .

# Jalankan air dengan konfigurasi dari file .air.toml
CMD ["air", "-c", ".air.toml"]