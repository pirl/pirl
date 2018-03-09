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
	"enode://e28ce45258c481f04be8f84fb9b3ce3906bb54251b17ef957e18afb22a4ac4f784bf39adc5069deec1ca99ee41f02d59d0cb64677bbe661e1e6979da2b7c5276@137.74.31.30:30303",
	"enode://33992f6c62498272d677ae721cdf606a2fbfeb7b1ed2d89a1b432f7945078f4de60cb7e130e1d74b59148863089e52c9c1c7cd1c0ad2e2fa8e95df9e3b858a26@213.32.72.24:30303",
	"enode://973a3d1a0ac3e69b75db4dc4d472cb719529d8c87746e0e0978d0fd40c184e23bca0d2ff01d8c74f5d16689f5a0cc6e95ccc7662148328d6e7c1cd9618573538@158.69.85.131:30303",
	"enode://34435d6908214d6c7d3d8ff218ac703654296a4af6bce21adbd7dbe40d81232389f7f4acf93dfb80fd2e3f3bb79512afddfff0dc43ef348a5b510b945387a24b@213.32.79.9:30303",
	"enode://a4cc2d78255f5eda16527c5566cbcb12f3bae7efe748c787206ad7c2028ad53690a634c9cc40067ad0c13547df721bea23862022817c330988f33fcba7ed03fe@139.59.244.73:30303",
	"enode://81635e82896e4c7687d1fe2ebe1fec7748d8c9af6f30d2ff30b7a2bb16974b85bf13d2d199f44ef13b8f466ff22edf4ca298b08bb58860db2e5b157f74e9a2b0@213.32.53.185:30303",
	"enode://81635e82896e4c7687d1fe2ebe1fec7748d8c9af6f30d2ff30b7a2bb16974b85bf13d2d199f44ef13b8f466ff22edf4ca298b08bb58860db2e5b157f74e9a2b0@213.32.53.186:30303",
	"enode://1c30efbd2a7ce2e6430bd0399a2c3bf6024701dee7b544fe514ee9157ca28596ef941531f43ea2c88eded14ccd38446518a603aa4e903c2db39232fea5c5ab2c@45.77.147.205:30303",
	"enode://662847c24913d76e3859502a7bc4134a8dd663e6349d1cd66af5d72dd32ecfbf8284db199332975c507b038fe11a5f1be4c44b382956fb02936274160cbf7df1@213.32.53.182:30303",
	"enode://53c90a798c90fc454d7a678d15f71bac8605abd408e33124cdb9338296ceb53865afb86af5373f7da4e732590a1822a80ba31398efcc7f569eabb890c9aab7b2@213.32.53.189:30303",
	"enode://fc8b1b0e4342385c22574cec5545f1b35573954231f60dd33b72b4bc1331f581adeec9a4c387d3d9700cf3c7bb0b818a293919153286f8aaf9ade7474906473c@45.77.186.124:30303",
	"enode://88748f08d48beaa629772872fa41f7fedbf53dc956d2ac14830b8816a58aea256ebdada37788472c3ea21098d4e09bc6cd14a38dad1abacbbd2e0c51a804068c@45.77.128.46:30303",
	"enode://8e3ef0a22f72127b3f024310d34fa0a123ab9034b8141afd141e098500f80f061e940626457ad47194ba9125e480ee42752aa088c94d2ee273299400980fe465@45.77.44.95:30303",
	"enode://f577e485d2ac69b297c60f3f34cb08c123b9c2316c2e54722d2e0e4055afc24bccefde1778d18cc73238b2cfc1bab28b0be6c89f8580176974c6e7c1e79edf47@45.32.190.9:30303",
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

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://e28ce45258c481f04be8f84fb9b3ce3906bb54251b17ef957e18afb22a4ac4f784bf39adc5069deec1ca99ee41f02d59d0cb64677bbe661e1e6979da2b7c5276@137.74.31.30:30303",
	"enode://33992f6c62498272d677ae721cdf606a2fbfeb7b1ed2d89a1b432f7945078f4de60cb7e130e1d74b59148863089e52c9c1c7cd1c0ad2e2fa8e95df9e3b858a26@213.32.72.24:30303",
	"enode://973a3d1a0ac3e69b75db4dc4d472cb719529d8c87746e0e0978d0fd40c184e23bca0d2ff01d8c74f5d16689f5a0cc6e95ccc7662148328d6e7c1cd9618573538@158.69.85.131:30303",
	"enode://34435d6908214d6c7d3d8ff218ac703654296a4af6bce21adbd7dbe40d81232389f7f4acf93dfb80fd2e3f3bb79512afddfff0dc43ef348a5b510b945387a24b@213.32.79.9:30303",
	"enode://a4cc2d78255f5eda16527c5566cbcb12f3bae7efe748c787206ad7c2028ad53690a634c9cc40067ad0c13547df721bea23862022817c330988f33fcba7ed03fe@139.59.244.73:30303",
}
