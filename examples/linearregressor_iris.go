package main

import (
    //mat "github.com/skelterjohn/go.matrix"
    base "golearn/base"
    util "golearn/utilities"
    linearclass "golearn/linear"
    //"fmt"
    )

func main(){
  //Parses the infamous Iris data.
  cols, rows, _, labels, data := base.ParseCsv("datasets/iris.csv", 2, []int{0,1})
  newlabels := util.ConvertLabelsToFloat(labels)
  linear := linearclass.LinearRegressor{}
  linear.New("Testing", newlabels, data, rows, cols)
  linear.Fit()
}
