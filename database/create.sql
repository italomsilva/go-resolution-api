DROP Table "user"

CREATE TABLE "user" (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    document TEXT NOT NULL UNIQUE,
    profile INT NOT NULL,
    login TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    token TEXT
);

CREATE TABLE problem (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    localization TEXT
);

CREATE TABLE solution (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    estimated_cost NUMERIC(12,2),
    likes INT DEFAULT 0,
    deslikes INT DEFAULT 0,
    problem_id UUID,
    CONSTRAINT fk_problem
      FOREIGN KEY(problem_id) REFERENCES problem(id)
      ON DELETE SET NULL
);

CREATE TABLE sector (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE problem_sector (
    id SERIAL PRIMARY KEY,
    problem_id UUID NOT NULL,
    sector_id INT NOT NULL,
    CONSTRAINT fk_problem_sector_problem FOREIGN KEY (problem_id) REFERENCES problem(id) ON DELETE CASCADE,
    CONSTRAINT fk_problem_sector_sector FOREIGN KEY (sector_id) REFERENCES sector(id) ON DELETE CASCADE
);
