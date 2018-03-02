package tokenstats

import "strconv"

type TokenStats struct {
	stats map[string]map[string]int
}

func NewTokenStats(stats map[string]map[string]int) *TokenStats {
	return &TokenStats{stats: stats}
}

func (s *TokenStats) GetStats() map[string]map[string]int {
	return s.stats
}

func (s *TokenStats) AsTable() [][]string {
	var table [][]string

	for user, labels := range s.stats {
		for label, count := range labels {
			table = append(table, []string{user, label, strconv.Itoa(count)})
		}
	}

	return table
}
