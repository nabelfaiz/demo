package node

type Pembeli struct {
	ID       int
	Nama     string
	SepatuID int
}

// Struktur untuk merepresentasikan node dalam linked list pembeli
type ListPembeli struct {
	Data Pembeli
	Link *ListPembeli
}
