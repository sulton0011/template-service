package postgres

import (
	pb "home_work/task-service/genproto/task"
	"time"

	"github.com/google/uuid"
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
	Task.Id = uuid.New().String()

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
            assignee,
            title,
            summary,
            deadline,
            status
        FROM tasks
        WHERE id = $1
    `
	var task pb.Task
	err := r.db.DB.QueryRow(query,
		id,
	).Scan(
		&task.Assignee,
		&task.Title,
		&task.Summary,
		&task.Deadline,
		&task.Status,
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
			assignee,
			title,
			summary,
			deadline,
			status
		FROM tasks
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
			&task.Assignee,
			&task.Title,
			&task.Summary,
			&task.Deadline,
			&task.Status,
		)
		if err != nil {
			return nil, err
		}

		resp.Tasks = append(resp.Tasks, &task)
	}
	query = `
		SELECT count(*) FROM tasks
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
        RETURNING id
    `
	Task.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	err := r.db.DB.QueryRow(query,
		Task.Assignee,
		Task.Title,
		Task.Summary,
		Task.Deadline,
		Task.Status,
		Task.UpdatedAt,

		Task.Id,
	).Scan(&Task.Id)
	if err != nil {
		return nil, err
	}
	return r.Get(Task.Id)
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
