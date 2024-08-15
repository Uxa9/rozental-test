# Rozental Test Backend

Проект содежит в себе тестовую сборку сервера
Для запуска проекта необходимо:

> `Go` версии 1.22
> `MySQL` версии 8.0 (может можно и другую, я не проверял)
____
Чтобы запустить проект необходимо:

> Скопировать себе директорию проекта

> Здесь находится файл для запуска сервера [text](internal/cmd/api/main.go)

> Файл конфигурации находится в [text](internal/config/config.yml)

____

Для удобства в репозитории расположены два файла:

> rg-test.postman_collection.json - Конфигурация для Postman, в котором прописаны все роуты, а также допустимые параметры для обращения к ним

> inventory_operations.sql - Дамп БД, использованной для разработки

____

Доступные роуты:

- /api
    - /user
        - POST   - создание пользователя [raw json] {"username": []string}
        - GET    - получение пользователей [params] {id: []int}
        - PUT    - изменение пользователя [raw json] {"Name": []string; "Id" : []int}
        - DELETE - удаление пользователя [params] {id: []int}
    - /inventory-operation
        - /getOperations POST - получение транзакций с оборудованием [raw json] 
            > {
            >   "id": []int; 
            >   "src_executor": []int; 
            >   "dst_executor": []int; 
            >   "request_time": ["ГГГГ-ММ-ДД ЧЧ-ММ-СС" "ГГГГ-ММ-ДД ЧЧ-ММ-СС"];
            >   "status_time": ["ГГГГ-ММ-ДД ЧЧ-ММ-СС" "ГГГГ-ММ-ДД ЧЧ-ММ-СС"]; 
            >   "status": ["accept","cancel","await","reject"]
            > }
        - /getInventory POST - получения оборудования с транзакциями [raw json]
            > {
            >   "id": []int; 
            >   "src_executor": []string; 
            >   "dst_executor": []int;
            >   "status": ["accept","cancel","await","reject"]
            > }
