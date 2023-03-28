# Project Description
Scrum estimation ceremony tool.

Provides back-end service implementation.

Supports multi-user concurrent access through WebSockets and OAuth2 authentication.

# Technology Stack
1. Go v1.20
2. nhooyr/websocket

# Project Structure
```bash
.
|-- cmd/gleos-estimation            # Provides application command line runner
|-- internal
|   |-- app                         # Features code location
|   |-- pkg                         # Shared code location
|-- .gitignore
|-- go.mod
|-- go.sum
|-- README.md
```

# TODO
1. Websockets core logic (IN_PROGRESS)
2. Keycloak server integration
