package model

import "GO/node"

var DaftarPembeli node.ListPembeli

// CREATE
func CreatePembeli(jab node.Pembeli) bool {
	tempLL := node.ListPembeli{
		Data: jab,
		Link: nil,
	}
	if DaftarPembeli.Link == nil {
		DaftarPembeli.Link = &tempLL
		return true
	} else {
		temp := &DaftarPembeli
		for temp.Link != nil {
			temp = temp.Link
		}
		temp.Link = &tempLL
		return true
	}
	//return false
}

// READ
func ReadPembeli() []node.Pembeli {
	daftarPembeli := []node.Pembeli{}
	temp := &DaftarPembeli
	for temp.Link != nil {
		daftarPembeli = append(daftarPembeli, temp.Link.Data)
		temp = temp.Link
	}
	return daftarPembeli
}

// UPDATE
func UpdatePembeli(jab node.Pembeli, ID int) bool {
	temp := DaftarPembeli.Link
	for temp != nil {
		if temp.Data.ID == ID {
			temp.Data = jab
			return true
		}
		temp = temp.Link
	}
	return false
}

// DELETE
func DeletePembeli(ID int) bool {
	temp := &DaftarPembeli
	for temp.Link != nil {
		if temp.Link.Data.ID == ID {
			temp.Link = temp.Link.Link
			return true
		}
		temp = temp.Link
	}
	return false
}

func SearchPembeli(ID int) bool {
	temp := &DaftarPembeli
	for temp.Link != nil {
		if temp.Link.Data.ID == ID {
			return true
		}
		temp = temp.Link
	}
	return false
}

func GetNama(ID int) string {
	temp := &DaftarPembeli
	for temp.Link != nil {
		if temp.Link.Data.ID == ID {
			return temp.Link.Data.Nama
		}
		temp = temp.Link
	}
	return ""
}
