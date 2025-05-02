# Документация и отказоустойчивость 

**Скриншот работающего API**
[Скриншот работающего API](/images/image.png)

**Ссылка на успешный запуск пайплайна в CI/CD** 
https://github.com/AnnaPanchekhina/book/actions/runs/14802306105

## Запуск проекта

**Первый вариант:** \
`docker-compose up --build -d` \
Проверка: \
`curl http://localhost:4000/v1/healthcheck` 

**Второй вариант:** \
`chmod +x deploy.sh` \
`./deploy.sh` 

Для проверки: \
`chmod +x healthcheck.sh` \
`healthcheck.sh`

## CI/CD и деплой

**Необходимо добавить следующие секреты в GitHub:** 
* SSH_USER
* SSH_PRIVATE_KEY 
* SSH_PORT
* SSH_HOST 
* DB_DSN 
* SONAR_TOKEN
* SONAR_HOST_URL

## Примеры запросов к API

[http://localhost:4000/v1/healthcheck](/images/curl.png)


**В случае, если приложение не подключается к БД:**

Посмотреть логи \
`docker logs app-db-1`

**В случае, если контейнер падает при запуске** 
Посмотреть логи \
`docker logs app-app-1`