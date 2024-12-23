package repositories

import (
    "database/sql"
    "fmt"
)

type Charge struct {
    ID            int     `json:"id"`
    Concept       string  `json:"concept"`
    Amount        float64 `json:"amount"`
    PaymentMethod string  `json:"payment_method"`
    Category      string  `json:"category"`
    Date          string  `json:"date"`
}

type ChargeRepository struct {
    DB *sql.DB
}

func NewChargeRepository(db *sql.DB) *ChargeRepository {
    return &ChargeRepository{DB: db}
}

func (r *ChargeRepository) CreateCharge(charge *Charge) error {
    query := `INSERT INTO charges (concept, amount, payment_method, category, date) VALUES (?, ?, ?, ?, ?)`
    result, err := r.DB.Exec(query, charge.Concept, charge.Amount, charge.PaymentMethod, charge.Category, charge.Date)
    if err != nil {
        return fmt.Errorf("failed to insert charge: %w", err)
    }
    id, _ := result.LastInsertId()
    charge.ID = int(id)
    return nil
}

func (r *ChargeRepository) GetAllCharges() ([]Charge, error) {
    query := `SELECT id, concept, amount, payment_method, category, date FROM charges`
    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, fmt.Errorf("failed to query charges: %w", err)
    }
    defer rows.Close()

    var charges []Charge
    for rows.Next() {
        var charge Charge
        if err := rows.Scan(&charge.ID, &charge.Concept, &charge.Amount, &charge.PaymentMethod, &charge.Category, &charge.Date); err != nil {
            return nil, fmt.Errorf("failed to scan charge: %w", err)
        }
        charges = append(charges, charge)
    }
    return charges, nil
}

func (r *ChargeRepository) GetChargeByID(id int) (*Charge, error) {
    query := `SELECT id, concept, amount, payment_method, category, date FROM charges WHERE id = ?`
    row := r.DB.QueryRow(query, id)

    var charge Charge
    if err := row.Scan(&charge.ID, &charge.Concept, &charge.Amount, &charge.PaymentMethod, &charge.Category, &charge.Date); err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("charge not found")
        }
        return nil, fmt.Errorf("failed to scan charge: %w", err)
    }
    return &charge, nil
}

func (r *ChargeRepository) UpdateCharge(charge *Charge) error {
    query := `UPDATE charges SET concept = ?, amount = ?, payment_method = ?, category = ?, date = ? WHERE id = ?`
    _, err := r.DB.Exec(query, charge.Concept, charge.Amount, charge.PaymentMethod, charge.Category, charge.Date, charge.ID)
    if err != nil {
        return fmt.Errorf("failed to update charge: %w", err)
    }
    return nil
}

func (r *ChargeRepository) DeleteCharge(id int) error {
    query := `DELETE FROM charges WHERE id = ?`
    _, err := r.DB.Exec(query, id)
    if err != nil {
        return fmt.Errorf("failed to delete charge: %w", err)
    }
    return nil
}
