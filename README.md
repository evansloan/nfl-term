# nfl-term


![](https://user-images.githubusercontent.com/8287750/72228447-a6d9ee00-3574-11ea-9278-cd6fe8b6990a.gif)

## Installing

Requires `go >= 1.12`

[How to install Go](https://golang.org/doc/install)

### Binary

Download the latest release for your system from the [releases page](https://github.com/evansloan/nfl-term/releases).

_or_

Use Go modules

```
go get github.com/evansloan/nfl-term
```

_or_

Use cURL/wget (Make sure to download the correct release for your system)

- Mac: `nfl-term_darwin_amd64`
- Linux: `nfl-term_linux_amd64`
- Windows: `nfl-term_windows_amd64.exe`

```
$ curl -o /usr/local/bin/nfl-term https://github.com/evansloan/nfl-term/releases/download/v0.0.3/nfl-term_darwin_amd64
$ chmod +x /usr/local/bin/nfl-term
```

### Build from source

```
$ git clone https://github.com/evansloan/nfl-term.git
$ cd nfl-term
$ make install
```

## How to use

```
$ nfl-term
```

### Arguments

```
-v, --version // Displays the version of nfl-term
-g, --games <game ids> // Loads specific game stats
```

#### Constructing game IDs

Game IDs follow this structure:
```
<year><month><day><game number>
```

Game numbers start at 00 so the first game played on a given day would be 00, the second 01, the third 02, etc...

**Example**: The first game played on September 13th of the 2015 season would be: `2015091300`, the second `2015091301`, the third `2015091302`.

**Finding a specific game number:**

1. Go  to `http://www.nfl.com/schedules/<season>/REG1` where `<season>` is the year of the season of the game you would like. The link for the 2015 season would be http://www.nfl.com/schedules/2015/REG1

2. Find the game within the list you would like to display stats for. Click on the Game Center link. The game ID is located within the Game Center URL. The ID within https://www.nfl.com/gamecenter/2015092009/2015/REG2/49ers@steelers is `2015092009`


## TODO

* Add responsive layout

* Better error handling

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details


