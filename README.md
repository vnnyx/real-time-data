# Real-Time Data

This project fetches real-time news data, providing up-to-the-minute news articles and updates.

## Prerequisites

- Go 1.22.0 or newer
- Kafka
- [Data processing](https://github.com/vnnyx/ai-hub)

## Installation

1. **Clone the repository**:

   ```sh
   git clone <repository-url>
   cd <repository-directory>
   ```

2. **Install package:**

   ```sh
   go mod download
   ```

3. **Set up environment variables:**

   Create a .env file in the root directory or use the provided .env.example as a template:

   ```
    NEWS_API_KEY="your-news-api-key-here"
    NEWS_API_BASE_URL="https://newsapi.org/v2"
    KAFKA_BROKER_URL="your-kafka-broker-url-here"
    KAFKA_USERNAME="your-kafka-username-here"
    KAFKA_PASSWORD="your-kafka-password-here"
    VECTOR_HOST="your-vector-host-here"
    NEWS_TOPIC="news"
   ```

## Usage

Run the application with the available commands:

```sh
go run main.go [command]
```

**Commands**

- `producer`: Fetches real-time news articles from various sources and publishes them to a Kafka topic.
- `consumer`: Consumes news articles from a Kafka topic and processes them for further analysis or storage.
