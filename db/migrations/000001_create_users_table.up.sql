CREATE TABLE IF NOT EXISTS public.author (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    age INTEGER NOT NULL
);

INSERT INTO author (name, age) VALUES ('Федор Достоевский', 55);
INSERT INTO author (name, age) VALUES ('Джоан Роулинг', 44);
INSERT INTO author (name, age) VALUES ('Джек Лондон', 66); 