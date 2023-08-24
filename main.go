package main

import (
	"bufio"
	"fmt"
	"os"
)

type Candidate struct {
	Name  string
	Party string
	Votes int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Sistema de Votación")

	var candidates map[int]Candidate
	candidates = make(map[int]Candidate)

	fmt.Print("Ingrese el número de candidatos: ")
	var numCandidates int
	fmt.Scanln(&numCandidates)

	for i := 1; i <= numCandidates; i++ {
		var candidateName, candidateParty string
		fmt.Printf("Ingrese el nombre del candidato #%d: ", i)
		scanner.Scan()
		candidateName = scanner.Text()

		fmt.Printf("Ingrese el número de partido para %s: ", candidateName)
		scanner.Scan()
		candidateParty = scanner.Text()

		candidates[i] = Candidate{Name: candidateName, Party: candidateParty}
	}

	blankVotes := 0
	nullVotes := 0
	totalVotes := 0

	votingOpen := true
	for votingOpen {
		fmt.Println("\nMenú de Votación:")
		for num, candidate := range candidates {
			fmt.Printf("%d - %s (%s)\n", num, candidate.Name, candidate.Party)
		}
		fmt.Println("0 - Voto en blanco")
		fmt.Println("-1 - Voto nulo")
		fmt.Println("-2 - Cerrar votación")

		var vote int
		fmt.Print("Ingrese su voto: ")
		fmt.Scanln(&vote)

		switch vote {
		case -2:
			votingOpen = false
		case -1:
			nullVotes++
		case 0:
			blankVotes++
		default:
			if _, ok := candidates[vote]; ok {
				candidate := candidates[vote]
				candidate.Votes++
				candidates[vote] = candidate
				totalVotes++
				fmt.Printf("%s (%s): %d votos\n\n", candidate.Name, candidate.Party, candidate.Votes)
			} else {
				fmt.Println("Opción de voto inválida.")
			}
		}
	}

	fmt.Println("\nResultados de la Votación:")
	for _, candidate := range candidates {
		voteCount := candidate.Votes
		percentage := float64(voteCount) / float64(totalVotes) * 100
		fmt.Printf("%s (%s): %d votos (%.2f%%)\n", candidate.Name, candidate.Party, voteCount, percentage)
	}

	blankPercentage := float64(blankVotes) / float64(totalVotes) * 100
	fmt.Printf("Votos en blanco: %d (%.2f%%)\n", blankVotes, blankPercentage)

	nullPercentage := float64(nullVotes) / float64(totalVotes) * 100
	fmt.Printf("Votos nulos: %d (%.2f%%)\n", nullVotes, nullPercentage)

	//TODO Imprimir el nombre del ganador
	//o los candidatos que empataron

}
