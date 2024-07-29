# Message Service

## Overview

This project is a message processing microservice with a web API. It handles messages by storing them in a PostgreSQL database and processing them with Kafka. It includes both an API service and a consumer service.

## Getting Started

### Prerequisites

- Docker and Docker Compose installed on your machine.

### Setup

1. **Clone the Repository**

```bash
git clone https://github.com/milkaxq/message_service.git
cd message_service
```

2. **Configure Environment Variables**

Copy the example environment file to .env:

``` bash
cp .env.example .env
```
Update the .env file with your PostgreSQL credentials if necessary.

3. **Build and Run Services**

Use Docker Compose to build and start the services:

``` bash
docker-compose up --build
```

# Api Endpoints u can find in Api folder