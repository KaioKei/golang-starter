package backend

import (
	"fmt"
	pb "golang_starter/internal/api/grpc/go-grpc/route-guide"
	"math"
)

var eiffelTourPoint = pb.Point{Latitude: 48858370, Longitude: 2294481}

var Features = []*pb.Feature{
	{
		Name:     "Eiffel Tour",
		Location: &eiffelTourPoint,
	},
}

var SavedRoute = []*pb.Point{
	{Latitude: 48862578, Longitude: 2287758},
	{Latitude: 48860672, Longitude: 2290730},
	&eiffelTourPoint,
	{Latitude: 48856155, Longitude: 2298176},
	{Latitude: 48852710, Longitude: 2302918},
}

var SavedNotes = []*pb.RouteNote{
	{Location: SavedRoute[0], Message: "First message"},
	{Location: SavedRoute[1], Message: "Second message"},
	{Location: SavedRoute[2], Message: "Eiffel tower !!"},
	{Location: SavedRoute[3], Message: "Fourth message"},
	{Location: SavedRoute[4], Message: "Fifth message"},
}

// InRange takes a Point struct and tells if the Point's location (Longitude and Latitude) is inside
// a given Rectangle (composed of 2 points)
func InRange(point *pb.Point, rect *pb.Rectangle) bool {
	left := math.Min(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	right := math.Max(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	top := math.Max(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))
	bottom := math.Min(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))

	if float64(point.Longitude) >= left &&
		float64(point.Longitude) <= right &&
		float64(point.Latitude) >= bottom &&
		float64(point.Latitude) <= top {
		return true
	}
	return false
}

func toRadians(num float64) float64 {
	return num * math.Pi / float64(180)
}

// CalcDistance calculates the distance between two points using the "haversine" formula.
// The formula is based on http://mathforum.org/library/drmath/view/51879.html.
func CalcDistance(p1 *pb.Point, p2 *pb.Point) int32 {
	const CordFactor float64 = 1e7
	const R = float64(6371000) // earth radius in metres
	lat1 := toRadians(float64(p1.Latitude) / CordFactor)
	lat2 := toRadians(float64(p2.Latitude) / CordFactor)
	lng1 := toRadians(float64(p1.Longitude) / CordFactor)
	lng2 := toRadians(float64(p2.Longitude) / CordFactor)
	dlat := lat2 - lat1
	dlng := lng2 - lng1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return int32(distance)
}

// Serialize function for a point is just a string to collapse the location into a text
func Serialize(point *pb.Point) string {
	return fmt.Sprintf("%d %d", point.Latitude, point.Longitude)
}
