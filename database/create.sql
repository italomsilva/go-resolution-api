DROP TABLE IF EXISTS problem_sector;

DROP TABLE IF EXISTS solution;

DROP TABLE IF EXISTS problem;

DROP TABLE IF EXISTS sector;

DROP TABLE IF EXISTS "user";

CREATE TABLE "user" (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
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
    location TEXT,
    status INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id UUID,
    CONSTRAINT fk_problem_user FOREIGN KEY (user_id) REFERENCES "user" (id) ON DELETE SET NULL
);

CREATE TABLE solution (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    estimated_cost NUMERIC(12, 2),
    approved BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    problem_id UUID,
    user_id UUID,
    CONSTRAINT fk_solution_problem FOREIGN KEY (problem_id) REFERENCES problem (id) ON DELETE SET NULL,
    CONSTRAINT fk_solution_user FOREIGN KEY (user_id) REFERENCES "user" (id) ON DELETE SET NULL
);

CREATE TABLE sector (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT
);

CREATE TABLE problem_sector (
    id SERIAL PRIMARY KEY,
    problem_id UUID NOT NULL,
    sector_id INT NOT NULL,
    CONSTRAINT fk_problem_sector_problem FOREIGN KEY (problem_id) REFERENCES problem (id) ON DELETE CASCADE,
    CONSTRAINT fk_problem_sector_sector FOREIGN KEY (sector_id) REFERENCES sector (id) ON DELETE CASCADE
);

CREATE TABLE solution_reaction (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    solution_id UUID NOT NULL,
    type INT NOT NULL, 
    CONSTRAINT fk_reaction_user FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE,
    CONSTRAINT fk_reaction_solution FOREIGN KEY (solution_id) REFERENCES solution(id) ON DELETE CASCADE
);
