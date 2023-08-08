# :satellite: Telefil

[![Go Reference](https://pkg.go.dev/badge/github.com/masih/telefil.svg)](https://pkg.go.dev/github.com/masih/telefil)
[![Go Test](https://github.com/masih/telefil/actions/workflows/go-test.yml/badge.svg)](https://github.com/masih/telefil/actions/workflows/go-test.yml)

> Filecoin API client in pure Go

Telefil offers a seamless way to engage with the Filecoin chain API using the JSON RPC 2.0 protocol, all while being written purely in Go. Notably, there are zero dependencies on [Lotus](https://github.com/filecoin-project/lotus), eliminating the need for C-bindings.

## Features

Filecoin API interactions provided by Telefil are:

- **`Filecoin.StateMinerInfo`**
- **`Filecoin.ChainHead`**
- **`Filecoin.ChainGetGenesis`**
- **`Filecoin.StateDealProviderCollateralBounds`**
- **`Filecoin.StateListMiners`**
- **`Filecoin.StateMinerPower`**

## Installation

```bash
go get github.com/masih/telefil@latest
```

## Status

:construction: **This repository is under active development.** :construction:

Please be aware that while we strive to keep the master branch stable, breaking changes may be introduced as we push
forward. We recommend using released versions for production systems and always checking the release notes before
updating.

## Documentation

For detailed usage and integration guidance, please refer
to [godoc documentation](https://pkg.go.dev/github.com/masih/telefil).

## Contribution

Contributions are welcome! Please see the [CONTRIBUTING.md](CONTRIBUTING.md) for more details.

## License

This project is licensed under the MIT and Apache 2.0 Licenses - see the [LICENSE.md](LICENSE.md) file for details.
