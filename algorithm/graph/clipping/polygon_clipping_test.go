package clipping

import (
	"testing"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/algorithm/buffer/simplify"
	"github.com/spatial-go/geoos/algorithm/graph/graphtests"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/space"
)

func TestPolygonClipping_Intersection(t *testing.T) {

	for _, tt := range graphtests.TestsPolygonIntersecation {
		if !geoos.GeoosTestTag &&
			tt.Name != "poly poly3" {
			continue
		}
		t.Run(tt.Name, func(t *testing.T) {
			p := &PolygonClipping{
				PointClipping: &PointClipping{tt.Fields[0], tt.Fields[1]},
			}
			got, err := p.Intersection()
			if (err != nil) != tt.WantErr {
				t.Errorf("PolygonClipping.Intersection() %v error = %v, wantErr %v", tt.Name, err, tt.WantErr)
				return
			}
			if !got.Proximity(tt.Want) {
				if gotPoly, ok := got.(matrix.PolygonMatrix); ok {
					if wantPoly, ok := tt.Want.(matrix.PolygonMatrix); ok {
						if measure.AreaOfPolygon(gotPoly) == measure.AreaOfPolygon(wantPoly) {
							return
						}
					}
				}
				t.Errorf("PolygonClipping.Intersection()%v = %v, \nwant %v type %T, want %T", tt.Name, got, tt.Want, got, tt.Want)
			}
		})
	}
}

func TestPolygonClipping_Union(t *testing.T) {

	for _, tt := range graphtests.TestsPolygonUnion {
		if !geoos.GeoosTestTag &&
			tt.Name != "poly x1" {
			continue
		}
		t.Run(tt.Name, func(t *testing.T) {
			p := &PolygonClipping{
				PointClipping: &PointClipping{tt.Fields[0], tt.Fields[1]},
			}
			got, err := p.Union()
			if (err != nil) != tt.WantErr {
				t.Errorf("PolygonClipping.Union() %v error = %v, wantErr %v", tt.Name, err, tt.WantErr)
				return
			}
			isEqual := got.Proximity(tt.Want[0])
			if len(tt.Want) > 1 {
				isEqual1 := got.Proximity(tt.Want[1])
				isEqual = isEqual || isEqual1
			}

			if !isEqual {
				t.Errorf("PolygonClipping.Union()%v = %v, \nwant %v", tt.Name, got, tt.Want)
			}
		})
	}
}

func TestPolygonClipping_Difference(t *testing.T) {

	for _, tt := range graphtests.TestsPolygonDifference {
		if !geoos.GeoosTestTag &&
			tt.Name != "poly poly2" {
			continue
		}
		t.Run(tt.Name, func(t *testing.T) {
			p := &PolygonClipping{
				PointClipping: &PointClipping{tt.Fields[0], tt.Fields[1]},
			}
			got, err := p.Difference()
			if (err != nil) != tt.WantErr {
				t.Errorf("PolygonClipping.Difference() error = %v, wantErr %v", err, tt.WantErr)
				return
			}
			if !got.Proximity(tt.Want) {
				t.Errorf("PolygonClipping.Difference()%v = %v, \nwant %v", tt.Name, got, tt.Want)
			}
		})
	}
}

func TestPolygonClipping_SymDifference(t *testing.T) {

	for _, tt := range graphtests.TestsPolygonSymDifference {
		if !geoos.GeoosTestTag &&
			tt.Name != "poly poly" {
			continue
		}
		t.Run(tt.Name, func(t *testing.T) {
			p := &PolygonClipping{
				PointClipping: &PointClipping{tt.Fields[0], tt.Fields[1]},
			}
			got, err := p.SymDifference()
			if (err != nil) != tt.WantErr {
				t.Errorf("PolygonClipping.SymDifference() error = %v, wantErr %v", err, tt.WantErr)
				return
			}
			if !got.Proximity(tt.Want) {
				t.Errorf("PolygonClipping.SymDifference() = %T\n%v, \nwant %T\n%v", got, got, tt.Want, tt.Want)
			}
		})
	}
}

func TestLargePolygonClipping_Union(t *testing.T) {

	m := make(matrix.Collection, len(graphtests.Tianjian)+len(graphtests.Hebei))
	i := 0
	for j, p := range graphtests.Tianjian {
		i = j
		m[i] = simplify.Simplify(matrix.PolygonMatrix(p), 0.006).(matrix.PolygonMatrix)
	}
	for j, p := range graphtests.Hebei {
		m[i+j+1] = simplify.Simplify(matrix.PolygonMatrix(p), 0.006).(matrix.PolygonMatrix)
	}
	got, err := Union(m[0], m[13])
	if (err != nil) != false {
		t.Errorf("PolygonClipping.Union() error = %v, wantErr %v", err, false)
		return
	}
	writeGeom(dir+"data_union.geojson", space.TransGeometry(got))
}

func TestLargePolygonClipping_UnaryUnion(t *testing.T) {

	m := make(matrix.Collection, len(graphtests.Tianjian)+len(graphtests.Hebei))
	i := 0
	for j, p := range graphtests.Tianjian {
		i = j
		m[i] = simplify.Simplify(matrix.PolygonMatrix(p), 0.008).(matrix.PolygonMatrix)
	}
	for j, p := range graphtests.Hebei {
		m[i+j+1] = simplify.Simplify(matrix.PolygonMatrix(p), 0.008).(matrix.PolygonMatrix)
	}
	got, err := UnaryUnion(m)
	if (err != nil) != false {
		t.Errorf("PolygonClipping.UnaryUnion() error = %v, wantErr %v", err, false)
		return
	}
	writeGeom(dir+"data_unaryunion.geojson", space.TransGeometry(got))
}
