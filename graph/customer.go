package graph

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"gymboss/graph/model"
)

// CreateCustomer is the resolver for the createCustomer field.
func (r *mutationResolver) CreateCustomer(ctx context.Context, customer model.CustomerCreate) (*model.Customer, error) {
	id := uuid.NewV4().String()
	c := model.CustomerDB{
		ID:    id,
		Name:  customer.Name,
		Email: customer.Email,
	}

	res := DBConn.Create(&c)
	if res.Error != nil {
		return nil, res.Error
	}

	return &model.Customer{
		ID:       c.ID,
		Name:     c.Name,
		Email:    c.Email,
		Register: nil,
	}, nil
}

// UpdateCustomer is the resolver for the updateCustomer field.
func (r *mutationResolver) UpdateCustomer(ctx context.Context, customer model.CustomerUpdate) (*model.Customer, error) {
	c := model.CustomerDB{ID: customer.ID}
	res := DBConn.First(&c)
	if res.Error != nil {
		return nil, res.Error
	}

	if customer.Name != nil {
		c.Name = *customer.Name
	}
	if customer.Email != nil {
		c.Email = *customer.Email
	}

	res = DBConn.Save(&c)
	if res.Error != nil {
		return nil, res.Error
	}

	var rs []model.PurchaseDB
	var register []*model.Purchase

	res = DBConn.Where("customer_id = ?", c.ID).Find(&rs)
	if res.Error != nil {
		return nil, res.Error
	}

	for _, p := range rs {
		register = append(register, &model.Purchase{
			ID:       p.ID,
			Training: nil,
			Customer: nil,
			Coast:    p.Coast,
			Income:   p.Income,
		})
	}

	return &model.Customer{
		ID:       c.ID,
		Name:     c.Name,
		Email:    c.Email,
		Register: register,
	}, nil
}

// DeleteCustomer is the resolver for the deleteCustomer field.
func (r *mutationResolver) DeleteCustomer(ctx context.Context, id string) (string, error) {
	c := model.CustomerDB{ID: id}
	res := DBConn.Delete(&c)
	if res.Error != nil {
		return "", res.Error
	}

	return "Succeed", nil
}

// ReadCustomer is the resolver for the readCustomer field.
func (r *queryResolver) ReadCustomer(ctx context.Context, id string) (*model.Customer, error) {
	c := model.CustomerDB{ID: id}
	res := DBConn.First(&c)
	if res.Error != nil {
		return nil, res.Error
	}

	var rs []model.PurchaseDB
	var register []*model.Purchase

	res = DBConn.Where("customer_id = ?", c.ID).Find(&rs)
	if res.Error != nil {
		return nil, res.Error
	}

	for _, p := range rs {
		register = append(register, &model.Purchase{
			ID:       p.ID,
			Training: nil,
			Customer: nil,
			Coast:    p.Coast,
			Income:   p.Income,
		})
	}

	return &model.Customer{
		ID:       c.ID,
		Name:     c.Name,
		Email:    c.Email,
		Register: register,
	}, nil
}
