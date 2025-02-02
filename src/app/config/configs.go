package config

import (
	"encoding/json"
	"kscan/src/lib/misc"
	"kscan/src/lib/slog"
	"os"
)

type config_ struct {
	Ports                    []int
	SslPorts                 []int
	HttpCode                 []int
	Path                     []string
	UserAgents               []string
	UnWebPorts               map[string]string
	Top, Threads, Timeout    int
	Host, FofaEmail, FofaKey string
}

var Config config_

var defaultConfigChar = `
{
  "fofaEmail": "",
  "fofaKey": "",
  "userAgents": [
    "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)",
    "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; AcooBrowser; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
    "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; Acoo Browser; SLCC1; .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.0.04506)",
    "Mozilla/4.0 (compatible; MSIE 7.0; AOL 9.5; AOLBuild 4337.35; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
    "Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)",
    "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 2.0.50727; Media Center PC 6.0)",
    "Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 1.0.3705; .NET CLR 1.1.4322)",
    "Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.2; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.2; .NET CLR 3.0.04506.30)",
    "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN) AppleWebKit/523.15 (KHTML, like Gecko, Safari/419.3) Arora/0.3 (Change: 287 c9dfb30)",
    "Mozilla/5.0 (X11; U; Linux; en-US) AppleWebKit/527+ (KHTML, like Gecko, Safari/419.3) Arora/0.6",
    "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.2pre) Gecko/20070215 K-Ninja/2.1.1",
    "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.9) Gecko/20080705 Firefox/3.0 Kapiko/3.0",
    "Mozilla/5.0 (X11; Linux i686; U;) Gecko/20070322 Kazehakase/0.4.5",
    "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.8) Gecko Fedora/1.9.0.8-1.fc10 Kazehakase/0.5.6",
    "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.20 (KHTML, like Gecko) Chrome/19.0.1036.7 Safari/535.20",
    "Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; fr) Presto/2.9.168 Version/11.52"
  ],
  "top": 400,
  "threads": 400,
  "httpCode": [],
  "path": [
    "/"
  ],
  "host": "",
  "ports": [
	22,
	23,
	21,
	2100,
	25,
	445,
	135,
	139,
	389,
	636,
	2049,
	1433,
	1434,
	1521,
	3306,
	3389,
	5432,
	50000,
    8080,
    80,
    81,
    8081,
    7001,
    8000,
    8088,
    8888,
    9090,
    8090,
    88,
    8001,
    82,
    9080,
    8082,
    8089,
    9000,
    8443,
    9999,
    8002,
    89,
    8083,
    8200,
    8008,
    90,
    8086,
    801,
    8011,
    8085,
    9001,
    9200,
    8100,
    8012,
    85,
    8084,
    8070,
    7002,
    8091,
    8003,
    99,
    7777,
    8010,
    443,
    8028,
    8087,
    83,
    7003,
    10000,
    808,
    38888,
    8181,
    800,
    18080,
    8099,
    8899,
    86,
    8360,
    8300,
    8800,
    8180,
    3505,
    7000,
    9002,
    8053,
    1000,
    7080,
    8989,
    28017,
    9060,
    888,
    3000,
    8006,
    41516,
    880,
    8484,
    6677,
    8016,
    84,
    7200,
    9085,
    5555,
    8280,
    7005,
    1980,
    8161,
    9091,
    7890,
    8060,
    6080,
    8880,
    8020,
    7070,
    889,
    8881,
    9081,
    8009,
    7007,
    8004,
    38501,
    1010,
    93,
    6666,
    7010,
    100,
    9003,
    6789,
    7060,
    8018,
    8022,
    4848,
    3050,
    8787,
    2000,
    9010,
    10001,
    8013,
    6888,
    8040,
    10021,
    1080,
    2011,
    6006,
    4000,
    5000,
    8055,
    4430,
    1723,
    6060,
    7788,
    8066,
    9898,
    6001,
    8801,
    10040,
    7006,
    9998,
    803,
    6688,
    10080,
    7008,
    8050,
    7011,
    7009,
    40310,
    18090,
    802,
    10003,
    8014,
    2080,
    7288,
    8044,
    9992,
    8005,
    8889,
    5644,
    8886,
    9500,
    58031,
    50000,
    9020,
    8015,
    50060,
    8887,
    8021,
    8700,
    91,
    9900,
    9191,
    3312,
    8186,
    8735,
    8380,
    1234,
    38080,
    9088,
    9988,
    2110,
    8007,
    21245,
    3333,
    2046,
    9061,
    8686,
    2375,
    9011,
    8061,
    8093,
    9876,
    8030,
    8282,
    60465,
    2222,
    98,
    9009,
    1100,
    18081,
    70,
    8383,
    5155,
    92,
    8188,
    2517,
    50070,
    8062,
    11324,
    2008,
    9231,
    999,
    28214,
    5001,
    16080,
    8092,
    8987,
    8038,
    809,
    2010,
    8983,
    7700,
    3535,
    7921,
    9093,
    11080,
    6778,
    805,
    9083,
    8073,
    10002,
    114,
    2012,
    701,
    8810,
    8400,
    9099,
    8098,
    9007,
    8808,
    20000,
    8065,
    8822,
    15000,
    9100,
    9901,
    11158,
    1107,
    28099,
    12345,
    2006,
    9527,
    51106,
    688,
    25006,
    8045,
    9006,
    8023,
    8029,
    9997,
    9043,
    7048,
    8580,
    8585,
    2001,
    8035,
    10088,
    20022,
    4001,
    9005,
    2013,
    20808,
    8095,
    106,
    3580,
    7742,
    8119,
    7071,
    6868,
    32766,
    50075,
    7272,
    3380,
    3220,
    7801,
    5256,
    5255,
    10086,
    1300,
    5200,
    8096,
    6198,
    1158,
    6889,
    3503,
    6088,
    9991,
    9008,
    806,
    7004,
    5050,
    8183,
    8688,
    1001,
    58080,
    1182,
    9025,
    8112,
    7776,
    7321,
    235,
    8077,
    8500,
    11347,
    7081,
    8877,
    8480,
    9182,
    58000,
    8026,
    11001,
    10089,
    5888,
    8196,
    8078,
    9995,
    2014,
    5656,
    8019,
    5003,
    8481,
    6002,
    9889,
    9015,
    8866,
    8182,
    8057,
    8399,
    10010,
    8308,
    511,
    12881,
    4016,
    8042,
    1039,
    28080,
    5678,
    7500,
    8051,
    18801,
    15018,
    15888,
    38443,
    8123,
    9004,
    8144,
    94,
    9070,
    1800,
    9112,
    8990,
    3456,
    2051,
    9098,
    444,
    9131,
    97,
    7100,
    7711,
    7180,
    11000,
    8037,
    6988,
    122,
    8885,
    14007,
    8184,
    7012,
    8079,
    9888,
    9301,
    59999,
    49705,
    1979,
    8900,
    5080,
    5013,
    1550,
    8844,
    4850,
    206,
    5156,
    8813,
    3030,
    1790,
    8802,
    9012,
    5544,
    3721,
    8980,
    10009,
    8043,
    8390,
    7943,
    8381,
    8056,
    7111,
    1500,
    7088,
    5881,
    9437,
    5655,
    8102,
    6000,
    65486,
    4443,
    3690,
    2181,
    10025,
    8024,
    8333,
    8666,
    103,
    8,
    9666,
    8999,
    9111,
    8071,
    9092,
    522,
    11381,
    20806,
    8041,
    1085,
    8864,
    7900,
    1700,
    8036,
    8032,
    8033,
    8111,
    60022,
    955,
    3080,
    8788,
    27017,
    7443,
    8192,
    6969,
    9909,
    5002,
    9990,
    188,
    8910,
    9022,
    50030,
    10004,
    866,
    8582,
    4300,
    9101,
    6879,
    8891,
    4567,
    4440,
    10051,
    10068,
    50080,
    8341,
    30001,
    6890,
    8168,
    8955,
    16788,
    8190,
    18060,
    6379,
    7041,
    42424,
    8848,
    15693,
    2521,
    19010,
    18103,
    6010,
    8898,
    9910,
    9190,
    9082,
    8260,
    8445,
    1680,
    8890,
    8649,
    30082,
    3013,
    30000,
    2480,
    7202,
    9704,
    5233,
    8991,
    11366,
    7888,
    8780,
    7129,
    6600,
    9443,
    47088,
    7791,
    18888,
    50045,
    15672,
    9089,
    2585,
    60,
    9494,
    31945,
    2060,
    8610,
    8860,
    58060,
    6118,
    2348,
    8097,
    38000,
    18880,
    13382,
    6611,
    8064,
    7101,
    5081,
    7380,
    7942,
    10016,
    8027,
    2093,
    403,
    9014,
    8133,
    6886,
    95,
    8058,
    9201,
    6443,
    5966,
    27000,
    7017,
    6680,
    8401,
    9036,
    8988,
    8806,
    6180,
    421,
    423,
    57880,
    7778,
    18881,
    812,
    3306,
    15004,
    9110,
    8213,
    8868,
    9300,
    87,
    1213,
    8193,
    8956,
    1108,
    778,
    65000,
    7020,
    1122,
    9031,
    17000,
    8039,
    8600,
    50090,
    1863,
    8191,
    65,
    6587,
    8136,
    9507,
    132,
    200,
    2070,
    308,
    5811,
    3465,
    8680,
    7999,
    7084,
    18082,
    3938,
    18001,
    8069,
    5902,
    9595,
    442,
    4433,
    7171,
    9084,
    7567,
    811,
    1128,
    6003,
    2125,
    6090,
    10007,
    7022,
    1949,
    6565,
    65001,
    1301,
    19244,
    10087,
    8025,
    5098,
    21080,
    1200,
    15801,
    1005,
    22343,
    7086,
    8601,
    6259,
    7102,
    10333,
    211,
    10082,
    18085,
    180,
    40000,
    7021,
    7702,
    66,
    38086,
    666,
    6603,
    1212,
    65493,
    96,
    9053,
    7031,
    23454,
    30088,
    6226,
    8660,
    6170,
    8972,
    9981,
    48080,
    9086,
    10118,
    40069,
    28780,
    20153,
    20021,
    20151,
    58898,
    10066,
    1818,
    9914,
    55351,
    8343,
    18000,
    6546,
    3880,
    8902,
    22222,
    19045,
    5561,
    7979,
    5203,
    8879,
    50240,
    49960,
    2007,
    1722,
    25,
    8913,
    8912,
    9504,
    8103,
    8567,
    1666,
    8720,
    8197,
    3012,
    8220,
    9039,
    5898,
    925,
    38517,
    8382,
    6842,
    8895,
    2808,
    447,
    3600,
    3606,
    9095,
    45177,
    19101,
    171,
    133,
    8189,
    7108,
    10154,
    47078,
    6800,
    8122,
    381,
    1443,
    15580,
    23352,
    3443,
    1180,
    268,
    2382,
    43651,
    10099,
    65533,
    7018,
    60010,
    60101,
    6699,
    2005,
    18002,
    2009,
    59777,
    591,
    1933,
    9013,
    8477,
    9696,
    9030,
    2015,
    7925,
    6510,
    18803,
    280,
    5601,
    2901,
    2301,
    5201,
    302,
    610,
    8031,
    5552,
    8809,
    6869,
    9212,
    17095,
    20001,
    8781,
    25024,
    5280,
    7909,
    17003,
    1088,
    7117,
    20052,
    1900,
    10038,
    30551,
    9980,
    9180,
    59009,
    28280,
    7028,
    61999,
    7915,
    8384,
    9918,
    9919,
    55858,
    7215,
    77,
    9845,
    20140,
    8288,
    7856,
    1982,
    1123,
    17777,
    8839,
    208,
    2886,
    877,
    6101,
    5100,
    804,
    983,
    5600,
    8402,
    5887,
    8322,
    5632,
    770,
    13333,
    7330,
    3216,
    31188,
    47583,
    8710,
    22580,
    1042,
    2020,
    34440,
    20,
    7703,
    65055,
    8997,
    6543,
    6388,
    8283,
    7201,
    4040,
    61081,
    12001,
    3588,
    7123,
    2490,
    4389,
    1313,
    19080,
    9050,
    6920,
    299,
    20046,
    8892,
    9302,
    7899,
    30058,
    7094,
    6801,
    321,
    1356,
    12333,
    11362,
    11372,
    6602,
    7709,
    45149,
    3668,
    517,
    9912,
    9096,
    8130,
    7050,
    7713,
    40080,
    8104,
    13988,
    18264,
    8799,
    7072,
    55070,
    23458,
    8176,
    9517,
    9541,
    9542,
    9512,
    8905,
    11660,
    1025,
    44445,
    44401,
    17173,
    436,
    560,
    733,
    968,
    602,
    3133,
    3398,
    16580,
    8488,
    8901,
    8512,
    10443,
    9113,
    9119,
    6606,
    22080,
    5560,
    7,
    5757,
    1600,
    8250,
    10024,
    10200,
    333,
    73,
    7547,
    8054,
    6372,
    223,
    3737,
    9800,
    9019,
    8067,
    45692,
    15400,
    15698,
    9038,
    37006,
    2086,
    1002,
    9188,
    8094,
    8201,
    8202,
    30030,
    2663,
    9105,
    10017,
    4503,
    1104,
    8893,
    40001,
    27779,
    3010,
    7083,
    5010,
    5501,
    309,
    1389,
    10070,
    10069,
    10056,
    3094,
    10057,
    10078,
    10050,
    10060,
    10098,
    4180,
    10777,
    270,
    6365,
    9801,
    1046,
    7140,
    1004,
    9198,
    8465,
    8548,
    108,
    2100,
    30015,
    8153,
    1020,
    50100,
    8391,
    34899,
    7090,
    6100,
    8777,
    8298,
    8281,
    7023,
    3377,
    8499,
    7501,
    4321,
    3437,
    9977,
    14338,
    843,
    7901,
    6020,
    6011,
    1988,
    4023,
    20202,
    20200,
    7995,
    18181,
    9836,
    586,
    2340,
    8110,
    9192,
    2525,
    6887,
    4005,
    8992,
    11212,
    2168,
    21,
    60080,
    6664,
    10005,
    956,
    1016,
    4453,
    8974,
    10101,
    58124,
    30025,
    7789,
    7280,
    8222,
    8068,
    11180,
    1984,
    873,
    5566,
    11211,
    1433,
    916,
    8828,
    17071,
    15080,
    8820,
    104,
    21900,
    5151,
    860,
    6286,
    5118,
    18765,
    7055,
    9989,
    807,
    7751,
    8684,
    1999,
    9333,
    55352,
    8681,
    19994,
    3033,
    8017,
    7093,
    7896,
    4242,
    58083,
    56688,
    6167,
    9922,
    3618,
    7082,
    1603,
    16929,
    198,
    8075,
    7044,
    8101,
    8232,
    12315,
    4570,
    4569,
    31082,
    8861,
    3680,
    3001,
    4455,
    8403,
    4497,
    4380,
    7273,
    8896,
    21188,
    22480,
    1445,
    20165,
    20142,
    9068,
    1083,
    59093,
    41474,
    9224,
    9718,
    23380,
    5225,
    18889,
    4237,
    30,
    14549,
    8052,
    911,
    19000,
    7799,
    7300,
    9168,
    29798,
    4480,
    22228,
    7903,
    810,
    68,
    31000,
    9103,
    20992,
    8049,
    2261,
    8105,
    10152,
    5780,
    10111,
    3003
  ],
  "sslPorts": [
    443,
    8443,
    7443,
    8043
  ],
  "unWebPorts": {
    "22": "ssh",
    "23": "telnet",
    "21": "ftp",
    "2100": "ftp",
    "25": "stmp",
    "445": "smb",
    "135": "rpc",
    "139": "NetBIOS",
    "389": "ldap",
    "636": "ldap",
    "2049":"nfs",
    "1433": "sqlserver",
    "1434": "sqlserver",
    "1521": "oracle",
    "3306": "mysql",
    "3389": "rdp",
    "5432": "postgresql",
    "50000": "db2"
  }
}
`

//读取配置文件
func LoadConfig() {
	if misc.FileIsExist("config.json") {
		slog.Warning("成功读取配置文件config.json")
		file, _ := os.Open("config.json")
		decoder := json.NewDecoder(file)
		conf := config_{}
		err := decoder.Decode(&conf)
		if err != nil {
			slog.Debug(err.Error())
		}
		file.Close()
		Config = conf
		//return conf
	} else {
		slog.Debug("未检测到存在config.json配置文件，正在创建...")
		file, _ := os.Create("config.json")
		slog.Info("成功创建配置文件：config.json")
		_, _ = file.WriteString(defaultConfigChar)
		file.Close()
		LoadConfig()
		//return loadConfig()
	}
}

//读取参数
