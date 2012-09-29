package main

import(
		"fmt"
		"math"
)

const EARTH_RADIUS = 6378137.0

func rad(d float64) (float64){
	return d*math.Pi/180.0
}

func get_distance(lat1,lng1,lat2,lng2 float64)(float64){
rl1:=rad(lat1)
		rl2:=rad(lat2)
		a:=rl1-rl2
		b:=rad(lng1)-rad(lng2)
		s:=EARTH_RADIUS * 2 * math.Asin(math.Sqrt(math.Pow(math.Sin(a/2),2) + math.Cos(rl1)*math.Cos(rl2)*math.Pow(math.Sin(b/2),2)))
		return s
}

func main(){
	fmt.Println(get_distance(38.043115,114.519145,38.0433,114.518696))
}
