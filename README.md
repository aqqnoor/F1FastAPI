# 🏁 F1FastAPI — Go API для гонщиков Формулы-1

Проект на Go, с помощью которого можно управлять базой данных гонщиков Формулы-1.

## 🚀 Возможности

- 📄 Получить список всех гонщиков: `GET /drivers`
- 🔍 Получить гонщика по ID: `GET /drivers/{id}`
- ➕ Добавить нового гонщика: `POST /drivers`
- ❌ Удалить гонщика: `DELETE /drivers/{id}`
- 💾 Данные хранятся в JSON-файле (или PostgreSQL)

## 🛠️ Используемые технологии

- Go 1.21+
- `gorilla/mux` — маршрутизация
- `encoding/json` — работа с JSON
- (опционально) PostgreSQL и pgx

## 🧱 Структура

```bash
.
├── main.go
├── handlers/
│   └── driver.go
├── models/
│   └── driver.go
├── data/
│   └── drivers.json
└── README.md
