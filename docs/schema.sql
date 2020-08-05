CREATE TABLE "chapter"
(
    "chapter_id"    uuid                                       NOT NULL,
    "chapter_no"    int4                                       NOT NULL,
    "chapter_title" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
    "create_at"     timestamptz(6),
    "feature_code"  int4,
    "out_source_id" uuid,
    "update_at"     timestamptz(6),
    "words_count"   int4,
    "novel_id"      uuid,
    CONSTRAINT "chapter_pkey" PRIMARY KEY ("chapter_id")
);
ALTER TABLE "chapter"
    OWNER TO "postgres";
COMMENT ON COLUMN "chapter"."chapter_id" IS '章节ID';
COMMENT ON COLUMN "chapter"."chapter_no" IS '章节序号';
COMMENT ON COLUMN "chapter"."chapter_title" IS '章节标题';
COMMENT ON COLUMN "chapter"."create_at" IS '创建时间';
COMMENT ON COLUMN "chapter"."feature_code" IS '标识 1 正确章 2 重复章 3 缺失章 4 错序章';
COMMENT ON COLUMN "chapter"."out_source_id" IS '外包编辑id';
COMMENT ON COLUMN "chapter"."update_at" IS '更新时间';
COMMENT ON COLUMN "chapter"."words_count" IS '字数';
COMMENT ON COLUMN "chapter"."novel_id" IS '小说ID';
COMMENT ON TABLE "chapter" IS '章节';

CREATE TABLE "episode"
(
    "episode_id"    uuid NOT NULL,
    "create_at"     timestamptz(6),
    "episode_title" varchar(50) COLLATE "pg_catalog"."default",
    "episode_no"    int4,
    "status"        int2,
    "update_at"     timestamptz(6),
    "words_count"   int4,
    "novel_id"      uuid,
    CONSTRAINT "episode_pkey" PRIMARY KEY ("episode_id")
);
ALTER TABLE "episode"
    OWNER TO "postgres";
COMMENT ON COLUMN "episode"."episode_id" IS '集数ID';
COMMENT ON COLUMN "episode"."create_at" IS '创建时间';
COMMENT ON COLUMN "episode"."episode_title" IS '集名';
COMMENT ON COLUMN "episode"."episode_no" IS '集序号';
COMMENT ON COLUMN "episode"."status" IS '状态  1 未审核 2审核中 4已审核 8已定稿';
COMMENT ON COLUMN "episode"."update_at" IS '更新时间';
COMMENT ON COLUMN "episode"."words_count" IS '字数';
COMMENT ON COLUMN "episode"."novel_id" IS '小说ID';
COMMENT ON TABLE "episode" IS '集数';

CREATE TABLE "novel"
(
    "novel_id"              uuid                                       NOT NULL,
    "chief_editor_id"       uuid,
    "create_at"             timestamptz(6),
    "novel_title"           varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
    "responsible_editor_id" uuid,
    "update_at"             timestamptz(6),
    "settings"              jsonb,
    CONSTRAINT "novel_pkey" PRIMARY KEY ("novel_id"),
    CONSTRAINT "uk_novel_novel_title" UNIQUE ("novel_title")
);
ALTER TABLE "novel"
    OWNER TO "postgres";
COMMENT ON COLUMN "novel"."novel_id" IS '小说ID';
COMMENT ON COLUMN "novel"."chief_editor_id" IS '主编';
COMMENT ON COLUMN "novel"."create_at" IS '创建时间';
COMMENT ON COLUMN "novel"."novel_title" IS '标题';
COMMENT ON COLUMN "novel"."responsible_editor_id" IS '责编';
COMMENT ON COLUMN "novel"."update_at" IS '更新时间';
COMMENT ON COLUMN "novel"."settings" IS '格式设置';
COMMENT ON TABLE "novel" IS '小说';

CREATE TABLE "novel_counts"
(
    "count_id"                uuid NOT NULL,
    "assigned_chapters_count" int4,
    "chapters_count"          int4,
    "words_count"             int4,
    "novel_id"                uuid,
    PRIMARY KEY ("count_id")
);
COMMENT ON COLUMN "novel_counts"."count_id" IS '计数ID';
COMMENT ON COLUMN "novel_counts"."assigned_chapters_count" IS '已指派的章节数';
COMMENT ON COLUMN "novel_counts"."chapters_count" IS '总章节数';
COMMENT ON COLUMN "novel_counts"."words_count" IS '总字数';

CREATE TABLE "novel_role"
(
    "role_id"    uuid NOT NULL,
    "age"        varchar(20) COLLATE "pg_catalog"."default",
    "characters" varchar(255) COLLATE "pg_catalog"."default",
    "create_at"  timestamptz(6),
    "gender"     varchar(10) COLLATE "pg_catalog"."default",
    "role_class" varchar(20) COLLATE "pg_catalog"."default",
    "role_name"  varchar(20) COLLATE "pg_catalog"."default",
    "update_at"  timestamptz(6),
    CONSTRAINT "novel_role_pkey" PRIMARY KEY ("role_id")
);
ALTER TABLE "novel_role"
    OWNER TO "postgres";
COMMENT ON COLUMN "novel_role"."role_id" IS '角色ID';
COMMENT ON COLUMN "novel_role"."age" IS '年纪';
COMMENT ON COLUMN "novel_role"."characters" IS '人设';
COMMENT ON COLUMN "novel_role"."create_at" IS '创建时间';
COMMENT ON COLUMN "novel_role"."gender" IS '性别';
COMMENT ON COLUMN "novel_role"."role_class" IS '角色类别';
COMMENT ON COLUMN "novel_role"."role_name" IS '角色名';
COMMENT ON COLUMN "novel_role"."update_at" IS '更新时间';
COMMENT ON TABLE "novel_role" IS '角色';

CREATE TABLE "paragraph"
(
    "paragraph_id" uuid NOT NULL,
    "content"      varchar(5000) COLLATE "pg_catalog"."default",
    "create_at"    timestamptz(6),
    "next"         uuid,
    "prev"         uuid,
    "update_at"    timestamptz(6),
    "words_count"  int4,
    "chapter_id"   uuid,
    "episode_id"   uuid,
    "role_id"      uuid,
    "novel_id"     uuid,
    CONSTRAINT "paragraph_pkey" PRIMARY KEY ("paragraph_id")
);
ALTER TABLE "paragraph"
    OWNER TO "postgres";
COMMENT ON COLUMN "paragraph"."paragraph_id" IS '段落ID';
COMMENT ON COLUMN "paragraph"."content" IS '段落内容';
COMMENT ON COLUMN "paragraph"."create_at" IS '创建时间';
COMMENT ON COLUMN "paragraph"."next" IS '下一个';
COMMENT ON COLUMN "paragraph"."prev" IS '上一个';
COMMENT ON COLUMN "paragraph"."update_at" IS '更新时间';
COMMENT ON COLUMN "paragraph"."words_count" IS '字数';
COMMENT ON COLUMN "paragraph"."chapter_id" IS '章节ID';
COMMENT ON COLUMN "paragraph"."episode_id" IS '集数ID';
COMMENT ON COLUMN "paragraph"."role_id" IS '角色ID';
COMMENT ON COLUMN "paragraph"."novel_id" IS '小说ID';
COMMENT ON TABLE "paragraph" IS '段落';

CREATE TABLE "permission"
(
    "pid"    uuid                                       NOT NULL,
    "name"   varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
    "uri"    varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
    "method" varchar(10) COLLATE "pg_catalog"."default" NOT NULL,
    "roles"  varchar[][] COLLATE "pg_catalog"."default",
    CONSTRAINT "permission_pk" PRIMARY KEY ("pid")
);
ALTER TABLE "permission"
    OWNER TO "postgres";
COMMENT ON COLUMN "permission"."pid" IS '权限ID';
COMMENT ON COLUMN "permission"."name" IS '权限名';
COMMENT ON COLUMN "permission"."uri" IS 'URI地址';
COMMENT ON COLUMN "permission"."method" IS '方法';
COMMENT ON COLUMN "permission"."roles" IS '角色列表';
COMMENT ON TABLE "permission" IS '权限';

CREATE TABLE "user"
(
    "user_id"              uuid                                       NOT NULL,
    "user_name"            varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
    "password"             varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
    "roles"                varchar[][] COLLATE "pg_catalog"."default" NOT NULL,
    "real_name"            varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
    "create_at"            timestamptz(6),
    "update_at"            timestamptz(6),
    "is_lock"              bool,
    "need_change_password" bool,
    CONSTRAINT "user_pkey" PRIMARY KEY ("user_id")
);
ALTER TABLE "user"
    OWNER TO "postgres";
CREATE UNIQUE INDEX "user_user_name_uindex" ON "user" USING btree (
                                                                   "user_name" COLLATE "pg_catalog"."default"
                                                                   "pg_catalog"."text_ops" ASC NULLS LAST
    );
COMMENT ON COLUMN "user"."user_id" IS '用户ID';
COMMENT ON COLUMN "user"."user_name" IS '用户名';
COMMENT ON COLUMN "user"."password" IS '密码';
COMMENT ON COLUMN "user"."roles" IS '角色列表';
COMMENT ON COLUMN "user"."real_name" IS '姓名';
COMMENT ON COLUMN "user"."create_at" IS '创建时间';
COMMENT ON COLUMN "user"."update_at" IS '更新时间';
COMMENT ON COLUMN "user"."is_lock" IS '是否锁定';
COMMENT ON COLUMN "user"."need_change_password" IS '是否需要修改密码';
COMMENT ON TABLE "user" IS '用户';

ALTER TABLE "chapter"
    ADD CONSTRAINT "fk_chapter_novel_id" FOREIGN KEY ("novel_id") REFERENCES "novel" ("novel_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "episode"
    ADD CONSTRAINT "fk_episode_novel_id" FOREIGN KEY ("novel_id") REFERENCES "novel" ("novel_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "novel_counts"
    ADD CONSTRAINT "fk_count_novel_id" FOREIGN KEY ("novel_id") REFERENCES "novel" ("novel_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "paragraph"
    ADD CONSTRAINT "fk_paragraph_chapter_id" FOREIGN KEY ("chapter_id") REFERENCES "chapter" ("chapter_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "paragraph"
    ADD CONSTRAINT "fk_paragraph_episode_id" FOREIGN KEY ("episode_id") REFERENCES "episode" ("episode_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "paragraph"
    ADD CONSTRAINT "fk_paragraph_role_id" FOREIGN KEY ("role_id") REFERENCES "novel_role" ("role_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "paragraph"
    ADD CONSTRAINT "fk_paragraph_novel_id" FOREIGN KEY ("novel_id") REFERENCES "novel" ("novel_id");

