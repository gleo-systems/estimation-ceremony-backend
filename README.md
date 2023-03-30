# Project Description
Scrum estimation ceremony back-end service implementation.

[TODO] Supports multi-user concurrent access through WebSockets and OAuth2 authentication.

# Run Server
```bash
go build -o estimation-ceremony cmd/estimation-ceremony/main.go
./estimation-ceremony --domain {domain} --port {port}
```

# Use Client
Install and run third party client.

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
