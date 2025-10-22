# 🛰 Proxy Request Server (Go)

## 📖 Overview

This is a simple HTTP proxy server built in Go.  
It accepts JSON requests from clients, forwards them to external services, receives responses, and returns results in JSON format.  
All requests and responses are stored locally in memory using `sync.Map`.

---

## 🚀 Features

- Accepts HTTP requests in JSON format
- Forwards requests to third-party services
- Returns structured JSON responses
- Saves requests and responses in memory
- Input validation
- Containerized with Docker
- Supports deployment via Render

---

## 🧩 Request Format

```json
{
  "method": "GET",
  "url": "http://google.com",
  "headers": {
    "User-Agent": "GoProxyClient"
  }
}
```
