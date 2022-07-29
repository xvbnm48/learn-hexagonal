package api

import "github.com/xvbnm48/learn-hexagonal/internal/ports"

type Adapter struct {
	db    ports.DBPort
	arith ports.ArithmaticPort
}

func NewAdapter(db ports.DBPort, arith ports.ArithmaticPort) *Adapter {
	return &Adapter{db: db, arith: arith}
}

func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apia.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddHistory(answer, "addition")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apia.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddHistory(answer, "multiplication")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (apia Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := apia.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddHistory(answer, "subtraction")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apia.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddHistory(answer, "division")
	if err != nil {
		return 0, err
	}

	return answer, nil
}
