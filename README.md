# Project Description
Scrum estimation ceremony back-end service implementation.

[TODO] Supports multi-user concurrent access through WebSockets and OAuth2 authentication.

# Run Application
Run commands in project root directory:

```bash
go run main.go run -h localhost -p 8000
```

# Run In Container
Run commands in project root directory. It creates `image version 0.0.1` of a server accepting connection on local machine `port 8000`.

```bash
docker build -f ./build/package/Dockerfile -t estimation-ceremony-backend:0.0.1 ./
docker run -p 8000:8000 estimation-ceremony-backend:0.0.1
```

Optionally pass container name option `--name estimation-ceremony-backend-v0.0.1`.

# Use Client
Install command line client as described in `https://github.com/vi/websocat`.

```bash
websocat ws://127.0.0.1:8000/{websockets-endpoint}
```

# Project Structure
```bash
.
|-- build/package                   # Provides Docker configuration files
|-- cmd/run                         # Run application command implementation
|-- deoployments                    # Provides CloudFormation/Kubernetes provisioning files 
|-- internal                        # Application business logic source code
|   |-- ...
|-- .gitignore
|-- go.mod
|-- go.sum
|-- README.md
```

# TODO
1. Keycloak server integration
