# Generate Stream Currency

Задача сервиса генерировать значения, хранить 100 последних и по GET-запросу отправлять их

Имеется:
    АПИ-метод(GET /health) для проверки статуса работы сервиса, возвращает 200 - если сервис доступен и работает.
    АПИ-метод(GET /values?target={currency_pair}) возвращает максимум 100 значений сделок

Для получения значений в необходимый промежуток времени для конкретной валютной паре реализован RPC метод  
    `internal/proto/server.proto`

Запуск с помощью docker-compose:  
    `make create_network` - если не настроен, необходим для взаимодействия сервисов  
    `make run` - старт сервиса  
    `make restart` - рестарт сервиса  
    `make kill` - стоп сервиса и удаление контейнеров  
    `make logs` - вывод логов


Запуск тестов:  
    `make test` - тесты  
    `make coverage` - генерация отчета о покрытии  
    `make lint` - проверка линтером

mockgen:  
    `mockgen -source={PATH_INTERFACE} -destination={PATH_DEST}`