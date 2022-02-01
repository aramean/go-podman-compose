[![Continuous integration and delivery](https://github.com/aramean/go-podman-compose/actions/workflows/releasement.yml/badge.svg)](https://github.com/aramean/go-podman-compose/actions/workflows/releasement.yml)
[![Continuous deployment](https://github.com/aramean/go-podman-compose/actions/workflows/deployment.yml/badge.svg)](https://github.com/aramean/go-podman-compose/actions/workflows/deployment.yml)<br><br>
<img src="/docs/logo.svg">

# Podman Compose
Podman Compose lets you define and run multi-container Podman applications. You configure the services using YAML files, and then start all the services with the same command.<br>

## Installing

### With Curl
#### MacOS (ARM64)
`curl -fsSL github.com/aramean/go-podman-compose/releases/download/v1.0.0/podman-compose-darwin-arm64.zip -O`

#### MacOS (AMD64)
`curl -fsSL github.com/aramean/go-podman-compose/releases/download/v1.0.0/podman-compose-darwin-amd64.zip -O`

#### Linux (AMD64)
`curl -fsSL github.com/aramean/go-podman-compose/releases/download/v1.0.0/podman-compose-linux-amd64.zip -O`


### From source
To compile from source, simply run `make && make install`
> **_NOTE:_**  Make sure you have the latest version of <a href="https://go.dev/dl/">Go</a> installed on your machine.

## Contributing

If you find an issue, please report it on the <a href="../../issues/new">issue tracker</a>.

## Donating

If you'd like to donate to the project, click the button bellow.<br><br>
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/donate/?hosted_button_id=T7A39PQ2YGZFE)
