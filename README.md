# To-Do List | REST API Service

## Описание
Этот проект представляет собой RESTful API для управления задачами (To-Do List), реализованный на языке Go с использованием фреймворка Echo и базой данных PostgreSQL.

## Установка и запуск
1. Склонируйте репозиторий:

    ```bash
    git clone https://github.com/yourusername/todo-list-go.git
    ```

2. Перейдите в директорию проекта:

    ```bash
    cd todo-list-go
    ```

3. Установите зависимости:

    ```bash
    go mod tidy
    ```

4. Настройте переменные окружения в файле `.env` или оставить по-умолчанию

5. Запустить docker контейнер с БД:

   ```bash
   docker compose  up -d
   ```

6. Запустите приложение:

    ```bash
    go run cmd/task-app/main.go
    ```

Приложение автоматически выполнит миграцию базы данных при первом запуске.