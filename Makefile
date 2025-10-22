# Сборка контейнера
build:
	docker-compose build

# Запуск контейнера
up:
	docker-compose up

# Перезапуск контейнера
restart:
	docker-compose down && docker-compose up --build

# Остановка контейнера
down:
	docker-compose down

# Очистка неиспользуемых ресурсов
clean:
	docker system prune -f
