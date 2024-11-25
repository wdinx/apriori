package util

func CountExactMatches(data [][]string, query []string) int {
	count := 0

	// Iterasi melalui setiap elemen dalam `data`
	for _, group := range data {
		matches := 0

		// Periksa apakah setiap elemen dalam `query` ada di `group`
		for _, q := range query {
			for _, val := range group {
				if q == val {
					matches++
					break // Lanjut ke elemen query berikutnya
				}
			}
		}

		// Jika semua elemen dalam `query` ada di `group`, tambahkan count
		if matches == len(query) {
			count++
		}
	}

	return count
}
