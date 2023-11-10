package collinear

type Point struct {
	X float64
	Y float64
}

type LineSegment struct {
	A     Point
	B     Point
	Slope float64
}

func NewLineSegment(a Point, b Point) LineSegment {
	result := LineSegment{
		A: a,
		B: b,
	}
	result.Slope = (b.Y - a.Y) / (b.X - a.X)
	return result
}

type CollinearReturnValue struct {
	NumberOfSegments int
	Segments         []LineSegment
}

func BruteCollinearPoints(points []Point) CollinearReturnValue {
	result := CollinearReturnValue{}
	for i := 0; i < len(points); i += 4 {
		p := points[i+0]
		q := points[i+1]
		segmentA := NewLineSegment(p, q)
		result.Segments = append(result.Segments, segmentA)
		r := points[i+2]
		s := points[i+3]
		segmentB := NewLineSegment(r, s)
		result.Segments = append(result.Segments, segmentB)
	}
	return result
}

func FastCollinearPoints(points []Point) CollinearReturnValue {
	return CollinearReturnValue{}
}
