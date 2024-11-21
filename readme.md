# Reverse Proxy Server

This project is a reverse proxy server designed for home networks with a single IP address. It allows routing to multiple services hosted within the internal network. The server uses Let's Encrypt to generate SSL certificates for each host address and redirects HTTP traffic to HTTPS.

## Features

- **Reverse Proxy**: Routes incoming requests to the appropriate backend service based on the host address.
- **Let's Encrypt Integration**: Automatically generates and renews SSL certificates for each host address.
- **HTTP to HTTPS Redirection**: Redirects all HTTP traffic to HTTPS to ensure secure communication.

## Configuration

The server configuration is defined in the `config.yaml` file. The configuration file should specify the services to be proxied, including their host addresses and backend URLs.

Example `config.yaml`:

```yaml
services:
  - host: example1.com
    backend: http://localhost:8081
  - host: example2.com
    backend: http://localhost:8082
```
## Usage

1. **Install Dependencies**: Ensure you have Go installed and run `go mod tidy` to install the required dependencies.
2. **Run the Server**: Start the reverse proxy server by running the following command:

    ```sh
    go run main.go serve
    or 
    go build *.go
    ./main serve
    ```

3. **Access Your Services**: Access your services using the configured host addresses. The server will handle the routing and SSL termination.

## Code Overview

- **Configuration Loading**: The `LoadConfig` function in [`cmd/proxy.go`](cmd/proxy.go) reads the configuration from `config.yaml` and unmarshals it into a `Config` struct.
- **Reverse Proxy Creation**: The `NewReverseProxy` function in [`cmd/proxy.go`](cmd/proxy.go) creates a new reverse proxy for a given target URL.
- **Server Initialization**: The `startProxyServer` function in [`cmd/proxy.go`](cmd/proxy.go) initializes and starts the HTTP and HTTPS servers, sets up the Let's Encrypt manager, and handles the routing logic.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Cobra](https://github.com/spf13/cobra) for command-line interface support.
- [Let's Encrypt](https://letsencrypt.org/) for providing free SSL certificates.
- [Go](https://golang.org/) for the programming language and standard library.

```

