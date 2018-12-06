create table products (
  id bigserial not null,
  name text not null,
  tax_code int4 not null,
  price int8 not null,
  constraint products_id primary key(id)
) with (OIDS=FALSE);