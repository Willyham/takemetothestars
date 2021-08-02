package model

import "github.com/go-spatial/geom"

// Candidate describes a potential place to send the user to see the stars.
// TODO:
//  - Include data on 'how dark' it is - still unclear what format this should be
//  - Include
type Candidate struct {
	Point geom.Point

	DistanceMeters    int
	TravelTimeMinutes int
}
