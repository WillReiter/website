DROP TABLE deals.companies;
DROP TABLE deals.discounts;
CREATE TABLE deals.companies(
    id              varchar(1024) primary key,
    name            varchar(1024),
    street_address  varchar(1024),
    city            varchar(1024),
    state           varchar(1024),
    zip_code        varchar(1024)
);


CREATE TABLE deals.discounts(
    id              varchar(1024) primary key,
    company         varchar(1024),
    day             varchar(2),
    discounts       varchar(1024)
);