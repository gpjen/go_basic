# Menggunakan Go image versi Alpine yang ringan
FROM golang:1.23

# Mengatur working directory di dalam container
WORKDIR /app

# Install air untuk reload otomatis
RUN go install github.com/air-verse/air@latest

# Copy seluruh file proyek ke dalam container
COPY . .

# Install dependensi Go (termasuk di dalam folder /app)
RUN go mod tidy

# Set air sebagai perintah utama container
CMD ["air"]
