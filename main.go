package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// PREVENTION SI YA UNE ERREUR causé par le ficher "hangman.txt" changer "CTRLF" en "LF" (en bas a droite du vs)
type Hangman struct {
	word       string
	hiddenWord []string
	attempt    int
	position   string
	tried      []string
}

func main() {
	var h1 Hangman
	h1.attempt = 10
	Bienvenue()
	h1.ReadFile()
	h1.creation()
	h1.rendu()
	for i := 0; i < 9999; i++ {
		h1.HangmanPositions()
		h1.joueur()
		h1.gagner()
		h1.perdu()

	}
}

func Bienvenue() {
	fmt.Printf("Bienvenue sur le pendu ynov ;) \n")
	fmt.Printf("les regles sont simples \n")
	fmt.Printf("un mot cacher et toi grand expert tu va devoir le découvrir! \n")
	fmt.Printf("pour cela tu disposera de 10essais \n")
	fmt.Printf("Bonne chance a toi l'ami! \n")
}

func (h *Hangman) ReadFile() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(38)
	data, err := ioutil.ReadFile("words.txt")
	content := string(data)
	words := strings.Split(content, "\n")
	if err != nil {
		fmt.Println(err)

	} else {
		for i := 0; i < n; i++ {
			h.word = words[i]
		}
	}
}

func (h *Hangman) creation() {
	for i := 0; i < len(h.word)+1; i++ {
		h.hiddenWord = append(h.hiddenWord, "_")

	}
}

func (h *Hangman) rendu() {
	var randomLetter int
	n := len(h.word)/2 - 1
	if n > 0 {
		for i := 1; i <= n; i++ {
			randomLetter = rand.Intn(len(h.word))
			for i, letter := range h.word {
				if i == randomLetter && h.hiddenWord[i] == "_" {
					h.hiddenWord[i] = string(letter)
				}
			}
		}
	}
}

//ascii art pendu
func (h *Hangman) HangmanPositions() {
	data, err := ioutil.ReadFile("hangman.txt")
	content := string(data)
	positions := strings.Split(content, "\n\n")
	var n int = 10 - h.attempt
	if err != nil {
		fmt.Println(err)

	} else {
		for i := 0; i < n; i++ {
			h.position = positions[i]
		}
		if n > 0 {
			fmt.Println(positions[n-1])
		}
	}

}

func (h *Hangman) joueur() {
	var s string
	test := false
	found := false
	fmt.Println("Tu doit trouver :", h.hiddenWord[:len(h.hiddenWord)-1])
	if h.attempt > 1 {
		fmt.Println("Il te reste", h.attempt, "essais.")
	} else {
		fmt.Println("Il te reste", h.attempt, "essais.")
	}
	if h.tried != nil {
		fmt.Println("tu as déjà essayé:", h.tried)
	}
	fmt.Println("Entre un mot ou une lettre.")
	fmt.Scanln(&s)
	strings.ToLower(s)
	if len(s) > 1 {
		if s == h.word {
			for i, letter := range h.word {
				h.hiddenWord[i] = string(letter)
			}
			h.gagner()
		} else {
			fmt.Println("faux, c'est pas ce mot")
			h.tried = append(h.tried, s)
			h.attempt -= 1
		}
	} else {
		for _, word := range h.tried {
			if s == word {
				fmt.Println("tu a  déjà essayé cette lettre, tente autre chose:")
				test = true
				fmt.Scanln(&s)
				h.tried = append(h.tried, s)
				h.attempt -= 1
			}
		}
		if !test {
			for i, letter := range h.word {
				if s == string(letter) {
					h.hiddenWord[i] = string(letter)
					found = true
				}
			}
			if !found {
				fmt.Println("Votre lettre n'est pas présente dans le mot.")
				h.tried = append(h.tried, s)
				h.attempt -= 1
			} else {
				h.tried = append(h.tried, s)
				fmt.Println("Bravo, tu a  deviné une lettre continue comme ça !")
			}
		}
	}
}

func (h *Hangman) test(s string) {
	for _, word := range h.tried {
		if s == word {
			fmt.Println("tu a  déjà essayé ", word)
			fmt.Println("Entrez un mot ou une lettre.")
			fmt.Scanln(&s)
			break
		}
	}
}

func (h *Hangman) gagner() {
	match := false
Loop:
	for i, letter := range h.word {

		if h.hiddenWord[i] == string(letter) {
			match = true
		} else {
			match = false
			break Loop
		}
	}
	if match {
		fmt.Println("Bravo, tu as gagnés ! Le mot était bien: ", h.word)
		os.Exit(0)
	}
}

//message de gagne + affichage du mot
func (h *Hangman) perdu() {
	if h.attempt == 0 {
		fmt.Println("Tu as perdu.")
		fmt.Println("Donc tu est mort")
		fmt.Println("Le mot a deviner était:", h.word)
		os.Exit(0)
	}
}

//message de fin si perte du joeur + affichage du mot
