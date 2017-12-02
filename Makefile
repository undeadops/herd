default: build

build:
	GOOS=0 go build -o herd.darwin main.go

release:
	rel=${RELEASE}
	if git rev-parse ${rel} >/dev/null 2>&1
	then
	#GOOS=0 go build -ldflags "-X main.version ${rel} -X main.buildDate `date -u +.%Y%m%d%.H%M%S`"
		echo "Building Release"
		goreleaser
	else
		echo "Tag not found, creating then building release"
		git tag -a ${rel} -m "Building Release for Version ${rel}"
		git push origin ${rel}
		#GOOS=0 go build -ldflags "-X main.version ${rel} -X main.buildDate `date -u +.%Y%m%d%.H%M%S`"
		gorleaser
	fi
