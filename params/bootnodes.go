// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Pirl Foundation Go Bootnodes
	"enode://33992f6c62498272d677ae721cdf606a2fbfeb7b1ed2d89a1b432f7945078f4de60cb7e130e1d74b59148863089e52c9c1c7cd1c0ad2e2fa8e95df9e3b858a26@213.32.72.24:30303",
	"enode://a4cc2d78255f5eda16527c5566cbcb12f3bae7efe748c787206ad7c2028ad53690a634c9cc40067ad0c13547df721bea23862022817c330988f33fcba7ed03fe@139.59.244.73:30303",
	"enode://e28ce45258c481f04be8f84fb9b3ce3906bb54251b17ef957e18afb22a4ac4f784bf39adc5069deec1ca99ee41f02d59d0cb64677bbe661e1e6979da2b7c5276@137.74.31.30:30303",
	"enode://59ddbcb6ea8f59f99c5436dafd0560434fa835f3f835060fc1c1babbbf56be1527b1669b50bdb30cb611d085a4a5c6e16e4e72bf721e7da5253cf8e2769e9e94@54.39.181.185:30303", // CA
	"enode://7744086e3ffa81f24b046c0d40de48f1b397174428f7e831aa66a3a7acb70cba9d27cb4fb490c002e576b64e21b202aa59179789b9af80183470af68dcd8f555@54.36.114.205:30303", // DE
	"enode://7374c07c42eb2767b4da977ee8cc0d63a760be0311a218513e5f98b96df9cfd4c7b105e5a772f4d642617974437636637f3047a65869116ebf0bb15559afe5f3@51.75.193.83:30303", // FR
	"enode://e3b014fef0914f599e7c8046068ef6f06aee457ee7ff94c4df3c395d9cc0f0469ec57d207a0a732d062e777604290652bd2753d2d32616a72f53418ad7acbd1f@51.75.169.35:30303", // UK
	"enode://805290942d7d2abd761356624aac90914a180195a6877826ec1549f54880501227fab28ee8649ae7d0c5067ec0000d935a90065f0bdcb97a13499709afcbb7e1@51.75.62.167:30303", // PL
	"enode://67572ef5994de853c57ff521146d24136412510b1fb9d4fd8b5c606d3acd326bab70251a352b6e5b4c05fccc964e41191ca6c1ac36eb7d0c2abc1cf83021a132@139.59.95.243:30303", // BENGLADORE
	"enode://095f192be56f68c11166f231fa270233599d7373df3ceadc578f09117d5f50bdbb13777bb52fb0392ddaadc1a889bc1003da52d2f2b8a6b0bce59dcdd23f2a13@139.59.244.73:30303", //SG

}


// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstrem bootnodes
	"enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	"enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	"enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	"enode://c1f8b7c2ac4453271fa07d8e9ecf9a2e8285aa0bd0c07df0131f47153306b0736fd3db8924e7a9bf0bed6b1d8d4f87362a71b033dc7c64547728d953e43e59b2@52.64.155.147:30303",
	"enode://f4a9c6ee28586009fb5a96c8af13a58ed6d8315a9eee4772212c1d4d9cebe5a8b8a78ea4434f318726317d04a3f531a1ef0420cf9752605a562cfe858c46e263@213.186.16.82:30303",

	// Ethereum Foundation bootnode
	"enode://573b6607cd59f241e30e4c4943fd50e99e2b6f42f9bd5ca111659d309c06741247f4f1e93843ad3e8c8c18b6e2d94c161b7ef67479b3938780a97134b618b5ce@52.56.136.200:30303",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	// Pirl Foundation Go Bootnodes
	"enode://33992f6c62498272d677ae721cdf606a2fbfeb7b1ed2d89a1b432f7945078f4de60cb7e130e1d74b59148863089e52c9c1c7cd1c0ad2e2fa8e95df9e3b858a26@213.32.72.24:30303",
	"enode://a4cc2d78255f5eda16527c5566cbcb12f3bae7efe748c787206ad7c2028ad53690a634c9cc40067ad0c13547df721bea23862022817c330988f33fcba7ed03fe@139.59.244.73:30303",
	"enode://e28ce45258c481f04be8f84fb9b3ce3906bb54251b17ef957e18afb22a4ac4f784bf39adc5069deec1ca99ee41f02d59d0cb64677bbe661e1e6979da2b7c5276@137.74.31.30:30303",
	"enode://59ddbcb6ea8f59f99c5436dafd0560434fa835f3f835060fc1c1babbbf56be1527b1669b50bdb30cb611d085a4a5c6e16e4e72bf721e7da5253cf8e2769e9e94@54.39.181.185:30303", // CA
	"enode://7744086e3ffa81f24b046c0d40de48f1b397174428f7e831aa66a3a7acb70cba9d27cb4fb490c002e576b64e21b202aa59179789b9af80183470af68dcd8f555@54.36.114.205:30303", // DE
	"enode://7374c07c42eb2767b4da977ee8cc0d63a760be0311a218513e5f98b96df9cfd4c7b105e5a772f4d642617974437636637f3047a65869116ebf0bb15559afe5f3@51.75.193.83:30303", // FR
	"enode://e3b014fef0914f599e7c8046068ef6f06aee457ee7ff94c4df3c395d9cc0f0469ec57d207a0a732d062e777604290652bd2753d2d32616a72f53418ad7acbd1f@51.75.169.35:30303", // UK
	"enode://805290942d7d2abd761356624aac90914a180195a6877826ec1549f54880501227fab28ee8649ae7d0c5067ec0000d935a90065f0bdcb97a13499709afcbb7e1@51.75.62.167:30303", // PL
	"enode://67572ef5994de853c57ff521146d24136412510b1fb9d4fd8b5c606d3acd326bab70251a352b6e5b4c05fccc964e41191ca6c1ac36eb7d0c2abc1cf83021a132@139.59.95.243:30303", // BENGLADORE
	"enode://095f192be56f68c11166f231fa270233599d7373df3ceadc578f09117d5f50bdbb13777bb52fb0392ddaadc1a889bc1003da52d2f2b8a6b0bce59dcdd23f2a13@139.59.244.73:30303", //SG
}

