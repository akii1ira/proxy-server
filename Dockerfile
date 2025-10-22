x# Используем официальный образ Go
FROM golang:1.22-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum (если есть) и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарник
RUN go build -o proxy-server .

# Указываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./proxy-server"]
