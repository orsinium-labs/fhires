package fhires

// The type of coordinates describing a 3D image region.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/imagingselection-3dgraphictype
type ImagingSelection3DGraphicType string

const (
	// A single location denoted by a single (x,y,z) triplet.
	ImagingSelection3DGraphicTypePoint ImagingSelection3DGraphicType = "point"
	// multiple locations each denoted by an (x,y,z) triplet; the points need not be coplanar.
	ImagingSelection3DGraphicTypeMultipoint ImagingSelection3DGraphicType = "multipoint"
	// a series of connected line segments with ordered vertices denoted by (x,y,z) triplets; the points need not be coplanar.
	ImagingSelection3DGraphicTypePolyline ImagingSelection3DGraphicType = "polyline"
	// a series of connected line segments with ordered vertices denoted by (x,y,z) triplets, where the first and last vertices shall be the same forming a polygon; the points shall be coplanar.
	ImagingSelection3DGraphicTypePolygon ImagingSelection3DGraphicType = "polygon"
	// an ellipse defined by four (x,y,z) triplets, the first two triplets specifying the endpoints of the major axis and the second two triplets specifying the endpoints of the minor axis.
	ImagingSelection3DGraphicTypeEllipse ImagingSelection3DGraphicType = "ellipse"
	// a three-dimensional geometric surface whose plane sections are either ellipses or circles and contains three intersecting orthogonal axes, "a", "b", and "c"; the ellipsoid is defined by six (x,y,z) triplets, the first and second triplets specifying the endpoints of axis "a", the third and fourth triplets specifying the endpoints of axis "b", and the fifth and sixth triplets specifying the endpoints of axis "c".
	ImagingSelection3DGraphicTypeEllipsoid ImagingSelection3DGraphicType = "ellipsoid"
)

func (code ImagingSelection3DGraphicType) String() string {
	return code.Code()
}

func (code ImagingSelection3DGraphicType) Code() string {
	return string(code)
}

func (code ImagingSelection3DGraphicType) Display() string {
	switch code {
	case ImagingSelection3DGraphicTypePoint:
		return "POINT"
	case ImagingSelection3DGraphicTypeMultipoint:
		return "MULTIPOINT"
	case ImagingSelection3DGraphicTypePolyline:
		return "POLYLINE"
	case ImagingSelection3DGraphicTypePolygon:
		return "POLYGON"
	case ImagingSelection3DGraphicTypeEllipse:
		return "ELLIPSE"
	case ImagingSelection3DGraphicTypeEllipsoid:
		return "ELLIPSOID"
	}
	return "<unknown>"
}

func (code ImagingSelection3DGraphicType) Definition() string {
	switch code {
	case ImagingSelection3DGraphicTypePoint:
		return "A single location denoted by a single (x,y,z) triplet."
	case ImagingSelection3DGraphicTypeMultipoint:
		return "multiple locations each denoted by an (x,y,z) triplet; the points need not be coplanar."
	case ImagingSelection3DGraphicTypePolyline:
		return "a series of connected line segments with ordered vertices denoted by (x,y,z) triplets; the points need not be coplanar."
	case ImagingSelection3DGraphicTypePolygon:
		return "a series of connected line segments with ordered vertices denoted by (x,y,z) triplets, where the first and last vertices shall be the same forming a polygon; the points shall be coplanar."
	case ImagingSelection3DGraphicTypeEllipse:
		return "an ellipse defined by four (x,y,z) triplets, the first two triplets specifying the endpoints of the major axis and the second two triplets specifying the endpoints of the minor axis."
	case ImagingSelection3DGraphicTypeEllipsoid:
		return "a three-dimensional geometric surface whose plane sections are either ellipses or circles and contains three intersecting orthogonal axes, \"a\", \"b\", and \"c\"; the ellipsoid is defined by six (x,y,z) triplets, the first and second triplets specifying the endpoints of axis \"a\", the third and fourth triplets specifying the endpoints of axis \"b\", and the fifth and sixth triplets specifying the endpoints of axis \"c\"."
	}
	return "<unknown>"
}
