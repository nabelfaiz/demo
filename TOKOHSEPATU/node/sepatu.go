package node

type Sepatu struct {
	IDSepatu    int
	NamaSepatu  string
	HargaSepatu int
}

// Struktur untuk merepresentasikan node dalam linked list Toko Sepatu
type ListSepatu struct {
	Data Sepatu
	Link *ListSepatu
}

