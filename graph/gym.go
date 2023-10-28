package graph

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"gymboss/graph/model"
)

// CreateGym is the resolver for the createGym field.
func (r *mutationResolver) CreateGym(ctx context.Context, gym model.GymCreate) (*model.Gym, error) {
	id := uuid.NewV4().String()
	res := DBConn.Create(&model.GymDB{
		ID:     id,
		Branch: gym.Branch,
		Admin:  gym.Admin,
		Phone:  gym.Phone,
		Slots:  gym.Slots,
	})
	if res.Error != nil {
		return nil, res.Error
	}

	return &model.Gym{
		ID:        id,
		Branch:    gym.Branch,
		Admin:     gym.Admin,
		Phone:     gym.Phone,
		Slots:     gym.Slots,
		Trainings: nil,
	}, nil
}

// UpdateGym is the resolver for the updateGym field.
func (r *mutationResolver) UpdateGym(ctx context.Context, gym model.GymUpdate) (*model.Gym, error) {
	g := model.GymDB{ID: gym.ID}
	res := DBConn.First(&g)
	if res.Error != nil {
		return nil, res.Error
	}

	if gym.Admin != nil {
		g.Admin = *gym.Admin
	}
	if gym.Branch != nil {
		g.Branch = *gym.Branch
	}
	if gym.Phone != nil {
		g.Phone = *gym.Phone
	}
	if gym.Slots != nil {
		g.Slots = *gym.Slots
	}

	res = DBConn.Save(&g)
	if res.Error != nil {
		return nil, res.Error
	}

	var ts []model.TrainingDB
	res = DBConn.Where("gym_id = ?", g.ID).Find(&ts)
	if res.Error != nil {
		return nil, res.Error
	}

	var trainings []*model.Training
	for _, t := range ts {
		trainings = append(trainings, &model.Training{
			ID:       t.ID,
			Category: t.Category,
			Coast:    t.Coast,
			Gym:      nil,
		})
	}

	return &model.Gym{
		ID:        g.ID,
		Branch:    g.Branch,
		Admin:     g.Admin,
		Phone:     g.Phone,
		Trainings: trainings,
		Slots:     g.Slots,
	}, nil
}

// DeleteGym is the resolver for the deleteGym field.
func (r *mutationResolver) DeleteGym(ctx context.Context, id string) (*model.Gym, error) {
	g := model.GymDB{ID: id}
	res := DBConn.Delete(&g)
	if res.Error != nil {
		return nil, res.Error
	}

	return nil, nil
}

// Gyms is the resolver for the gyms field.
func (r *queryResolver) Gyms(ctx context.Context) ([]*model.Gym, error) {
	var gs []model.GymDB
	res := DBConn.Find(&gs)
	if res.Error != nil {
		return nil, res.Error
	}

	var gyms []*model.Gym
	for _, g := range gs {
		gyms = append(gyms, &model.Gym{
			ID:        g.ID,
			Branch:    g.Branch,
			Admin:     g.Admin,
			Phone:     g.Phone,
			Trainings: nil,
			Slots:     g.Slots,
		})
	}

	return gyms, nil
}

// ReadGym is the resolver for the readGym field.
func (r *queryResolver) ReadGym(ctx context.Context, id string) (*model.Gym, error) {
	g := model.GymDB{ID: id}
	res := DBConn.First(&g)
	if res.Error != nil {
		return nil, res.Error
	}

	var ts []model.TrainingDB
	res = DBConn.Where("gym_id = ?", g.ID).Find(&ts)
	if res.Error != nil {
		return nil, res.Error
	}

	var trainings []*model.Training
	for _, t := range ts {
		trainings = append(trainings, &model.Training{
			ID:       t.ID,
			Category: t.Category,
			Coast:    t.Coast,
			Gym:      nil,
		})
	}

	return &model.Gym{
		ID:        g.ID,
		Branch:    g.Branch,
		Admin:     g.Admin,
		Phone:     g.Phone,
		Trainings: trainings,
		Slots:     g.Slots,
	}, nil
}
