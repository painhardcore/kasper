Вариант №1
Используя только средства стандартной библиотеки, разработать HTTP-сервер. Сервер должен:
	На каждый запрос возвращать значение счетчика, считающего общее количество обработанных сервером запросов за последние 60 секунд;
	Продолжать возвращать корректное значение счетчика после перезапуска приложения (используя персистентное хранилище.