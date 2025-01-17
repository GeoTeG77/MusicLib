# MusicLib (Тестовое задание)

Проект для управление своей музыкальной библиотекой.

## Оглавление

- [Описание](#описание)
- [Функциональные возможности](#функциональные-возможности)
- [Установка](#установка)
- [Использование](#использование)
---

## Описание

Приложение для управления музыкальной библиотекой в локальном хранилище при помощи браузера.

## Функциональные возможности

- Возможность добавлять, удалять и изменять записанные песни.
- Возможность получить требуемый куплет запрашиваемой песни.
- Возможность получить данные библиотеки с фильтрацией по основным полям и пагинацией.
- Возможность добавить новую песню по запросу к API определенного сервиса
---

## Установка

Для установки приложения необходимо выполнить следующие шаги:
1. Установить Docker и Docker-compose.
2. Установить git.
2. Открыть терминал в требуемой папке и скопировать проект с gitHub
git clone https://github.com/GeoTeG77/MusicLib.git
4. Запустить в этом же терминале команду docker-compose up --build

## Использование

После запуска приложения API в формате Swagger будет доступен по адресу:
http://localhost:8080/swagger/index.html#/
