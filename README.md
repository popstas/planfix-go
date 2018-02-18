# planfix-go
[![Build Status](https://travis-ci.org/popstas/planfix-go.svg?branch=master)](https://travis-ci.org/popstas/planfix-go)
[![Coverage Status](https://coveralls.io/repos/github/popstas/planfix-go/badge.svg?branch=master)](https://coveralls.io/github/popstas/planfix-go?branch=master)

Клиент API для Планфикса на Go

Все запросы и ответы типизированы, по возможности сохранено именование из [документации](https://planfix.ru/docs/Список_функций)

Структуры запросов и ответов лежат в [planfix/functions.go](planfix/functions.go)

Примеры ответов лежат в [tests/fixtures](tests/fixtures)

Вызовы к API лежат в [planfix/functions.go](planfix/functions.go)

Дополнительные функции на основе основных лежат в [planfix/functions_extra.go](planfix/functions_extra.go)

## Реализованные функции
- auth.login
- action.get
- action.getList
- action.add
- analitic.getList
- task.get

## Дополнительные функции
- GetAnaliticByName

Пример использования в [main.go](main.go)
