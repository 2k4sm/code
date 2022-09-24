package main

import (
	"fmt"
	"os"
)

const (
	MAX_voters     int = 100
	MAX_candidates     = 9
)

var (
	preferences     [MAX_voters][MAX_candidates]int
	candidates      [MAX_candidates]candidate
	voter_count     int
	candidate_count int
	i               int
	j               int
	k               int
	name            string
)

type candidate struct {
	name       string
	votes      int
	eliminated bool
}

func main() {
	strarg := argc()
	fmt.Println(argc())
	//Populate array of candidates.
	candidate_count = len(argc()) - 1
	if candidate_count > MAX_candidates {
		fmt.Printf("Maximum number of candidates is %d\n", MAX_candidates)
		os.Exit(2)
	}
	for i = 0; i < candidate_count; i++ {
		candidates[i].name = strarg[i]
		candidates[i].votes = 0
		candidates[i].eliminated = false

	}
	fmt.Println("Number of voters:")
	fmt.Scanf("%d", &voter_count)
	if voter_count > MAX_voters {
		fmt.Printf("Maximum number of voters is %d\n", MAX_voters)
		os.Exit(3)
	}

	//Keep querying for votes.
	for i = 0; i < voter_count; i++ {

		//Query for each rank.
		for j = 0; j < candidate_count; j++ {

			fmt.Printf("Rank %d:", j+1)
			fmt.Scanf("%s", &name)

			// Record vote, unless it's invalid
			if !vote(i, j, name) {
				fmt.Printf("invalid vote.\n")
				os.Exit(4)

			}

		}
		fmt.Printf("\n")
	}

	//Keep holding runoffs untill winner exists.
	for true {

		//calculate votes given remaining candidates.
		tabulate()

		//check if election has been won.
		won := print_winner()

		if won {
			break
		}

		//eliminate last-place candidates
		min := find_min()
		tie := is_tie(min)

		// If tie,everyone wins
		if tie {
			for i = 0; i < candidate_count; i++ {
				if !candidates[i].eliminated {
					fmt.Printf("%s\n", candidates[i].name)

				}
			}
			break
		}
		//eliminate anyone with minimum number of votes.
		eliminate(min)

		//Reset vote count to zero.
		for i = 0; i < candidate_count; i++ {
			candidates[i].votes = 0
		}

	}

}
func argc() []string {
	//checks for invalid usage.
	var strarg []string
	strarg = os.Args[1:]
	if len(strarg) < 2 {
		fmt.Println("Usage: runoff [candidate ...]")
		os.Exit(1)
	}
	return strarg
}

//Record preference if vote is valid.
func vote(voter int, rank int, name string) bool {
	for i < candidate_count {
		if name == candidates[i].name {
			preferences[voter][rank] = i
			fmt.Println(preferences)
			i++
			continue
		}
		return true

	}
	return false
}

//tabulate votes for non eliminated candidates.
func tabulate() {
	for i = 0; i < voter_count; i++ {
		for j = 0; j < candidate_count; j++ {
			k := preferences[i][j]
			if !(candidates[k].eliminated) {
				candidates[k].votes += 1
				break
			}
		}
	}
}

//Print the winner of the election if there is one.
func print_winner() bool {
	mth := voter_count / 2
	for k := range candidates {
		if candidates[k].votes > mth {
			winner := candidates[k].name
			fmt.Printf("The winner is %s", winner)
			return true
		}
	}
	return false
}
func find_min() int {
	min := 100
	for k := range candidates {
		temp_min := candidates[k].votes
		if temp_min < min {
			min = temp_min
		}

	}
	if min != 100 {
		return min
	} else {
		return 0
	}
}
func is_tie(min int) bool {
	i := 0
	for k := range candidates {
		if candidates[k].votes == min {
			i++
			continue
		} else {
			break
		}
	}
	if i != 0 {
		return true
	}

	return false
}

func eliminate(min int) {
	for k := range candidates {
		if candidates[k].votes == min {
			candidates[k].eliminated = true
		}
	}
}
