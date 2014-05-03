package linear

import (
		mat "github.com/skelterjohn/go.matrix"
		//"math"
		"fmt"
		//util "golearn/utilities"
		base "golearn/base"
		)


//A linear Regressor. Consists of a data matrix, associated result variables in the same order as the matrix, and a name.
type LinearRegressor struct {
	base.BaseRegressor
}

//The parameters for the equation xCoef * x + yCoef * y + c = 0
type LinearParameters struct {
	xCoef float64
	yCoef float64
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
	rownumbers := make(map[int]float64)
	//labels := make([]float64, 1)
	//sum := 0.0
	fmt.Println("Inside Regressor fit method")
	fmt.Println(rows)
	fmt.Println(rownumbers)
	coefficients := LinearParameters{}
	return coefficients
}
