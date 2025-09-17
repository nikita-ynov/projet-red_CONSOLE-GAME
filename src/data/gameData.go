package data

type CoefWeapon struct {
	Name           string
	Coef           float32
	DropRate       int
	CumulativeDrop int
}

// permet de creer des armes et influer leurs stats
var WeaponsData = []CoefWeapon{
	{Name: "Excalibur", Coef: 1.0, DropRate: 50},
	{Name: "Excaliburne", Coef: 1.5, DropRate: 80},
	{Name: "Excalibatard", Coef: 1.7, DropRate: 100},
}
