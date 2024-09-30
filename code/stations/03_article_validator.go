package stations

import (
	"errors"
)

type ArticleValidator struct{}

func (this *ArticleValidator) Do(input any, output func(any)) {
	// TODO: given a contracts.Article, validate the Slug and the Title fields and emit the contracts.Article (or an error)
	// input: contracts.Article
	// output: contracts.Article (or error)
}

var errInvalidContent = errors.New("invalid content")
