package pgx

import (
	"context"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ResolvedRepository struct {
	log   *zap.Logger
	store *Store
}

func newResolvedRepository(store *Store) *ResolvedRepository {
	return &ResolvedRepository{
		log:   store.log.Named("resolved-repository"),
		store: store,
	}
}

func (r *ResolvedRepository) CreateResolved(ctx context.Context, data dto.Resolved) (*dto.Resolved, error) {
	const queryUpdateLastResolved = `UPDATE resolved SET is_active = false WHERE user_id = $1 AND resolved_type = $2`
	const queryCreateResolved = `INSERT INTO resolved (id, user_id, resolved_type, is_active, created_at, passed_at)VALUES($1, $2, $3, $4, $5)`
	const queryCreateResolvedQuestion = `INSERT INTO resolved_question(form_id, question_text, image_location, mark)`

	tx, err := r.store.pool.Begin(ctx)
	if err != nil {
		r.log.Error("failed to begin transaction", zap.Error(err))
		return nil, r.store.pgErr(err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, queryUpdateLastResolved,
		data.UserID,
		data.ResolvedType,
	)
	if err != nil {
		r.log.Error("failed to update last resolved question", zap.Error(err))
		return nil, r.store.pgErr(err)
	}

	_, err = tx.Exec(ctx, queryCreateResolved,
		data.ID,
		data.UserID,
		data.ResolvedType,
		data.IsActive,
		data.CreatedAt,
		data.PassedAt,
	)
	if err != nil {
		r.log.Error("failed to update last resolved question", zap.Error(err))
		return nil, r.store.pgErr(err)
	}

	for _, q := range data.Questions {
		_, err = tx.Exec(ctx, queryCreateResolvedQuestion,
			q.FromID,
			q.Issue,
			q.ImageLocation,
			q.Mark,
		)
		if err != nil {
			r.log.Error("failed to update last resolved question", zap.Error(err))
			return nil, r.store.pgErr(err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		r.log.Error("failed to commit transaction", zap.Error(err))
		return nil, r.store.pgErr(err)
	}

	return &data, nil
}

func (r *ResolvedRepository) GetAllActiveResolvedByUserID(ctx context.Context, id uuid.UUID) ([]dto.Resolved, error) {
	const query = `SELECT 
   		 r.*, 
    	rq.*
		FROM 
    	resolved r
		LEFT JOIN 
    	resolved_question rq ON r.id = rq.form_id
		WHERE 
    	r.user_id = $1 
    	AND r.is_active = TRUE; `

	var res []dto.Resolved
	rows, err := r.store.pool.Query(ctx, query, id)
	if err != nil {
		r.log.Error("failed to retrieve resolved questions", zap.Error(err))
		return nil, r.store.pgErr(err)
	}
	defer rows.Close()

	for rows.Next() {
		var resolved dto.Resolved
		err = rows.Scan(
			&resolved.ID,
			&resolved.Version,
			&resolved.UserID,
			&resolved.IsActive,
			&resolved.CreatedAt,
			&resolved.PassedAt,
			&resolved.Questions,
		)
		if err != nil {
			r.log.Debug("failed to retrieve resolved questions", zap.Error(err))
			return nil, r.store.pgErr(err)
		}
		res = append(res, resolved)
	}
	if err := rows.Err(); err != nil {
		r.log.Debug("failed to retrieve resolved questions", zap.Error(err))
		return nil, r.store.pgErr(err)
	}

	return res, nil
}

func (r *ResolvedRepository) GetResolvedByUserID(ctx context.Context, id uuid.UUID, resolved_type string, isActive bool) (*dto.Resolved, error) {
	const query = `SELECT 
   		 r.*, 
    	rq.*
		FROM 
    	resolved r
		LEFT JOIN 
    	resolved_question rq ON r.id = rq.form_id
		WHERE 
    	r.user_id = $1 
    	AND r.is_active = $2
		AND r.resolved_type = $3;`

	var res dto.Resolved

	err := r.store.pool.QueryRow(ctx, query, id, isActive, resolved_type).Scan(
		&res.ID,
		&res.Version,
		&res.UserID,
		&res.ResolvedType,
		&res.IsActive,
		&res.CreatedAt,
		&res.PassedAt,
		&res.Questions,
	)

	if err != nil {
		r.log.Debug("failed to retrieve resolved questions", zap.Error(err))
		return nil, r.store.pgErr(err)
	}
	return &res, nil
}

func (r *ResolvedRepository) GetResolvedByID(ctx context.Context, id uuid.UUID) (*dto.Resolved, error) {
	const query = `SELECT * FROM resolved WHERE id = $1;`
	var res dto.Resolved
	err := r.store.pool.QueryRow(ctx, query, id).Scan(
		&res.ID,
		&res.Version,
		&res.UserID,
		&res.ResolvedType,
		&res.IsActive,
		&res.CreatedAt,
		&res.PassedAt,
		&res.Questions)

	if err != nil {
		r.log.Debug("failed to retrieve resolved questions", zap.Error(err))
		return nil, r.store.pgErr(err)
	}

	return &res, nil
}
