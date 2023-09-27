package union_find

import "testing"

func TestUnionFileTiny(t *testing.T) {
	_, err := CreateUnionFile("./data/tinyUF.txt")
	if err != nil {
		t.Fatalf("failed to create union file find")
	}
}

func TestUnionFileMedium(t *testing.T) {
	_, err := CreateUnionFile("./data/mediumUF.txt")
	if err != nil {
		t.Fatalf("failed to create union file find")
	}
}

func TestUnionFileLarge(t *testing.T) {
	_, err := CreateUnionFile("./data/largeUF.txt")
	if err != nil {
		t.Fatalf("failed to create union file find")
	}
}
