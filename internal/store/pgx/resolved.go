package pgx

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
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

func (r *ResolvedRepository) CreateResolved(ctx context.Context, data dto.Resolved) error {
	const queryUpdateLastResolved = `UPDATE resolved
SET is_active = false
WHERE user_id = $1
  AND resolved_type = $2`
	const queryCreateResolved = `INSERT INTO resolved
    (id, user_id, resolved_type, is_active, created_at, passed_at)
VALUES ($1, $2, $3, $4, $5, $6)`
	const queryCreateResolvedQuestion = `INSERT INTO resolved_questions
(resolved_id, question_order, question_text, question_answer, image_location, mark)
VALUES ($1, $2, $3, $4, $5, $6)`

	tx, err := r.store.pool.Begin(ctx)
	if err != nil {
		r.log.Error("failed to begin transaction", zap.Error(err))
		return r.store.pgErr(err)
	}
	defer tx.Rollback(ctx)

	if _, err = tx.Exec(ctx, queryUpdateLastResolved,
		data.UserID,
		data.ResolvedType,
	); err != nil {
		r.log.Error("failed to update last resolved", zap.Error(err))
		return r.store.pgErr(err)
	}

	if _, err = tx.Exec(ctx, queryCreateResolved,
		data.ID,
		data.UserID,
		data.ResolvedType,
		data.IsActive,
		data.CreatedAt,
		data.PassedAt,
	); err != nil {
		r.log.Error("failed to create resolved", zap.Error(err))
		return r.store.pgErr(err)
	}

	batch := &pgx.Batch{}
	for _, q := range data.Questions {
		batch.Queue(queryCreateResolvedQuestion,
			data.ID,
			q.QuestionOrder,
			q.Issue,
			q.QuestionAnswer,
			q.ImageLocation,
			q.Mark,
		)
	}

	br := tx.SendBatch(ctx, batch)
	defer br.Close()

	err = multierr.Combine(br.Close())
	if err != nil {
		r.log.Error("failed to update last resolved question", zap.Error(err))
		return r.store.pgErr(err)
	}

	if err = tx.Commit(ctx); err != nil {
		r.log.Error("failed to commit transaction", zap.Error(err))
		return r.store.pgErr(err)
	}

	r.log.Debug("successfully updated last resolved question and added new")
	return nil
}

func (r *ResolvedRepository) GetAllActiveResolvedByUserID(ctx context.Context, id uuid.UUID) ([]dto.Resolved, error) {
	const query = `SELECT r.id,
       r.user_id,
       r.resolved_type,
       r.is_active,
       r.created_at,
       r.passed_at,
       rq.resolved_id,
       rq.question_order,
       rq.question_text,
       rq.question_answer,
       rq.image_location,
       rq.mark
FROM resolved r
         LEFT JOIN
     resolved_questions rq ON r.id = rq.form_id
WHERE r.user_id = $1
  AND r.is_active = TRUE;`

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
			&resolved.UserID,
			&resolved.IsActive,
			&resolved.CreatedAt,
			&resolved.PassedAt,
			&resolved.Questions,
		)
		if err != nil {
			r.log.Error("failed to retrieve resolved questions", zap.Error(err))
			return nil, r.store.pgErr(err)
		}
		res = append(res, resolved)
	}
	if err := rows.Err(); err != nil {
		r.log.Error("failed to retrieve resolved questions", zap.Error(err))
		return nil, r.store.pgErr(err)
	}

	r.log.Debug("successfully retrieved all active resolved questions")
	return res, nil
}

func (r *ResolvedRepository) GetResolvedByUserID(ctx context.Context, id uuid.UUID, resolved_type string, isActive bool) (*dto.Resolved, error) {
	const query = `SELECT 
   		r.id, r.user_id, r.resolved_type, r.is_active, r.created_at, r.passed_at, 
    	rq.resolved_id, rq.question_order, rq.question_text, rq.answer, rq.image_location, rq.mark
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
		&res.UserID,
		&res.ResolvedType,
		&res.IsActive,
		&res.CreatedAt,
		&res.PassedAt,
		&res.Questions,
	)

	if err != nil {
		r.log.Error("failed to retrieve resolved questions", zap.Error(err))
		return nil, r.store.pgErr(err)
	}

	r.log.Debug("successfully retrieved resolved by user", zap.Any("data", res))
	return &res, nil
}

func (r *ResolvedRepository) GetResolvedByID(ctx context.Context, id uuid.UUID) (*dto.Resolved, error) {
	const query = `SELECT * FROM resolved WHERE id = $1;`
	var res dto.Resolved
	err := r.store.pool.QueryRow(ctx, query, id).Scan(
		&res.ID,
		&res.UserID,
		&res.ResolvedType,
		&res.IsActive,
		&res.CreatedAt,
		&res.PassedAt,
		&res.Questions)

	if err != nil {
		r.log.Error("failed to retrieve resolved questions", zap.Error(err))
		return nil, r.store.pgErr(err)
	}

	return &res, nil
}
