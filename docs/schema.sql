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


create table if not exists public.novel
(
    novel_id uuid not null
        constraint novel_pkey
            primary key,
    assigned_chapters integer,
    chapters_count integer,
    chief_editor_id varchar(255),
    create_at timestamp,
    novel_title varchar(255)
        constraint uk_novel_novel_title
            unique,
    responsible_editor_id varchar(255),
    update_at timestamp,
    words_count integer
);

alter table public.novel owner to postgres;

create table if not exists public.chapter
(
    chapter_id uuid not null
        constraint chapter_pkey
            primary key,
    chapter_no integer,
    chapter_title varchar(255),
    create_at timestamp,
    flag integer,
    out_source_id uuid,
    update_at timestamp,
    words_count integer,
    novel_id uuid
        constraint fk_chapter_novel_id
            references public.novel
);

alter table public.chapter owner to postgres;

create table if not exists public.episode
(
    episode_id uuid not null
        constraint episode_pkey
            primary key,
    create_at timestamp,
    episode_name varchar(255),
    episode_no integer,
    status integer,
    update_at timestamp,
    words_count integer,
    novel_id uuid
        constraint fk_episode_novel_id
            references public.novel
);

alter table public.episode owner to postgres;

create table if not exists public.novel_role
(
    role_id uuid not null
        constraint novel_role_pkey
            primary key,
    age varchar(255),
    characters varchar(255),
    create_at timestamp,
    gender varchar(255),
    role_class varchar(255),
    role_name varchar(255),
    update_at timestamp
);

alter table public.novel_role owner to postgres;

create table if not exists public.novel_setting
(
    setting_id uuid not null
        constraint novel_setting_pkey
            primary key,
    chapter_has_prefix boolean,
    chapter_no_format integer,
    chapter_prefix varchar(255),
    episode_words_max integer,
    episode_words_min integer,
    novel_id uuid
        constraint fk_novel_setting_novel_id
            references public.novel
);

alter table public.novel_setting owner to postgres;

create table if not exists public.paragraph
(
    paragraph_id uuid not null
        constraint paragraph_pkey
            primary key,
    content varchar(5000),
    create_at timestamp,
    next integer,
    prev integer,
    update_at timestamp,
    word_count integer,
    chapter_id uuid
        constraint fk_paragraph_chapter_id
            references public.chapter,
    episode_id uuid
        constraint fk_paragraph_episode_id
            references public.episode,
    role_id uuid
        constraint fk_paragraph_role_id
            references public.novel_role
);

alter table public.paragraph owner to postgres;

