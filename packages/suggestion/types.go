package suggestion

type Suggestion struct {
	Suggestions    []string
	RequestTokens  int
	ResponseTokens int
	Timestamp      int64 // Unix
}
