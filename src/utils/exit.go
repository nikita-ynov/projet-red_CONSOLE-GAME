package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Exit() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press Enter to exit...")

	// On nettoie le buffer (cas où un "\n" est déjà présent)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Si jamais ça lit directement une ligne vide, on relance une vraie attente
	if input == "" {
		reader.ReadString('\n')
	}
}
