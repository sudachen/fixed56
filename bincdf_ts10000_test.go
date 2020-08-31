package fixed56

import "testing"

var bincdfTestCases_10000 = []BinCDFCase{
	{
		n:   2824,
		x:   294,
		p:   0.2483100999519666,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   2824,
		x:   1913,
		p:   0.4446421476297646,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   2824,
		x:   1695,
		p:   0.42049903523430154,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   2824,
		x:   2640,
		p:   0.5784359135409691,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   2824,
		x:   2498,
		p:   0.818098502489808,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   1520,
		x:   159,
		p:   0.11873904797231849,
		cdf: 0.04620028268945807,
		s:   "+0'0bd3c81f3798ef/56",
	},
	{
		n:   1520,
		x:   998,
		p:   0.5010710576206725,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   1520,
		x:   519,
		p:   0.5122490125877226,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   1520,
		x:   1378,
		p:   0.5299768875559047,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   1520,
		x:   1229,
		p:   0.7839039641923319,
		cdf: 0.9916878564399366,
		s:   "+0'fddf415f3ff788/56",
	},
	{
		n:   8359,
		x:   1438,
		p:   0.15563814164613257,
		cdf: 0.999978954555814,
		s:   "+0'fffe9eea8167e0/56",
	},
	{
		n:   8359,
		x:   4232,
		p:   0.5517614734259535,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   8359,
		x:   3439,
		p:   0.5396278789976454,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   8359,
		x:   6965,
		p:   0.45557426833432835,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   8359,
		x:   7304,
		p:   0.8914426144413563,
		cdf: 2.0893519712083095e-07,
		s:   "+0'000003815eadbe/56",
	},
	{
		n:   6514,
		x:   755,
		p:   0.11854916867602958,
		cdf: 0.261625452704057,
		s:   "+0'42f9e2bb2a4548/56",
	},
	{
		n:   6514,
		x:   3653,
		p:   0.47179587609692575,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   6514,
		x:   3579,
		p:   0.5207452062733783,
		cdf: 0.999998354326247,
		s:   "+0'ffffe463e292b0/56",
	},
	{
		n:   6514,
		x:   4519,
		p:   0.5459463573387636,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   6514,
		x:   5314,
		p:   0.7249652325706419,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   7201,
		x:   800,
		p:   0.21040812625464542,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   7201,
		x:   6997,
		p:   0.5257279598988888,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   7201,
		x:   3881,
		p:   0.5154704290513524,
		cdf: 0.9999685434091704,
		s:   "+0'fffdf03ef89cd8/56",
	},
	{
		n:   7201,
		x:   5084,
		p:   0.40916487673113244,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   7201,
		x:   6248,
		p:   0.8546136681577384,
		cdf: 0.9992980781206945,
		s:   "+0'ffd1ffb47be9c8/56",
	},
	{
		n:   2307,
		x:   448,
		p:   0.14655817727220605,
		cdf: 0.9999999997929367,
		s:   "+0'ffffffff1c54e0/56",
	},
	{
		n:   2307,
		x:   1359,
		p:   0.4760252450078333,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   2307,
		x:   1233,
		p:   0.5271368888528801,
		cdf: 0.7658373074373637,
		s:   "+0'c40de9ed800b70/56",
	},
	{
		n:   2307,
		x:   1911,
		p:   0.43253081943121696,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   2307,
		x:   1987,
		p:   0.7419014061542976,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   5374,
		x:   610,
		p:   0.22182620113339765,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   5374,
		x:   3387,
		p:   0.5068279754059416,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   5374,
		x:   2292,
		p:   0.4326804987523857,
		cdf: 0.1838209840343222,
		s:   "+0'2f0ee45abef6b6/56",
	},
	{
		n:   5374,
		x:   4359,
		p:   0.45398956535685564,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   5374,
		x:   4613,
		p:   0.836922850197975,
		cdf: 0.9999933534862768,
		s:   "+0'ffff907d70e090/56",
	},
	{
		n:   1916,
		x:   249,
		p:   0.2643605188013617,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   1916,
		x:   1782,
		p:   0.46309060961181636,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   1916,
		x:   912,
		p:   0.4132377210423095,
		cdf: 0.9999999871660684,
		s:   "+0'ffffffc8e0f168/56",
	},
	{
		n:   1916,
		x:   1856,
		p:   0.5752735252945338,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   1916,
		x:   1645,
		p:   0.7425253088108282,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   9179,
		x:   1322,
		p:   0.27693663077734054,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   9179,
		x:   8347,
		p:   0.4285743180884125,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   9179,
		x:   3630,
		p:   0.4493255015387967,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   9179,
		x:   8417,
		p:   0.5077954579334581,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   9179,
		x:   7498,
		p:   0.8169171980447082,
		cdf: 0.49904593676568776,
		s:   "+0'7fc179799c3d78/56",
	},
	{
		n:   7543,
		x:   1124,
		p:   0.14386415183145668,
		cdf: 0.9011110841688819,
		s:   "+0'e6af374c918600/56",
	},
	{
		n:   7543,
		x:   4337,
		p:   0.5019052587352929,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   7543,
		x:   2886,
		p:   0.5511564315067475,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   7543,
		x:   5477,
		p:   0.4305682638696522,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   7543,
		x:   6626,
		p:   0.8584158728725929,
		cdf: 0.9999998337735162,
		s:   "+0'fffffd36100c58/56",
	},
	{
		n:   10000,
		x:   1,
		p:   0.1,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   1001,
		p:   0.1,
		cdf: 0.5217047111595491,
		s:   "+0'858e70a0997198/56",
	},
	{
		n:   10000,
		x:   2001,
		p:   0.1,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   3001,
		p:   0.1,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   4001,
		p:   0.1,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   5001,
		p:   0.1,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   6001,
		p:   0.1,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   7001,
		p:   0.1,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   8001,
		p:   0.1,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   9001,
		p:   0.1,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   1,
		p:   0.2,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   1001,
		p:   0.2,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   2001,
		p:   0.2,
		cdf: 0.5159518893397638,
		s:   "+0'84156c4b0610f8/56",
	},
	{
		n:   10000,
		x:   3001,
		p:   0.2,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   4001,
		p:   0.2,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   5001,
		p:   0.2,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   6001,
		p:   0.2,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   7001,
		p:   0.2,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   8001,
		p:   0.2,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   9001,
		p:   0.2,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   1,
		p:   0.3,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   1001,
		p:   0.3,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   2001,
		p:   0.3,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   3001,
		p:   0.3,
		cdf: 0.5136354443606586,
		s:   "+0'837d9ccb986fd0/56",
	},
	{
		n:   10000,
		x:   4001,
		p:   0.3,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   5001,
		p:   0.3,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   6001,
		p:   0.3,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   7001,
		p:   0.3,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   8001,
		p:   0.3,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   9001,
		p:   0.3,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   1,
		p:   0.4,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   1001,
		p:   0.4,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   2001,
		p:   0.4,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   3001,
		p:   0.4,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   4001,
		p:   0.4,
		cdf: 0.5124841213347535,
		s:   "+0'833228ccda20c0/56",
	},
	{
		n:   10000,
		x:   5001,
		p:   0.4,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   6001,
		p:   0.4,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   7001,
		p:   0.4,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   8001,
		p:   0.4,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   9001,
		p:   0.4,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   1,
		p:   0.5,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   1001,
		p:   0.5,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   2001,
		p:   0.5,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   3001,
		p:   0.5,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   4001,
		p:   0.5,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   5001,
		p:   0.5,
		cdf: 0.5119663737953237,
		s:   "+0'83103a701a45c0/56",
	},
	{
		n:   10000,
		x:   6001,
		p:   0.5,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   7001,
		p:   0.5,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   8001,
		p:   0.5,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   9001,
		p:   0.5,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   1,
		p:   0.6,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   1001,
		p:   0.6,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   2001,
		p:   0.6,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   3001,
		p:   0.6,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   4001,
		p:   0.6,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   5001,
		p:   0.6,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   6001,
		p:   0.6,
		cdf: 0.5119419673303732,
		s:   "+0'830ea0f7224438/56",
	},
	{
		n:   10000,
		x:   7001,
		p:   0.6,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   8001,
		p:   0.6,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   9001,
		p:   0.6,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   1,
		p:   0.7,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   1001,
		p:   0.7,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   2001,
		p:   0.7,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   3001,
		p:   0.7,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   4001,
		p:   0.7,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   5001,
		p:   0.7,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   6001,
		p:   0.7,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   7001,
		p:   0.7,
		cdf: 0.5124764954612413,
		s:   "+0'8331a8dbf98780/56",
	},
	{
		n:   10000,
		x:   8001,
		p:   0.7,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   9001,
		p:   0.7,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   1,
		p:   0.8,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   1001,
		p:   0.8,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   2001,
		p:   0.8,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   3001,
		p:   0.8,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   4001,
		p:   0.8,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   5001,
		p:   0.8,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   6001,
		p:   0.8,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   7001,
		p:   0.8,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   8001,
		p:   0.8,
		cdf: 0.513961242136345,
		s:   "+0'8392f6c6631b88/56",
	},
	{
		n:   10000,
		x:   9001,
		p:   0.8,
		cdf: 0.9999999999999999,
		s:   "+0'fffffffffffff8/56",
	},
	{
		n:   10000,
		x:   1,
		p:   0.9,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   1001,
		p:   0.9,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   2001,
		p:   0.9,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   3001,
		p:   0.9,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   4001,
		p:   0.9,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   5001,
		p:   0.9,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   6001,
		p:   0.9,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   7001,
		p:   0.9,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   8001,
		p:   0.9,
		cdf: 0,
		s:   "+0'00000000000000/56",
	},
	{
		n:   10000,
		x:   9001,
		p:   0.9,
		cdf: 0.5181713946056972,
		s:   "+0'84a6e1698ddec0/56",
	},
}

func Benchmark_Fixed_BinCDF_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tc := bincdfTestCases_10000[i%len(bincdfTestCases_10000)]
		bincdfResultFix = BinCDF(tc.n, From(tc.p), tc.x)
	}

	bincdfResultFix.lo++
}

func Benchmark_Float_BinCDF_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tc := bincdfTestCases_10000[i%len(bincdfTestCases_10000)]
		bincdfResultFlt = bincdf_(tc.n, tc.p, tc.x)
	}

	bincdfResultFlt++
}

func TestFixed_BinCDF_10000(t *testing.T) {
	acc := accuracy{Epsilon: 1e-10}
	for i, tc := range bincdfTestCases_10000 {
		p := From(tc.p)
		got := BinCDF(tc.n, p, tc.x)
		if ok := acc.update(got, tc.cdf); !ok {
			t.Errorf("%d: BinCDF(%v,%v,%v) => got %v|%v, want %v|%v", i, tc.n, tc.p, tc.x, got, got.Float(), From(tc.cdf), tc.cdf)
		}
	}
	t.Log(acc)
}