# Project Description
Scrum estimation ceremony back-end service implementation.

[TODO] Supports multi-user concurrent access through WebSockets and OAuth2 authentication.

# Run Server
Run commands in project root directory:
```bash
go build -o estimation-ceremony-backend cmd/estimation-ceremony/main.go
./estimation-ceremony-backend --domain {domain} --port {port}
```

# Run Container
Run commands in project root directory:
```bash
docker build -f ./build/package/Dockerfile -t estimation-ceremony-backend:{version} .
docker run -p 0.0.0.0:{port}:8000 estimation-ceremony-backend:{version}
```
Optionally run container with a name `--name estimation-ceremony-backend-v{version}`.

# Use Client
Install globally third party client and run command:
```bash
go install github.com/hashrocket/ws@latest
ws ws://{domain}:{port} 
```

# Project Structure
```bash
.
|-- build/package                   # Provides Docker container definition
|-- cmd/gleos-estimation            # Provides application command line runner
|-- deoployments                    # Provides CloudFormation/Kubernetes deployment definitions
|-- internal
|   |-- app                         # Features code location
|   |-- pkg                         # Shared code location
|-- .gitignore
|-- go.mod
|-- go.sum
|-- README.md
```

# TODO
1. Keycloak server integration
