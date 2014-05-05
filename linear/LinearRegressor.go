package linear

import (
		mat "github.com/skelterjohn/go.matrix"
		//"math"
		//"fmt"
		//util "golearn/utilities"
		base "golearn/base"
		)


//A linear Regressor. Consists of a data matrix, associated result variables in the same order as the matrix, and a name.
type LinearRegressor struct {
	base.BaseRegressor
}

//The parameters for the equation y = slope * x + c
type LinearParameters struct {
	slope float64
	c float64
}

//Mints a new classifier.
func (LinearModel *LinearRegressor) New(name string, labels []float64, numbers []float64, x int, y int) {

	LinearModel.Data = *mat.MakeDenseMatrix(numbers, x, y)
	LinearModel.Name = name
	LinearModel.Labels = labels
}

func (LinearModel *LinearRegressor) Fit() LinearParameters{
	rows := LinearModel.Data.Rows()
	sumX, sumY, sumXSquared, sumXY := 0.0, 0.0, 0.0, 0.0


	for i := 0; i < rows; i++{
		row := LinearModel.Data.GetRowVector(i)
		sumX += row.Array()[0]
		sumY += row.Array()[1]
		sumXSquared += row.Array()[0] * row.Array()[0]
		sumXY += row.Array()[0] * row.Array()[1]
	}

	coefficients := LinearParameters{}
	coefficients.slope = ((float64(rows) * sumXY) - (sumX)*(sumY)) / ((float64(rows) * sumXSquared)-(sumX*sumX))
	coefficients.c = (sumY - (coefficients.slope * sumX)) / float64(rows)
	return coefficients
}
