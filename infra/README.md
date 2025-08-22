# Infrastructure Setup

This directory contains all the infrastructure configuration for the Event Manager application.

## Directory Structure

```
infra/
├── docker-compose.yml    # Main Docker Compose configuration
└── nginx/
    ├── nginx.conf      # Nginx main configuration
    ├── certs/          # SSL certificates (not version controlled)
    └── conf.d/         # Additional Nginx configuration files
```

## Getting Started

1. **SSL Certificates**
   - Place your SSL certificates in the `nginx/certs` directory:
     - `yourdomain.com.crt`
     - `yourdomain.com.key`
   - For development, you can use self-signed certificates or tools like `mkcert`.

2. **Environment Variables**
   - Copy `.env.example` to `.env` and update the values as needed.
   - The API service expects certain environment variables to be set.

3. **Running the Application**
   From the project root, you can use these commands:
   ```bash
   # Start all services
   docker-compose -f infra/docker-compose.yml up -d
   
   # View logs
   docker-compose -f infra/docker-compose.yml logs -f
   
   # Stop all services
   docker-compose -f infra/docker-compose.yml down
   ```

   For convenience, you can create an alias:
   ```bash
   alias dcp='docker-compose -f infra/docker-compose.yml'
   # Then use: dcp up -d
   ```

## Services

- **web**: Frontend application (Vite/React)
  - Access: http://localhost:3000

- **api**: Backend API (Go)
  - Access: http://localhost:8080
  - API Base URL: /api

- **nginx-proxy**: Reverse proxy with SSL termination
  - HTTP: http://yourdomain.com (redirects to HTTPS)
  - HTTPS: https://yourdomain.com

## Development Workflow

- The frontend and backend services are configured for development with hot-reloading.
- Changes to the source code will automatically trigger rebuilds in development mode.
- For production builds, set the appropriate environment variables and build flags.

## Troubleshooting

- If you encounter permission issues with Docker volumes, ensure your user has the necessary permissions.
- Check container logs for errors: `docker-compose -f infra/docker-compose.yml logs <service_name>`
- To completely reset the environment:
  ```bash
  docker-compose -f infra/docker-compose.yml down -v
  docker system prune -f
  ```
