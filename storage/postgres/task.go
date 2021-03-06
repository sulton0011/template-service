package postgres

import (
	"database/sql"
	pb "home_work/task-service/genproto/task"
	"time"

	
	"github.com/jmoiron/sqlx"
)

type TaskRepo struct {
	db *sqlx.DB
}

//NewTaskRepo ...
func NewTaskRepo(db *sqlx.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

// Create task ...
func (r *TaskRepo) Create(Task *pb.Task) (*pb.Task, error) {
	query := `
        INSERT INTO 
			tasks (
                id,
                assignee,
                title,
                summary,
                deadline,
                status,
                created_at
            )
        VALUES(
            $1,
            $2,
            $3,
            $4,
            $5,
            $6,
            $7
        )
        RETURNING id
    `
	Task.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	err := r.db.DB.QueryRow(query,
		Task.Id,
		Task.Assignee,
		Task.Title,
		Task.Summary,
		Task.Deadline,
		Task.Status,
		Task.CreatedAt,
	).Scan(&Task.Id)
	if err != nil {
		return nil, err
	}
	return r.Get(Task.Id)
}

// Get task ...
func (r *TaskRepo) Get(id string) (*pb.Task, error) {
	query := `
        SELECT
			id,
            assignee,
            title,
            summary,
            deadline,
            status,
			created_at
        FROM tasks
        WHERE id = $1
		AND deleted_at IS NULL
    `
	var task pb.Task
	err := r.db.DB.QueryRow(query,
		id,
	).Scan(
		&task.Id,
		&task.Assignee,
		&task.Title,
		&task.Summary,
		&task.Deadline,
		&task.Status,
		&task.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// List task ...
func (r *TaskRepo) List(req *pb.ListReq) (*pb.ListResp, error) {
	offset := (req.Page - 1) * req.Limit
	var resp pb.ListResp
	query := `
		SELECT 
			id,
			assignee,
			title,
			summary,
			deadline,
			status,
			created_at
		FROM tasks
		WHERE deleted_at IS NULL
		LIMIT $1
		OFFSET $2
	`
	rows, err := r.db.DB.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task pb.Task
		err = rows.Scan(
			&task.Id,
			&task.Assignee,
			&task.Title,
			&task.Summary,
			&task.Deadline,
			&task.Status,
			&task.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Tasks = append(resp.Tasks, &task)
	}
	query = `
		SELECT count(*) FROM tasks
		WHERE deleted_at IS NULL
	`
	err = r.db.DB.QueryRow(query).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Update task ...
func (r *TaskRepo) Update(Task *pb.Task) (*pb.Task, error) {
	query := `
		UPDATE tasks
		SET
            assignee = $1,
            title = $2,
            summary = $3,
            deadline = $4,
            status = $5,
            updated_at = $6
        WHERE id = $7
		AND deleted_at IS NULL
        RETURNING id
    `
	Task.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	result,err := r.db.DB.Exec(query,
		Task.Assignee,
		Task.Title,
		Task.Summary,
		Task.Deadline,
		Task.Status,
		Task.UpdatedAt,

		Task.Id,
	)
	if err != nil {
		return nil, err
	}
	i,_:=result.RowsAffected();
	if i==0 {
		return nil,sql.ErrNoRows
	}
	task,err := r.Get(Task.Id)
	if err != nil {
		return nil,err
	}
	return task,nil
}

// Delete task ...
func (r *TaskRepo) Delete(id *pb.IdReq) (*pb.EmptyResp, error) {
	query := `
		UPDATE tasks
		SET 
			deleted_at = $1
		WHERE id = $2
	`
	newTime := time.Now().Format("2006-01-02 15:04:05")
	_, err := r.db.DB.Exec(query, newTime, id.Id)
	if err != nil {
		return nil, err
	}
	return &pb.EmptyResp{}, nil
}

// ListOverdue task ...
func (r *TaskRepo) ListOverdue(req *pb.ListOverReq) (*pb.ListOverResp, error) {
	offset := (req.Page - 1) * req.Limit
	var resp pb.ListOverResp
	query := `
		SELECT
			id,
			assignee,
			title,
			summary,
			deadline,
			status,
			created_at
		FROM tasks
		WHERE deadline < $1
		AND deleted_at IS NULL
		LIMIT $2
		OFFSET $3
	`
	rows, err := r.db.DB.Query(query, req.Time, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task pb.Task
		err = rows.Scan(
			&task.Id,
			&task.Assignee,
			&task.Title,
			&task.Summary,
			&task.Deadline,
			&task.Status,
			&task.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Tasks = append(resp.Tasks, &task)
	}
	query = `
		SELECT count(*) FROM tasks 
		WHERE deadline < $1
		AND deleted_at IS NULL
	`
	err = r.db.DB.QueryRow(query, req.Time).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}