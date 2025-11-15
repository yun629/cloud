# Cloud SDK (v1)

An early-stage, vendor-agnostic Go SDK for managing **clusterable, GPU-accelerated compute** across cloud providers.

---

## Project Goals

- Define a clean, minimal interface for cloud compute primitives:
  - `Instance`
  - `Storage`
  - `FirewallRule`
  - `InstanceType`
  - `Location`

- Enable **clusterable GPU workloads** across multiple providers, with shared semantics and L3 network guarantees. (WIP)

---

## Security

All cloud integrations must follow our [Security Requirements](SECURITY.md), which define:

- **Network Security**: Default "deny all inbound, allow all outbound" model
- **Cluster Security**: Internal instance communication with external isolation
- **Data Protection**: Encryption requirements for data at rest and in transit
- **Implementation Guidelines**: Security checklists for cloud provider integrations

See [SECURITY.md](docs/SECURITY.md) for complete security specifications and implementation requirements.

---

## Status

- Version: `v1` — internal interface, open-sourced
- Current scope: core types + interfaces + tests
- Cloud provider implementations are internal-only for now
- `v2` will be shaped by feedback and contributions from the community

## Platform Support

- **Operating System**: Currently supports Ubuntu 22 only
- **Architecture**: Designed for GPU-accelerated compute workloads
- **Access Method**: Requires SSH server and SSH key-based authentication
- **System Requirements**: Requires systemd to be running and accessible

---

## Who This Is For

- **NVIDIA Cloud Partners (NCPs)** looking to offer Brev-compatible GPU compute
- **Infra teams** building cluster-aware systems or abstractions on raw compute
- **Cloud providers** interested in contributing to a shared interface for accelerated compute
- **Compute brokers & marketplaces (aggregators)** offering multi-cloud compute

## Documentation

- **[V1 Design Notes](pkg/v1/V1_DESIGN_NOTES.md)**: Design decisions, known quirks, and AWS-inspired patterns in the v1 API
- **[Architecture Overview](docs/ARCHITECTURE.md)**: How the Cloud SDK fits into Brev's overall architecture
- **[Security Requirements](docs/SECURITY.md)**: Security specifications and implementation requirements
- **[How to Add a Provider](docs/how-to-add-a-provider.md)**: Step-by-step guide to implement a new cloud provider using the Lambda Labs example

---

## Getting Started

### 1. Install prerequisites

- **Go 1.24+** – download from [go.dev/dl](https://go.dev/dl/) or use your OS package manager. The error `make: go: No such file or directory` means Go is not installed or not on your `PATH`.
  - **Linux/macOS**: follow the [official install guide](https://go.dev/doc/install) and make sure `$HOME/go/bin` (or the custom install path) is on `PATH`.
  - **Windows**: install Go via the MSI installer *or* use WSL2 and install Go inside the Linux distribution. Commands such as `make deps` must run inside the environment where Go is installed.
- **make & git** – already shipped on most Linux/macOS distros. On Windows, install them through WSL, MSYS2, or Git for Windows.

### 2. Bootstrap the repository

```bash
make deps            # downloads Go modules (go mod download/tidy)
make install-tools   # installs golangci-lint, gofumpt, gosec, etc.
```

If you prefer manual control, you can replace those targets with `go mod download` and the individual `go install` commands listed in the `Makefile`.

### 3. Provide credentials through environment variables

All provider validation tests look for API keys in environment variables. The recommended workflow is to create a `.env` file in the repo root (see [`docs/example-dot-env`](docs/example-dot-env)):

```bash
cat <<'EOF' > .env
LAMBDALABS_API_KEY=your-lambda-key
VASTAI_API_KEY=your-vast-key
# add more provider variables as needed
EOF
```

The Makefile automatically loads `.env` so every target can read those values. Alternatively, `export` the variables directly in your shell before running commands.

### 4. Run tests

```bash
make test            # unit tests (skips validation)
make test-validation # runs validation suite – requires real provider API keys
make test-all        # convenience target that runs everything
```

For debugging a single provider you can run `go test -v -run TestValidationFunctions ./internal/{provider}/v1/` after exporting the matching API key. Validation tests hit real cloud APIs, so expect them to take longer and incur usage on your account.

---

## Get Involved

This is a foundation — we're opening it early to **learn with the community** and shape a clean, composable `v2`. If you're building GPU compute infrastructure or tooling, we'd love your input.

