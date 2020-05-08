package projection

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"math"

	"gonum.org/v1/gonum/mat"
)

func TransformationMatrix(xAngle, yAngle, zAngle float64) *mat.Dense {
	xMatr := mat.NewDense(3, 3, []float64{
		1, 0, 0,
		0, math.Cos(xAngle), -math.Sin(xAngle),
		0, math.Sin(xAngle), math.Cos(xAngle),
	})

	yMatr := mat.NewDense(3, 3, []float64{
		math.Cos(xAngle), 0, math.Sin(xAngle),
		0, 1, 0,
		-math.Sin(xAngle), 0, math.Cos(xAngle),
	})

	zMatr := mat.NewDense(3, 3, []float64{
		math.Cos(xAngle), -math.Sin(xAngle), 0,
		math.Sin(xAngle), math.Cos(xAngle), 0,
		0, 0, 1,
	})

	res := mat.NewDense(3, 3, nil)
	res.Mul(xMatr, yMatr)
	res.Mul(res, zMatr)

	return res
}

func makeTurn(vector *mat.VecDense, matr *mat.Dense) *mat.VecDense {
	vector.MulVec(matr, vector)
	return vector
}

func makePointProjection(point *gui.QVector3D, angle float64) *core.QPointF {
	return core.NewQPointF3(
		float64(point.X()) + float64(point.Z())*math.Cos(angle)/2,
		float64(point.Y()) + float64(point.Z())*math.Sin(angle)/2,
	)
}

func GetCubeProjection(center *gui.QVector3D, size float64, angle float64, matr *mat.Dense) []*core.QLineF {
	var (
		vert000 = makePointProjection(sumQVec3D(center, toQVec3d(makeTurn(mat.NewVecDense(3, []float64{-size, -size, -size}), matr))), angle)
		vert001 = makePointProjection(sumQVec3D(center, toQVec3d(makeTurn(mat.NewVecDense(3, []float64{-size, -size, size}), matr))), angle)
		vert010 = makePointProjection(sumQVec3D(center, toQVec3d(makeTurn(mat.NewVecDense(3, []float64{-size, size, -size}), matr))), angle)
		vert011 = makePointProjection(sumQVec3D(center, toQVec3d(makeTurn(mat.NewVecDense(3, []float64{-size, size, size}), matr))), angle)
		vert100 = makePointProjection(sumQVec3D(center, toQVec3d(makeTurn(mat.NewVecDense(3, []float64{size, -size, -size}), matr))), angle)
		vert101 = makePointProjection(sumQVec3D(center, toQVec3d(makeTurn(mat.NewVecDense(3, []float64{size, -size, size}), matr))), angle)
		vert110 = makePointProjection(sumQVec3D(center, toQVec3d(makeTurn(mat.NewVecDense(3, []float64{size, size, -size}), matr))), angle)
		vert111 = makePointProjection(sumQVec3D(center, toQVec3d(makeTurn(mat.NewVecDense(3, []float64{size, size, size}), matr))), angle)
	)

	return []*core.QLineF {
		core.NewQLineF2(vert000, vert001),
		core.NewQLineF2(vert000, vert010),
		core.NewQLineF2(vert000, vert100),
		core.NewQLineF2(vert001, vert011),
		core.NewQLineF2(vert001, vert101),
		core.NewQLineF2(vert010, vert110),
		core.NewQLineF2(vert010, vert011),
		core.NewQLineF2(vert100, vert101),
		core.NewQLineF2(vert100, vert110),
		core.NewQLineF2(vert110, vert111),
		core.NewQLineF2(vert101, vert111),
		core.NewQLineF2(vert011, vert111),
	}
}

func sumQVec3D(a, b *gui.QVector3D) *gui.QVector3D {
	return gui.NewQVector3D3(
		a.X()+b.X(),
		a.Y()+b.Y(),
		a.Z()+b.Z(),
	)
}

func toQVec3d(vec *mat.VecDense) *gui.QVector3D {
	return gui.NewQVector3D3(
		float32(vec.AtVec(0)),
		float32(vec.AtVec(1)),
		float32(vec.AtVec(2)),
	)
}
