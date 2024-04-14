# Используем официальный образ Go в качестве базового
FROM golang:1.19-alpine as builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем остальные файлы проекта
COPY . .

# Собираем приложение
RUN go build -o main ./cmd/main.go

# Выставляем порт, на котором будет слушать приложение
EXPOSE 8080

# Запускаем приложение при старте контейнера
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]