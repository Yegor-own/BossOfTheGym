package graph

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"gymboss/graph/model"
)

// CreateTraining is the resolver for the createTraining field.
func (r *mutationResolver) CreateTraining(ctx context.Context, training model.TrainingCreate) (*model.Training, error) {

	id := uuid.NewV4().String()

	res := DBConn.Create(&model.TrainingDB{
		ID:       id,
		Category: training.Category,
		Coast:    training.Coast,
		GymID:    training.GymID,
	})
	if res.Error != nil {
		return nil, res.Error
	}

	g := model.GymDB{ID: training.GymID}
	res = DBConn.First(&g)
	if res.Error != nil {
		return nil, res.Error
	}

	return &model.Training{
		ID:       id,
		Category: training.Category,
		Coast:    training.Coast,
		Gym: &model.Gym{
			ID:        "",
			Branch:    g.Branch,
			Admin:     g.Admin,
			Phone:     g.Phone,
			Trainings: nil,
			Slots:     g.Slots,
		},
	}, nil
}

// UpdateTraining is the resolver for the updateTraining field.
func (r *mutationResolver) UpdateTraining(ctx context.Context, training model.TrainingUpdate) (*model.Training, error) {
	t := model.TrainingDB{ID: training.ID}
	res := DBConn.First(&t)
	if res.Error != nil {
		return nil, res.Error
	}

	if training.Coast != nil {
		t.Coast = *training.Coast
	}
	if training.Category != nil {
		t.Category = *training.Category
	}
	if training.GymID != nil {
		t.GymID = *training.GymID
	}

	res = DBConn.Save(&t)
	if res.Error != nil {
		return nil, res.Error
	}

	g := model.GymDB{ID: t.GymID}
	res = DBConn.First(&g)
	if res.Error != nil {
		return nil, res.Error
	}

	return &model.Training{
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
	}, nil
}

// DeleteTraining is the resolver for the deleteTraining field.
func (r *mutationResolver) DeleteTraining(ctx context.Context, id string) (*model.Training, error) {
	t := model.TrainingDB{ID: id}
	res := DBConn.Delete(&t)
	if res.Error != nil {
		return nil, res.Error
	}
	return nil, nil
}

// Trainings is the resolver for the trainings field.
func (r *queryResolver) Trainings(ctx context.Context) ([]*model.Training, error) {
	ts := []model.TrainingDB{}
	res := DBConn.First(&ts)
	if res.Error != nil {
		return nil, res.Error
	}

	var trainings []*model.Training
	for _, t := range ts {
		g := model.GymDB{ID: t.GymID}
		DBConn.First(&g)
		trainings = append(trainings, &model.Training{
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
		})
	}

	return trainings, nil

}

// ReadTraining is the resolver for the readTraining field.
func (r *queryResolver) ReadTraining(ctx context.Context, id string) (*model.Training, error) {
	t := model.TrainingDB{ID: id}
	res := DBConn.First(&t)
	if res.Error != nil {
		return nil, res.Error
	}

	g := model.GymDB{ID: t.GymID}
	res = DBConn.First(&g)
	if res.Error != nil {
		return nil, res.Error
	}

	return &model.Training{
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
	}, nil
}
