# Cara menjalankan 
## backend
### inisiasi database di psql
- `psql -U <user> -d <db name>`
- `cd /path/to/backend`
- `\i initDb.sql`
- `\c random_db`
- `\i initTable.sql`
- `\i initDb.sql`
### build binary
- edit .env.dev
- `cd /path/to/backend`
- `go build -o run.exe`

## frontend
- `cd /path/to/frontend`
- `npm install`
- `npm run dev`
