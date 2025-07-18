
# 🚀 Mars Rover Challenge

A project to simulate the control of multiple rovers exploring the surface of Mars, developed as a solution for the **Backend Test**. 💥

---

## 📦 About the Project

The system reads mission instruction files, executes the movement of multiple rovers in parallel across a virtual plateau, and returns their final positions in **plain-text** format.

The complete flow:

1. 📂 Upload the instruction file.
2. 🧠 Parse the file → validate plateau and rovers.
3. ⚙️ Concurrently execute instructions with the Orchestrator.
4. 📦 Return HTTP response with the final positions in the same input order.

---

## 🚀 Tech Stack

- **Go 1.21+** (default)
- **Stdlib only:** no external libraries in core (as required by the test)
- **Testing:** `stretchr/testify`, `httptest`, `mock`
- **Build/Test Tools:** `go mod`, `go test`

---

## 📂 Project Structure

```
marsrover/
├── orchestrator/       # Orchestrates concurrent execution of rovers
├── rover/              # Logic for rover movement and direction
├── plateau/            # Validates and manages plateau boundaries
├── internal/http/      # Handlers, parser, formatter, HTTP errors, and server
├── cmd/server.go       # API HTTP entry point
└── README.md
```

---

## 💻 Makefile Commands

The project includes a Makefile with useful commands. You can view all available targets by running `make help`:

```
Usage:
  make [target]

Targets:
help                Display this help
install-tools       Install gofumpt, gocritic, and swaggo
lint                Run golangci-lint
format              Format code
test                Run all tests
test/unit           Run unit tests
test/coverage       Run tests, generate coverage report, and open in browser
test/coverage-browser  Open coverage report in browser
swagger             Generate Swagger docs
run                 Run HTTP server
clean               Remove cache files
```

---

## 🎲 How to Run the Project

```bash
# Clone this repository
git clone <https://github.com/uesleicarvalhoo/marsover>

# Navigate to the project directory
cd marsover

# Build and start the development dependencies with Docker Compose
docker compose build && docker compose up -d

# This will start the container:
# backend       localhost:5000  -> application backend
```

---

## 🧪 Testing

### Unit and Integration Tests

```bash
make test/unit
```

### Full unit and integration test (upload → response)

```bash
make test
```

---

## 📡 API

### POST `/missions`

Accepts a `.txt` file containing mission instructions.

#### Example file

```
5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
```

#### cURL

```bash
curl -X POST -F 'file=@mission.txt' http://localhost:8080/missions
```

#### Response

```
1 3 N
5 1 E
```

---

## ⚙️ Architecture

- **Plateau:** defines boundaries and validates coordinates.
- **Rover:** moves according to commands (`L`, `R`, `M`).
- **Orchestrator:** executes rovers concurrently (manages workers).
- **HTTP Layer:** parses file, handles requests, and returns plain-text responses.

---

## 🤝 Contributing

1. Fork this repository
2. Create your branch: `git checkout -b feature/your-feature`
3. Commit your changes: `git commit -m 'Add your feature'`
4. Push to your branch: `git push origin feature/your-feature`
5. Open a Pull Request

---

## 📜 Requirements Met

✅ No external libraries in core
✅ External libraries only for testing
✅ Concurrent execution supported
✅ Functional HTTP API with file upload
✅ Unit + integration tests (full flow coverage)

---

## 📝 License

This project is open-source and licensed under the MIT License.
