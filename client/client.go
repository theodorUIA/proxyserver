package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8200")
	if err != nil {
		log.Fatal(err)
	}

	message := []rune("Møte i Ålesund 1. juni kl. 25:59")
	kryptertMelding := minKrypteringsfunksjon(message, 4) //Kryptering av meldingen
	_, err = conn.Write([]byte(string(kryptertMelding)))
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	response := string(buf[:n])

	deKryptertMelding := minDekrypteringsfunksjon([]rune(response), 4)
	log.Printf("reply from proxy: %s", string(deKryptertMelding))
}

//Funksjon som krypterer meldinger med en cæsar kryptering
func minKrypteringsfunksjon(message []rune, krypteringsTall int) []rune {

	kryptertMelding := make([]rune, len(message)) // Lager en rune array med lengden av meldingen

	for i := 0; i < len(message); i++ { // For loop som kjører så lenge "i" er mindre enn lengden av meldingen
		asciiNummer := int(message[i]) // Henter ASCII verdien til hvert tegn i meldingen

		kryptertMelding[i] = rune(asciiNummer + krypteringsTall) // Legger til krypteringsverdien til ASCII verdien
	}
	return kryptertMelding // Returnerer kryptert melding
}

//Funksjon som dekrypterer meldinger med en cæsar kryptering
func minDekrypteringsfunksjon(message []rune, krypteringsTall int) []rune {

	kryptertMelding := make([]rune, len(message)) // Lager en rune array med lengden av meldingen

	for i := 0; i < len(message); i++ { // For loop som kjører så lenge "i" er mindre enn lengden av meldingen
		asciiNummer := int(message[i]) // Henter ASCII verdien til hvert tegn i meldingen

		kryptertMelding[i] = rune(asciiNummer - krypteringsTall) // Fjerner krypteringsverdien til ASCII verdien
	}
	return kryptertMelding // Returnerer ukryptert melding
}
