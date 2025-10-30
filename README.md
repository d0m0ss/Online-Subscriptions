# Online-Subscriptions

# 1) поднять только БД в фоне
docker compose up -d db

# 2) запустить миграции
docker compose run --rm migrate

# 3) запустить сервис
docker compose up -d --build app
