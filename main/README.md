# Weather CLI Tool with Redis Caching

This CLI tool fetches weather data and caches it using Redis for improved performance.

## Prerequisites

1. **Docker**: Ensure Docker is installed and running on your system.
2. **Go**: Install Go if you want to build the binary from source.

## Setup

### Start Redis with Docker
Run the following command to start a Redis container:
```bash
docker run -d --name wet-redis -p 6379:6379 redis:latest
```

## Build the Binary
If you want to build the binary from source, run:
```bash
go build -o wet .
```

# Move the Binary to Local Path
Move the binary to a directory in your PATH for easy access:
```bash
mv wet /usr/local/bin
```

# Usage
## Display Help
Run the following command to display the help menu:
```bash
wet help
```
## Fetch Weather Data
To fetch weather data for a specific location, run:
```bash
wet <location>
```
Replace <location> with the desired city name (e.g., wet Berlin). You can set your desired location in the main file. Change var location = <yourCity>

## Extended Menu
To access the extended menu, run:
```bash
wet ext
```

## Troubleshooting
# Redis Connection Issues
If you encounter a connection refused error, ensure the Redis container is running:
```bash
docker ps
```

If the container is not running, start it again:
```bash
docker start wet-redis
```

Verify Redis is Working
You can verify Redis is working by running:
```bash
redis-cli ping
```
It should return PONG.

## License
This project is licensed under the MIT License.