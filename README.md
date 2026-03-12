# POLIGLOTIM

> Foreign language courses API

## 🚀 Quick Start

### With Docker

## build from directory
```bash
docker build --tag ghcr.io/imirjar/poliglotim-api:latest .
```

## run from GitHub registry
```bash
docker run -d -p 3000:3000 ghcr.io/imirjar/poliglotim-api:latest
```

### With Docker Compose
```yaml
services:
  api:
    image: ghcr.io/imirjar/poliglotim-api:latest
    ports: ["3000:3000"]
    environment:
      - PORT=3000
      - DB_CONN=${DB_CONN}
```

## 📖 What is Poliglotim?

Poliglotim API Gateway provides secure access to educational content - courses, chapters, and lessons. It handles authentication, authorization, and serves as the main entry point for the Poliglotim learning platform.


## ⚙️ Configuration

| Env Variable | Purpose | Default |
|--------------|---------|---------|
| `PORT` | Server port | `3000` |
| `DB_CONN` | Database connection  | `postgresql://{login}:{password}@{db_host}/{db_name}
` |

## 📚 Swagger Docs

Interactive API documentation available at:
```
http://localhost:3000/swagger/
```

## 🛠 For Developers

This is an open-source project! Contributions welcome.

### Build & Deploy
- Automatic Docker builds via GitHub Actions
- Images published to GitHub Container Registry
- Swagger docs auto-generated from code annotations


## 📄 License

MIT © reDBeaver Team

---

**Getting Help**: Open an issue or check `/swagger` for API details