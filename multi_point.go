package geoos

// A MultiPoint represents a set of points in the 2D Eucledian or Cartesian plane.
type MultiPoint []Point

// GeoJSONType returns the GeoJSON type for the object.
func (mp MultiPoint) GeoJSONType() string {
	return TypeMultiPoint
}

// Dimensions returns 0 because a MultiPoint is a 0d object.
func (mp MultiPoint) Dimensions() int {
	return 0
}

// Nums num of multiPoint.
func (mp MultiPoint) Nums() int {
	return len(mp)
}

// Bound returns a bound around the points. Uses rectangular coordinates.
func (mp MultiPoint) Bound() Bound {
	if len(mp) == 0 {
		return emptyBound
	}

	b := Bound{mp[0], mp[0]}
	for _, p := range mp {
		b = b.Extend(p)
	}

	return b
}

// EqualMultiPoint compares two MultiPoint objects. Returns true if lengths are the same
// and all points are Equal, and in the same order.
func (mp MultiPoint) EqualMultiPoint(multiPoint MultiPoint) bool {
	if len(mp) != len(multiPoint) {
		return false
	}
	for i, v := range mp.ToPointArray() {
		if !v.Equal(Point(multiPoint[i])) {
			return false
		}
	}
	return true
}

// Equal checks if the MultiPoint represents the same Geometry or vector.
func (mp MultiPoint) Equal(g Geometry) bool {
	if g.GeoJSONType() != mp.GeoJSONType() {
		return false
	}
	return mp.EqualMultiPoint(g.(MultiPoint))
}

// EqualsExact Returns true if the two Geometrys are exactly equal,
// up to a specified distance tolerance.
// Two Geometries are exactly equal within a distance tolerance
func (mp MultiPoint) EqualsExact(g Geometry, tolerance float64) bool {
	if mp.GeoJSONType() != g.GeoJSONType() {
		return false
	}
	for i, v := range mp {
		if v.EqualsExact((g.(MultiPoint)[i]), tolerance) {
			return false
		}
	}
	return true
}

// Area returns the area of a polygonal geometry. The area of a multipoint is 0.
func (mp MultiPoint) Area() (float64, error) {
	return 0.0, nil
}

// ToPointArray returns the PointArray
func (mp MultiPoint) ToPointArray() (pa []Point) {
	return []Point(mp)
}

// IsEmpty returns true if the Geometry is empty.
func (mp MultiPoint) IsEmpty() bool {
	return mp == nil || len(mp) == 0
}
