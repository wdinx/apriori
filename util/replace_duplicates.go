package util

func RemoveDuplicate(data []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, item := range data {
		if !seen[item] {
			result = append(result, item) // hanya ambil pertama kali
			seen[item] = true
		}
	}
	return result
}
