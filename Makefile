default: build

build:
	GOOS=0 go build -o herd.darwin main.go

release:

	if git rev-parse $(ver) >/dev/null 2>&1
	then
	#GOOS=0 go build -ldflags "-X main.version ${rel} -X main.buildDate `date -u +.%Y%m%d%.H%M%S`"
		echo "Building Release"
		goreleaser
	else
		echo "Tag not found, creating then building release"
		git tag -a $(ver) -m "Building Release for Version $(ver)"
		git push origin $(ver)
		#GOOS=0 go build -ldflags "-X main.version ${rel} -X main.buildDate `date -u +.%Y%m%d%.H%M%S`"
		gorleaser
	fi
