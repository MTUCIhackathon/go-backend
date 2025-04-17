package pgx

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

type ResultsRepository struct {
	log   *zap.Logger
	store *Store
}

func newResultsRepository(store *Store) *ResultsRepository {
	return &ResultsRepository{
		log:   store.log.Named("results-repository"),
		store: store,
	}
}

func (r *ResultsRepository) GetResultByResolvedIDAndUserID(ctx context.Context, userID uuid.UUID, resolvedID uuid.UUID) (*dto.Result, error) {
	const query = `SELECT
       t.id,
       t.user_id,
       t.resolved_id,
       t.image_location,
       t.profession,
       t.created_at
FROM test_results t
         JOIN resolved r ON r.id = t.resolved_id
WHERE t.user_id = $1
  AND t.resolved_id = $2
  AND r.is_active = true;`

	var data dto.Result
	err := r.store.pool.QueryRow(ctx, query, userID, resolvedID, true).Scan(
		&data.ID,
		&data.UserID,
		&data.ResolvedID,
		&data.ImageLocation,
		&data.Profession,
		&data.CreatedAt,
	)
	if err != nil {
		r.log.Error("failed to retrieve last result", zap.Error(err))
		return nil, r.store.pgErr(err)
	}
	r.log.Debug("retrieved last result", zap.Any("data", data))
	return &data, nil
}

func (r *ResultsRepository) GetResultByUserID(ctx context.Context, userID uuid.UUID) ([]dto.Result, error) {
	const query = `SELECT
       t.id,
       t.user_id,
       t.resolved_id,
       t.image_location,
       t.profession,
       t.created_at
FROM test_results t
         JOIN resolved r ON r.id = t.resolved_id
WHERE t.user_id = $1
  AND r.is_active = $2;`

	var data []dto.Result

	rows, err := r.store.pool.Query(ctx, query, userID, true)
	if err != nil {
		r.log.Error("failed to retrieve last result", zap.Error(err))
		return nil, r.store.pgErr(err)
	}
	defer rows.Close()
	for rows.Next() {
		var result dto.Result
		err = rows.Scan(
			&result.ID,
			&result.UserID,
			&result.ResolvedID,
			&result.ImageLocation,
			&result.Profession,
			&result.CreatedAt,
		)
		if err != nil {
			r.log.Error("failed to retrieve last result", zap.Error(err))
			return nil, r.store.pgErr(err)
		}
		r.log.Debug("retrieved last result", zap.Any("result", result))

		data = append(data, result)
	}

	if err := rows.Err(); err != nil {
		r.log.Error("failed to retrieve last result", zap.Error(err))
		return nil, r.store.pgErr(err)
	}

	r.log.Debug("retrieved last result", zap.Any("data", data))
	return data, nil
}

func (r *ResultsRepository) DeleteResult(ctx context.Context, resultID uuid.UUID) error {
	const query = `DELETE FROM test_results WHERE id = $1;`
	result, err := r.store.pool.Exec(ctx, query, resultID)
	if err != nil {
		r.log.Error("failed to delete last result", zap.Error(err))
		return r.store.pgErr(err)
	}

	if result.RowsAffected() != 1 {
		r.log.Error("failed to delete last result", zap.Any("result", result))
		return r.store.pgErr(err)
	}
	return nil
}

func (r *ResultsRepository) CreateResult(ctx context.Context, result dto.Result) error {
	const query = `INSERT INTO test_results (id, user_id, resolved_id, image_location, profession, created_at)
VALUES ($1, $2, $3, $4, $5, $6);`
	commandTag, err := r.store.pool.Exec(ctx, query,
		result.ID,
		result.UserID,
		result.ResolvedID,
		result.ImageLocation,
		result.Profession,
		result.CreatedAt,
	)
	if err != nil {
		r.log.Error("failed to insert last result", zap.Error(err))
		return r.store.pgErr(err)
	}
	if commandTag.RowsAffected() != 1 {
		r.log.Error("failed to insert last result", zap.Any("result", result))
		return r.store.pgErr(err)
	}
	r.log.Debug("inserted last result", zap.Any("result", result))
	return nil

}

func (r *ResultsRepository) GetAllResults(ctx context.Context, userID uuid.UUID) ([][]string, error) {
	const query = `SELECT t.profession
FROM test_results t
         JOIN resolved r ON r.id = t.resolved_id
WHERE t.user_id = $1 AND r.is_active = true;`
	var (
		allProfessions [][]string
		professions    []string
	)

	rows, err := r.store.pool.Query(ctx, query, userID)
	if err != nil {
		r.log.Error("failed to retrieve all results", zap.Error(err))
		return nil, r.store.pgErr(err)
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&professions)
		if err != nil {
			r.log.Error("failed to retrieve all results", zap.Error(err))
			return nil, r.store.pgErr(err)
		}
		allProfessions = append(allProfessions, professions)
	}

	if err := rows.Err(); err != nil {
		r.log.Error("failed to retrieve all results", zap.Error(err))
		return nil, r.store.pgErr(err)
	}
	r.log.Debug("retrieved all results", zap.Any("results", allProfessions))
	return allProfessions, nil
}

func (r *ResultsRepository) SetImageToResult(ctx context.Context, imageLocation string, resultID uuid.UUID) (bool, error) {
	const query = `UPDATE test_results SET image_location = $1 WHERE id = $2;`

	commandTag, err := r.store.pool.Exec(ctx, query, imageLocation, resultID)
	if err != nil {
		r.log.Error("failed to update image to result", zap.Error(err))
		return false, r.store.pgErr(err)
	}

	if commandTag.RowsAffected() == 0 {
		r.log.Error("failed to update image to result", zap.Any("result", resultID))
		return false, ErrZeroRowsAffected
	}

	return true, nil
}
