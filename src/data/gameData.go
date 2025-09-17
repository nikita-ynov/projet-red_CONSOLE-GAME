package data

type CoefWeapon struct {
	Name string
	Coef float32
}

var WeaponsData = []CoefWeapon{
	{Name: "Excalibur", Coef: 1.0},
	{Name: "Excaliburne", Coef: 1.5},
	{Name: "Excalibatard", Coef: 1.7},
}
