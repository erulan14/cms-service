package geozone

import (
	"cms-service/internal/model"
	def "cms-service/internal/repository"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ def.UnitRepository = (*repository)(nil)

type repository struct {
	Conn *pgxpool.Pool
}

func NewRepository(conn *pgxpool.Pool) def.UnitRepository {
	return &repository{Conn: conn}
}

func (r *repository) Create(ctx context.Context, geozoneUUID string, model *model.CreateUnitDTO) error {

	return nil
}

func (r *repository) List(ctx context.Context) ([]model.Unit, error) {
	rows, _ := r.Conn.Query(ctx, `SELECT * FROM geozone`)
	geozones, err := pgx.CollectRows(rows, pgx.RowToStructByPos[model.Unit])
	if err != nil {
		return nil, err
	}
	return geozones, nil
}

func (r *repository) Get(ctx context.Context, geozoneUUID string) (*model.Unit, error) {
	row, _ := r.Conn.Query(ctx, `SELECT * FROM geozone WHERE uuid = $1`, geozoneUUID)
	geozone, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[model.Unit])

	if err != nil {
		return nil, err
	}
	return &geozone, nil
}

func (r *repository) Update(ctx context.Context, geozoneUUID string, model *model.UpdateUnitDTO) error {

	return nil
}

func (r *repository) Delete(ctx context.Context, geozoneUUID string) error {
	_, err := r.Conn.Exec(ctx, `DELETE FROM geozone WHERE uuid = $1`, geozoneUUID)
	if err != nil {
		return err
	}
	return nil
}
