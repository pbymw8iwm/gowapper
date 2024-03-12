package utilswapper

import (
	"math"
)

const earthR = 6371000.0

func Distance(latA, lngA, latB, lngB float64) float64 {
	var x, y, s, alpha float64
	x = math.Cos(latA*math.Pi/180) * math.Cos(latB*math.Pi/180) * math.Cos((lngA-lngB)*math.Pi/180)
	y = math.Sin(latA*math.Pi/180) * math.Sin(latB*math.Pi/180)
	s = x + y
	if s > 1 {
		s = 1
	}
	if s < -1 {
		s = -1
	}
	alpha = math.Acos(s)
	return alpha * earthR
}

func getDistanceFrom(lat1, lon1, lat2, lon2 float64) float64 {
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}
