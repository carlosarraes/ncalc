package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

type (
	EnergyKJ            float64
	SugarGram           float64
	SaturatedFattyAcids float64
	SodiumMilligram     float64
	FruitsPercent       float64
	FibreGram           float64
	ProteinGram         float64
)

type NutritionalData struct {
	Energy              EnergyKJ
	Sugars              SugarGram
	SaturatedFattyAcids SaturatedFattyAcids
	Sodium              SodiumMilligram
	Fruits              FruitsPercent
	Fibre               FibreGram
	Proteins            ProteinGram
	IsWater             bool
}
