package model

import "math/big"

type Direction string

const (
	DirectionUp    Direction = "up"
	DirectionDown  Direction = "down"
	DirectionLeft  Direction = "left"
	DirectionRight Direction = "right"
)

type level int

type Person struct {
	Hp             *big.Int
	Attack         *big.Int
	Defence        *big.Int
	Shield         *big.Int
	Exp            int
	ExpToNextLevel int
	Level          *Level
	YellowKeys     int
	BlueKeys       int
	RedKeys        int
	GreenKeys      int
	Axe            int
	Earthquake     int
	Fly            int
	Location       [2]int
	Facing         Direction
}

type Level struct {
	Name string
	Exp  int
}

func (p *Person) UseEarthquake() {

}

func (p *Person) UseFly() {

}

var (
	L1 = &Level{"陀神之境", 0}
	L2 = &Level{"鸿沟神境", 150}
)

/*
[
  {"need": "0", "title": "陀神之境", "clear": true, "action": [

  ]},
  {"need": "150", "title": "鸿沟神境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加10",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "10"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "10"},
  ]},
  {"need": "150", "title": "道天神境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加20",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "20"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "20"},
  ]},
  {"need": "150", "title": "造化神境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加30",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "30"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "30"},
  ]},
  {"need": "150", "title": "人级真神境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加50",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "50"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "50"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "50"},
  ]},
  {"need": "150", "title": "地级真神境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加100",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "100"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "100"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "100"},
  ]},
  {"need": "150", "title": "天级真神境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加200",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "200"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "200"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "200"},
  ]},
  {"need": "150", "title": "王级真神境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加2000",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "2000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "2000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "4000"},
  ]},
  {"need": "150", "title": "通天真神境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加4000",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "4000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "4000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "4000"},
  ]},
  {"need": "150", "title": "玄天真神境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加8000",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "8000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "8000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "8000"},
  ]},
  {"need": "150", "title": "帝级真神境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加15000",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "15000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "15000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "15000"},
  ]},
  {"need": "150", "title": "混沌真神境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加30000",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "30000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "30000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "30000"},
  ]},
  {"need": "150", "title": "人级真主境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加50000，生命增加500万",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "50000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "50000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "50000"},
    {"type": "setValue", "name": "status:hp", "operator": "+=", "value": "5000000"},
  ]},
  {"need": "150", "title": "地级真主境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加100000",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "100000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "100000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "100000"},
  ]},
  {"need": "150", "title": "天级真主境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加150000",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "150000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "150000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "150000"},
  ]},
  {"need": "150", "title": "玄天真主境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加200000",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "200000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "200000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "200000"},
  ]},
  {"need": "150", "title": "通天真主境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加300000",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "300000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "300000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "300000"},
  ]},
  {"need": "150", "title": "造化真主境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加400000",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "400000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "400000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "400000"},
  ]},
  {"need": "150", "title": "帝级真主境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加500000",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "500000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "500000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "500000"},
  ]},
  {"need": "150", "title": "圣级真主境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加600000",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "600000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "600000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "600000"},
  ]},
  {"need": "150", "title": "混沌真主境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加3百万",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "3000000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "3000000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "3000000"},
  ]},
  {"need": "200", "title": "红尘法则境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加5百万，生命增加10亿",
    {"type": "setValue", "name": "status:hp", "operator": "+=", "value": "1000000000"},
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "5000000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "5000000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "5000000"},
  ]},
  {"need": "300", "title": "脱法则境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加 2000万",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "20000000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "20000000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "20000000"},
  ]},
  {"need": "300", "title": "踏天法则境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加4000万",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "40000000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "40000000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "40000000"},
  ]},
  {"need": "300", "title": "三千法则境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加 2亿",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "200000000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "200000000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "200000000"},
  ]},
  {"need": "300", "title": "归一法则境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加 4亿",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "400000000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "400000000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "400000000"},
  ]},
  {"need": "300", "title": "神级法则境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加 8 亿",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "800000000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "800000000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "800000000"},
  ]},
  {"need": "300", "title": "下位纪元境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加 50 亿，生命增加10万亿",
    {"type": "setValue", "name": "status:hp", "operator": "+=", "value": "10000000000000"},
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "5000000000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "5000000000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "5000000000"},
  ]},
  {"need": "300", "title": "上位纪元境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加 100 亿",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "10000000000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "10000000000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "10000000000"},
  ]},
  {"need": "300", "title": "天位纪元境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加 500 亿",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "50000000000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "50000000000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "50000000000"},
  ]},
  {"need": "300", "title": "仙位纪元境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加 1000 亿",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "100000000000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "100000000000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "100000000000"},
  ]},
  {"need": "300", "title": "圣位纪元境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加 3000 亿",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "300000000000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "300000000000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "300000000000"},
  ]},
  {"need": "300", "title": "祖位纪元境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加 7000 亿",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "700000000000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "700000000000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "700000000000"},
  ]},
  {"need": "300", "title": "神位纪元境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加 1.2 兆",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "1200000000000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "1200000000000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "1200000000000"},
  ]},
  {"need": "300", "title": "超纪元境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加 2 兆",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "2000000000000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "2000000000000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "2000000000000"},
  ]},
  {"need": "300", "title": "不朽纪元境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加 3 兆",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "3000000000000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "3000000000000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "3000000000000"},
  ]},
  {"need": "300", "title": "下位转生境", "clear": true, "action": [
    {"type": "animate", "name": "yongchang"},
    "恭喜升级，全属性增加 4.5 兆",
    {"type": "setValue", "name": "status:atk", "operator": "+=", "value": "4500000000000"},
    {"type": "setValue", "name": "status:def", "operator": "+=", "value": "4500000000000"},
    {"type": "setValue", "name": "status:mdef", "operator": "+=", "value": "4500000000000"},
  ]},
]

*/
