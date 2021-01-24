CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE go_userlist (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v1(),
    email VARCHAR(150) UNIQUE NOT NULL,
    password VARCHAR(150) NOT NULL
);