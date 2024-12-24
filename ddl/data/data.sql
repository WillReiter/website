INSERT INTO deals.companies(id, name, street_address, city, state, zip_code) VALUES
    ('141bc259-b446-4990-8c40-93c57add542e','Bankhead Brewing Co.','611 University Drive','Fort Worth','TX','76107'),
    ('ac797001-0377-4d6d-8901-6aba33fcfd8c','Billy Bobs Texas','2520 Rodeo Plaza','Fort Worth','TX','76164'),
    ('979a5ef9-e4f1-468f-98d8-cca2e5a849d7','Club Ritzy','1201 Oakland Boulevard','Fort Worth','TX','76103'),
    ('30bd25f1-3421-42f7-acd0-9f78b4026b8e','Downtown Cowtown at the Isis','2401 North Main Street','Fort Worth,','TX','76164')
ON CONFLICT DO NOTHING;

INSERT INTO deals.discounts(id, company, day, discounts) VALUES
    ('c73939d5-28a2-4a8b-b7ad-83316d4639ca', 'Bankhead Brewing Co.', 'Th', 'First Beer Free!'),
    ('d31bb509-f2bd-4ce4-925d-f85a5e8df5df', 'Billy Bobs Texas', 'Fr', 'All Day Happy Hour!'),
    ('8a7d84f5-c9fb-476e-8ce0-854a0620d3e4', 'Club Ritzy', 'Sa', 'Buy One Get One Free Bottles UNDER $50'),
    ('6958d919-f47a-4eb6-ae7c-e8900a8f3823', 'Downtown Cowtown at the Isis', 'Fr', 'Half Priced House Liquor'),
    ('6c0d006a-3910-4eb3-b29d-c6f68447d30e', 'Downtown Cowtown at the Isis', 'Sa', '$5 Large One-topping Pizza')
ON CONFLICT DO NOTHING;