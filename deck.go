package main

import (
	"fmt"
	"os"
	"math/rand"
	"io/ioutil"
	"strings"
	"time"
)

// Create a new type of deck
// wich is a slice of strings

type deck []string

//retorna un nuevo deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace","Two","Three","Four"}

	// si la clave como i o j no se usa se remplaza por _
	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards =  append(cards, value + " of " + suit)
		}
	}

	return cards
}

func (d deck) print(){
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// (deck, deck) - retorna dos valores de tipo deck
func deal(d deck, handSize int) (deck, deck){
	// d[:handSize] de 0 hasta la posicion y d[handSize:] despues de la posicion hasta el final del listado
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
	
}

// []byte(d.toString()) convierte string a bytes - 0666 anyone can read and write this file 
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}