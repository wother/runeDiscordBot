package workers

// ChunkArray chunks an array into smaller arrays of a specified size.
func ChunkArray(arr []string, size int) [][]string {
	var chunks [][]string
	for i := 0; i < len(arr); i += size {
		end := i + size
		if end > len(arr) {
			end = len(arr)
		}
		chunks = append(chunks, arr[i:end])
	}
	return chunks
}
