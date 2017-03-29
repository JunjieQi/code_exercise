package message

import "time"

const (
	regexMention   = `@(\w+)`
	regexEmoticons = `\(([0-9A-Za-z]{1,15})\)`
)

const (
	maxIdleConnection int = 10
	timeout               = time.Second
)

const (
	MaxGorountine int = 10
)
