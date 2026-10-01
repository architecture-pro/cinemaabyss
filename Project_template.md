## Изучите [README.md](README.md) файл и структуру проекта.

## Задание 1

1. Спроектируйте to be архитектуру КиноБездны, разделив всю систему на отдельные домены и организовав интеграционное взаимодействие и единую точку вызова сервисов.
   Результат представьте в виде контейнерной диаграммы в нотации С4.
   Добавьте ссылку на файл в этот шаблон
   [ссылка на файл](docs/c4/container/to-be/to-be-container.puml)

## Задание 2

### 1. Proxy

DONE

### 2. Kafka

![tests 1](docs/screenshots/tests-1.png)
![tests 2](docs/screenshots/tests-2.png)
![Kafka topics](docs/screenshots/kafka-topics.png)

## Задание 3

Команда начала переезд в Kubernetes для лучшего масштабирования и повышения надежности.
Вам, как архитектору осталось самое сложное:

- реализовать CI/CD для сборки прокси сервиса
- реализовать необходимые конфигурационные файлы для переключения трафика.

### CI/CD

DONE

### Proxy в Kubernetes

#### Шаг 1

DONE

#### Шаг 2

![event-service](docs/screenshots/events-kafka.png)
![event-service](docs/screenshots/1-kubernetes-tests.png)
![event-service](docs/screenshots/2-kubernetes-tests.png)
![event-service](docs/screenshots/3-kubernetes-tests.png)

#### Шаг 3

![movies-api-kubernetes](docs/screenshots/movies-api-kubernetes.png)
![movies-api-kubernetes](docs/screenshots/proxy-migration-0-monolith.png)
![movies-api-kubernetes](docs/screenshots/proxy-migration-100-monolith.png)

## Задание 4

![api-movies](docs/screenshots/api-movies.png)
![helm-list](docs/screenshots/helm-list.png)

# Задание 5

![circuit-breaker-1](docs/screenshots/circuit-breaker-1.png)
![circuit-breaker-2](docs/screenshots/circuit-breaker-2.png)
