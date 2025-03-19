# <div align="center">🌍 Reliable Proxy</div>

<div align="center">

[![License: Dual](https://img.shields.io/badge/License-Dual%20License-orange.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue.svg)](https://go.dev/dl/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?logo=docker&logoColor=white)](https://www.docker.com/)

*A simple, reliable proxy service for accessing region-restricted APIs from anywhere in the world.*
</div>

---

## 🤔 Why Reliable Proxy?

I created this tool while working remotely from Nigeria for a US company, where I faced challenges accessing US-restricted APIs needed for my work. Many popular paid proxies suffer from reliability issues because their IPs get blocked by target services.

### ✨ Reliable Proxy solves this by allowing you to:

- 🚀 Deploy your own proxy in your target region
- 🔒 Access region-restricted APIs reliably
- 💸 Run it for free on platforms like Render

---

## 🚀 Quick Start

## 🎮 Live Demo: Access Region-Restricted APIs Instantly!

### Try this yourself: 

**1️⃣ First, try accessing a US-restricted API directly:**
```bash
# Try accessing a US-only service API directly - you'll get blocked
curl -X GET "https://api.hulu.com/v1/shows/popular"
```

**Result:** ❌ *Access denied - This content is not available in your country.*

**2️⃣ Now, try again using our already deployed proxy in US East:**
```bash
# The same request, but through our US-based proxy on Render
curl -X GET "https://reliable-proxy.onrender.com/v1/shows/popular" \
  -H "target-api-url: https://api.hulu.com"
```

**Result:** ✅ *Success! The API responds with content as if you were accessing from the US.*

> 💡 **Without installing anything, you can immediately use our hosted proxy to bypass region restrictions.** Just replace the example API with your actual target API to instantly access region-restricted content from anywhere.

---


### 💻 Running Locally

```bash
go run main.go
```

The server runs on port 8080 by default.

### 🎯 Specifying Target APIs

The most flexible way to use Reliable Proxy is by specifying your target API URL in the request header:

```bash
curl -X GET http://localhost:8080/some/path -H "target-api-url: https://target-api.com"
```

This allows you to use a single proxy instance for multiple target APIs without any configuration changes.

> 💡 The request will be forwarded to `https://target-api.com/some/path` with all headers, query parameters, and body preserved.

### ⚙️ Environment Configuration

Alternatively, you can configure a default target API using a `.env` file in the project root:

```bash
# Copy the example file
cp .env.example .env

# Edit with your configuration
nano .env
```

Example `.env` file contents:
```
TARGET_API_URL=https://api.example.com
PORT=8080
```

When a default target API is configured, you can omit the header:

```bash
curl -X GET http://localhost:8080/some/path
```

### 🐳 Running with Docker

<details>
<summary>Click to expand Docker options</summary>

Basic usage:
```bash
docker build -t reliable-proxy .
docker run -p 8080:8080 reliable-proxy
```

With environment variables:
```bash
docker run -p 8080:8080 -e TARGET_API_URL="https://target-api.com" reliable-proxy
```

With persistent region cache and custom .env file:
```bash
docker run -p 8080:8080 \
  -v $(pwd)/data:/app/data \
  -v $(pwd)/.env:/app/.env \
  reliable-proxy
```

#### Including .env File in Docker Image

The Dockerfile includes a commented line for copying your .env file directly into the image:

```dockerfile
# COPY .env ./
```

You can uncomment this line if you want to build a Docker image with your configuration baked in:

1. Create your `.env` file with your settings
2. Uncomment the line in Dockerfile
3. Build your image: `docker build -t reliable-proxy .`

This is useful for creating pre-configured images for specific APIs or regions, but remember that environment variables in the `.env` file will be visible to anyone who has access to the image.
</details>

### 🔧 Setting Default Target API

```bash
TARGET_API_URL="https://target-api.com" go run main.go
```

This is useful when you want to use the proxy primarily for a specific API.

### ☁️ Deploy to the Cloud (Free Options)

| Platform | Description | Region Selection |
|----------|-------------|------------------|
| **Render** | Deploy as a Web Service | ✅ Region selection available |
| **Railway** | Deploy from GitHub repo | ✅ Region selection available |
| **Fly.io** | Deploy with their free tier | ✅ Regional selection available |

---

## 📚 Usage Examples

### ⚡ Basic Usage

Make requests to the proxy with the Target API URL in the header:

```bash
curl -X GET http://localhost:8080/some/path -H "target-api-url: https://target-api.com"
```

The request will be forwarded to `https://target-api.com/some/path` with all headers, query parameters, and body preserved.

If you've set a default `TARGET_API_URL` environment variable, you can omit the header:

```bash
curl -X GET http://localhost:8080/some/path
```

### 🔍 Example Curl Commands

<details open>
<summary><b>Check proxy status and region</b></summary>

```bash
curl http://localhost:8080/
```

Response: 
```json
{"status":"Reliable Proxy server is running","region":"New York, New York, US"}
```
</details>

<details open>
<summary><b>Make a GET request through the proxy</b></summary>

```bash
curl -X GET "http://localhost:8080/search?q=test" \
  -H "target-api-url: https://www.google.com"
```
</details>

<details open>
<summary><b>Make a POST request with JSON data</b></summary>

```bash
curl -X POST "http://localhost:8080/api/data" \
  -H "target-api-url: https://api.example.com" \
  -H "Content-Type: application/json" \
  -d '{"key": "value"}'
```
</details>

<details open>
<summary><b>Test using a deployed instance on Render (US West)</b></summary>

```bash
curl -X GET "https://reliable-proxy.onrender.com/users" \
  -H "target-api-url: https://jsonplaceholder.typicode.com"
```

This will proxy your request through a US West region, allowing you to access US-restricted APIs.
</details>

---

## 🧪 Running Tests

Run the tests with:

```bash
go test -v
```

This will run all the unit tests, including tests for the proxy handler, environment loading, and error handling.

---

## 📜 License

Reliable Proxy is available under a dual license:

- **Non-Commercial Use**: 
  > ✅ Free to use, modify, and contribute to for non-commercial purposes

- **Commercial Use**: 
  > 💼 Requires a separate license agreement with royalty terms

For full license details, see the [LICENSE](LICENSE) file. If you're interested in using Reliable Proxy for commercial purposes, please contact the copyright holder to arrange suitable terms.

---

## ✨ Features

| Feature | Description |
|---------|-------------|
| 🔄 **Full HTTP Support** | Forwards all HTTP methods (GET, POST, PUT, DELETE, PATCH) |
| 📋 **Preserves Request Data** | Maintains headers and query parameters |
| 📦 **Content Type Handling** | Handles various content types seamlessly |
| 🌎 **Region-specific Deployment** | For accessing geo-restricted APIs |
| 🌐 **Auto Region Detection** | With persistent caching for performance |
| ⚡ **Concurrent Processing** | For optimal performance under load |
| ⚙️ **Flexible Configuration** | Environment variables or header-based setup |
| 📝 **Simple Implementation** | Single-file core for easy deployment |

---

## 📋 Requirements

- Go 1.23+ (for direct execution)
- Docker (optional)

---

<div align="center">
<p>Made with ❤️ by <a href="https://github.com/providenceifeosame">Providence Ifeosame</a></p>

<a href="#-reliable-proxy">⬆️ Back to Top</a>
</div>