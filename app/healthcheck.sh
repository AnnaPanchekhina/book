#!/bin/bash

APP_URL="http://localhost:4000/v1/healthcheck"
MAX_RETRIES=10
RETRY_INTERVAL=5

echo "Проверка здоровья приложения..."

for ((i=1; i<=$MAX_RETRIES; i++)); do
    response=$(curl -s -o /dev/null -w "%{http_code}" $APP_URL || true)
    
    if [ "$response" = "200" ]; then
        echo "Приложение доступно. HTTP Status: $response"
        exit 0
    else
        echo "Попытка $i/$MAX_RETRIES: Приложение не отвечает (Status: ${response:-none})"
        sleep $RETRY_INTERVAL
    fi
done

echo "Ошибка: приложение не стало доступно после $MAX_RETRIES попыток"
exit 1