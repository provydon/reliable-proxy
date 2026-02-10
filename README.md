# <div align="center">🌍 Reliable Proxy</div>

<div align="center">

[![License: Dual](https://img.shields.io/badge/License-Dual%20License-orange.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue.svg)](https://go.dev/dl/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?logo=docker&logoColor=white)](https://www.docker.com/)

*A simple, reliable proxy service for accessing region-restricted APIs from anywhere in the world.*
</div>

---

## 🤔 Why Reliable Proxy?

I created this tool while working remotely from My country for a US company, where I faced challenges accessing US-restricted APIs needed for my work. Many popular paid proxies suffer from reliability issues because their IPs get blocked by target services.

### ✨ Reliable Proxy solves this by allowing you to:

- 🚀 Deploy your own proxy in your target region
- 🔒 Access region-restricted APIs reliably
- 💸 Run it for free on platforms like Render

---
### See It In Action

**Without Proxy (Access Denied):**
```bash
# Try accessing the US-only API directly
curl -H "Accept: application/json" "https://us-only-api.onrender.com"
```

Result:
```json
{"error":"Access restricted to US only."}
```

**With Proxy (Success):**
```bash
# Same request through our proxy (query param)
curl "https://reliable-proxy.onrender.com/?target=https://us-only-api.onrender.com"
```

Result:
```json
{"message":"Hello from the US-only API!"}
```

## ✨ Key Features

- 🌎 **Region-specific deployment** for accessing geo-restricted APIs
- 🔄 **Full HTTP support** (GET, POST, PUT, DELETE, PATCH)
- 📋 **Preserves headers and query parameters** 
- 🌐 **Auto region detection** with caching
- ⚙️ **Flexible configuration** via query param, header, or env (all use `target` / `TARGET`)

## 🚀 Quick Installation

### 1️⃣ Pre-built Binaries (Fastest Method)

```bash
# For Linux x86_64
curl -L https://github.com/provydon/reliable-proxy/releases/latest/download/reliable-proxy_Linux_x86_64.tar.gz -o reliable-proxy.tar.gz

# For Linux ARM64
curl -L https://github.com/provydon/reliable-proxy/releases/latest/download/reliable-proxy_Linux_arm64.tar.gz -o reliable-proxy.tar.gz

# For macOS Intel (x86_64)
curl -L https://github.com/provydon/reliable-proxy/releases/latest/download/reliable-proxy_Darwin_x86_64.tar.gz -o reliable-proxy.tar.gz

# For macOS Apple Silicon (ARM64)
curl -L https://github.com/provydon/reliable-proxy/releases/latest/download/reliable-proxy_Darwin_arm64.tar.gz -o reliable-proxy.tar.gz

# For Windows x86_64
curl -L https://github.com/provydon/reliable-proxy/releases/latest/download/reliable-proxy_Windows_x86_64.zip -o reliable-proxy.zip

# For Windows ARM64
curl -L https://github.com/provydon/reliable-proxy/releases/latest/download/reliable-proxy_Windows_arm64.zip -o reliable-proxy.zip
```

```bash
# Extract and install (Linux/macOS)
tar -xzf reliable-proxy.tar.gz
chmod +x reliable-proxy
sudo mv reliable-proxy /usr/local/bin/

# Start the proxy
reliable-proxy
```

### 2️⃣ Using Docker

```bash
# Build and run
docker build -t reliable-proxy .
docker run -p 8080:8080 reliable-proxy
```

### 3️⃣ From Source

```bash
# Clone and enter the repository
git clone https://github.com/provydon/reliable-proxy.git && cd reliable-proxy

# Run directly
go run main.go
```
## 💻 Usage

### Basic Usage

> Pass the API to proxy via the **query param** `target`, the **header** `target`, or the **env var** `TARGET`. Query is checked first.

```bash
# Simple: query param (e.g. links, browser, any client)
curl "http://localhost:8080/?target=https://us-only-api.onrender.com"

# Or via header
curl -X GET "http://localhost:8080/" -H "target: https://us-only-api.onrender.com" -H "Accept: application/json"

# Or set a default with env
TARGET="https://us-only-api.onrender.com" reliable-proxy
```

## 🛠️ Troubleshooting

If you see `exec format error`, you downloaded the wrong binary for your system:

```bash
# Find your architecture
uname -m

# Download the correct version (example for macOS ARM64)
curl -L https://github.com/provydon/reliable-proxy/releases/latest/download/reliable-proxy_Darwin_arm64.tar.gz -o reliable-proxy.tar.gz
```

## 📜 License

- **Non-Commercial**: Free to use and modify
- **Commercial**: Requires license agreement

---

<div align="center">
<p>Made with ❤️ by <a href="https://github.com/providenceifeosame">Providence Ifeosame</a></p>
</div>
