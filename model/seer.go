package model

const (
	SeerElementRestraintResultNormal      float64 = 1
	SeerElementRestraintResultRestraint   float64 = 2
	SeerElementRestraintResultRestrainted         = 1 / SeerElementRestraintResultRestraint
	SeerElementRestraintResultZero        float64 = 0
	SeerElementRestraintResultMax         float64 = 4
)

const (
	SeerElementTypeGrass = iota + 1
	SeerElementTypeWater
	SeerElementTypeFire
	SeerElementTypeFly
	SeerElementTypeElectricity
	SeerElementTypeMachine
	SeerElementTypeGround
	SeerElementTypeOrdinary
	SeerElementTypeIce
	SeerElementTypeSuper
	SeerElementTypeWar
	SeerElementTypeLight
	SeerElementTypeDark
	SeerElementTypeMystery
	SeerElementTypeDragon
	SeerElementTypeHoly
	SeerElementTypeDimension
	SeerElementTypeAncient
	SeerElementTypeEvil
	SeerElementTypeNature
	SeerElementTypeKing
	SeerElementTypeChaos
	SeerElementTypeGod
	SeerElementTypeReincarnation
	SeerElementTypeInsect
	SeerElementTypeVoid
)

var SeerElementTypeRestraint = map[int][]int{
	SeerElementTypeGrass:         {SeerElementTypeWater, SeerElementTypeGround, SeerElementTypeLight},
	SeerElementTypeWater:         {SeerElementTypeFire, SeerElementTypeGround},
	SeerElementTypeFire:          {SeerElementTypeGrass, SeerElementTypeMachine, SeerElementTypeIce},
	SeerElementTypeFly:           {SeerElementTypeGrass, SeerElementTypeWar, SeerElementTypeInsect},
	SeerElementTypeElectricity:   {SeerElementTypeWater, SeerElementTypeFly, SeerElementTypeDark, SeerElementTypeDimension, SeerElementTypeChaos, SeerElementTypeVoid},
	SeerElementTypeMachine:       {SeerElementTypeIce, SeerElementTypeWar, SeerElementTypeAncient, SeerElementTypeEvil, SeerElementTypeGod},
	SeerElementTypeGround:        {SeerElementTypeFire, SeerElementTypeElectricity, SeerElementTypeMachine, SeerElementTypeKing, SeerElementTypeReincarnation},
	SeerElementTypeIce:           {SeerElementTypeGrass, SeerElementTypeFly, SeerElementTypeGround, SeerElementTypeAncient, SeerElementTypeDimension, SeerElementTypeReincarnation, SeerElementTypeInsect},
	SeerElementTypeSuper:         {SeerElementTypeWar, SeerElementTypeMystery, SeerElementTypeNature},
	SeerElementTypeWar:           {SeerElementTypeMachine, SeerElementTypeIce, SeerElementTypeDragon, SeerElementTypeHoly},
	SeerElementTypeLight:         {SeerElementTypeSuper, SeerElementTypeDark, SeerElementTypeInsect},
	SeerElementTypeDark:          {SeerElementTypeSuper, SeerElementTypeDark, SeerElementTypeDimension},
	SeerElementTypeMystery:       {SeerElementTypeElectricity, SeerElementTypeMystery, SeerElementTypeHoly, SeerElementTypeNature, SeerElementTypeKing, SeerElementTypeGod, SeerElementTypeReincarnation},
	SeerElementTypeDragon:        {SeerElementTypeIce, SeerElementTypeDragon, SeerElementTypeHoly, SeerElementTypeEvil},
	SeerElementTypeHoly:          {SeerElementTypeGrass, SeerElementTypeFire, SeerElementTypeWater, SeerElementTypeElectricity, SeerElementTypeIce, SeerElementTypeAncient, SeerElementTypeVoid},
	SeerElementTypeDimension:     {SeerElementTypeFly, SeerElementTypeMachine, SeerElementTypeSuper, SeerElementTypeEvil, SeerElementTypeNature, SeerElementTypeInsect, SeerElementTypeVoid},
	SeerElementTypeAncient:       {SeerElementTypeGrass, SeerElementTypeFly, SeerElementTypeMystery, SeerElementTypeDragon, SeerElementTypeVoid},
	SeerElementTypeEvil:          {SeerElementTypeLight, SeerElementTypeDark, SeerElementTypeMystery, SeerElementTypeDimension, SeerElementTypeNature},
	SeerElementTypeNature:        {SeerElementTypeGrass, SeerElementTypeWater, SeerElementTypeFire, SeerElementTypeFly, SeerElementTypeElectricity, SeerElementTypeGround, SeerElementTypeLight, SeerElementTypeKing, SeerElementTypeReincarnation},
	SeerElementTypeKing:          {SeerElementTypeWar, SeerElementTypeDark, SeerElementTypeDimension, SeerElementTypeEvil},
	SeerElementTypeChaos:         {SeerElementTypeFly, SeerElementTypeIce, SeerElementTypeMystery, SeerElementTypeDimension, SeerElementTypeEvil, SeerElementTypeNature, SeerElementTypeGod},
	SeerElementTypeGod:           {SeerElementTypeWater, SeerElementTypeFire, SeerElementTypeGrass, SeerElementTypeElectricity, SeerElementTypeIce, SeerElementTypeAncient, SeerElementTypeEvil, SeerElementTypeChaos},
	SeerElementTypeReincarnation: {SeerElementTypeLight, SeerElementTypeDark, SeerElementTypeHoly, SeerElementTypeDimension, SeerElementTypeEvil, SeerElementTypeChaos},
	SeerElementTypeInsect:        {SeerElementTypeGrass, SeerElementTypeGround, SeerElementTypeWar, SeerElementTypeChaos, SeerElementTypeInsect},
	SeerElementTypeVoid:          {SeerElementTypeSuper, SeerElementTypeWar, SeerElementTypeLight, SeerElementTypeMystery, SeerElementTypeNature, SeerElementTypeReincarnation},
}

var SeerElementTypeRestrainted = map[int][]int{
	SeerElementTypeGrass:         {SeerElementTypeFire, SeerElementTypeFly, SeerElementTypeMachine, SeerElementTypeGod, SeerElementTypeHoly, SeerElementTypeGrass, SeerElementTypeChaos, SeerElementTypeAncient},
	SeerElementTypeWater:         {SeerElementTypeGrass, SeerElementTypeWater, SeerElementTypeHoly, SeerElementTypeNature, SeerElementTypeChaos, SeerElementTypeGod},
	SeerElementTypeFire:          {SeerElementTypeWater, SeerElementTypeFire, SeerElementTypeHoly, SeerElementTypeNature, SeerElementTypeChaos, SeerElementTypeGod},
	SeerElementTypeFly:           {SeerElementTypeElectricity, SeerElementTypeMachine, SeerElementTypeDimension, SeerElementTypeEvil, SeerElementTypeNature, SeerElementTypeChaos},
	SeerElementTypeElectricity:   {SeerElementTypeGrass, SeerElementTypeElectricity, SeerElementTypeMystery, SeerElementTypeHoly, SeerElementTypeNature, SeerElementTypeGod},
	SeerElementTypeMachine:       {SeerElementTypeWater, SeerElementTypeFire, SeerElementTypeElectricity, SeerElementTypeMachine, SeerElementTypeDimension},
	SeerElementTypeGround:        {SeerElementTypeGrass, SeerElementTypeSuper, SeerElementTypeDark, SeerElementTypeDragon, SeerElementTypeHoly, SeerElementTypeNature, SeerElementTypeGod, SeerElementTypeInsect},
	SeerElementTypeIce:           {SeerElementTypeWater, SeerElementTypeFire, SeerElementTypeMachine, SeerElementTypeIce, SeerElementTypeGod, SeerElementTypeChaos, SeerElementTypeHoly},
	SeerElementTypeSuper:         {SeerElementTypeMachine, SeerElementTypeSuper, SeerElementTypeInsect},
	SeerElementTypeWar:           {SeerElementTypeSuper, SeerElementTypeWar, SeerElementTypeDark, SeerElementTypeEvil, SeerElementTypeKing},
	SeerElementTypeLight:         {SeerElementTypeMachine, SeerElementTypeIce, SeerElementTypeLight, SeerElementTypeHoly, SeerElementTypeEvil, SeerElementTypeNature, SeerElementTypeGod, SeerElementTypeReincarnation, SeerElementTypeVoid},
	SeerElementTypeDark:          {SeerElementTypeMachine, SeerElementTypeIce, SeerElementTypeLight, SeerElementTypeHoly, SeerElementTypeEvil, SeerElementTypeGod},
	SeerElementTypeMystery:       {SeerElementTypeGround, SeerElementTypeWar, SeerElementTypeEvil, SeerElementTypeChaos, SeerElementTypeInsect},
	SeerElementTypeDragon:        {SeerElementTypeGrass, SeerElementTypeWater, SeerElementTypeFire, SeerElementTypeElectricity, SeerElementTypeAncient, SeerElementTypeInsect},
	SeerElementTypeHoly:          {SeerElementTypeWar, SeerElementTypeMystery, SeerElementTypeDragon, SeerElementTypeReincarnation},
	SeerElementTypeDimension:     {SeerElementTypeIce, SeerElementTypeKing, SeerElementTypeChaos, SeerElementTypeGod, SeerElementTypeReincarnation},
	SeerElementTypeAncient:       {SeerElementTypeMachine, SeerElementTypeIce, SeerElementTypeKing, SeerElementTypeReincarnation},
	SeerElementTypeEvil:          {SeerElementTypeMachine, SeerElementTypeIce, SeerElementTypeSuper, SeerElementTypeHoly, SeerElementTypeKing, SeerElementTypeChaos, SeerElementTypeReincarnation},
	SeerElementTypeNature:        {SeerElementTypeMachine, SeerElementTypeSuper, SeerElementTypeWar, SeerElementTypeDark, SeerElementTypeMystery, SeerElementTypeDimension, SeerElementTypeEvil, SeerElementTypeChaos, SeerElementTypeVoid},
	SeerElementTypeKing:          {SeerElementTypeSuper, SeerElementTypeNature, SeerElementTypeInsect},
	SeerElementTypeChaos:         {SeerElementTypeElectricity, SeerElementTypeMachine, SeerElementTypeWar, SeerElementTypeReincarnation},
	SeerElementTypeGod:           {SeerElementTypeMachine, SeerElementTypeWar, SeerElementTypeDragon},
	SeerElementTypeReincarnation: {SeerElementTypeIce, SeerElementTypeSuper, SeerElementTypeNature, SeerElementTypeVoid},
	SeerElementTypeInsect:        {SeerElementTypeWater, SeerElementTypeFire, SeerElementTypeIce, SeerElementTypeLight},
	SeerElementTypeVoid:          {SeerElementTypeFly, SeerElementTypeDark, SeerElementTypeHoly, SeerElementTypeDimension},
}

var SeerElementTypeZero = map[int]int{
	SeerElementTypeLight:       SeerElementTypeGrass,
	SeerElementTypeEvil:        SeerElementTypeGod,
	SeerElementTypeElectricity: SeerElementTypeGround,
	SeerElementTypeGround:      SeerElementTypeFly,
	SeerElementTypeDimension:   SeerElementTypeDark,
	SeerElementTypeSuper:       SeerElementTypeLight,
	SeerElementTypeChaos:       SeerElementTypeVoid,
}
