
CREATE TABLE metadata (
    id serial PRIMARY KEY,
    mime_type varchar(255),
    path varchar(255),
    sub_domain varchar(255),
    table_reference varchar(255),
    mark_as_delete boolean,
    create_date bigint
);
