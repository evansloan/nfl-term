.PHONY: install release

VERSION=v0.0.3
ARTIFACTS_DIR=build/artifacts/$(VERSION)
GITHUB_USERNAME=evansloan
FLAGS="-X main.version=$(VERSION)"

install:
	go build -ldflags=$(FLAGS) -o ${GOPATH}/bin/nfl-term

release:
	GOOS=windows GOARCH=amd64 go build -ldflags=$(FLAGS) -o $(ARTIFACTS_DIR)/nfl-term_windows_amd64.exe
	GOOS=darwin GOARCH=amd64 go build -ldflags=$(FLAGS) -o $(ARTIFACTS_DIR)/nfl-term_darwin_amd64
	GOOS=linux GOARCH=amd64 go build -ldflags=$(FLAGS) -o $(ARTIFACTS_DIR)/nfl-term_linux_amd64
	ghr -u $(GITHUB_USERNAME) --replace $(VERSION) $(ARTIFACTS_DIR)