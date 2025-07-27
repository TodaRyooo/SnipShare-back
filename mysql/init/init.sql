create table if not exists snippet(
    id bigserial primary key,
    name varchar(255) notnull,
    body text not null,
    type_id bigint not null,
    created_by bigint not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

INSERT INTO snippet(id,name,body,type_id,created_at,updated_at,created_by)
VALUES(1,'nxkzm','var body: some View { codeView(code) }',1,now(),now(),1);
