// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"testing"
)

// Parse
// https://parse.com/docs/rest#summary
var parseAPI = []route{
	// Objects
	{"POST", "/1/classes/:className"},
	{"GET", "/1/classes/:className/:objectId"},
	{"PUT", "/1/classes/:className/:objectId"},
	{"GET", "/1/classes/:className"},
	{"DELETE", "/1/classes/:className/:objectId"},

	// Users
	{"POST", "/1/users"},
	{"GET", "/1/login"},
	{"GET", "/1/users/:objectId"},
	{"PUT", "/1/users/:objectId"},
	{"GET", "/1/users"},
	{"DELETE", "/1/users/:objectId"},
	{"POST", "/1/requestPasswordReset"},

	// Roles
	{"POST", "/1/roles"},
	{"GET", "/1/roles/:objectId"},
	{"PUT", "/1/roles/:objectId"},
	{"GET", "/1/roles"},
	{"DELETE", "/1/roles/:objectId"},

	// Files
	{"POST", "/1/files/:fileName"},

	// Analytics
	{"POST", "/1/events/:eventName"},

	// Push Notifications
	{"POST", "/1/push"},

	// Installations
	{"POST", "/1/installations"},
	{"GET", "/1/installations/:objectId"},
	{"PUT", "/1/installations/:objectId"},
	{"GET", "/1/installations"},
	{"DELETE", "/1/installations/:objectId"},

	// Cloud Functions
	{"POST", "/1/functions"},
}

var (
	parseAce         http.Handler
	parseBadger      http.Handler
	parseBear        http.Handler
	parseDenco       http.Handler
	parseEcho        http.Handler
	parseGin         http.Handler
	parseGoJsonRest  http.Handler
	parseGorillaMux  http.Handler
	parseHttpRouter  http.Handler
	parseHttpTreeMux http.Handler
	parseLARS        http.Handler
	parseMartini     http.Handler
	parsePossum      http.Handler
	parseR2router    http.Handler
	parseRivet       http.Handler
	parseVulcan      http.Handler
	// parseZeus        http.Handler
)

func init() {
	println("#ParseAPI Routes:", len(parseAPI))

	calcMem("Ace", func() {
		parseAce = loadAce(parseAPI)
	})
	calcMem("Badger", func() {
		parseBadger = loadBadger(parseAPI)
	})
	calcMem("Bear", func() {
		parseBear = loadBear(parseAPI)
	})
	calcMem("Denco", func() {
		parseDenco = loadDenco(parseAPI)
	})
	calcMem("Echo", func() {
		parseEcho = loadEcho(parseAPI)
	})
	calcMem("Gin", func() {
		parseGin = loadGin(parseAPI)
	})
	calcMem("GoJsonRest", func() {
		parseGoJsonRest = loadGoJsonRest(parseAPI)
	})
	calcMem("GorillaMux", func() {
		parseGorillaMux = loadGorillaMux(parseAPI)
	})
	calcMem("HttpRouter", func() {
		parseHttpRouter = loadHttpRouter(parseAPI)
	})
	calcMem("HttpTreeMux", func() {
		parseHttpTreeMux = loadHttpTreeMux(parseAPI)
	})
	calcMem("LARS", func() {
		parseLARS = loadLARS(parseAPI)
	})
	calcMem("Martini", func() {
		parseMartini = loadMartini(parseAPI)
	})
	calcMem("Possum", func() {
		parsePossum = loadPossum(parseAPI)
	})
	calcMem("R2router", func() {
		parseR2router = loadR2router(parseAPI)
	})
	calcMem("Rivet", func() {
		parseRivet = loadRivet(parseAPI)
	})
	calcMem("Vulcan", func() {
		parseVulcan = loadVulcan(parseAPI)
	})
	// calcMem("Zeus", func() {
	// 	parseZeus = loadZeus(parseAPI)
	// })

	println()
}

// Static
func BenchmarkAce_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseAce, req)
}
func BenchmarkBadger_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseBadger, req)
}
func BenchmarkBear_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseBear, req)
}
func BenchmarkDenco_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseDenco, req)
}
func BenchmarkEcho_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseEcho, req)
}
func BenchmarkGin_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseGin, req)
}
func BenchmarkGoJsonRest_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseGoJsonRest, req)
}
func BenchmarkGorillaMux_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseGorillaMux, req)
}
func BenchmarkHttpRouter_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseHttpRouter, req)
}
func BenchmarkHttpTreeMux_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseHttpTreeMux, req)
}
func BenchmarkLARS_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseLARS, req)
}
func BenchmarkMartini_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseMartini, req)
}
func BenchmarkPossum_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parsePossum, req)
}
func BenchmarkR2router_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseR2router, req)
}
func BenchmarkRivet_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseRivet, req)
}
func BenchmarkVulcan_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseVulcan, req)
}

// func BenchmarkZeus_ParseStatic(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/1/users", nil)
// 	benchRequest(b, parseZeus, req)
// }

// One Param
func BenchmarkAce_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseAce, req)
}
func BenchmarkBadger_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseBadger, req)
}
func BenchmarkBear_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseBear, req)
}
func BenchmarkDenco_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseDenco, req)
}
func BenchmarkEcho_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseEcho, req)
}
func BenchmarkGin_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseGin, req)
}
func BenchmarkGoJsonRest_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseGoJsonRest, req)
}
func BenchmarkGorillaMux_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseGorillaMux, req)
}
func BenchmarkHttpRouter_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseHttpRouter, req)
}
func BenchmarkHttpTreeMux_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseHttpTreeMux, req)
}
func BenchmarkLARS_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseLARS, req)
}
func BenchmarkMartini_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseMartini, req)
}
func BenchmarkPossum_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parsePossum, req)
}
func BenchmarkR2router_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseR2router, req)
}
func BenchmarkRivet_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseRivet, req)
}
func BenchmarkVulcan_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseVulcan, req)
}

// func BenchmarkZeus_ParseParam(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
// 	benchRequest(b, parseZeus, req)
// }

// Two Params
func BenchmarkAce_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseAce, req)
}
func BenchmarkBadger_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseBadger, req)
}
func BenchmarkBear_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseBear, req)
}
func BenchmarkDenco_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseDenco, req)
}
func BenchmarkEcho_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseEcho, req)
}
func BenchmarkGin_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseGin, req)
}
func BenchmarkGoJsonRest_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseGoJsonRest, req)
}
func BenchmarkGorillaMux_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseGorillaMux, req)
}
func BenchmarkHttpRouter_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseHttpRouter, req)
}
func BenchmarkHttpTreeMux_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseHttpTreeMux, req)
}
func BenchmarkLARS_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseLARS, req)
}
func BenchmarkMartini_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseMartini, req)
}
func BenchmarkPossum_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parsePossum, req)
}
func BenchmarkR2router_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseR2router, req)
}
func BenchmarkRivet_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseRivet, req)
}
func BenchmarkVulcan_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseVulcan, req)
}

// func BenchmarkZeus_Parse2Params(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
// 	benchRequest(b, parseZeus, req)
// }

// All Routes
func BenchmarkAce_ParseAll(b *testing.B) {
	benchRoutes(b, parseAce, parseAPI)
}
func BenchmarkBadger_ParseAll(b *testing.B) {
	benchRoutes(b, parseBadger, parseAPI)
}
func BenchmarkBear_ParseAll(b *testing.B) {
	benchRoutes(b, parseBear, parseAPI)
}
func BenchmarkDenco_ParseAll(b *testing.B) {
	benchRoutes(b, parseDenco, parseAPI)
}
func BenchmarkEcho_ParseAll(b *testing.B) {
	benchRoutes(b, parseEcho, parseAPI)
}
func BenchmarkGin_ParseAll(b *testing.B) {
	benchRoutes(b, parseGin, parseAPI)
}
func BenchmarkGoJsonRest_ParseAll(b *testing.B) {
	benchRoutes(b, parseGoJsonRest, parseAPI)
}
func BenchmarkGorillaMux_ParseAll(b *testing.B) {
	benchRoutes(b, parseGorillaMux, parseAPI)
}
func BenchmarkHttpRouter_ParseAll(b *testing.B) {
	benchRoutes(b, parseHttpRouter, parseAPI)
}
func BenchmarkHttpTreeMux_ParseAll(b *testing.B) {
	benchRoutes(b, parseHttpTreeMux, parseAPI)
}
func BenchmarkLARS_ParseAll(b *testing.B) {
	benchRoutes(b, parseLARS, parseAPI)
}
func BenchmarkMartini_ParseAll(b *testing.B) {
	benchRoutes(b, parseMartini, parseAPI)
}
func BenchmarkPossum_ParseAll(b *testing.B) {
	benchRoutes(b, parsePossum, parseAPI)
}
func BenchmarkR2router_ParseAll(b *testing.B) {
	benchRoutes(b, parseR2router, parseAPI)
}
func BenchmarkRivet_ParseAll(b *testing.B) {
	benchRoutes(b, parseRivet, parseAPI)
}
func BenchmarkVulcan_ParseAll(b *testing.B) {
	benchRoutes(b, parseVulcan, parseAPI)
}

// func BenchmarkZeus_ParseAll(b *testing.B) {
// 	benchRoutes(b, parseZeus, parseAPI)
// }
