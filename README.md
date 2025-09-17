# 🎮 PROJETRED – Jeu RPG en console (Go)

## 📖 Description
**PROJETRED** est un petit RPG textuel développé en **Go (Golang)**.  
Le joueur incarne un personnage, choisit sa classe, explore des menus, combat des monstres, achète des objets chez le marchand, forge de l’équipement et fait évoluer son héros.

Tout se joue directement dans le terminal avec des choix numériques.

---

## 🚀 Fonctionnalités principales
- 👤 **Création de personnage** : choix du nom et de la classe (Humain, Elfe, Nain).  
- 📊 **Statistiques** : HP, Mana, Argent, Expérience, Niveau.  
- ⚔ **Combat** : affrontements contre des gobelins d’entraînement, avec attaques, compétences et potions.  
- 🧪 **Potions** : possibilité de consommer des potions de vie et de mana.  
- 🔥 **Compétences** : acquisition et utilisation de sorts comme *Fire Ball*.  
- 🛒 **Marchand** : achat de potions, matériaux, amélioration de l’inventaire, nouvelles compétences.  
- ⚒ **Forgeron** : fabrication d’armures à partir de matériaux collectés.  
- 🛡 **Équipement** : gestion des emplacements (casque, plastron, jambières).  
- 📦 **Inventaire** : limité en taille (extensible via le marchand).  
- ☠ **Mort & renaissance** : en cas de défaite, le joueur renaît avec la moitié de ses HP.  
- 📈 **Progression** : gain d’expérience et montée de niveau.  

---

## 🗂 Structure du projet
   ```bash
PROJETRED/
│── main.go # Point d’entrée du jeu
│
├── structure/ # Définitions des structures (personnage, inventaire, monstres…)
│ └── structure.go
│
├── functions/ # Logique principale du jeu
│ ├── menu.go # Menu principal
│ ├── characterCreation.go# Création du personnage
│ ├── initCharacter.go # Initialisation des stats du héros
│ ├── initGoblin.go # Initialisation des monstres
│ ├── traningFight.go # Combat d’entraînement
│ ├── characterTurn.go # Choix des actions en combat
│ ├── takePot.go # Utilisation des potions
│ ├── takeSkill.go # Utilisation des compétences
│ ├── displayInfo.go # Affichage des stats du joueur
│ ├── DisplayEquipment.go # Menu équipement
│ ├── equipment.go # Gestion de l’équipement
│ ├── merchant.go # Marchand (achat d’objets/compétences)
│ └── blacksmith.go # Forgeron (craft d’armures)
│
└── utils/ # Fonctions utilitaires
├── add.go # Ajout (objets, compétences, HP, mana…)
├── remove.go # Retrait (objets, HP, mana…)
├── inventory.go # Gestion de l’inventaire
├── exit.go # Pause / Attente touche entrée
└── isWasted.go # Mort et renaissance
---

## ▶️ Lancer le projet
1. Cloner le dépôt :  
   ```bash
   git clone https://github.com/ton-compte/PROJETRED.git
   cd PROJETRED

2. Lancer le jeu :
   ```bash
    cd src
    go run main.go
