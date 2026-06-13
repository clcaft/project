FROM golang:1.22-alpine

# Устанавливаем git
RUN apk add --no-cache git

WORKDIR /app

# Копируем все файлы с модулями
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем остальной код
COPY . .

# Собираем приложение (с явным указанием модулей)
RUN go build -mod=mod -o api ./cmd/main.go

EXPOSE 8080

CMD ["./api"]