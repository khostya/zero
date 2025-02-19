package dto

import (
	"errors"
)

type (
	Page struct {
		Size uint
		Page uint
	}

	CreateNewsParam struct {
		Title      string
		Content    string
		Categories []uint
	}

	GetNewsParam struct {
		Page *Page
	}
)

func (p Page) Offset() (uint, error) {
	if p.Page <= 0 {
		return 0, errors.New("page out of range")
	}
	return (p.Page - 1) * p.Size, nil
}

func (p Page) Limit() uint {
	return p.Size
}
