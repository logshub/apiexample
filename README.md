# API Example

### Installation

Copy `config.example.toml` to `config.toml` and provide valid configuration.

```
make install
nano config.toml
make test
```

Running PostgreSQL: `make postgres-run`

### Running

The application will listen on the port provided in configuration

```
make run
```

### Requests

1. Adding new user.

```
curl -i -XPOST 127.0.0.1:1234/users \
    -d '{"name":"Greg", "active":true}' \
    -H "Content-Type: application/json"
```

2. Deleting user.

```
curl -i -XDELETE 127.0.0.1:1234/users/96e1a704-d2cf-4d9a-8fcb-c28b6d1a6f16
```

3. Reading users list.

```
curl -i 127.0.0.1:1234/users
curl -i 127.0.0.1:1234/users?active=0
curl -i 127.0.0.1:1234/users?active=1
```
