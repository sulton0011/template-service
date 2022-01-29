CREATE TABLE tasks (
    id UUID PRIMARY KEY,
    assignee VARCHAR(124),
    title VARCHAR(124),
    summary VARCHAR(124),
    deadline TIMESTAMP,
    status VARCHAR(124),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- create or replace function modify_employee_info() returns trigger
-- language plpgsql as
-- $$
--   begin 
--   new.modifyed_date = current_date;
--   return new;
--   end;
-- $$;

-- create trigger bring_modification before update on hl_employees
-- for each row execute function modify_employee_info();