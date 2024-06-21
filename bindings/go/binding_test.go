package tree_sitter_prolog_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-prolog"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_prolog.Language())
	if language == nil {
		t.Errorf("Error loading Prolog grammar")
	}
}
