package main_test

import (
	once_more_into_the_breach "github.com/ajgajg1134/test-golang-failure-with-long-path-windows/folder-level-a/folder-level-b-much-longer-than-a/folder-level-c-with-even-more-characters-much-longer-than-b/folder_level_c_with_even_more_characters_and_further_more_a_longer_directory_name_much_longer_than_b/once-more-into-the-breach"
	"github.com/stretchr/testify/assert"
	"testing"
)

//This tests seems to not be required to force vet to fail
func TestMain_Potato_Test(t *testing.T) {
	//_ = folder_level_c_with_even_more_characters_and_further_more_a_longer_directory_name_much_longer_than_b.Blah{}

	_ = once_more_into_the_breach.BigBlah{}

	assert.True(t, true)
}
