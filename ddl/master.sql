CREATE SCHEMA auth;

CREATE TABLE auth.users(
    id          SERIAL primary key,
    username    varchar(1024) NOT NULL UNIQUE,
    password    varchar(1024) NOT NULL
);

CREATE SCHEMA deals;

CREATE TABLE deals.companies(
    id              SERIAL primary key,
    name            varchar(1024),
    street_address  varchar(1024),
    city            varchar(1024),
    state           varchar(1024),
    zip_code        varchar(1024)
);


CREATE TABLE deals.discounts(
    id              SERIAL primary key,
    company         varchar(1024),
    day             varchar(2),
    discounts       varchar(1024)
);

INSERT INTO deals.companies(name, street_address, city, state, zip_code) VALUES
    ('Bankhead Brewing Co.','611 University Drive','Fort Worth','TX','76107'),
    ('Billy Bobs Texas','2520 Rodeo Plaza','Fort Worth','TX','76164'),
    ('Club Ritzy','1201 Oakland Boulevard','Fort Worth','TX','76103'),
    ('Downtown Cowtown at the Isis','2401 North Main Street','Fort Worth,','TX','76164')
ON CONFLICT DO NOTHING;

INSERT INTO deals.discounts(company, day, discounts) VALUES
    ('Bankhead Brewing Co.', 'Th', 'First Beer Free!'),
    ('Billy Bobs Texas', 'Fr', 'All Day Happy Hour!'),
    ('Club Ritzy', 'Sa', 'Buy One Get One Free Bottles UNDER $50'),
    ('Downtown Cowtown at the Isis', 'Fr', 'Half Priced House Liquor'),
    ('Downtown Cowtown at the Isis', 'Sa', '$5 Large One-topping Pizza')
ON CONFLICT DO NOTHING;