CREATE TABLE team
(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE game
(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,

    location TEXT NOT NULL, -- can be extended into separate table

    team_a UUID NOT NULL,
    team_b UUID NOT NULL,

    score_a INT NOT NULL DEFAULT 0,
    score_b INT NOT NULL DEFAULT 0,

    starts_at TIMESTAMP,
    finished_at TIMESTAMP,

    CONSTRAINT fk_team_a FOREIGN KEY(team_a) REFERENCES team(id),
    CONSTRAINT fk_team_b FOREIGN KEY(team_b) REFERENCES team(id)
);

