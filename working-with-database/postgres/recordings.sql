CREATE TABLE album(
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(255),
  artist VARCHAR(255),
  price INT
);

INSERT INTO
  album (title, artist, price)
VALUES
  ('Hari Yang Cerah', 'Peterpan', 50000);

INSERT INTO
  album (title, artist, price)
VALUES
  ('Sebuah Nama Sebuah Cerita', 'Peterpan', 50000);

INSERT INTO
  album (title, artist, price)
VALUES
  ('Bintang Di surga', 'Peterpan', 50000);