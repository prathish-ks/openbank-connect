# 🏦 OpenBank Connect

A mock Open Banking Gateway built in Go using Chi, simulating OAuth2 login, account aggregation, and internal microservice integration.

## 🌐 Features

- OAuth2-based login simulation
- Chi-based router with clean middleware
- Protected endpoints: `/accounts`, `/transactions`
- Risk Scoring Microservice integration
- Dockerized and ready to deploy

## 🚀 Tech Stack

- Go 1.24
- Chi router
- In-memory token store
- Internal HTTP microservices
- Docker + Compose

## 🧪 Run Locally

```bash
docker-compose up --build

Visit:

GET /auth/callback?code=test

GET /accounts

GET /transactions

Project Structure: 

cmd/            → app entry point
internal/
  handler/      → HTTP route handlers
  middleware/   → Auth middleware
  client/       → Risk engine client
  service/      → Token store

🎯 Demo Use Case
Simulates a user securely logging in and viewing their bank transactions, each enriched with a risk score using an internal scoring engine.

👨‍💻 Author
Built by Prathish Kumar Srinivasan to demonstrate production-quality Go APIs and microservices.