package pgx

import (
	"context"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"github.com/google/uuid"
	"go.uber.org/zap"
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

func (r *ResultsRepository) GetLastResultByFormId(ctx context.Context, userID uuid.UUID, formID uuid.UUID) (*dto.Result, error) {
	const query = `SELECT 
    	t.user_id, 
    	t.resolved_id, 
    	t.resolved_version, 
    	t.profession,
    	t.created_at
		FROM test_results t JOIN resolved r ON r.id = t.resolved_id
		WHERE t.user_id = $1 AND t.resolved_id = $2 AND r.is_active = true;`

	var data dto.Result
	err := r.store.pool.QueryRow(ctx, query, userID, formID, true).Scan(
		&data.UserID,
		&data.ResolvedID,
		&data.ResolvedVersion,
		&data.Profession,
		&data.CreatedAt,
	)
	if err != nil {
		r.log.Debug("failed to retrieve last result", zap.Error(err))
		return nil, r.store.pgErr(err)
	}
	r.log.Debug("retrieved last result", zap.Any("data", data))
	return &data, nil
}

func (r *ResultsRepository) GetLastResults(ctx context.Context, userID uuid.UUID) ([]dto.Result, error) {
	const query = `SELECT 
    	t.user_id, 
    	t.resolved_id, 
    	t.resolved_version, 
    	t.profession,
    	t.created_at
		FROM test_results t JOIN resolved r ON r.id = t.resolved_id
		WHERE t.user_id = $1 AND t.is_active = $2;`

	var data []dto.Result

	rows, err := r.store.pool.Query(ctx, query, userID, true)
	if err != nil {
		r.log.Debug("failed to retrieve last result", zap.Error(err))
		return nil, r.store.pgErr(err)
	}
	defer rows.Close()
	for rows.Next() {
		var result dto.Result
		err = rows.Scan(
			&result.UserID,
			&result.ResolvedID,
			&result.ResolvedVersion,
			&result.Profession,
			&result.CreatedAt,
		)
		if err != nil {
			r.log.Debug("failed to retrieve last result", zap.Error(err))
			return nil, r.store.pgErr(err)
		}
		r.log.Debug("retrieved last result", zap.Any("result", result))

		data = append(data, result)
	}

	if err := rows.Err(); err != nil {
		r.log.Debug("failed to retrieve last result", zap.Error(err))
		return nil, r.store.pgErr(err)
	}

	r.log.Debug("retrieved last result", zap.Any("data", data))
	return data, nil
}

func (r *ResultsRepository) DeleteResult(ctx context.Context, resultID uuid.UUID) error {
	const query = `DELETE FROM test_results WHERE id = $1;`
	result, err := r.store.pool.Exec(ctx, query, resultID)
	if err != nil {
		r.log.Debug("failed to delete last result", zap.Error(err))
		return r.store.pgErr(err)
	}

	if result.RowsAffected() != 1 {
		r.log.Debug("failed to delete last result", zap.Any("result", result))
		return r.store.pgErr(err)
	}
	return nil
}

func (r *ResultsRepository) InsertResult(ctx context.Context, result dto.Result) error {
	const query = `INSERT INTO test_results (user_id, resolved_id, resolved_version, profession, created_at) VALUES ($1, $2, $3, $4);`
	commandTag, err := r.store.pool.Exec(ctx, query,
		result.UserID,
		result.ResolvedID,
		result.ResolvedVersion,
		result.Profession,
		result.CreatedAt,
	)
	if err != nil {
		r.log.Debug("failed to insert last result", zap.Error(err))
		return r.store.pgErr(err)
	}
	if commandTag.RowsAffected() != 1 {
		r.log.Debug("failed to insert last result", zap.Any("result", result))
		return r.store.pgErr(err)
	}
	r.log.Debug("inserted last result", zap.Any("result", result))
	return nil

}
