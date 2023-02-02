package model

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/babyname/fate/ent"
	"github.com/babyname/fate/logger"
)

type Model struct {
	*ent.Client
}

func (m Model) Initialize(ctx context.Context, luckies <-chan *ent.WuGeLucky) error {
	err := m.Schema.Create(ctx)
	if err != nil {
		return err
	}
	var tmp []*ent.WuGeLuckyCreate
	var id int
	var count int
	for lucky := range luckies {
		id = WuGeLuckyID(lucky.LastStroke1, lucky.LastStroke2, lucky.FirstStroke1, lucky.FirstStroke2)
		tmp = append(tmp, m.WuGeLucky.Create().SetWuGeLuckyWithOptional(lucky).SetID(id))
		count++
		if len(tmp) >= PerInitStep {
			log.Info("insert into wugelucky", "count", count)
			_, err := m.insertWuGeLucky(ctx, tmp)
			if err != nil {
				return err
			}
			tmp = nil
		}
	}
	if len(tmp) != 0 {
		log.Info("insert into wugelucky", "count", count)
		_, err := m.insertWuGeLucky(ctx, tmp)
		if err != nil {
			return err
		}
		tmp = nil
	}
	return nil
}

func (m Model) insertWuGeLucky(ctx context.Context, tmp []*ent.WuGeLuckyCreate) ([]*ent.WuGeLucky, error) {
	return m.WuGeLucky.CreateBulk(tmp...).Save(ctx)
}

// ID ...
func ID(name string) string {
	sum := md5.Sum([]byte(name))
	return hex.EncodeToString(sum[:])
}

// New ...
// @param *ent.Client
// @return *Model
func New(client *ent.Client) *Model {
	log = logger.WithGroup("Model")
	return &Model{Client: client}
}