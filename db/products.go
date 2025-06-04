package db

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"slices"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"        validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price"       validate:"required,gt=0"`
	SKU         string  `json:"sku"         validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

var productList = Products{
	{
		ID:          1,
		Name:        "Clean Water",
		Description: "Real water",
		Price:       1000,
		SKU:         "MW-001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Dirty Water",
		Description: "Not so real water",
		Price:       500,
		SKU:         "DW-001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[A-Z]+-[0-9]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}

func (p *Product) Validate() error {
	val := validator.New()
	val.RegisterValidation("sku", validateSKU)

	return val.Struct(p)
}

func (p *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(p)
}

func (p *Product) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

func GetProductById(id int) (*Product, error) {
	pos, err := findProduct(id)
	if err != nil {
		return nil, err
	}

	return productList[pos], nil
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	p.CreatedOn = time.Now().UTC().String()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	p.CreatedOn = productList[pos].CreatedOn
	p.UpdatedOn = time.Now().UTC().String()
	productList[pos] = p
	return nil
}

func DeleteProduct(id int) error {
	i, err := findProduct(id)
	if err != nil {
		return ErrProductNotFound
	}

	productList = slices.Delete(productList, i, i+1)

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (int, error) {
	for i, p := range productList {
		if p.ID == id {
			return i, nil
		}
	}
	return 0, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}
