package main

// Floor 楼层属性
type Floor struct {
	// 文件名和floorId需要保持完全一致
	// 楼层唯一标识符仅能由字母、数字、下划线组成，且不能由数字开头
	// 推荐用法：第20层就用MT20，第38层就用MT38，地下6层就用MT_6（用下划线代替负号），隐藏3层用MT3h（h表示隐藏），等等
	// 楼层唯一标识符，需要和名字完全一致
	// 这里不能更改floorId，请通过另存为来实现
	FloorId            string   `json:"floor_id"`
	Title              string   `json:"title"`                // 楼层中文名，将在切换楼层和浏览地图时显示
	Name               string   `json:"name"`                 // 显示在状态栏中的层数
	Width              int      `json:"width"`                // 地图x方向大小,这里不能更改,仅能在新建地图时设置,默认为13
	Height             int      `json:"height"`               // 地图y方向大小,这里不能更改,仅能在新建地图时设置,默认为13
	CanFlyTo           bool     `json:"can_fly_to"`           // 该楼能否被楼传器飞到（不能的话在该楼也不允许使用楼传器）
	CanUseQuickShop    bool     `json:"can_use_quick_shop"`   // 该层是否允许使用快捷商店
	CannotViewMap      bool     `json:"cannot_view_map"`      // 该层是否不允许被浏览地图看到；如果勾上则浏览地图会跳过该层
	CannotMoveDirectly bool     `json:"cannot_move_directly"` // 该层是否不允许瞬间移动；如果勾上则不可在此层进行瞬移
	FirstArrive        any      `json:"first_arrive"`         // 第一次到该楼层触发的事件，可以双击进入事件编辑器。
	EachArrive         any      `json:"each_arrive"`          // 每次到该楼层触发的事件，可以双击进入事件编辑器；该事件会在firstArrive执行后再执行。
	ParallelDo         any      `json:"parallel_do"`          // 在该层楼时执行的并行事件处理。可以在这里写上任意需要自动执行的脚本，比如打怪自动开门等。详见文档-事件-并行事件处理。
	UpFloor            [2]uint8 `json:"up_floor"`             // 该层上楼点，如[2,3]。如果此项不为null，则楼层转换时的stair:upFloor，以及楼传器的落点会被替换成该点而不是该层的上楼梯。
	DownFloor          [2]uint8 `json:"down_floor"`           // 该层下楼点，如[2,3]。如果此项不为null，则楼层转换时的stair:downFloor，以及楼传器的落点会被替换成该点而不是该层的下楼梯。
	DefaultGround      string   `json:"default_ground"`       // 默认地面的图块ID，此项修改后需要刷新才能看到效果。
	Images             []string `json:"images"`               // 背景/前景图；你可以选择若干张图片来作为背景/前景素材。
	Color              string   `json:"color"`                // 该层的默认画面色调。本项可不写（代表无色调），如果写需要是一个RGBA数组如[255,0,0,0.3]
	Weather            *Weather `json:"weather"`              // 该层的默认天气。本项可忽略表示晴天，如果写则第一项为rain;snow;fog代表雨雪雾，第二项为1-10之间的数代表强度。如[8]代表8级雨天。
	Bgm                string   `json:"bgm"`                  // 到达该层后默认播放的BGM。本项可忽略，或者为一个定义过的背景音乐如&quot;bgm.mp3&quot;。
	ItemRatio          int      `json:"item_ratio"`           // 每一层的宝石/血瓶效果，即获得宝石和血瓶时框内的值。
	UnderGround        bool     `json:"underground"`          // 是否是地下层；如果该项为true则同层传送将传送至上楼梯
}

type Weather struct {
	Type     WeatherType `json:"type"`
	Strength int         `json:"strength"`
}

type WeatherType string

const (
	WeatherNormal WeatherType = "sunny"
	WeatherRain   WeatherType = "rain"
	WeatherSnow   WeatherType = "snow"
	WeatherFog    WeatherType = "fog"
)

type Player struct {
	Name       string `json:"name"`
	Lv         int    `json:"lv"`
	HPMax      int    `json:"hp_max"`
	HP         int    `json:"hp"`
	ManaMax    int    `json:"mana_max"`
	Mana       int    `json:"mana"`
	Atk        int    `json:"atk"`
	Def        int    `json:"def"`
	MDef       int    `json:"m_def"`
	Money      int    `json:"money"`
	Experience int    `json:"experience"`
	Equipment  []any  `json:"equipment"`
	Items      struct {
		Keys struct {
			YellowKey int `json:"yellowKey"`
			BlueKey   int `json:"blueKey"`
			RedKey    int `json:"redKey"`
		} `json:"keys"`
		Constants map[string]any `json:"constants"`
		Tools     map[string]any `json:"tools"`
		Equips    map[string]any `json:"equips"`
	} `json:"items"`
	Loc struct {
		Direction Direction `json:"direction"`
		X         int       `json:"x"`
		Y         int       `json:"y"`
	} `json:"loc"`
	Flags map[string]any `json:"flags"`
}

type Direction string

const (
	DirectionUp    Direction = "up"
	DirectionDown  Direction = "down"
	DirectionLeft  Direction = "left"
	DirectionRight Direction = "right"
)
