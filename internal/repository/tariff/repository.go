package tariff

import (
	"cms-service/internal/model"
	def "cms-service/internal/repository"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ def.TariffRepository = (*repository)(nil)

type repository struct {
	Conn *pgxpool.Pool
}

func NewRepository(conn *pgxpool.Pool) def.TariffRepository {
	return &repository{Conn: conn}
}

func (r *repository) Create(ctx context.Context, tariffUUID string, model *model.CreateTariffDTO) error {
	_, err := r.Conn.Exec(ctx, `INSERT INTO tariff (uuid, name, cost, curr) VALUES ($1, $2, $3, $4)`,
		tariffUUID, model.Name, model.Cost, model.Curr)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) List(ctx context.Context) ([]model.Tariff, error) {
	rows, _ := r.Conn.Query(ctx, `SELECT * FROM tariff`)
	tariffs, err := pgx.CollectRows(rows, pgx.RowToStructByPos[model.Tariff])
	if err != nil {
		return nil, err
	}
	return tariffs, nil
}

func (r *repository) Get(ctx context.Context, tariffUUID string) (*model.Tariff, error) {
	row, _ := r.Conn.Query(ctx, `SELECT * FROM tariff WHERE uuid = $1`, tariffUUID)
	tariff, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[model.Tariff])
	if err != nil {
		return nil, err
	}
	return &tariff, nil
}

func (r *repository) Update(ctx context.Context, tariffUUID string, model *model.UpdateTariffDTO) error {
	_, err := r.Conn.Exec(ctx, `UPDATE tariff SET name = $2, cost = $3, curr = $4 WHERE uuid = $1`,
		tariffUUID, model.Name, model.Cost, model.Curr)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, tariffUUID string) error {
	_, err := r.Conn.Exec(ctx, `DELETE FROM tariff WHERE uuid = $1`, tariffUUID)
	if err != nil {
		return err
	}
	return nil
}
