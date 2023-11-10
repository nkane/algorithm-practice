package collinear

type Point struct {
	X int
	Y int
}

type LineSegment struct {
	P     Point
	Q     Point
	Slope float64
}

type CollinearReturnValue struct {
	NumberOfSegments int
	Segments         []LineSegment
}

func BruteCollinearPoints(points []Point) CollinearReturnValue {
	return CollinearReturnValue{}
}

func FastCollinearPoints(points []Point) CollinearReturnValue {
	return CollinearReturnValue{}
}
