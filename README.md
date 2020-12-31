# go-whosonfirst-search

Go package for Who's On First search interfaces.

## Important

This is work in progress. Documentation to follow.

## Interfaces

### filter.Filter

```
type Filter interface {
	HasPlacetypes(flags.PlacetypeFlag) bool
	IsCurrent(flags.ExistentialFlag) bool
	IsDeprecated(flags.ExistentialFlag) bool
	IsCeased(flags.ExistentialFlag) bool
	IsSuperseded(flags.ExistentialFlag) bool
	IsSuperseding(flags.ExistentialFlag) bool
	IsAlternateGeometry(flags.AlternateGeometryFlag) bool
	HasAlternateGeometry(flags.AlternateGeometryFlag) bool
}
```

### fulltext.FullTextDatabase

```
type FullTextDatabase interface {
	IndexFeature(context.Context, geojson.Feature) error
	QueryString(context.Context, string, ...filter.Filter) (spr.StandardPlacesResults, error)
	Close(context.Context) error
}
```

### spatial.SpatialDatabase

```
type SpatialDatabase interface {
	IndexFeature(context.Context, geojson.Feature) error
	PointInPolygon(context.Context, *geom.Coord, ...filter.Filter) (spr.StandardPlacesResults, error)
	PointInPolygonCandidates(context.Context, *geom.Coord, ...filter.Filter) ([]*spatial.PointInPolygonCandidate, error)
	PointInPolygonWithChannels(context.Context, chan spr.StandardPlacesResult, chan error, chan bool, *geom.Coord, ...filter.Filter)
	PointInPolygonCandidatesWithChannels(context.Context, chan *spatial.PointInPolygonCandidate, chan error, chan bool, *geom.Coord, ...filter.Filter)
	Close(context.Context) error
}
```

_As of this writing the `spatial` interfaces are currently defined and implemented in the [go-whosonfirst-spatial](https://github.com/whosonfirst/go-whosonfirst-spatial), and related, packages. Some or all of that work may be migrated in to this package but nothing final has been decided yet._

## Implementations

* https://github.com/whosonfirst/go-whosonfirst-search-sqlite

## See also

* https://github.com/whosonfirst/go-whosonfirst-geojson-v2
* https://github.com/whosonfirst/go-whosonfirst-spr
* https://github.com/whosonfirst/go-whosonfirst-flags