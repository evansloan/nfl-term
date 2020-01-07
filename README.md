# nfl-term


![](https://user-images.githubusercontent.com/8287750/71776885-be0e4f80-2f66-11ea-9760-9da1e87db370.png)

## Installing

### Binary

Download the latest release for your system from the [releases page](https://github.com/evansloan/nfl-term/releases).

_or_

Use cURL/wget (Make sure to download the correct release for your system)

- Mac: `nfl-term_darwin_amd64`
- Linux: `nfl-term_linux_amd64`

```
$ curl -o /usr/local/bin/nfl-term https://github.com/evansloan/nfl-term/releases/download/v0.0.1/nfl-term_darwin_amd64

$ chmod +x /usr/local/bin/nfl-term
```

### Build from source

Requires `go >= 1.12.7`

[How to install Go](https://golang.org/doc/install)

```
$ git clone https://github.com/evansloan/nfl-term.git
$ cd nfl-term
$ make install
```

## How to use

```
$ nfl-term
```

## TODO

* Add responsive layout

* Display stats from games of user's choice

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details


