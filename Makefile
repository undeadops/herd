default: build

build:
	GOOS=0 go build -o herd.darwin main.go
	#GOOS=0 go build -ldflags "-X main.version ${rel} -X main.buildDate `date -u +.%Y%m%d%.H%M%S`" \

release:
	@if git rev-parse -q --verify "refs/tags/$(ver)" >/dev/null; then \
		(echo "Building Release"); \
		/usr/local/bin/goreleaser; \
	else \
		(echo "Tag not found, creating then building release"); \
		git tag -a $(ver) -m "Building Release for Version $(ver)"; \
		git push origin $(ver); \
		/usr/local/bin/goreleaser; \
	fi
