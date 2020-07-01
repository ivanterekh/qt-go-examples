prepare:
	go install -v -tags=no_env github.com/therecipe/qt/cmd/...
	rm -rf vendor/github.com/therecipe/env_linux_amd64_513
	git clone https://github.com/therecipe/env_linux_amd64_513.git vendor/github.com/therecipe/env_linux_amd64_513
	`go env GOPATH`/bin/qtsetup

build.bezier:
	./scripts/build.sh bezier

run.bezier: build.bezier
	./bin/bezier

build.cabinet:
	./scripts/build.sh cabinet

run.cabinet: build.cabinet
	./bin/cabinet

build.convexhull:
	./scripts/build.sh convexhull

run.convexhull: build.convexhull
	./bin/convexhull

build.points:
	./scripts/build.sh points

run.points: build.points
	./bin/points
