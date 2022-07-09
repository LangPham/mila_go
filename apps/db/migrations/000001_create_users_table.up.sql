CREATE TABLE IF NOT EXISTS accounts(
   id bigserial PRIMARY KEY,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP,
   user_name VARCHAR (255) UNIQUE NOT NULL,
   password VARCHAR (255) NOT NULL,
   role VARCHAR (255) UNIQUE NOT NULL
);