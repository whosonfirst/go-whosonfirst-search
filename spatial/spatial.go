package spatial

import (
	"github.com/skelterjohn/geom"
)

type PointInPolygonCandidate struct {
	Id        string
	FeatureId string
	IsAlt     bool
	AltLabel  string
	Bounds    *geom.Rect
}
