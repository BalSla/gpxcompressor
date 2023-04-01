package tests

import (
	"testing"

	"github.com/tkrajina/gpxgo/gpx"
)

func Test_Returns_All_Points_In_Gpx_From_Route(t *testing.T) {
	gpxFile, _ := gpx.ParseFile("test.gpx")
	var points []gpx.GPXPoint
	for _, route := range gpxFile.Routes {
		points = append(points, route.Points...)
	}
	if len(points) != 2 {
		t.Errorf("Expected 2 points, got %d", len(points))
	}
}
func Test_Returns_All_Points_In_Gpx_From_Tracks(t *testing.T) {
	gpxFile, _ := gpx.ParseFile("test1.gpx")
	var points []gpx.GPXPoint
	for _, route := range gpxFile.Tracks {
		for _, segment := range route.Segments {
			points = append(points, segment.Points...)
		}
	}
	if len(points) != 2 {
		t.Errorf("Expected 2 points, got %d", len(points))
	}
}
