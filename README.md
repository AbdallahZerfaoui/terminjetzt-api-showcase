# TerminJetzt API – Public Showcase
![Uptime](https://github.com/AbdallahZerfaoui/terminjetzt-api-showcase/workflows/health-check/badge.svg)
![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)
![Python](https://img.shields.io/badge/python-3.11-blue)

A **read-only REST API** that exposes real-time immigration-office appointment data for Heilbronn, Germany.  
This repository contains **client examples, infrastructure as code, and documentation**—all you need to consume or replicate the service.

---
## 🔍 Live Docs

Swagger UI → https://api.terminjetzt.com/docs  
ReDoc → https://api.terminjetzt.com/redoc

---

## 🔍 Live Endpoints
| Endpoint | Description |
|----------|-------------|
| `GET /docs` | Swagger UI |
| `GET /redoc` | ReDoc |
| `GET /appointments` | All appointments |
| `GET /appointments/latest` | Latest batch |
| `GET /appointments/{id}` | Single appointment |
| `GET /summary/by-date` | Daily counts |
| `GET /summary/by-month` | Monthly counts |
| `GET /status` | System heartbeat |

Base URL: `https://api.terminjetzt.com`

---

## ⚡ Quickstart

### cURL
```bash
curl https://api.terminjetzt.com/ | jq
```

