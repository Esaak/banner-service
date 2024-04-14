````
banner-service/
├── cmd/
│   └── main.go
├── configs/
│   └── config.go
├── internal/
│   ├── banner/
│   │   ├── delivery/
│   │   │   └── http/
│   │   │       └── handler.go
│   │   ├── repository/
│   │   │   └── postgres/
│   │   │       ├── repository_mocks
│   │   │           └── repository_mocks.go
│   │   │       └── repository.go
│   │   └── usecase/
│   │       └── usecase.go
│   ├── models/
│   │   └── models.go
│   └── server/
│       └── server.go
├── pkg/
│   ├── auth/
│   │   ├── auth_mocks/
│   │   │   └──auth_mock.go
│   │   └── auth.go
│   ├── logger/
│   │   └── logger.go
│   └── postgres/
│       └── postgres.go
├── migrations/
├── Makefile
├── README.md
├── go.mod
└── go.sum
````
1. Клонируйте репозиторий:
````
git clone https://github.com/your-username/banner-service.git
````
2. Перейдите в директорию проекта:
````
cd banner-service
````
3. Установите зависимости:
````
go mod download
````
4. Создайте базу данных и примените миграции:
````
make migrate-up
````
5. Соберите приложение:
````
make build
````
6. Запустите приложение:
````
make run
````
7. Для запуска тестов выполните:
````
make test
````
**Для развертывания приложения с использованием Docker выполните следующие шаги:**

1. Соберите Docker-образ:
````
make docker-build
````
2. Запустите контейнеры приложения и базы данных:
````
docker-compose up -d
````