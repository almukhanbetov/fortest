# Используем официальный образ Go
FROM golang:1.24

# Создаем директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod ./
COPY go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем остальные файлы
COPY . ./

# Собираем приложение
RUN go build -o main .

# Указываем порт
EXPOSE 8889

# Команда запуска
CMD ["./main"]
