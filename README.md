# Qt Go Examples

This repository contains geometrical algorithms and their visualisation made with [Qt binding for Go](https://github.com/therecipe/qt). All the code was written as a part of the course "Computational geometry and geometric modeling" during my master studies in Belarusian State University.

## Requirements

Currently only linux is supported. You should have git and Go installed on your computer. It is also necessary to install g++ >= 5 and some OpenGL dependencies. Visit https://github.com/therecipe/qt/wiki/Installation-on-Linux to get the way to do this on your linux distribution.

## Getting started

To install Qt and the binding you may run
```
make prepare
```

To take a look at the example run one of the following
```
make run.bezier
make run.cabinet
make run.convexhull
make run.points
```

Binaries can be built with the following commands
```
make build.bezier
make build.cabinet
make build.convexhull
make build.points
```

## Examples description

### bezier

This application paints [Bézier curve](https://en.wikipedia.org/wiki/Bézier_curve) by given points. [De Casteljau's algorithm](https://en.wikipedia.org/wiki/De_Casteljau%27s_algorithm) is used.

Points are defined with a mouse. Points are draggable.

### cabinet

This application creates [Cabinet projection](https://en.wikipedia.org/wiki/Oblique_projection#Cabinet_projection) of a cube.

The cube can be rotated by any of the axis.

### convexhull

This application builds a [convex hull](https://en.wikipedia.org/wiki/Convex_hull) for a set of points.

One of 3 methods can be used:
- [Graham scan](https://en.wikipedia.org/wiki/Graham_scan)
- [Gift wrapping algorithm](https://en.wikipedia.org/wiki/Gift_wrapping_algorithm) also known as Jarvis march
- [Quickhull](https://en.wikipedia.org/wiki/Quickhull)

The application works normally with up to 100 000 points.

### points

This application performs search of points inside of a given rectangle with O(log n) complexity.

On of the following data structures can be used to organise points:
- [2-d tree](https://en.wikipedia.org/wiki/K-d_tree)
- Point net. In this data structure plane is splited with a uniform grid. It allows to get points of cells that are fully located inside a rectangle instanly. 
