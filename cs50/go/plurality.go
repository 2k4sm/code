package main

import (
	"bufio"
	"fmt"
	"os"
)

// Max number of candidates
const MAX int = 9

// Candidates have name and vote count
type candidate struct {
	name  string
	votes int
}

var (
	candidates      [MAX]candidate
	candidate_count int
	i               int = 0
	voter_count     int
)

func main() {

	//Populate array of candidates.
	candidate_count = len(candiarg())
	if candidate_count > MAX {
		fmt.Printf("Maximum number of candidates is %d\n", MAX)
		os.Exit(2)
	}
	strarg := candiarg()
	for i < candidate_count {
		candidates[i].name = strarg[i]
		candidates[i].votes = 0
		i++
	}

	//Take number of voters.
	fmt.Printf("Number of voters: ")
	fmt.Scanf("%d", &voter_count)

	//loop over all voters.
	for i = 0; i < voter_count; i++ {
		name := inp("Vote: ")
		if !vote(name) {
			fmt.Println("Invalid vote.")
		}
	}

	//Display winner of election
	print_winner()
}
func candiarg() []string {
	var strarg []string
	strarg = os.Args[1:]
	if len(strarg) < 2 {
		fmt.Println("Usage: plurality [candidate ...]")
		os.Exit(1)
	}
	return strarg
}
func inp(givetxt string) (text string) {
	fmt.Printf(givetxt)
	//takes the text as user input.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text = scanner.Text()
	return text
}

//Update vote totals given a new vote.
func vote(name string) bool {
	for j := range candidates {
		if name == candidates[j].name {
			candidates[j].votes += 1
			return true
		}
	}
	return false

}

//print the winner of the election.
func print_winner() {
	winner_votes := 0
	var winner_class []string
	for j := range candidates {
		if candidates[j].votes > winner_votes {
			winner_votes = candidates[j].votes
			continue
		}

	}
	for k := range candidates {
		if candidates[k].votes == winner_votes {
			winner_class = append(winner_class, candidates[k].name)
		}
	}
	fmt.Printf("The winner is:%v\n", winner_class)
}
