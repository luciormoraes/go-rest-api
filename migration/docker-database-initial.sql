create table personalities(
    id serial primary key,
    name varchar,
    history varchar
);

INSERT INTO personalities(name, history) VALUES
('Sir Peter Jackson', 'The mastermind behind the Lord of the Rings trilogy, Sir Peter Jackson, was born in the beautiful Pukerua Bay in Porirua, a 30 minute drive from Wellington.'),
('Sir Edmund Hillary ', 'With regards to global achievements, they don’t get much bigger than Sir Ed’s feat of being the first person to conquer the summit of the mighty Mount Everest in 1953.');