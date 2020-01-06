.PHONY: install release

VERSION=v0.0.1
ARTIFACTS_DIR=build/artifacts/$(VERSION)
GITHUB_USERNAME=evansloan

install:
	go build -ldflags="-X main.version=${version}" -o /usr/local/bin

release:
	GOOS=windows GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/nfl-term_windows_amd64
	GOOS=darwin GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/nfl-term_darwin_amd64
	GOOS=linux GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/nfl-term_linux_amd64
	ghr -u $(GITHUB_USERNAME) --replace $(VERSION) $(ARTIFACTS_DIR)