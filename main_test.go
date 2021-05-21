package main_test

import (
	folder_level_c_with_even_more_characters_and_further_more_a_longer_directory_name_much_longer_than_b "github.com/ajgajg1134/test-golang-failure-with-long-path-windows/folder-level-a/folder-level-b-much-longer-than-a/folder-level-c-with-even-more-characters-much-longer-than-b/github.com/ajgajg1134/test-golang-failure-with-long-path-windows/folder-level-a/folder-level-b-much-longer-than-a/folder-level-c-with-even-more-characters-and-further-more-a-longer-directory-name-much-longer-than-b"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain_Potato_Test(t *testing.T) {
	_ = folder_level_c_with_even_more_characters_and_further_more_a_longer_directory_name_much_longer_than_b.Blah{}

	assert.True(t, true)
}
