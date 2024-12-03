package entity

type Network struct {
	IP           string `db:"ip"`
	Mask         string `db:"mask"`
	BinaryPrefix string `db:"binary_prefix"`
}
