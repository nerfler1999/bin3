package storage

import (
	"bin_3/bins"
	"bin_3/file"
	"encoding/json"
	"fmt"
	"time"
)

type BinList struct {
	Bins      []bins.Bin `json:"bins"`
	UpdatedAt time.Time  `json:"upadtedat"`
}

func (bin *BinList) ToBytes() ([]byte, error) {
	file, err := json.Marshal(bin)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewBinList() *BinList {
	file, err := file.ReadFile("data.json")
	if err != nil {
		return &BinList{
			Bins:      []bins.Bin{},
			UpdatedAt: time.Now(),
		}
	}
	var binlist BinList
	err = json.Unmarshal(file, &binlist)
	if err != nil {
		fmt.Println("Не удалось разобрать файл data.json")
	}
	return &binlist
}

func (bin *BinList) AddBinToBinList(newbin *bins.Bin) {
	bin.Bins = append(bin.Bins, acc)
	bin.UpdatedAt = time.Now()
	data, err := bin.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать файл")
	}
	file.WriteFile(data, "data.json")
}
