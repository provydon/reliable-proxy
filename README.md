# Reliable Proxy

A simple, reliable proxy service for accessing region-restricted APIs from anywhere in the world.

## Why Reliable Proxy?

I created this tool while working remotely from Nigeria for a US company, where I faced challenges accessing US-restricted APIs needed for my work. Many popular paid proxies suffer from reliability issues because their IPs get blocked by target services. Reliable Proxy solves this by allowing you to:

- Deploy your own proxy in your target region
- Access region-restricted APIs reliably
- Run it for free on platforms like Render

## Quick Start

### Running Locally

```
go run main.go
```

The server runs on port 8080 by default.

### Specifying Target APIs

The most flexible way to use Reliable Proxy is by specifying your target API URL in the request header:

```
curl -X GET http://localhost:8080/some/path -H "target-api-url: https://target-api.com"
```

This allows you to use a single proxy instance for multiple target APIs without any configuration changes.

The request will be forwarded to `https://target-api.com/some/path` with all headers, query parameters, and body preserved.

### Environment Configuration

Alternatively, you can configure a default target API using a `.env` file in the project root. See the provided `.env.example` file for available options:

```
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

```
curl -X GET http://localhost:8080/some/path
```

### Running with Docker

Basic usage:
```
docker build -t reliable-proxy .
docker run -p 8080:8080 reliable-proxy
```

With environment variables:
```
docker run -p 8080:8080 -e TARGET_API_URL="https://target-api.com" reliable-proxy
```

With persistent region cache and custom .env file:
```
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

### Setting Default Target API

You can set a default target API URL using an environment variable:

```
TARGET_API_URL="https://target-api.com" go run main.go
```

This is useful when you want to use the proxy primarily for a specific API.

### Deploy to the Cloud (Free Options)

- **Render**: Deploy as a Web Service in your target region
- **Railway**: Deploy from GitHub repo with region selection
- **Fly.io**: Deploy with regional selection using their free tier

## Usage

Make requests to the proxy with the Target API URL in the header:

```
curl -X GET http://localhost:8080/some/path -H "target-api-url: https://target-api.com"
```

The request will be forwarded to `https://target-api.com/some/path` with all headers, query parameters, and body preserved.

If you've set a default `TARGET_API_URL` environment variable, you can omit the header:

```
curl -X GET http://localhost:8080/some/path
```

## Example Curl Commands

### Check proxy status and region
```bash
curl http://localhost:8080/
```
Response: `{"status":"Reliable Proxy server is running","region":"New York, New York, US"}`

### Make a GET request through the proxy
```bash
curl -X GET "http://localhost:8080/search?q=test" \
  -H "target-api-url: https://www.google.com"
```

### Make a POST request with JSON data
```bash
curl -X POST "http://localhost:8080/api/data" \
  -H "target-api-url: https://api.example.com" \
  -H "Content-Type: application/json" \
  -d '{"key": "value"}'
```

### Test using a deployed instance on Render (US East)
```bash
curl -X GET "https://reliable-proxy.onrender.com/users" \
  -H "target-api-url: https://jsonplaceholder.typicode.com"
```

This will proxy your request through a US East region, allowing you to access US-restricted APIs.

## Running Tests

Run the tests with:

```bash
go test -v
```

This will run all the unit tests, including tests for the proxy handler, environment loading, and error handling.

## License

Reliable Proxy is available under a dual license:

- **Non-Commercial Use**: Free to use, modify, and contribute to for non-commercial purposes
- **Commercial Use**: Requires a separate license agreement with royalty terms

For full license details, see the [LICENSE](LICENSE) file. If you're interested in using Reliable Proxy for commercial purposes, please contact the copyright holder to arrange suitable terms.

## Features

- Forwards all HTTP methods (GET, POST, PUT, DELETE, PATCH)
- Preserves headers and query parameters
- Handles various content types
- Region-specific deployment for accessing geo-restricted APIs
- Automatic region detection with persistent caching
- Concurrent processing for optimal performance
- Default target API configuration via environment variable or .env file
- Simple, single-file implementation

## Requirements

- Go 1.16+ (for direct execution)
- Docker (optional)