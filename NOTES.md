To set up the database, create a `example.go` file.

Next, enter the sqlite3 shell and create the `users` table along with a unique index for the `email` column.

```
sqlite3 example.db
sqlite> create table users (id integer primary key autoincrement, email text not null, name text not null);
sqlite> create unique index idx_users_email on users (email);
```
