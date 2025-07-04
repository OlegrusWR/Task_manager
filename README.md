#  Task Manager API

**Минималистичный HTTP-сервер для управления задачами на Go**  
_Асинхронное выполнение задач с отслеживанием прогресса_ (сс возможностью на масштабируемость)

## Быстрый старт

### Требования
- Go 1.21 или новее
- Утилита `curl` для тестирования API

### Установка
```bash
git clone https://github.com/yourusername/task-manager.git
cd task-manager
go run cmd/main.go
```
Сервер запустится на:
http://localhost:8080

## API Endpoints

### POST /tasks
*Создание новой задачи*

```bash
curl -X POST http://localhost:8080/tasks
```

Ответ (201 Created):

```json
{
  "id": "3f5e303d-0915-4332-9fe2-2ab087f20712",
  "status": "pending",
  "time_at_created": "2025-07-04T11:13:14.168567+03:00",
  "time_at_started": null,
  "time_at_completed": null,
  "result": "",
  "error": "",
  "in_process": "0s"
}
```

### GET /tasks
 *Получение списка всех задач*

 ```bash
curl http://localhost:8080/tasks
```

Ответ (200 OK):

```json
{
    "id": "3f5e303d-0915-4332-9fe2-2ab087f20712",
    "status": "running",
    "time_at_created": "2025-07-04T11:13:14.168567+03:00",
    "time_at_started": "2025-07-04T11:13:14.168603+03:00",
    "time_at_completed": null,
    "in_process": "1m32s"
  }
```
### GET /tasks/{id}

*Получение конкретной задачи*
``` bash
curl http://localhost:8080/tasks/{конкретный ID}
```

Ответ (200 OK):

```json
{
  "id": "3f5e303d-0915-4332-9fe2-2ab087f20712",
  "status": "completed",
  "time_at_created": "2025-07-04T11:13:14.168567+03:00",
  "time_at_started": "2025-07-04T11:13:14.168603+03:00",
  "time_at_completed": "2025-07-04T11:15:22.453821+03:00",
  "result": "Task completed successfully",
  "in_process": "2m8s"
}
```

### DELETE /tasks/{id}
*Удаление задачи*
```bash
curl -X DELETE http://localhost:8080/tasks/{конкретный ID}
```

## Технические детали
### Жизненный цикл задачи
- pending → Создана, ожидает в очереди

- running → Взята воркером на выполнение

- completed/failed → Финальный статус

### Особенности реализации
- Ограничение на 5 одновременных воркеров

- 20% вероятность ошибки выполнения

- Время выполнения: 3-5 минут
