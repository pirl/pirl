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
// the main Pirl network.
var MainnetBootnodes = []string{
	// Pirl Foundation Go Bootnodes
	"enode://59ddbcb6ea8f59f99c5436dafd0560434fa835f3f835060fc1c1babbbf56be1527b1669b50bdb30cb611d085a4a5c6e16e4e72bf721e7da5253cf8e2769e9e94@54.39.181.185:30303",
	"enode://21b56c94b2da60f16c10b4f1231a22c33beb6a52022f7e681e94d57781ce9ee9c76e36c2e20cb9a40cf969b940b81da297de4296087d86d9d4e3f059dbf08dec@54.36.114.205:30303",
	"enode://7374c07c42eb2767b4da977ee8cc0d63a760be0311a218513e5f98b96df9cfd4c7b105e5a772f4d642617974437636637f3047a65869116ebf0bb15559afe5f3@51.75.193.83:30303",
	"enode://e3b014fef0914f599e7c8046068ef6f06aee457ee7ff94c4df3c395d9cc0f0469ec57d207a0a732d062e777604290652bd2753d2d32616a72f53418ad7acbd1f@51.75.169.35:30303",
	"enode://805290942d7d2abd761356624aac90914a180195a6877826ec1549f54880501227fab28ee8649ae7d0c5067ec0000d935a90065f0bdcb97a13499709afcbb7e1@51.75.62.167:30303",
	"enode://4598ccd54d60c9143b7ee80ece7d0694841b9b4e577413669544403d265e84977f32fbd7e0a1c832caa0b4b1863a2d686173cd174d5bffc84c9036353a6b799f@104.248.13.155:30303",
	"enode://67572ef5994de853c57ff521146d24136412510b1fb9d4fd8b5c606d3acd326bab70251a352b6e5b4c05fccc964e41191ca6c1ac36eb7d0c2abc1cf83021a132@139.59.95.243:30303",
	"enode://8d2ceae7a0f3f34276d103ca342c5af62d4f29ed5032fbff50d9cf0b097dd46db74d2631ce1477b76a279adf9282174f35ca12bc2ef4838230d259f009ea0d6d@68.183.236.30:30303",
}


// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure pirl
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

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://59ddbcb6ea8f59f99c5436dafd0560434fa835f3f835060fc1c1babbbf56be1527b1669b50bdb30cb611d085a4a5c6e16e4e72bf721e7da5253cf8e2769e9e94@54.39.181.185:30303",
	"enode://21b56c94b2da60f16c10b4f1231a22c33beb6a52022f7e681e94d57781ce9ee9c76e36c2e20cb9a40cf969b940b81da297de4296087d86d9d4e3f059dbf08dec@54.36.114.205:30303",
	"enode://7374c07c42eb2767b4da977ee8cc0d63a760be0311a218513e5f98b96df9cfd4c7b105e5a772f4d642617974437636637f3047a65869116ebf0bb15559afe5f3@51.75.193.83:30303",
	"enode://e3b014fef0914f599e7c8046068ef6f06aee457ee7ff94c4df3c395d9cc0f0469ec57d207a0a732d062e777604290652bd2753d2d32616a72f53418ad7acbd1f@51.75.169.35:30303",
}

