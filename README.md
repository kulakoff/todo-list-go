# To-Do List | REST API Service

## Описание
Это REST API сервис для управления задачами (To-Do List), разработанный на языке Go с использованием фреймворка Echo. Сервис предоставляет возможность создавать, просматривать, обновлять и удалять задачи.

## Функциональность

### 1. Создание задачи
- **Метод:** `POST /tasks`
- **Описание:** Создает новую задачу.
- **Запрос:**
  ```json
  {
    "title": "string",
    "description": "string",
    "due_date": "string (RFC3339 format)"
  }
