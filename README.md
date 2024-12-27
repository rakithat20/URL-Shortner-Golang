# Go URL Shortener

A high-performance URL shortening service built with Go and Fiber framework.

## Features

- URL shortening with custom slug generation
- PostgreSQL for persistent storage
- In-memory caching with 12-hour TTL
- Automatic cleanup of URLs older than 72 hours
- URL analytics tracking
- RESTful API endpoints

## Planned Improvements

- [ ] Docker containerization
- [ ] Load balancing
- [ ] Metrics and monitoring
- [ ] Rate limiting

## Tech Stack

- Go 1.21+
- Fiber (web framework)
- PostgreSQL
- In-memory cache

## API Endpoints

```
GET  /:slug         - Retrieve and redirect to original URL
POST /api/urls      - Create shortened URL
GET  /api/urls/:slug - Get URL analytics (total clicks)
```

### Example Responses

```json
// POST /api/urls
Request:
{
    "url": "https://example.com/very-long-url"
}

Response:
{
    "Id": "2",
    "slug": "abc123",
    "shortUrl": "http://domain.com/abc123",
    "long_url": "https://example.com/very-long-url",
    "expires_at": "2024-12-30T15:04:05Z"
}

// GET /api/urls/:slug
Response:
{   "Id": "2",
    "shortUrl": "abc123",
    "long_url": "https://example.com/very-long-url",
    "created_at": "2024-12-27T15:04:05Z",
    "total_clicks": 42
}
```

## Quick Start

```bash
# Clone repository
git clone https://github.com/yourusername/url-shortener.git

# Install dependencies
go mod download

# Set up PostgreSQL
createdb url_shortener

# Configure environment variables
export DB_HOST=localhost
export DB_USER=postgres
export DB_PASSWORD=yourpassword
export DB_NAME=url_shortener
export DB_PORT=5432

# Run application
go run main.go
```

## Architecture

The service uses a layered architecture:

- Router Layer (HTTP handlers)
- Service Layer (business logic)
- Storage Layer (PostgreSQL + cache)

Cache implementation uses an in-memory store with mutex locks for thread safety, backed by PostgreSQL for persistence.


## License

MIT
