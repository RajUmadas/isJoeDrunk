-- Make sure we're using our `blog` database
\c ijd_db;

-- We can create our user table
CREATE TABLE IF NOT EXISTS ans (
  id SERIAL PRIMARY KEY,
  date_created timestamp NOT NULL DEFAULT now(),
  ans VARCHAR, 
  lit_level VARCHAR
);

INSERT INTO ans (
    "ans",
    "lit_level"
)VALUES(
    'yes',
    '10'
);

INSERT INTO ans (
    "ans",
    "lit_level"
)VALUES(
    'yes',
    '7'
);

INSERT INTO ans (
    "ans",
    "lit_level"
)VALUES(
    'no',
    '0'
);

INSERT INTO ans (
    "ans",
    "lit_level"
)VALUES(
    'yes',
    '9000'
);

