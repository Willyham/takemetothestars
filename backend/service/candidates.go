package service

import (
	"context"

	"github.com/Willyham/takemetothestars/model"
	"github.com/go-spatial/geom"
)

// CandidateService provides the ability to generate candidates for travel to see the stars.
type CandidateService interface {
	FindCandidates(context.Context, geom.Point, int) ([]model.Candidate, error)
}

// Candidates implements the CandidateService interface.
type Candidates struct {
}

// FindCandidates generates a list of candidates for users to travel.
// It roughly works by:
// - Finding when it will be dark (i.e. astronomical twilight) at the origin point.
// - Finding a rough set of points where there is little light-pollution
// - For each of those points, check the weather in the location and only select clear skies
// - For remaining points, compute the travel distance with some kind of maps api
// - Produce a set of candidate values which describe options for the user.
func (c Candidates) FindCandidates(ctx context.Context, origin geom.Point, limit int) ([]model.Candidate, error) {
	return []model.Candidate{}, nil
}
