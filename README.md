# Peerkat

<p align="center">
    <img src="https://yalinpala.dev/projects/peerkat.png" alt="peerkat logo"  width="250" height="250">
</p>

Peerkat is a peer-to-peer file sharing and chat application example that allows users to easily share files and chat.
Also see the [Peerkat Relay](https://github.com/yplog/peerkat-relay) repository.


## Table of Contents

1. [Installation](#installation)
2. [Usage](#usage)
3. [Key Features](#key-features)
4. [Contribution](#contribution)
5. [License](#license)

## Installation

To use the Peerkat application, you can follow the steps below:


### Clone the repository:

```bash
    git clone git@github.com:yplog/peerkat.git
```

### Install the dependencies:

```bash
    go mod tidy
```

### Build the application:

```bash
    go build -o peerkat cmd/peerkat/main.go
```

## Usage

Using the Peerkat application is straightforward. You can follow these steps:

Command line arguments:

```bash
    -help         Show help message
    -relay        Relay server address
    -mode         file-transfer/chat
    -peer         Peer address
```

Usage:

> The -peer parameter is the address of the peer you want to connect to. Do not use if you are a first user.

```bash
    go run cmd/peerkat/main.go -relay <RELAY_ADDRESS> -mode <MODE>
```

```bash
    go run cmd/peerkat/main.go -relay <RELAY_ADDRESS> -mode <MODE> -peer <PEER_ADDRESS>
```

Or build the application and run:

```bash
    ./peerkat -relay <RELAY_ADDRESS> -mode <MODE> -peer <PEER_ADDRESS>
```

## Key Features

Peerkat offers users the following key features:

- **Fast and Secure File Sharing:** Share your files quickly and securely with other users.
- **Peer-to-Peer Communication:** Direct file transfer between users without the need for a server.
- **Chat:** Communicate with other users through the chat feature.

## Contribution

If you want to contribute to the development of Peerkat, please follow these steps:

 1. Fork this repository.
 2. Add a new feature or fix a bug.
 3. Create a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
