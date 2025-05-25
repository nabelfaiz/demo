package model

import "GO/node"

var DaftarSepatu node.ListSepatu

func CreateSepatu(emp node.Sepatu) bool {
	tempLL := node.ListSepatu{
		Data: emp,
		Link: nil,
	}
	if DaftarSepatu.Link == nil {
		DaftarSepatu.Link = &tempLL
		return true
	} else {
		temp := &DaftarSepatu
		for temp.Link != nil {
			temp = temp.Link
		}
		temp.Link = &tempLL
		return true
	}
	//return false
}

func ReadSepatu() []node.Sepatu {
    var daftarSepatu []node.Sepatu // Menggunakan slice untuk menyimpan hasil
    temp := DaftarSepatu.Link // Mulai dari node pertama
    for temp != nil {
        daftarSepatu = append(daftarSepatu, temp.Data) // Menambahkan data ke slice
        temp = temp.Link // Melanjutkan ke node berikutnya
    }
    return daftarSepatu
}

func UpdateSepatu(emp node.Sepatu, id int) bool {
	temp := DaftarSepatu.Link
	for temp != nil {
		if temp.Data.IDSepatu == id {
			temp.Data = emp
			return true
		}
		temp = temp.Link
	}
	return false
}

func DeleteSepatu(id int) bool {
	temp := &DaftarSepatu
	for temp.Link != nil {
		if temp.Link.Data.IDSepatu == id {
			temp.Link = temp.Link.Link
			return true
		}
		temp = temp.Link
	}
	return false
}


func SearchSepatu(ID int) bool {
	temp := &DaftarSepatu
	for temp.Link != nil {
		if temp.Link.Data.IDSepatu == ID {
			return true
		}
		temp = temp.Link
	}
	return false
}

func GetSepatuByID(id int) (node.Sepatu, bool) {
	temp := DaftarSepatu.Link
	for temp != nil {
		if temp.Data.IDSepatu == id {
			return temp.Data, true
		}
		temp = temp.Link
	}
	return node.Sepatu{}, false
}

func GetNamaSepatu(ID int) string {
    temp := &DaftarSepatu
    for temp.Link != nil {
        if temp.Link.Data.IDSepatu == ID {
            return temp.Link.Data.NamaSepatu
        }
        temp = temp.Link
    }
    return ""
}
