// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

// Package crawler holds the whole crawler service. It includes crawler, db component and GraphQL
package crawler

import (
	"eth2-crawler/crawler/crawl"
	ipResolver "eth2-crawler/resolver"
	"eth2-crawler/store/peerstore"
	"eth2-crawler/store/record"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

//nolint:lll
var GBCBootnodes = []string{
	"enr:-IS4QGmLwm7gFd0L0CEisllrb1op3v-wAGSc7_pwSMGgN3bOS9Fz7m1dWbwuuPHKqeETz9MbhjVuoWk0ohkyRv98kVoBgmlkgnY0gmlwhGjtlgaJc2VjcDI1NmsxoQLMdh0It9fJbuiLydZ9fpF6MRzgNle0vODaDiMqhbC7WIN1ZHCCIyg",
	"enr:-IS4QFUVG3dvLPCUEI7ycRvFm0Ieg_ITa5tALmJ9LI7dJ6ieT3J4fF9xLRjOoB4ApV-Rjp7HeLKzyTWG1xRdbFBNZPQBgmlkgnY0gmlwhErP5weJc2VjcDI1NmsxoQOBbaJBvx0-w_pyZUhQl9A510Ho2T0grE0K8JevzES99IN1ZHCCIyg",
	"enr:-Ku4QOQk8V-Hu2gxFzRXmLYIO4AvWDZhoMFwTf3n3DYm_mbsWv0ZitoqiN6JZUUj6Li6e1Jk1w2zFSVHKPMUP1g5tsgBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD5Jd3FAAAAZP__________gmlkgnY0gmlwhC1PTpmJc2VjcDI1NmsxoQL1Ynt5PoA0UOcHa1Rfn98rmnRlLzNuWTePPP4m4qHVroN1ZHCCKvg",
	"enr:-Ku4QFaTwgoms-EiiRIfHUH3FXprWUFgjHg4UuWvilqoUQtDbmTszVIxUEOwQUmA2qkiP-T9wXjc_rVUuh9cU7WgwbgBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD5Jd3FAAAAZP__________gmlkgnY0gmlwhC0hBmCJc2VjcDI1NmsxoQOpsg1XCrXmCwZKcSTcycLwldoKUMHPUpMEVGeg_EEhuYN1ZHCCKvg",
}

// Start starts the crawler service
func Start(peerStore peerstore.Provider, historyStore record.Provider, ipResolver ipResolver.Provider) {
	h := log.CallerFileHandler(log.StdoutHandler)
	log.Root().SetHandler(h)

	handler := log.MultiHandler(
		log.LvlFilterHandler(log.LvlInfo, h),
	)
	log.Root().SetHandler(handler)

	err := crawl.Initialize(peerStore, historyStore, ipResolver, params.V5Bootnodes)
	if err != nil {
		panic(err)
	}
}
