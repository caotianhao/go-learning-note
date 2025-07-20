package model

type SeerElemKezhiRes float64

const (
	SeerElemKezhiResNormal SeerElemKezhiRes = 1
	SeerElemKezhiResKe     SeerElemKezhiRes = 2
	SeerElemKezhiResKed    SeerElemKezhiRes = 0.5
	SeerElemKezhiResZero   SeerElemKezhiRes = 0
	SeerElemKezhiResMax    SeerElemKezhiRes = 4
)

type SeerElem uint16

const (
	SeerElemGrass SeerElem = iota + 1
	SeerElemWater
	SeerElemFire
	SeerElemFly
	SeerElemElectricity
	SeerElemMachine
	SeerElemGround
	SeerElemOrdinary
	SeerElemIce
	SeerElemSuper
	SeerElemWar
	SeerElemLight
	SeerElemDark
	SeerElemMystery
	SeerElemDragon
	SeerElemHoly
	SeerElemDimension
	SeerElemAncient
	SeerElemEvil
	SeerElemNature
	SeerElemKing
	SeerElemChaos
	SeerElemGod
	SeerElemReincarnation
	SeerElemInsect
	SeerElemVoid
)

var SeerElemKe = map[SeerElem][]SeerElem{
	SeerElemGrass:         {SeerElemWater, SeerElemGround, SeerElemLight},
	SeerElemWater:         {SeerElemFire, SeerElemGround},
	SeerElemFire:          {SeerElemGrass, SeerElemMachine, SeerElemIce},
	SeerElemFly:           {SeerElemGrass, SeerElemWar, SeerElemInsect},
	SeerElemElectricity:   {SeerElemWater, SeerElemFly, SeerElemDark, SeerElemDimension, SeerElemChaos, SeerElemVoid},
	SeerElemMachine:       {SeerElemIce, SeerElemWar, SeerElemAncient, SeerElemEvil, SeerElemGod},
	SeerElemGround:        {SeerElemFire, SeerElemElectricity, SeerElemMachine, SeerElemKing, SeerElemReincarnation},
	SeerElemIce:           {SeerElemGrass, SeerElemFly, SeerElemGround, SeerElemAncient, SeerElemDimension, SeerElemReincarnation, SeerElemInsect},
	SeerElemSuper:         {SeerElemWar, SeerElemMystery, SeerElemNature},
	SeerElemWar:           {SeerElemMachine, SeerElemIce, SeerElemDragon, SeerElemHoly},
	SeerElemLight:         {SeerElemSuper, SeerElemDark, SeerElemInsect},
	SeerElemDark:          {SeerElemSuper, SeerElemDark, SeerElemDimension},
	SeerElemMystery:       {SeerElemElectricity, SeerElemMystery, SeerElemHoly, SeerElemNature, SeerElemKing, SeerElemGod, SeerElemReincarnation},
	SeerElemDragon:        {SeerElemIce, SeerElemDragon, SeerElemHoly, SeerElemEvil},
	SeerElemHoly:          {SeerElemGrass, SeerElemFire, SeerElemWater, SeerElemElectricity, SeerElemIce, SeerElemAncient, SeerElemVoid},
	SeerElemDimension:     {SeerElemFly, SeerElemMachine, SeerElemSuper, SeerElemEvil, SeerElemNature, SeerElemInsect, SeerElemVoid},
	SeerElemAncient:       {SeerElemGrass, SeerElemFly, SeerElemMystery, SeerElemDragon, SeerElemVoid},
	SeerElemEvil:          {SeerElemLight, SeerElemDark, SeerElemMystery, SeerElemDimension, SeerElemNature},
	SeerElemNature:        {SeerElemGrass, SeerElemWater, SeerElemFire, SeerElemFly, SeerElemElectricity, SeerElemGround, SeerElemLight, SeerElemKing, SeerElemReincarnation},
	SeerElemKing:          {SeerElemWar, SeerElemDark, SeerElemDimension, SeerElemEvil},
	SeerElemChaos:         {SeerElemFly, SeerElemIce, SeerElemMystery, SeerElemDimension, SeerElemEvil, SeerElemNature, SeerElemGod},
	SeerElemGod:           {SeerElemWater, SeerElemFire, SeerElemGrass, SeerElemElectricity, SeerElemIce, SeerElemAncient, SeerElemEvil, SeerElemChaos},
	SeerElemReincarnation: {SeerElemLight, SeerElemDark, SeerElemHoly, SeerElemDimension, SeerElemEvil, SeerElemChaos},
	SeerElemInsect:        {SeerElemGrass, SeerElemGround, SeerElemWar, SeerElemChaos, SeerElemInsect},
	SeerElemVoid:          {SeerElemSuper, SeerElemWar, SeerElemLight, SeerElemMystery, SeerElemNature, SeerElemReincarnation},
}

var SeerElemKed = map[SeerElem][]SeerElem{
	SeerElemGrass:         {SeerElemFire, SeerElemFly, SeerElemMachine, SeerElemGod, SeerElemHoly, SeerElemGrass, SeerElemChaos, SeerElemAncient},
	SeerElemWater:         {SeerElemGrass, SeerElemWater, SeerElemHoly, SeerElemNature, SeerElemChaos, SeerElemGod},
	SeerElemFire:          {SeerElemWater, SeerElemFire, SeerElemHoly, SeerElemNature, SeerElemChaos, SeerElemGod},
	SeerElemFly:           {SeerElemElectricity, SeerElemMachine, SeerElemDimension, SeerElemEvil, SeerElemNature, SeerElemChaos},
	SeerElemElectricity:   {SeerElemGrass, SeerElemElectricity, SeerElemMystery, SeerElemHoly, SeerElemNature, SeerElemGod},
	SeerElemMachine:       {SeerElemWater, SeerElemFire, SeerElemElectricity, SeerElemMachine, SeerElemDimension},
	SeerElemGround:        {SeerElemGrass, SeerElemSuper, SeerElemDark, SeerElemDragon, SeerElemHoly, SeerElemNature, SeerElemGod, SeerElemInsect},
	SeerElemIce:           {SeerElemWater, SeerElemFire, SeerElemMachine, SeerElemIce, SeerElemGod, SeerElemChaos, SeerElemHoly},
	SeerElemSuper:         {SeerElemMachine, SeerElemSuper, SeerElemInsect},
	SeerElemWar:           {SeerElemSuper, SeerElemWar, SeerElemDark, SeerElemEvil, SeerElemKing},
	SeerElemLight:         {SeerElemMachine, SeerElemIce, SeerElemLight, SeerElemHoly, SeerElemEvil, SeerElemNature, SeerElemGod, SeerElemReincarnation, SeerElemVoid},
	SeerElemDark:          {SeerElemMachine, SeerElemIce, SeerElemLight, SeerElemHoly, SeerElemEvil, SeerElemGod},
	SeerElemMystery:       {SeerElemGround, SeerElemWar, SeerElemEvil, SeerElemChaos, SeerElemInsect},
	SeerElemDragon:        {SeerElemGrass, SeerElemWater, SeerElemFire, SeerElemElectricity, SeerElemAncient, SeerElemInsect},
	SeerElemHoly:          {SeerElemWar, SeerElemMystery, SeerElemDragon, SeerElemReincarnation},
	SeerElemDimension:     {SeerElemIce, SeerElemKing, SeerElemChaos, SeerElemGod, SeerElemReincarnation},
	SeerElemAncient:       {SeerElemMachine, SeerElemIce, SeerElemKing, SeerElemReincarnation},
	SeerElemEvil:          {SeerElemMachine, SeerElemIce, SeerElemSuper, SeerElemHoly, SeerElemKing, SeerElemChaos, SeerElemReincarnation},
	SeerElemNature:        {SeerElemMachine, SeerElemSuper, SeerElemWar, SeerElemDark, SeerElemMystery, SeerElemDimension, SeerElemEvil, SeerElemChaos, SeerElemVoid},
	SeerElemKing:          {SeerElemSuper, SeerElemNature, SeerElemInsect},
	SeerElemChaos:         {SeerElemElectricity, SeerElemMachine, SeerElemWar, SeerElemReincarnation},
	SeerElemGod:           {SeerElemMachine, SeerElemWar, SeerElemDragon},
	SeerElemReincarnation: {SeerElemIce, SeerElemSuper, SeerElemNature, SeerElemVoid},
	SeerElemInsect:        {SeerElemWater, SeerElemFire, SeerElemIce, SeerElemLight},
	SeerElemVoid:          {SeerElemFly, SeerElemDark, SeerElemHoly, SeerElemDimension},
}

var SeerElemInvalid = map[SeerElem]SeerElem{
	SeerElemLight:       SeerElemGrass,
	SeerElemEvil:        SeerElemGod,
	SeerElemElectricity: SeerElemGround,
	SeerElemGround:      SeerElemFly,
	SeerElemDimension:   SeerElemDark,
	SeerElemSuper:       SeerElemLight,
	SeerElemChaos:       SeerElemVoid,
}
