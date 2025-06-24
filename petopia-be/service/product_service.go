
package service

import (
    "time"
    "errors"
    "petopia-be/dao"
    "petopia-be/db"
    "petopia-be/dto"
)

func CreateProduct(productRequestDTO dto.ProductRequestDTO) (dto.ProductResponseDTO, error) {
	database, err := db.Connect()
	if err != nil {
		return dto.ProductResponseDTO{}, err
	}
	product := dao.Product{
		Name:        productRequestDTO.Name,
		Description: productRequestDTO.Description,
		Price:       productRequestDTO.Price,
	}
	if err := database.Create(&product).Error; err != nil {
		return dto.ProductResponseDTO{}, err
	}

	productResponseDTO := dto.ProductResponseDTO{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
	}

	return productResponseDTO, nil
}

func UpdateProduct(id int, productRequestDTO dto.ProductRequestDTO) (dto.ProductResponseDTO, error) {
	database, err := db.Connect()
	if err != nil {
		return dto.ProductResponseDTO{}, err
	}

	var product dao.Product
	if err := database.First(&product, id).Error; err != nil {
		return dto.ProductResponseDTO{}, errors.New("product not found")
	}

	product.Name = productRequestDTO.Name
	product.Description = productRequestDTO.Description
	product.Price = productRequestDTO.Price

	if err := database.Save(&product).Error; err != nil {
		return dto.ProductResponseDTO{}, err
	}

	productResponseDTO := dto.ProductResponseDTO{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
	}

	return productResponseDTO, nil
}

func DeleteProduct(id int) error {
	database, err := db.Connect()
	if err != nil {
		return err
	}

	if err := database.Delete(&dao.Product{}, id).Error; err != nil {
		return err
	}

	return nil
}

func ListProducts() ([]dto.ProductResponseDTO, error) {
	database, err := db.Connect()
	if err != nil {
		return nil, err
	}

	var products []dao.Product
	if err := database.Find(&products).Error; err != nil {
		return nil, err
	}

	var productResponseDTOs []dto.ProductResponseDTO
	for _, product := range products {
		productResponseDTOs = append(productResponseDTOs, dto.ProductResponseDTO{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			CreatedAt:   product.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
		})
	}

	return productResponseDTOs, nil
}
