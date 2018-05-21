// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"testing"
)

// Google+
// https://developers.google.com/+/api/latest/
// (in reality this is just a subset of a much larger API)
var gplusAPI = []route{
	// People
	{"GET", "/people/:userId"},
	{"GET", "/people"},
	{"GET", "/activities/:activityId/people/:collection"},
	{"GET", "/people/:userId/people/:collection"},
	{"GET", "/people/:userId/openIdConnect"},

	// Activities
	{"GET", "/people/:userId/activities/:collection"},
	{"GET", "/activities/:activityId"},
	{"GET", "/activities"},

	// Comments
	{"GET", "/activities/:activityId/comments"},
	{"GET", "/comments/:commentId"},

	// Moments
	{"POST", "/people/:userId/moments/:collection"},
	{"GET", "/people/:userId/moments/:collection"},
	{"DELETE", "/moments/:id"},
}

var (
	gplusAce         http.Handler
	gplusBadger      http.Handler
	gplusBear        http.Handler
	gplusDenco       http.Handler
	gplusEcho        http.Handler
	gplusGin         http.Handler
	gplusGoJsonRest  http.Handler
	gplusGorillaMux  http.Handler
	gplusHttpRouter  http.Handler
	gplusHttpTreeMux http.Handler
	gplusLARS        http.Handler
	gplusMartini     http.Handler
	gplusPossum      http.Handler
	gplusR2router    http.Handler
	gplusRivet       http.Handler
	gplusVulcan      http.Handler
	// gplusZeus        http.Handler
)

func init() {
	println("#GPlusAPI Routes:", len(gplusAPI))

	calcMem("Ace", func() {
		gplusAce = loadAce(gplusAPI)
	})
	calcMem("Badger", func() {
		gplusBadger = loadBadger(gplusAPI)
	})
	calcMem("Bear", func() {
		gplusBear = loadBear(gplusAPI)
	})
	calcMem("Denco", func() {
		gplusDenco = loadDenco(gplusAPI)
	})
	calcMem("Echo", func() {
		gplusEcho = loadEcho(gplusAPI)
	})
	calcMem("Gin", func() {
		gplusGin = loadGin(gplusAPI)
	})
	calcMem("GoJsonRest", func() {
		gplusGoJsonRest = loadGoJsonRest(gplusAPI)
	})
	calcMem("GorillaMux", func() {
		gplusGorillaMux = loadGorillaMux(gplusAPI)
	})
	calcMem("HttpRouter", func() {
		gplusHttpRouter = loadHttpRouter(gplusAPI)
	})
	calcMem("HttpTreeMux", func() {
		gplusHttpTreeMux = loadHttpTreeMux(gplusAPI)
	})
	calcMem("LARS", func() {
		gplusLARS = loadLARS(gplusAPI)
	})
	calcMem("Martini", func() {
		gplusMartini = loadMartini(gplusAPI)
	})
	calcMem("Possum", func() {
		gplusPossum = loadPossum(gplusAPI)
	})
	calcMem("R2router", func() {
		gplusR2router = loadR2router(gplusAPI)
	})
	calcMem("Rivet", func() {
		gplusRivet = loadRivet(gplusAPI)
	})
	calcMem("Vulcan", func() {
		gplusVulcan = loadVulcan(gplusAPI)
	})
	// calcMem("Zeus", func() {
	// 	gplusZeus = loadZeus(gplusAPI)
	// })

	println()
}

// Static
func BenchmarkAce_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusAce, req)
}
func BenchmarkBadger_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusBadger, req)
}
func BenchmarkBear_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusBear, req)
}
func BenchmarkDenco_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusDenco, req)
}
func BenchmarkEcho_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusEcho, req)
}
func BenchmarkGin_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGin, req)
}
func BenchmarkGoJsonRest_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGoJsonRest, req)
}
func BenchmarkGorillaMux_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGorillaMux, req)
}
func BenchmarkHttpRouter_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusHttpRouter, req)
}
func BenchmarkHttpTreeMux_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusHttpTreeMux, req)
}
func BenchmarkLARS_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusLARS, req)
}
func BenchmarkMartini_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusMartini, req)
}
func BenchmarkPossum_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusPossum, req)
}
func BenchmarkR2router_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusR2router, req)
}
func BenchmarkRivet_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusRivet, req)
}
func BenchmarkVulcan_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusVulcan, req)
}

// func BenchmarkZeus_GPlusStatic(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/people", nil)
// 	benchRequest(b, gplusZeus, req)
// }

// One Param
func BenchmarkAce_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusAce, req)
}
func BenchmarkBadger_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusBadger, req)
}
func BenchmarkBear_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusBear, req)
}
func BenchmarkDenco_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusDenco, req)
}
func BenchmarkEcho_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusEcho, req)
}
func BenchmarkGin_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGin, req)
}
func BenchmarkGoJsonRest_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGoJsonRest, req)
}
func BenchmarkGorillaMux_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGorillaMux, req)
}
func BenchmarkHttpRouter_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusHttpRouter, req)
}
func BenchmarkHttpTreeMux_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusHttpTreeMux, req)
}
func BenchmarkLARS_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusLARS, req)
}
func BenchmarkMartini_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusMartini, req)
}
func BenchmarkPossum_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusPossum, req)
}
func BenchmarkR2router_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusR2router, req)
}
func BenchmarkRivet_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusRivet, req)
}
func BenchmarkVulcan_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusVulcan, req)
}

// func BenchmarkZeus_GPlusParam(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
// 	benchRequest(b, gplusZeus, req)
// }

// Two Params
func BenchmarkAce_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusAce, req)
}
func BenchmarkBadger_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusBadger, req)
}
func BenchmarkBear_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusBear, req)
}
func BenchmarkDenco_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusDenco, req)
}
func BenchmarkEcho_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusEcho, req)
}
func BenchmarkGin_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGin, req)
}
func BenchmarkGoJsonRest_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGoJsonRest, req)
}
func BenchmarkGorillaMux_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGorillaMux, req)
}
func BenchmarkHttpRouter_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusHttpRouter, req)
}
func BenchmarkHttpTreeMux_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusHttpTreeMux, req)
}
func BenchmarkLARS_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusLARS, req)
}
func BenchmarkMartini_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusMartini, req)
}
func BenchmarkPossum_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusPossum, req)
}
func BenchmarkR2router_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusR2router, req)
}
func BenchmarkRivet_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusRivet, req)
}
func BenchmarkVulcan_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusVulcan, req)
}

// func BenchmarkZeus_GPlus2Params(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
// 	benchRequest(b, gplusZeus, req)
// }

// All Routes
func BenchmarkAce_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusAce, gplusAPI)
}
func BenchmarkBadger_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusBadger, gplusAPI)
}
func BenchmarkBear_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusBear, gplusAPI)
}
func BenchmarkDenco_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusDenco, gplusAPI)
}
func BenchmarkEcho_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusEcho, gplusAPI)
}
func BenchmarkGin_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGin, gplusAPI)
}
func BenchmarkGoJsonRest_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGoJsonRest, gplusAPI)
}
func BenchmarkGorillaMux_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGorillaMux, gplusAPI)
}
func BenchmarkHttpRouter_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusHttpRouter, gplusAPI)
}
func BenchmarkHttpTreeMux_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusHttpTreeMux, gplusAPI)
}
func BenchmarkLARS_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusLARS, gplusAPI)
}
func BenchmarkMartini_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusMartini, gplusAPI)
}
func BenchmarkPossum_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusPossum, gplusAPI)
}
func BenchmarkR2router_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusR2router, gplusAPI)
}
func BenchmarkRivet_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusRivet, gplusAPI)
}
func BenchmarkVulcan_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusVulcan, gplusAPI)
}

// func BenchmarkZeus_GPlusAll(b *testing.B) {
// 	benchRoutes(b, gplusZeus, gplusAPI)
// }
