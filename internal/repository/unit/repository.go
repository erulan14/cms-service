package unit

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

func (r *repository) Create(ctx context.Context, deviceUUID string, model *model.CreateUnitDTO) error {
	_, err := r.Conn.Exec(ctx, `INSERT INTO 
    device (uuid, name, hw, phone, phone2, 
            tariff, icon, creator, sensors, ) VALUES ($1, $2, $3, $4)`,
		deviceUUID, model.Name, model.Company, model.Port)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) List(ctx context.Context) ([]model.Unit, error) {
	rows, _ := r.Conn.Query(ctx, `SELECT * FROM device`)
	devices, err := pgx.CollectRows(rows, pgx.RowToStructByPos[model.Unit])
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (r *repository) Get(ctx context.Context, deviceUUID string) (*model.Unit, error) {
	row, _ := r.Conn.Query(ctx, `SELECT * FROM device WHERE uuid = $1`, deviceUUID)
	device, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[model.Unit])

	if err != nil {
		return nil, err
	}
	return &device, nil
}

func (r *repository) Update(ctx context.Context, deviceUUID string, model *model.UpdateUnitDTO) error {
	_, err := r.Conn.Exec(ctx, `UPDATE device SET name = $2, company = $3, port = $4 WHERE uuid = $1`,
		deviceUUID, model.Name, model.Company, model.Port)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, deviceUUID string) error {
	_, err := r.Conn.Exec(ctx, `DELETE FROM device WHERE uuid = $1`, deviceUUID)
	if err != nil {
		return err
	}
	return nil
}
