# TimeManagement
Is an application that will help you to do your own time management or to manage the working time of your employees.

Here the server is written in golang and the client uses Avalonia for the interface. You can bolt on your own interface if you like.

## How to build the project (server and client)
1. Install [Go](https://go.dev/doc/install).
2. Install [PostgreSQL](https://www.postgresql.org/download/).
3. Install [Visual Studio](https://visualstudio.microsoft.com/ru/ru/vs/) and wpf development tools with avalonia.
4. Copy and rename `example.env` to `.env` and fill in the necessary information.
5. Install all the dependencies of the server, by typing in the console command `go mod tidy`.
6. Type `./build.sh` in the root directory of the project.
7. In the build directory there are compiled files of the project.

## How to run the server in development mode
1. First 5 steps are the same as in the build section.
2. Type `go run ./src/server/main.go` in the root directory of the project.

# Russian
Это приложение, которое поможет вам сделать собственную управление временем или управлять рабочим временем сотрудников.

Здесь написан сервер на golang и клиент на Avalonia для интерфейса. Вы можете подключить свой интерфейс, если вам этого хочется.

## Как собрать (сервер и клиент)
1. Установите [Go](https://go.dev/doc/install).
2. Установите [PostgreSQL](https://www.postgresql.org/download/).
3. Установите [Visual Studio](https://visualstudio.microsoft.com/ru/ru/vs/) и инструменты для разработки WPF и avalonia.
4. Скопируйте и переименуйте `example.env` в `.env` и заполните необходимую информацию.
5. Установите все зависимости сервера, прописав в консоль команду `go mod tidy`.
6. Введите `./build.sh` в корневой директории проекта.
7. В директории build находятся скомпилированные файлы проекта.

## Как запустить сервер в режиме разработки
1. Первые 5 шагов как в секции сборки.
2. Введите `go run ./src/server/main.go` в корневой директории проекта.

## We have russian and english documentation

### Russian
[russian docs](./docs/ru/README.md)

### English
[english docs](./docs/en/README.md)