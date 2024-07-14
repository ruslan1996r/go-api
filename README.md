# Start
- `go mod download` - install dependencies
- `docker run --rm --name go-api -d -p 5434:5432 -e POSTGRES_HOST_AUTH_METHOD=trust postgres` - run Database
- `make run-hot` - run server in hot-reload mode
- `make docs` - regenerate docs
    - Docs - `http://localhost:8080/swagger/index.html`

