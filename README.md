Golang REST API с использованием Echo и GORM
Этот проект реализует базовый CRUD REST API для управления предметами, построенный на Go с использованием веб-фреймворка Echo и ORM GORM для взаимодействия с базой данных. Ниже приведено руководство по настройке и запуску проекта.

Функционал
Операции создания, чтения, обновления и удаления (CRUD) для модели Item.
Основная база данных — PostgreSQL.
Автоматическая миграция для модели Item с использованием GORM.

Стек технологий
Go: Язык программирования, используемый для бэкенда.
Gin: Легковесный веб-фреймворк для создания быстрых и масштабируемых API.
GORM: ORM-библиотека для Go, использующаяся для взаимодействия с базой данных.
PostgreSQL: База данных для хранения данных о предметах.

Установка
git clone https://github.com/maximusprimeavenger/golang-rest-api