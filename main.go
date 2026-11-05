package main

import "fmt"

type Date struct {
	day, month, year int
}

func daysInMonth(month, year int) int {
	if month == 2 {
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			return 29
		}
		return 28
	}
	if month == 4 || month == 6 || month == 9 || month == 11 {
		return 30
	}
	return 31
}

type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate Date
	shelfLife       int
}

func (m *Medicine) SetName(name string) error {
	if name == "" {
		return fmt.Errorf("el nombre no puede estar vacío")
	}
	m.name = name
	return nil
}

func (m *Medicine) SetManufacturer(manufacturer string) error {
	if manufacturer == "" {
		return fmt.Errorf("el fabricante no puede estar vacío")
	}
	m.manufacturer = manufacturer
	return nil
}

func (m *Medicine) SetManufactureDate(day, month, year int) error {
	if year < 1 || month < 1 || month > 12 || day < 1 || day > daysInMonth(month, year) {
		return fmt.Errorf("la fecha de fabricación %d/%d/%d no es válida", day, month, year)
	}
	m.manufactureDate = Date{day, month, year}
	return nil
}

func (m *Medicine) SetShelfLife(shelfLife int) error {
	if shelfLife <= 0 {
		return fmt.Errorf("la vida útil debe ser un número positivo")
	}
	m.shelfLife = shelfLife
	return nil
}

func NewMedicine(name, manufacturer string, day, month, year, shelfLife int) (*Medicine, error) {
	m := &Medicine{}
	if err := m.SetName(name); err != nil {
		return nil, err
	}
	if err := m.SetManufacturer(manufacturer); err != nil {
		return nil, err
	}
	if err := m.SetManufactureDate(day, month, year); err != nil {
		return nil, err
	}
	if err := m.SetShelfLife(shelfLife); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *Medicine) Name() string         { return m.name }
func (m *Medicine) Manufacturer() string { return m.manufacturer }
func (m *Medicine) ManufactureDate() string {
	return fmt.Sprintf("%02d/%02d/%d", m.manufactureDate.day, m.manufactureDate.month, m.manufactureDate.year)
}
func (m *Medicine) ShelfLife() int { return m.shelfLife }
