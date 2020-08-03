# Go API Server

Account Detail API

- API version: 2.0
- Build date: 2020-04-10T14:51:30.493Z

## Pre-requisite
1. Install golang v1.13 or above.
2. Basic understanding of the golang syntax.
3. Basic understanding of SQL query.
4. Code Editor (I recommend to use VS Code with Go extension by Microsoft installed)
5. Postman for calling APIs
  
## PostgreSQL Table

```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email TEXT NOT NULL,
  dashboard BOOLEAN NOT NULL,
  spc BOOLEAN NOT NULL,
  adjustment_guiding BOOLEAN NOT NULL,
  parameter_failure_pediction BOOLEAN NOT NULL,
  process_status_analysis BOOLEAN NOT NULL,
  company TEXT NOT NULL
);
```


### Running the server
To run the server, follow these simple steps:

```
go run main.go
```

