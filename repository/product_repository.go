package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, name, price, created_at FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
			&productObj.CreatedAt)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO product" + "(name, price)" + "VALUES($1, $2) RETURNING id")	
	if err != nil{
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil{
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil

}


func(pr *ProductRepository) GetProductsByID(id_product int) (*model.Product, error){
	query, err := pr.connection.Prepare("SELECT * from product where id = $1")
	if err != nil{
		fmt.Println(err)
		return nil, err
	}

	var product model.Product

	err = query.QueryRow(id_product).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.CreatedAt,
	)
	if(err != nil){
		if(err == sql.ErrNoRows){
			return nil, nil
		}
		return nil, err
	}
	query.Close()
	return &product, nil
}


func(pr *ProductRepository) DeleteProduct(id_product int) error {
    query, err := pr.connection.Prepare("DELETE from product where id = $1")
    if err != nil {
        fmt.Println(err)
        return err
    }
    defer query.Close()

    result, err := query.Exec(id_product)
    if err != nil {
        fmt.Println(err)
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        fmt.Println(err)
        return err
    }

    if rowsAffected == 0 {
        return sql.ErrNoRows
    }

    return nil
}
