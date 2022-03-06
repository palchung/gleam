CREATE TABLE users (
  id            serial primary key,
  firstName     varchar(255) NOT NULL,
  lastName      varchar(355) NOT NULL,
  password      varchar(355) NOT NULL,
  email         varchar(255) UNIQUE NOT NULL,
  createdAt     timestamptz NOT NULL,
  updatedAt     timestamptz NOT NULL
);