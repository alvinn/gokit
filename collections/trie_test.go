package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()

	// Test Insert and Search
	trie.Insert("apple")
	assert.True(t, trie.Search("apple"), "Expected Search to return true for 'apple'")
	assert.False(t, trie.Search("app"), "Expected Search to return false for 'app'")

	// Test StartsWith
	assert.True(t, trie.StartsWith("app"), "Expected StartsWith to return true for 'app'")
	assert.False(t, trie.StartsWith("apl"), "Expected StartsWith to return false for 'apl'")

	// Test Insert and Search for prefix
	trie.Insert("app")
	assert.True(t, trie.Search("app"), "Expected Search to return true for 'app'")

	// Test Insert and Search for another word
	trie.Insert("banana")
	assert.True(t, trie.Search("banana"), "Expected Search to return true for 'banana'")
	assert.False(t, trie.Search("ban"), "Expected Search to return false for 'ban'")
	assert.True(t, trie.StartsWith("ban"), "Expected StartsWith to return true for 'ban'")
}
