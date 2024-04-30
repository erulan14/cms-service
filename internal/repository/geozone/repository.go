package geozone

import (
	"cms-service/internal/model"
	def "cms-service/internal/repository"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ def.GeozoneRepository = (*repository)(nil)

type repository struct {
	Conn *pgxpool.Pool
}

func NewRepository(conn *pgxpool.Pool) def.GeozoneRepository {
	return &repository{Conn: conn}
}

func (r *repository) Create(ctx context.Context, geozoneUUID string, model *model.CreateGeozoneDTO) error {
	_, err := r.Conn.Exec(ctx, `INSERT INTO geozone (uuid, name, positions, createdat,updatedat,type,color,radius,width,description,perimeter,area)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
		geozoneUUID, model.Name, model.Positions, model.CreatedAt, model.UpdatedAt, model.Type, model.Color, model.Radius,
		model.Width, model.Description, model.Perimeter, model.Area)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) List(ctx context.Context) ([]model.Geozone, error) {
	rows, _ := r.Conn.Query(ctx, `SELECT * FROM geozone`)
	geozones, err := pgx.CollectRows(rows, pgx.RowToStructByPos[model.Geozone])
	if err != nil {
		return nil, err
	}
	return geozones, nil
}

func (r *repository) Get(ctx context.Context, geozoneUUID string) (*model.Geozone, error) {
	row, _ := r.Conn.Query(ctx, `SELECT * FROM geozone WHERE uuid = $1`, geozoneUUID)
	geozone, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[model.Geozone])

	if err != nil {
		return nil, err
	}
	return &geozone, nil
}

func (r *repository) Update(ctx context.Context, geozoneUUID string, model *model.UpdateGeozoneDTO) error {
	_, err := r.Conn.Exec(ctx, `UPDATE geozone SET name = $2, positions = $3, updatedat = $4, type = $5, color = $6, radius = $7, width = $8, description = $9, perimeter = $10, area = $11 WHERE uuid = $1`,
		geozoneUUID, model.Name, model.Positions, model.UpdatedAt, model.Type, model.Color, model.Radius,
		model.Width, model.Description, model.Perimeter, model.Area)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, geozoneUUID string) error {
	_, err := r.Conn.Exec(ctx, `DELETE FROM geozone WHERE uuid = $1`, geozoneUUID)
	if err != nil {
		return err
	}
	return nil
}
