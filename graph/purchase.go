package graph

import (
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	"gymboss/graph/model"
)

// PurchaseTraining is the resolver for the purchaseTraining field.
func (r *mutationResolver) PurchaseTraining(ctx context.Context, trainingID string, customerId string) (*model.Purchase, error) {
	t := model.TrainingDB{ID: trainingID}
	res := DBConn.First(&t)
	if res.Error != nil {
		return nil, res.Error
	}

	g := model.GymDB{ID: t.GymID}
	res = DBConn.First(&g)
	if res.Error != nil {
		return nil, res.Error
	}

	if g.Slots-1 < 0 {
		return nil, errors.New("Failed: no slots avalible")
	} else {
		g.Slots -= 1
		res = DBConn.Save(&g)
		if res.Error != nil {
			return nil, res.Error
		}
	}

	c := model.Customer{ID: customerId}
	res = DBConn.First(&c)
	if res.Error != nil {
		return nil, res.Error
	}

	id := uuid.NewV4().String()
	purchase := &model.Purchase{
		ID: id,
		Training: &model.Training{
			ID:       t.ID,
			Category: t.Category,
			Coast:    t.Coast,
			Gym: &model.Gym{
				ID:        g.ID,
				Branch:    g.Branch,
				Admin:     g.Admin,
				Phone:     g.Phone,
				Trainings: nil,
				Slots:     g.Slots,
			},
		},
		Customer: &c,
		Coast:    t.Coast,
		Income:   float64(t.Coast) * 0.8,
	}

	return purchase, nil
}

// DeletePurchase is the resolver for the deletePurchase field.
func (r *mutationResolver) DeletePurchase(ctx context.Context, id string) (*model.Purchase, error) {
	p := model.PurchaseDB{ID: id}
	res := DBConn.Delete(&p)
	if res.Error != nil {
		return nil, res.Error
	}

	return nil, nil
}

// Purchases is the resolver for the purchases field.
func (r *queryResolver) Purchases(ctx context.Context, customerID string) ([]*model.Purchase, error) {
	var ps []model.PurchaseDB
	res := DBConn.Where("customer_id = ?", customerID).Find(&ps)
	if res.Error != nil {
		return nil, res.Error
	}

	var purchases []*model.Purchase
	for _, p := range ps {
		purchases = append(purchases, &model.Purchase{
			ID:       p.ID,
			Training: &model.Training{ID: p.TrainingID},
			Customer: &model.Customer{ID: customerID},
			Coast:    p.Coast,
			Income:   p.Income,
		})
	}

	return purchases, nil
}

// ReadPurchase is the resolver for the readPurchase field.
func (r *queryResolver) ReadPurchase(ctx context.Context, id string) (*model.Purchase, error) {
	p := model.PurchaseDB{ID: id}
	res := DBConn.First(&p)
	if res.Error != nil {
		return nil, res.Error
	}

	c := model.CustomerDB{ID: p.CustomerID}
	res = DBConn.First(&c)
	if res.Error != nil {
		return nil, res.Error
	}

	t := model.TrainingDB{ID: p.TrainingID}
	res = DBConn.First(&t)
	if res.Error != nil {
		return nil, res.Error
	}

	return &model.Purchase{
		ID: p.ID,
		Training: &model.Training{
			ID:       t.ID,
			Category: t.Category,
			Coast:    t.Coast,
			Gym:      nil,
		},
		Customer: &model.Customer{
			ID:       c.ID,
			Name:     c.Name,
			Email:    c.Email,
			Register: nil,
		},
		Coast:  p.Coast,
		Income: p.Income,
	}, nil
}
