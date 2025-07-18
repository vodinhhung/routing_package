package internal

import (
	"math"
	"routing/algorithm/src/dependency"
	"sort"
)

type Coord struct {
	Lat float64
	Lng float64
}

// HaversineDistance returns the distance in km between 2 geo points
func HaversineDistance(p1, p2 Coord) float64 {
	const R = 6371 // Earth radius in KM
	lat1 := p1.Lat * math.Pi / 180
	lat2 := p2.Lat * math.Pi / 180
	dLat := (p2.Lat - p1.Lat) * math.Pi / 180
	dLng := (p2.Lng - p1.Lng) * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLng/2)*math.Sin(dLng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func EuclideanDistance(c1, c2 Coord) float64 {
	dx := c2.Lng - c1.Lng
	dy := c2.Lat - c1.Lat
	return math.Sqrt(dx*dx + dy*dy)
}

func minimalDistance(start Coord, orders []dependency.Order, distanceFunc func(Coord, Coord) float64) ([]dependency.Order, float64) {
	n := len(orders)

	// Precompute all coordinates: 0 is start, 1..n are orders
	all := make([]Coord, n+1)
	all[0] = start
	for i := 0; i < n; i++ {
		all[i+1] = Coord{
			Lat: orders[i].DeliveryLatitude,
			Lng: orders[i].DeliveryLongitude,
		}
	}

	// Distance matrix
	dist := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = make([]float64, n+1)
		for j := 0; j <= n; j++ {
			dist[i][j] = distanceFunc(all[i], all[j])
		}
	}

	// DP setup
	size := 1 << n
	dp := make([][]float64, size)
	path := make([][]int, size)
	for i := range dp {
		dp[i] = make([]float64, n)
		path[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxFloat64
		}
	}

	// Base case: start -> each point
	for i := 0; i < n; i++ {
		dp[1<<i][i] = dist[0][i+1]
	}

	// DP transitions
	for mask := 0; mask < size; mask++ {
		for u := 0; u < n; u++ {
			if (mask>>u)&1 == 0 {
				continue
			}
			for v := 0; v < n; v++ {
				if u == v || (mask>>v)&1 == 1 {
					continue
				}
				nextMask := mask | (1 << v)
				newCost := dp[mask][u] + dist[u+1][v+1]
				if newCost < dp[nextMask][v] {
					dp[nextMask][v] = newCost
					path[nextMask][v] = u
				}
			}
		}
	}

	// Reconstruct path
	endMask := size - 1
	minCost := math.MaxFloat64
	last := -1
	for i := 0; i < n; i++ {
		if dp[endMask][i] < minCost {
			minCost = dp[endMask][i]
			last = i
		}
	}

	// Trace back the order
	order := []int{}
	mask := endMask
	for last != -1 {
		order = append([]int{last}, order...)
		prev := path[mask][last]
		mask ^= 1 << last
		last = prev
		if mask == 0 {
			break
		}
	}

	// Build final route
	finalOrders := []dependency.Order{}
	for _, idx := range order {
		finalOrders = append(finalOrders, orders[idx])
	}

	return finalOrders, minCost
}

func planOptimalRoute(start Coord, orders []dependency.Order) ([]string, map[string][]dependency.Order, []dependency.Order, float64) {
	// Group orders by region
	regions := map[string][]dependency.Order{}
	for _, order := range orders {
		regions[order.RegionCode] = append(regions[order.RegionCode], order)
	}

	// Build list of regions & their centers
	type regionWithOrders struct {
		Code   string
		Center Coord
	}

	regionList := []regionWithOrders{}
	for code := range regions {
		if center, ok := RegionAddress[code]; ok {
			regionList = append(regionList, regionWithOrders{Code: code, Center: center})
		}
	}
	sort.Slice(regionList, func(i, j int) bool { return regionList[i].Code < regionList[j].Code }) // stable

	// Get region visit order using Haversine
	regionOrders := []dependency.Order{}
	for _, r := range regionList {
		regionOrders = append(regionOrders, dependency.Order{
			ID:                0,
			DeliveryLatitude:  r.Center.Lat,
			DeliveryLongitude: r.Center.Lng,
			RegionCode:        r.Code,
		})
	}
	orderedRegionOrders, interRegionCost := minimalDistance(start, regionOrders, HaversineDistance)

	// Build region order list
	orderedRegionCodes := []string{}
	for _, r := range orderedRegionOrders {
		orderedRegionCodes = append(orderedRegionCodes, r.RegionCode)
	}

	// Visit orders within each region using Euclidean
	orderedOrdersPerRegion := map[string][]dependency.Order{}
	mergedOrderList := []dependency.Order{}
	intraRegionTotal := 0.0

	for _, code := range orderedRegionCodes {
		center := RegionAddress[code]
		orderedOrders, intraCost := minimalDistance(center, regions[code], EuclideanDistance)
		orderedOrdersPerRegion[code] = orderedOrders
		mergedOrderList = append(mergedOrderList, orderedOrders...)
		intraRegionTotal += intraCost
	}

	totalDistance := interRegionCost + intraRegionTotal

	return orderedRegionCodes, orderedOrdersPerRegion, mergedOrderList, totalDistance
}
