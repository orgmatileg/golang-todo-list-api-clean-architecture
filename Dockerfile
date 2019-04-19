# Luqmanul Hakim / arlhba@gmail.com

# Step 1 membuat binary
FROM golang:alpine AS builder

# Install git (Ga usah di tanya lah ya kalau ini haha), g++ (untuk build), tzdata (Untuk set timezone)
RUN apk update && apk add --no-cache git g++ tzdata

# Mengganti working directory (kalau di linux/mac seperti command cd)
WORKDIR $GOPATH/src/myapp/

# Melakukan copy file dari folder saat ini ke folder working directory
COPY . .

# eksekusi go get untuk mendapatkan semua library / depedensi yang kita gunakan
RUN go get -d -v

# Set environment spesifik untuk build
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 

# Melakukan build binary apps 
RUN go build -o /go/bin/app

# Step 2 - membuat image baru hanya untuk running apps kita dari hasil build di atas
# ini begunakan agar image container kita size nya kecil
FROM scratch

# Melakukan copy binary dari hasil build image sebelumnya ke image scratch ini
COPY --from=builder /usr/share/zoneinfo/Asia/Jakarta /etc/localtime
COPY --from=builder /go/bin/app /app

# Melakukan eksekusi binary apps. goodluck!
ENTRYPOINT ["/app"]