package database

type Article struct {
	EntryName string
	Title     string
	Issued    string
	Modified  string
	Author    string
	TOC       string
	HTMLText  []byte
}
