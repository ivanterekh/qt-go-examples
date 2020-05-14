package bezier

import "github.com/ivanterekh/qt-go-examples/internal/geometry"

func Build(points []*geometry.Point, resolution int) []*geometry.Point {
	n := len(points)
	if n < 3 {
		return nil
	}

	xs := make([]float64, n)
	ys := make([]float64, n)
	for i, p := range points {
		xs[i] = float64(p.X)
		ys[i] = float64(p.Y)
	}

	res := make([]*geometry.Point, resolution+1)
	for i := 0; i <= resolution; i++ {
		t := float64(i) / float64(resolution)
		res[i] = &geometry.Point{
			X: int(getCoord(xs, t)),
			Y: int(getCoord(ys, t)),
		}
	}

	return res
}

func getCoord(xs []float64, t float64) float64 {
	n := len(xs)

	b := make([][]float64, n)
	for i := 0; i < n; i++ {
		b[i] = make([]float64, n)
	}

	copy(b[0], xs)

	for r := 1; r < n; r++ {
		for i := 0; i < n-r; i++ {
			b[r][i] = (1-t)*b[r-1][i] + t*b[r-1][i+1]
		}
	}

	return b[n-1][0]
}
