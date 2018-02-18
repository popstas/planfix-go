# planfix-go
Клиент API для Планфикса на Go

Все запросы и ответы типизированы, по возможности сохранено именование из [документации](https://planfix.ru/docs/Список_функций)

Структуры запросов и ответов лежат в [planfix/functions.go](planfix/functions.go)

Примеры ответов лежат в [tests/fixtures](tests/fixtures)

Вызовы к API лежат в [planfix/functions.go](planfix/functions.go)

## Реализованные функции
- auth.login
- action.get
- action.getList

Пример использования в [main.go](main.go)