package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/golang/geo/s2"
	"github.com/tkrajina/gpxgo/gpx"
)

var input_file = flag.String("input", "", "Input GPX file")
var output_file = flag.String("output", "", "Output GPX file")
var epsilon = flag.Float64("epsilon", 0.0000045, "Epsilon value (about 50 m by default)")

func pointToSegmentDistance(point, start, end s2.Point) float64 {
	projected := s2.Project(point, start, end)
	return point.Distance(projected).Radians()
}

func simplifyPoints(points []s2.Point, epsilon float64) []s2.Point {
	if len(points) < 3 {
		return points
	}

	var (
		firstPoint   = points[0]
		lastPoint    = points[len(points)-1]
		indexWithMax = -1
		maxDist      = epsilon
	)

	for i, point := range points[1 : len(points)-1] {
		dist := pointToSegmentDistance(point, firstPoint, lastPoint)

		if dist > maxDist {
			indexWithMax = i + 1
			maxDist = dist
		}
	}

	if indexWithMax == -1 {
		return []s2.Point{firstPoint, lastPoint}
	}

	left := simplifyPoints(points[:indexWithMax+1], epsilon)
	right := simplifyPoints(points[indexWithMax:], epsilon)

	return append(left[:len(left)-1], right...)
}

func simplifyGPXFile(inputFile, outputFile string, epsilon float64) error {
	gpxFile, err := gpx.ParseFile(inputFile)
	if err != nil {
		return fmt.Errorf("could not parse input GPX file: %w", err)
	}

	routeGpx := false
	var originalPoints []gpx.GPXPoint
	for _, route := range gpxFile.Routes {
		originalPoints = append(originalPoints, route.Points...)
		routeGpx = true
	}
	for _, route := range gpxFile.Tracks {
		for _, segment := range route.Segments {
			originalPoints = append(originalPoints, segment.Points...)
		}
	}

	if len(originalPoints) == 0 {
		return fmt.Errorf("no points found in GPX file")
	}

	// Convert GPX points to s2.Points
	points := make([]s2.Point, len(originalPoints))
	for i, gpxPoint := range originalPoints {
		points[i] = s2.PointFromLatLng(s2.LatLngFromDegrees(gpxPoint.Latitude, gpxPoint.Longitude))
	}

	// Simplify the s2.Points
	simplifiedPoints := simplifyPoints(points, epsilon)

	// Convert simplified s2.Points back to GPX points
	simplifiedGPXPoints := make([]gpx.GPXPoint, len(simplifiedPoints))
	for k, point := range simplifiedPoints {
		latLng := s2.LatLngFromPoint(point)
		simplifiedGPXPoints[k] = gpx.GPXPoint{
			Point: gpx.Point{
				Latitude:  latLng.Lat.Degrees(),
				Longitude: latLng.Lng.Degrees()},
		}
	}

	// Replace the original points with the simplified points
	if routeGpx {
		gpxFile.Routes[0].Points = simplifiedGPXPoints
	} else {
		gpxFile.Tracks[0].Segments[0].Points = simplifiedGPXPoints
	}

	gpxbytes, _ := gpxFile.ToXml(gpx.ToXmlParams{Version: "1.1", Indent: true})

	wrfile, errsf := os.Create(outputFile)
	if errsf != nil {
		return fmt.Errorf("could not create file: %s %w", outputFile, errsf)
	}
	defer wrfile.Close()

	_, errwd := io.Copy(wrfile, bytes.NewReader(gpxbytes))
	if errwd != nil {
		return fmt.Errorf("unable to write data to file %s: %w", outputFile, errwd)
	}
	return nil
}

func main() {

	flag.Parse()

	if *input_file == "" || *output_file == "" {
		flag.Usage()
		os.Exit(1)
	}

	err := simplifyGPXFile(*input_file, *output_file, *epsilon)
	if err != nil {
		fmt.Println("Error simplifying GPX file:", err)
		os.Exit(1)
	}

	fmt.Printf("Simplified GPX file saved to %s\n", *output_file)
	os.Exit(0)
}
