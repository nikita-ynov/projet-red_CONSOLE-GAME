package structure

type Inventory struct {
	Name      string
	ChangeHp  int
	Quantity  int
	UniqueObj int
}

type Skills struct {
	Name   string
	Damage int
}
type Character struct {
	Name            string
	Class           string
	Lvl             int
	HpMax           int
	CurrentHp       int
	Money           int
	Damage          int
	InventoryLimit  int
	Inventory       []Inventory
	Skills          []Skills
	PlayerEquipment []Equipment
}

type MerchantItems struct {
	Name      string
	ChangeHp  int
	Quantity  int
	UniqueObj int
	Price     int
}

type Monster struct {
	Name      string
	HpMax     int
	CurrentHp int
	Damage    int
}
type Equipment struct {
	Helmet      int
	Breastplate int
	Legwarmer   int
}

type BlacksmithItems struct {
	Name            string
	Price           int
	CrowFeather     int
	WildBoarLeather int
	TrollSkin       int
	WolfFur         int
}
