[![CircleCI](https://circleci.com/gh/painhardcore/kasper.svg?style=svg)](https://circleci.com/gh/painhardcore/kasper)
### build

Use docker to build image

```
docker build . -t "counter" 
```
### run

Use docker to run on 8080 port

```
docker run -p 8080:8080 counter 
```


Вариант №1

Используя только средства стандартной библиотеки, разработать HTTP-сервер. Сервер должен:

* На каждый запрос возвращать значение счетчика, считающего общее количество обработанных сервером запросов за последние 60 секунд;

* Продолжать возвращать корректное значение счетчика после перезапуска приложения (используя персистентное хранилище.
