CREATE TABLE IF NOT EXISTS album(
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(255),
  artist VARCHAR(255),
  price INT
);