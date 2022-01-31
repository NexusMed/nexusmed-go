// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package products

import (
	"fmt"
	"io"
	"strconv"
)

type IProduct interface {
	IsIProduct()
}

type Product interface {
	IsProduct()
}

type Dosage struct {
	Quantity *float64    `json:"quantity,omitempty"`
	Unit     *DosageUnit `json:"unit,omitempty"`
}

type Medication struct {
	Name     *string `json:"name,omitempty"`
	Dosage   *Dosage `json:"dosage,omitempty"`
	Quantity *int    `json:"quantity,omitempty"`
}

type MedicinalProduct struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Price      int         `json:"price"`
	Medication *Medication `json:"medication,omitempty"`
	ConsultFee *int        `json:"consult_fee,omitempty"`
}

func (MedicinalProduct) IsProduct()  {}
func (MedicinalProduct) IsIProduct() {}

type Products struct {
	Items     []Product `json:"items,omitempty"`
	NextToken *string   `json:"next_token,omitempty"`
}

type DosageUnit string

const (
	DosageUnitMg  DosageUnit = "mg"
	DosageUnitMcg DosageUnit = "mcg"
	DosageUnitMl  DosageUnit = "ml"
	DosageUnitIu  DosageUnit = "iu"
)

var AllDosageUnit = []DosageUnit{
	DosageUnitMg,
	DosageUnitMcg,
	DosageUnitMl,
	DosageUnitIu,
}

func (e DosageUnit) IsValid() bool {
	switch e {
	case DosageUnitMg, DosageUnitMcg, DosageUnitMl, DosageUnitIu:
		return true
	}
	return false
}

func (e DosageUnit) String() string {
	return string(e)
}

func (e *DosageUnit) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DosageUnit(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DosageUnit", str)
	}
	return nil
}

func (e DosageUnit) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProductType string

const (
	ProductTypeMedicinalProduct ProductType = "MedicinalProduct"
)

var AllProductType = []ProductType{
	ProductTypeMedicinalProduct,
}

func (e ProductType) IsValid() bool {
	switch e {
	case ProductTypeMedicinalProduct:
		return true
	}
	return false
}

func (e ProductType) String() string {
	return string(e)
}

func (e *ProductType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProductType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProductType", str)
	}
	return nil
}

func (e ProductType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
