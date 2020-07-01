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
