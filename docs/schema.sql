create table public."user"
(
    user_id uuid not null
        constraint user_pkey
            primary key,
    user_name varchar(32) not null,
    password varchar(64) not null,
    roles varchar(64) [] not null,
    real_name varchar(16) not null,
    create_at timestamp(6) with time zone,
    update_at timestamp(6) with time zone,
    is_lock boolean,
    need_change_password boolean
);

alter table public."user" owner to postgres;

create unique index user_user_name_uindex
    on public."user" (user_name);

create table public.permission
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

create table public.novel
(
    novel_id uuid not null
        constraint novel_pkey
            primary key,
    assigned_chapters_count integer,
    chapters_count integer,
    chief_editor_id uuid,
    create_at timestamp(6) with time zone,
    novel_title varchar(50) not null
        constraint uk_novel_novel_title
            unique,
    responsible_editor_id uuid,
    update_at timestamp(6) with time zone,
    words_count integer,
    settings jsonb
);

comment on table public.novel is '小说';

comment on column public.novel.novel_id is '小说ID';

comment on column public.novel.assigned_chapters_count is '已指派的章节数';

comment on column public.novel.chapters_count is '总章节数量';

comment on column public.novel.chief_editor_id is '主编';

comment on column public.novel.create_at is '创建时间';

comment on column public.novel.novel_title is '标题';

comment on column public.novel.responsible_editor_id is '责编';

comment on column public.novel.update_at is '更新时间';

comment on column public.novel.words_count is '总字数';

comment on column public.novel.settings is '格式设置';

alter table public.novel owner to postgres;

create table public.chapter
(
    chapter_id uuid not null
        constraint chapter_pkey
            primary key,
    chapter_no integer not null,
    chapter_title varchar(50) not null,
    create_at timestamp(6) with time zone,
    flag integer,
    out_source_id uuid,
    update_at timestamp(6) with time zone,
    words_count integer,
    novel_id uuid
        constraint fk_chapter_novel_id
            references public.novel
);

comment on table public.chapter is '章节';

comment on column public.chapter.chapter_id is '章节ID';

comment on column public.chapter.chapter_no is '章节序号';

comment on column public.chapter.chapter_title is '章节标题';

comment on column public.chapter.create_at is '创建时间';

comment on column public.chapter.flag is '标识 1 正确章 2 重复章 3 缺失章 4 错序章';

comment on column public.chapter.out_source_id is '外包编辑id';

comment on column public.chapter.update_at is '更新时间';

comment on column public.chapter.words_count is '字数';

comment on column public.chapter.novel_id is '小说ID';

alter table public.chapter owner to postgres;

create table public.episode
(
    episode_id uuid not null
        constraint episode_pkey
            primary key,
    create_at timestamp(6) with time zone,
    episode_name varchar(50),
    episode_no integer,
    status smallint,
    update_at timestamp(6) with time zone,
    words_count integer,
    novel_id uuid
        constraint fk_episode_novel_id
            references public.novel
);

comment on table public.episode is '集数';

comment on column public.episode.episode_id is '集数ID';

comment on column public.episode.create_at is '创建时间';

comment on column public.episode.episode_name is '集名';

comment on column public.episode.episode_no is '集序号';

comment on column public.episode.status is '状态  1 未审核 2审核中 4已审核 8已定稿';

comment on column public.episode.update_at is '更新时间';

comment on column public.episode.words_count is '字数';

comment on column public.episode.novel_id is '小说ID';

alter table public.episode owner to postgres;

create table public.novel_role
(
    role_id uuid not null
        constraint novel_role_pkey
            primary key,
    age varchar(20),
    characters varchar(255),
    create_at timestamp(6) with time zone,
    gender varchar(10),
    role_class varchar(20),
    role_name varchar(20),
    update_at timestamp(6) with time zone
);

comment on table public.novel_role is '角色';

comment on column public.novel_role.role_id is '角色ID';

comment on column public.novel_role.age is '年纪';

comment on column public.novel_role.characters is '人设';

comment on column public.novel_role.create_at is '创建时间';

comment on column public.novel_role.gender is '性别';

comment on column public.novel_role.role_class is '角色类别';

comment on column public.novel_role.role_name is '角色名';

comment on column public.novel_role.update_at is '更新时间';

alter table public.novel_role owner to postgres;

create table public.paragraph
(
    paragraph_id uuid not null
        constraint paragraph_pkey
            primary key,
    content varchar(5000),
    create_at timestamp(6) with time zone,
    next uuid,
    prev uuid,
    update_at timestamp(6) with time zone,
    word_count integer,
    chapter_id uuid
        constraint fk_paragraph_chapter_id
            references public.chapter,
    episode_id uuid
        constraint fk_paragraph_episode_id
            references public.episode,
    role_id uuid
        constraint fk_paragraph_role_id
            references public.novel_role,
    novel_id uuid
        constraint fk_paragraph_novle_id
            references public.novel
);

comment on table public.paragraph is '段落';

comment on column public.paragraph.paragraph_id is '段落ID';

comment on column public.paragraph.content is '段落内容';

comment on column public.paragraph.create_at is '创建时间';

comment on column public.paragraph.next is '下一个';

comment on column public.paragraph.prev is '上一个';

comment on column public.paragraph.update_at is '更新时间';

comment on column public.paragraph.word_count is '字数';

comment on column public.paragraph.chapter_id is '章节ID';

comment on column public.paragraph.episode_id is '集数ID';

comment on column public.paragraph.role_id is '角色ID';

comment on column public.paragraph.novel_id is '小说ID';

alter table public.paragraph owner to postgres;

