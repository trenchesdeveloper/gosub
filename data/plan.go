package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/upper/db/v4"
)

// Plan is the type for subscription plans
type Plan struct {
	ID                  int
	PlanName            string
	PlanAmount          int
	PlanAmountFormatted string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (p *Plan) GetAll() ([]*Plan, error) {
	_, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	planCollection := dbInstance.Collection("plans")

	var plans []*Plan

	err := planCollection.Find().OrderBy("plan_amount").All(&plans)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return plans, nil
}

// GetOne returns one plan by id
func (p *Plan) GetOne(id int) (*Plan, error) {
	_, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var plan Plan

	plansCollection := dbInstance.Collection("plans")

	err := plansCollection.Find(db.Cond{"id": id}).One(&plan)

	if err != nil {
		return nil, err
	}

	return &plan, nil
}

// SubscribeUserToPlan subscribes a user to one plan by insert
// values into user_plans table
func (p *Plan) SubscribeUserToPlan(user User, plan Plan) error {
	_, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	userPlanCollection := dbInstance.Collection("user_plans")

	err := userPlanCollection.Find(db.Cond{"user_id": user.ID}).Delete()

	if err != nil {
		return err
	}

	_, err = userPlanCollection.Insert(map[string]interface{}{
		"user_id": user.ID,
		"plan_id": plan.ID,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	})

	if err != nil {
		return err
	}

	return nil
}

// AmountForDisplay formats the price we have in the DB as a currency string
func (p *Plan) AmountForDisplay() string {
	amount := float64(p.PlanAmount) / 100.0
	return fmt.Sprintf("$%.2f", amount)
}
