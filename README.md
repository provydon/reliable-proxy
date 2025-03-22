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

## 🚀 Deployment Methods

Choose the deployment method that works best for your needs:

### 1️⃣ Using Pre-built Binaries (Quickest Method)

Download a pre-built binary for your platform:

```bash
# Download the latest release for your OS (examples below)
# For Linux x86_64
curl -L https://github.com/provydon/reliable-proxy/releases/latest/download/reliable-proxy_Linux_x86_64.tar.gz -o reliable-proxy.tar.gz

# For macOS Intel
curl -L https://github.com/provydon/reliable-proxy/releases/latest/download/reliable-proxy_Darwin_x86_64.tar.gz -o reliable-proxy.tar.gz

# For macOS Apple Silicon
curl -L https://github.com/provydon/reliable-proxy/releases/latest/download/reliable-proxy_Darwin_arm64.tar.gz -o reliable-proxy.tar.gz

# For Windows
curl -L https://github.com/provydon/reliable-proxy/releases/latest/download/reliable-proxy_Windows_x86_64.zip -o reliable-proxy.zip
```

Then install it:

```bash
# For Linux/macOS:
tar -xzf reliable-proxy.tar.gz
chmod +x reliable-proxy
sudo mv reliable-proxy /usr/local/bin/

# For Windows:
# Extract the ZIP file and add the executable to your PATH
```

Run the proxy:

```bash
# Basic usage
reliable-proxy

# With a specific target API
reliable-proxy --target-api-url=https://target-api.com
```

### 2️⃣ Using Docker (Simple & Portable)

Pull and run the Docker image:

```bash
# Pull the image
docker pull ghcr.io/provydon/reliable-proxy:latest

# Run the container
docker run -p 8080:8080 ghcr.io/provydon/reliable-proxy:latest

# With a target API specified
docker run -p 8080:8080 -e TARGET_API_URL="https://target-api.com" ghcr.io/provydon/reliable-proxy:latest
```

### 3️⃣ Building from Source (For Development)

```bash
# Clone the repository
git clone https://github.com/provydon/reliable-proxy.git
cd reliable-proxy

# Build the executable
go build -o reliable-proxy

# Run the proxy
./reliable-proxy
```

### 4️⃣ Cloud Deployment Options

#### Deploy to Render (Free Tier Available)

1. Sign up for [Render](https://render.com/)
2. Create a new Web Service
3. Connect your GitHub repository or use the pre-built binary
4. Configure environment variables if needed
5. Deploy and get your proxy URL

#### Deploy to Railway

1. Sign up for [Railway](https://railway.app/)
2. Create a new project from GitHub
3. Select your repository
4. Configure environment variables
5. Deploy and get your proxy URL

#### Deploy to Fly.io (Regional selection available)

1. Install the Fly CLI: `curl -L https://fly.io/install.sh | sh`
2. Log in: `fly auth login`
3. Create an app: `fly launch`
4. Choose your region during setup
5. Deploy: `fly deploy`

---

## 💻 Using Reliable Proxy

Once deployed, you can use Reliable Proxy in several ways:

### Basic Usage

Access your deployed proxy with:

```bash
# Replace with your actual deployment URL or localhost:8080 if running locally
curl -X GET "https://your-proxy-url.com/some-path" \
  -H "target-api-url: https://target-api.com"
```

The request will be forwarded to `https://target-api.com/some-path` with all headers, query parameters, and body preserved.

### Setting a Default Target API

You can configure a default target API using environment variables:

```bash
# When running locally
TARGET_API_URL="https://target-api.com" reliable-proxy

# For cloud deployments, set the TARGET_API_URL environment variable
# in your deployment platform's settings
```

When a default target API is configured, you can omit the header:

```bash
curl -X GET "https://your-proxy-url.com/some-path"
```

### 🎮 Live Demo: Access Region-Restricted APIs Instantly!

### Try this yourself: 

**1️⃣ First, try accessing a US-restricted website directly:**
```bash
# Try accessing PeacockTV's sports page directly - you'll get blocked outside the US
curl -X GET "https://www.peacocktv.com/sports"
```

**Result:** ❌ *Access denied - "Unavailable In Your Region" page appears*

**2️⃣ Now, try again using our already deployed proxy in US East:**
```bash
# The same request, but through our US-based proxy on Render
curl -X GET "https://reliable-proxy.onrender.com/sports" \
  -H "target-api-url: https://www.peacocktv.com"
```

**Result:** ✅ *Success! You'll get the full PeacockTV sports page with upcoming events, Premier League, Big Ten basketball, and more - as if you were in the US.*

> 💡 **Without installing anything, you can immediately use our hosted proxy to bypass region restrictions.** Just replace the example with your actual target website to instantly access region-restricted content from anywhere.

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