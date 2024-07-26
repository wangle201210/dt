package json

import (
	"testing"
)

func TestT2J(t *testing.T) {
	data, err := TomlToJsonFile(`[panel]
iconDef = [
    # 图标编号      对应中N个图标的赔率
    { icon = 1, payout = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0] },  # scatter 棒棒糖
    { icon = 2, payout = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0] },  # 炸弹
    { icon = 3, payout = [0, 0, 0, 0, 0, 0, 0, 1000, 1000, 2500, 2500, 5000] },  # 红心
    { icon = 4, payout = [0, 0, 0, 0, 0, 0, 0, 250, 250, 1000, 1000, 2500] },  # 紫方块
    { icon = 5, payout = [0, 0, 0, 0, 0, 0, 0, 200, 200, 500, 500, 1500] },  # 绿五边形
    { icon = 6, payout = [0, 0, 0, 0, 0, 0, 0, 150, 150, 200, 200, 1200] },  # 蓝长方形
    { icon = 7, payout = [0, 0, 0, 0, 0, 0, 0, 100, 100, 150, 150, 1000] },  # 红苹果
    { icon = 8, payout = [0, 0, 0, 0, 0, 0, 0, 80, 80, 120, 120, 800] },  # 紫桃子
    { icon = 9, payout = [0, 0, 0, 0, 0, 0, 0, 50, 50, 100, 100, 500] },  # 绿西瓜
    { icon = 10, payout = [0, 0, 0, 0, 0, 0, 0, 40, 40, 90, 90, 400] },  # 蓝葡萄
    { icon = 11, payout = [0, 0, 0, 0, 0, 0, 0, 25, 25, 75, 75, 200] },  # 黄香蕉
]`)
	if err != nil {
		t.Errorf("json err:%s", err.Error())
		return
	}
	t.Log(data)
	///Users/apple/work/nova/game-play/sweet/configs
}

func TestT2JF(t *testing.T) {
	data, err := TomlFileToJsonFile("/Users/apple/work/nova/game-play/sweet/configs/play_common_demo.toml")
	if err != nil {
		t.Errorf("json err:%s", err.Error())
		return
	}
	t.Log(data)
}

func TestJ2S(t *testing.T) {
	// 定义一个 JSON 字符串
	jsonString := `{
  "glossary": {
    "title": "example glossary",
    "GlossDiv": {
      "title": "S",
      "GlossList": {
        "GlossEntry": {
          "ID": 1,
          "SortAs": "SGML",
          "GlossTerm": "Standard Generalized Markup Language",
          "Acronym": "SGML",
          "Abbrev": "ISO 8879:1986",
          "GlossDef": {
            "para": "A meta-markup language, used to create markup languages such as DocBook.",
            "GlossSeeAlso": [
              "10",
              1
            ]
          },
          "GlossSee": "markup"
        }
      }
    }
  }
}
`
	// 生成 Golang struct 定义
	// 生成 Golang struct 定义
	res, err := ProcessByFile(jsonString)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}

func TestJ2S2(t *testing.T) {
	// 定义一个 JSON 字符串
	jsonString := `{
		"name": "Alice",
		"age": 30,
		"email": "alice@example.com",
		"address": {
			"city": "Wonderland",
			"postcode": 12345
		},
		"friends": [
			{
				"name": "Bob",
				"age": 25
			},
			{
				"name": "Charlie",
				"age": 35
			}
		]
	}
`
	// 生成 Golang struct 定义
	res, err := ProcessByFile(jsonString)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}
