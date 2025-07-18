
# 🚀 Mars Rover Challenge

Um projeto para simular o controle de múltiplos rovers explorando a superfície de Marte, desenvolvido como solução para o **Teste Backend**. 💥

---

## 📦 Sobre o projeto

O sistema lê arquivos com instruções de missão, executa o movimento de vários rovers em paralelo sobre um plateau virtual, e retorna suas posições finais em formato **plain-text**.

O fluxo completo:

1. 📂 Upload do arquivo com as instruções.
2. 🧠 Parsing do arquivo → validação do plateau e rovers.
3. ⚙️ Execução concorrente das instruções com o Orchestrator.
4. 📦 Resposta HTTP com as posições finais na mesma ordem de entrada.

---

## 🚀 Stack utilizada

- **Go 1.21+** (padrão)
- **Stdlib**: sem libs externas no core (conforme exigência do teste)
- **Testes:** `stretchr/testify`, `httptest`, `mock`
- **Build/Test Tools:** go mod, go test

---

## 📂 Estrutura do projeto

```
marsrover/
├── orchestrator/       # Orquestra execução concorrente dos rovers
├── rover/              # Lógica de movimento e direção dos rovers
├── plateau/            # Valida e gerencia limites do plateau
├── internal/http/      # Handlers, parser, formatter erros HTTP e server
├── cmd/server.go       # Entry point da API HTTP
└── README.md
```

---

## Comandos

O projeto conta com um Makefile com alguns comandos úteis, você pode consultar os comandos disponíveis através do `make help`

```
Usage:
  make [target]

Targets:
help                Display this help
install-tools       Install gofumpt, gocritic and swaggo
lint                Run golangci-lint
format              Format code
test                Run all tests
test/unit           Run unit tests
test/coverage       Run tests, make coverage report and display it into browser
test/coverage-browser  Open coverage report in browser
swagger             Generate swagger docs
run                 Run http server
clean               Remove cache files
```

## 🎲 Como rodar o projeto

```bash
# Clone este repositório
$ git clone <https://github.com/uesleicarvalhoo/marsover>

# Acesse a pasta do projeto no terminal
$ cd marsover

# Você pode facilmente iniciar as dependencias de desenvolvimento com o comando
$ docker compose build && docker compose up -d

# Isso vai iniciar o container:
# backend       localhost:5000  -> backend da aplicação
```
---

## 🧪 Testes

### Unitários e de integração

```bash
make test/unit
```

### Teste unitários e de integração full (upload → resposta)

```bash
make test
```

---

## 📡 API

### POST `/missions`

Recebe um arquivo `.txt` com as instruções da missão.

#### Exemplo de arquivo

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

#### Resposta

```
1 3 N
5 1 E
```

---

## ⚙️ Arquitetura

- **Plateau:** define os limites e valida coordenadas.
- **Rover:** movimenta-se com comandos (`L`, `R`, `M`).
- **Orchestrator:** executa rovers em paralelo (controla workers).
- **HTTP Layer:** parser do arquivo, handlers e resposta plain-text.

---

## 🤝 Contribuindo

1. Fork este repositório
2. Crie uma branch: `git checkout -b feature/sua-feature`
3. Commit: `git commit -m 'Minha feature'`
4. Push: `git push origin feature/sua-feature`
5. Abra um Pull Request

---

## 📜 Requisitos atendidos

✅ Sem libs externas no core
✅ Libs externas apenas para testes
✅ Concurrent execution
✅ API HTTP funcional com upload
✅ Testes unitários + integração (100% flow)

---

## 📝 Licença

Este projeto é open-source e está sob a licença MIT.
