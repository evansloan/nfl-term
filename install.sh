#!/bin/bash

version="v0.0.1"
commit=`git rev-parse --short HEAD`

go build -ldflags="-X main.version=${version}  -X main.commit=${commit}"
sudo cp nfl-term $HOME/bin
rm nfl-term
