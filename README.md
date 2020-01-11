# nfl-term


![](https://user-images.githubusercontent.com/8287750/71860176-d645b200-30bf-11ea-9ea5-3693ce7dd328.gif)

## Installing

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

```
$ curl -o /usr/local/bin/nfl-term https://github.com/evansloan/nfl-term/releases/download/v0.0.2/nfl-term_darwin_amd64
$ chmod +x /usr/local/bin/nfl-term
```

### Build from source

Requires `go >= 1.12`

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

2. Find the game within the list you would like to display stats for. Using the link above, we'll get the stats for Colts vs. Bills. It was the 4th game played on September 13, 2015 so its game ID would be `20150901303`. Game numbers reset to 00 for each different day. So the Thursday night game of that same week (NE vs. PIT) would have a game number of 00.



## TODO

* Add responsive layout

* Better error handling

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details


