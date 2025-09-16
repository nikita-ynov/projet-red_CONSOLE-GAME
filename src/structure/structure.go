package structure

type Inventory struct {
	Name        string
	ChangeHp    int
	ChangeManna int
	Quantity    int
	UniqueObj   int
}

type Skills struct {
	Name   string
	Damage int
}
type Character struct {
	Name            string
	Class           string
	Lvl             int
	Exp             int
	HpMax           int
	CurrentHp       int
	MannaMax        int
	CurrentManna    int
	Money           int
	Damage          int
	Initiative      int
	InventoryLimit  int
	Inventory       []Inventory
	Skills          []Skills
	PlayerEquipment []Equipment
}

type MerchantItems struct {
	Name        string
	ChangeHp    int
	ChangeManna int
	Quantity    int
	UniqueObj   int
	Price       int
}

type Monster struct {
	Name       string
	HpMax      int
	CurrentHp  int
	Damage     int
	Initiative int
}

type Equipment struct {
	Helmet      []Helmet
	Breastplate []Breastplate
	Legwarmer   []Legwarmer
}

type Helmet struct {
	Name       string
	Protection int
}

type Breastplate struct {
	Name       string
	Protection int
}

type Legwarmer struct {
	Name       string
	Protection int
}

type BlacksmithItems struct {
	Name            string
	Price           int
	CrowFeather     int
	WildBoarLeather int
	TrollSkin       int
	WolfFur         int
}
