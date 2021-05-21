package main

import (
	folder_level_b_much_longer_than_a "github.com/ajgajg1134/test-golang-failure-with-long-path-windows/folder-level-a/folder-level-b-much-longer-than-a"
	folder_level_c_with_even_more_characters_much_longer_than_b "github.com/ajgajg1134/test-golang-failure-with-long-path-windows/folder-level-a/folder-level-b-much-longer-than-a/folder-level-c-with-even-more-characters-much-longer-than-b"
	folder_level_c_with_even_more_characters_and_further_more_a_longer_directory_name_much_longer_than_b "github.com/ajgajg1134/test-golang-failure-with-long-path-windows/folder-level-a/folder-level-b-much-longer-than-a/folder-level-c-with-even-more-characters-much-longer-than-b/github.com/ajgajg1134/test-golang-failure-with-long-path-windows/folder-level-a/folder-level-b-much-longer-than-a/folder-level-c-with-even-more-characters-and-further-more-a-longer-directory-name-much-longer-than-b"
)

func main() {
	_ = folder_level_b_much_longer_than_a.TestStruct{}

	_ = folder_level_c_with_even_more_characters_much_longer_than_b.TestStruct2{}

	_ = folder_level_c_with_even_more_characters_and_further_more_a_longer_directory_name_much_longer_than_b.Blah{}

	println("Test")
}
