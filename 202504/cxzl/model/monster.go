package model

type Monster struct {
	Name    string
	Avatar  string
	Hp      int
	Attack  int
	Defense int
	Feature Feature
	Exp     int
}

type Feature int

const (
	MonsterFeatureAtkFast Feature = 1 << iota // 先攻
	MonsterFeatureJianGu                      // 坚固
	MonsterFeatureManaAtk                     // 魔攻
	MonsterFeature2Atk                        // 2 连击
	MonsterFeature3Atk                        // 3 连击
)
