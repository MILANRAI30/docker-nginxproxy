# 🔁 Multi-Service App with NGINX Reverse Proxy

This project sets up two microservices:

- 🟦 Service 1: a Go-based API
- 🟨 Service 2: a Python Flask-based API

An NGINX container acts as a reverse proxy to route incoming requests based on URL paths. All services are containerized using Docker and orchestrated with Docker Compose.

---

## 📦 Services Overview

| Path             | Description             | Technology |
|------------------|-------------------------|------------|
| `/service1`      | Go-based API            | Go (net/http) |
| `/service2`      | Python Flask API        | Python (Flask) |

Each service exposes:
- `GET /ping`: Health check endpoint
- `GET /hello`: Test message endpoint

---

## 🧱 Project Structure

.
├── docker-compose.yml
├── .gitignore
├── README.md
├── nginx
│ ├── nginx.conf
│ └── Dockerfile
├── service_1
│ ├── main.go
│ └── Dockerfile
├── service_2
│ ├── app.py
│ └── Dockerfile


🔁 How Routing Works

NGINX handles all requests and routes based on the path:

location /service1/ {
    proxy_pass http://service_1:8001/;
}
location /service2/ {
    proxy_pass http://service_2:8002/;
}

/service1/* → routed to Go service (port 8001)

/service2/* → routed to Flask service (port 8002)


❤️ Health Checks

Each service defines a health check on /ping in docker-compose.yml:

healthcheck:
  test: ["CMD", "wget", "-q", "--spider", "http://localhost:800X/ping"]

Services must pass health checks before NGINX is allowed to route traffic to them.


📜 Logging

NGINX logs every request with timestamp and path:

log_format main '$remote_addr - [$time_local] "$request" status=$status ...';

To view logs:

docker-compose logs nginx

📦 One-Command Setup

docker-compose up --build
