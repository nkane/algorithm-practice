package collinear

import "testing"

func TestBruteCollinearPoints(t *testing.T) {
	points := []Point{
		{
			X: 10000, Y: 0,
		},
		{
			X: 0, Y: 10000,
		},
		{
			X: 3000,
			Y: 7000,
		},
		{
			X: 7000,
			Y: 3000,
		},
		{
			X: 20000,
			Y: 21000,
		},
		{
			X: 3000,
			Y: 4000,
		},
		{
			X: 14000,
			Y: 15000,
		},
		{
			X: 6000,
			Y: 7000,
		},
	}
	BruteCollinearPoints(points)
}

func TestFastCollinearPoints(t *testing.T) {
}
