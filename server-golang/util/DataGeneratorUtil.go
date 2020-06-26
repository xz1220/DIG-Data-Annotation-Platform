package util

import (
	"labelproject-back/model"
	"math"
)

func GetBox(data model.DataForResponse) ([]float64, error) {
	var bbox []float64
	maxWidth := math.MaxFloat64
	minWidth := -math.MaxFloat64
	maxHeight := math.MaxFloat64
	minHeight := -math.MaxFloat64

	for _, point := range data.Point {
		maxWidth = math.Max(maxWidth, point.X)
		minWidth = math.Min(minWidth, point.X)

		maxHeight = math.Max(maxHeight, point.Y)
		minHeight = math.Min(minHeight, point.Y)
	}

	bbox = append(bbox, maxHeight, minHeight, maxWidth, minWidth)
	return bbox, nil
}

func CalculateArea(points []model.Points) float64 {
	var area float64
	for i := 0; i < len(points); i++ {
		area += points[i].X*points[i+1].Y - points[i].Y*points[i+1].X
	}
	area = 0.5 * math.Abs(area+points[len(points)-1].X*points[0].Y-points[len(points)-1].Y*points[0].X)
	return area
}

func GenPolygonData(points []model.Points) []float64 {
	var seg []float64
	for _, point := range points {
		seg = append(seg, point.X, point.Y)
	}
	return seg
}
