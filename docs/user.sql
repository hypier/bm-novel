create table if not exists "user"
(
	user_id uuid not null
		constraint user_pk
			primary key,
	user_name varchar(16) not null,
	password varchar(64) not null,
	role_code varchar(64) not null,
	real_name varchar(16),
	need_change_password boolean,
	is_lock boolean,
	create_at timestamp with time zone,
	update_at timestamp with time zone
);

alter table "user" owner to postgres;

create unique index if not exists user_user_name_uindex
	on "user" (user_name);
