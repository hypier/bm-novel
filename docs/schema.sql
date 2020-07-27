create table if not exists public."user"
(
    user_id uuid not null
        constraint user_pkey
            primary key,
    user_name varchar(32) not null,
    password varchar(64) not null,
    role_code varchar(64) [] not null,
    real_name varchar(16) not null,
    create_at timestamp(6) with time zone,
    update_at timestamp(6) with time zone,
    is_lock boolean,
    need_change_password boolean
);

alter table public."user" owner to postgres;

create unique index if not exists user_user_name_uindex
    on public."user" (user_name);

create table if not exists public.permission
(
    pid uuid not null
        constraint permission_pk
            primary key,
    name varchar(20) not null,
    uri varchar(50) not null,
    method varchar(10) not null,
    roles varchar(128) []
);

alter table public.permission owner to postgres;

