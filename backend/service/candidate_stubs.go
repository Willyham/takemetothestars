package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/Willyham/takemetothestars/model"
	"github.com/go-spatial/geom"
	"github.com/umahmood/haversine"
)

// StubCandidates is a stub for the candidate service which returns random data.
type StubCandidates struct{}

// FindCandidates returns random data around the origin for testing.
func (c StubCandidates) FindCandidates(ctx context.Context, origin geom.Point, maxCandidates int) ([]model.Candidate, error) {
	variance := 0.002
	r := rand.New(rand.NewSource(time.Now().Unix()))

	candidates := []model.Candidate{}
	for i := 0; i < maxCandidates; i++ {

		// Create another point within +- variance of each coordinate.
		latVar := float64Range(r, 1-variance, 1+variance)
		lonVar := float64Range(r, 1-variance, 1+variance)
		out := origin.Multiply(geom.Point{latVar, lonVar})

		_, km := haversine.Distance(
			haversine.Coord{Lat: origin.X(), Lon: origin.Y()},
			haversine.Coord{Lat: out.X(), Lon: out.Y()},
		)

		candidates = append(candidates, model.Candidate{
			Point:             out,
			DistanceMeters:    int(km * 1000),
			TravelTimeMinutes: int(km) * 2, //roughly average 0.5km/min
		})
	}
	return candidates, nil
}

func float64Range(r *rand.Rand, min, max float64) float64 {
	if min == max {
		return min
	}
	return r.Float64()*(max-min) + min
}
