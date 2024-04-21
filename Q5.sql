CREATE TABLE employee (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    salary DECIMAL(10, 2) NOT NULL,
    department_id INT NOT NULL
);

CREATE TABLE department (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

ALTER TABLE employee
ADD CONSTRAINT fk_department_id
FOREIGN KEY (department_id)
REFERENCES department(id);

INSERT into department (name) VALUES
  ('IT'),
  ('Sales');

INSERT INTO employee (name, salary, department_id) VALUES
('Joe', 85000.00, 1),
('Henry', 80000.00, 2),
('Same', 60000.00, 1),
('Max', 90000.00, 2),
('Janet', 69000.00, 1),
('Randy', 85000.00, 2),
('Will', 70000.00, 2);



WITH ranked_salaries AS (
    SELECT name, salary, departmentId ,
           ROW_NUMBER() OVER (PARTITION BY department_id ORDER BY salary DESC) AS rank
    FROM employee
)
SELECT rs.*, d.department_name
FROM ranked_salaries rs
JOIN departments d ON rs.department_id = d.department_id
WHERE rs.rank <= 3;
