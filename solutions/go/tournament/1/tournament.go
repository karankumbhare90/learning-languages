package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// TeamStats stores the performance stats for each team
type TeamStats struct {
	Name           string
	MatchesPlayed  int
	Wins           int
	Draws          int
	Losses         int
	Points         int
}

// Tally reads match results and writes the standings table.
func Tally(reader io.Reader, writer io.Writer) error {
	stats := make(map[string]*TeamStats)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue // skip empty or comment lines
		}

		parts := strings.Split(line, ";")
		if len(parts) != 3 {
			return fmt.Errorf("invalid input line: %q", line)
		}

		team1, team2, result := parts[0], parts[1], parts[2]

		// Ensure both teams exist in the map
		if _, ok := stats[team1]; !ok {
			stats[team1] = &TeamStats{Name: team1}
		}
		if _, ok := stats[team2]; !ok {
			stats[team2] = &TeamStats{Name: team2}
		}

		// Update stats based on result
		switch result {
		case "win":
			stats[team1].Wins++
			stats[team1].Points += 3
			stats[team2].Losses++
		case "loss":
			stats[team2].Wins++
			stats[team2].Points += 3
			stats[team1].Losses++
		case "draw":
			stats[team1].Draws++
			stats[team2].Draws++
			stats[team1].Points++
			stats[team2].Points++
		default:
			return fmt.Errorf("invalid result: %s", result)
		}

		stats[team1].MatchesPlayed++
		stats[team2].MatchesPlayed++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Collect and sort teams
	var teams []*TeamStats
	for _, t := range stats {
		teams = append(teams, t)
	}
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].Points == teams[j].Points {
			return teams[i].Name < teams[j].Name
		}
		return teams[i].Points > teams[j].Points
	})

	// Print header
	fmt.Fprintf(writer, "%-31s| MP |  W |  D |  L |  P\n", "Team")
	for _, t := range teams {
		fmt.Fprintf(writer, "%-31s| %2d | %2d | %2d | %2d | %2d\n",
			t.Name, t.MatchesPlayed, t.Wins, t.Draws, t.Losses, t.Points)
	}

	return nil
}
