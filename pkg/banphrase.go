package pkg

type Banphrase interface {
	Triggers(text string) bool
	IsCaseSensitive() bool

	// IsAdvanced decides whether or not the banphrase should be run on all variations or only the first one
	IsAdvanced() bool
}
