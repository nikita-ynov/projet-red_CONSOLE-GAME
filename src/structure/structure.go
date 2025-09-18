package structure

type Achievement struct {
	Name        string
	Description string
	Unlocked    bool
}

type Inventory struct {
	Name       string
	ChangeHp   int
	ChangeMana int
	Protection int
	Quantity   int
	IsWeapon   bool
}

type Skills struct {
	Name   string
	Damage int
}

type Character struct {
	Name                      string
	Class                     string
	Lvl                       int
	Exp                       int
	HpMax                     int
	CurrentHp                 int
	ManaMax                   int
	CurrentMana               int
	Money                     int
	Damage                    int
	Initiative                int
	InventoryLimit            int
	Inventory                 []Inventory
	Skills                    []Skills
	Equipment                 Equipment // maintenant c’est bien défini
	Achievements              []Achievement
	AchievementsMenuVisible   bool
	GoblinsKilledWithoutDying int
	Deaths                    int
}

type Weapon struct {
	Name        string
	Damage      int
	LvlRequired int
}

type MerchantItems struct {
	Name       string
	ChangeHp   int
	ChangeMana int
	Quantity   int
	UniqueObj  int
	Price      int
	IsSpell    bool
}

type Monster struct {
	Name       string
	HpMax      int
	CurrentHp  int
	Damage     int
	Initiative int
}

type Listequipment struct {
	Name       string
	Protection int
}

type Equipment struct {
	Helmet      Helmet
	Breastplate Breastplate
	Legwarmer   Legwarmer
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
