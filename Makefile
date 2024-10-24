docker-up:
	docker compose up

docker-down:
	docker compose down

rebuild:
	@docker system prune -f
	docker compose up
