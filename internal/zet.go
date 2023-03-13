package zet

import (
	"fmt"
	"time"
)

type Zettel struct {
	ID        int
	Content   string
	Tags      []string
	CreatedAt time.Time
}

func New(content string, tags []string) *Zettel {
	z := &Zettel{
		Content:   content,
		Tags:      tags,
		CreatedAt: time.Now(),
	}
	// TODO: gen uuid for each zet
	return z
}

func (z *Zettel) Add() error {
	// TODO: implement this
	fmt.Printf("Here's your zet: %s\n", z.Content)
	return nil
}

type ZettelRepo struct {
	// TODO
}

func NewZettelRepo() *ZettelRepo {
	// TODO
	return &ZettelRepo{}
}

func (zr *ZettelRepo) FindByID(id int) (*Zettel, error) {
	// TODO
	return nil, fmt.Errorf("no zettel with ID %d found", id)
}

func (zr *ZettelRepo) FindByTag(tag string) ([]*Zettel, error) {
	// TODO
	return nil, fmt.Errorf("no zettls with tag %s found", tag)
}
