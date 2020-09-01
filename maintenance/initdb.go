package maintenance

import (
	"log"
	"os"
	"rdm/dam"
)

const (
	sql = `
create table inc
(
    name text,
    val  integer
);

` + `

create table permission_code
(
    pcid         integer
        constraint permission_code_pk
            primary key,
    code         text    not null,
    permission   integer not null,
    deadline     integer not null,
    generated_by integer not null,
    used_by      integer default 0,
    deleted      integer default 0
);

create unique index permission_code_code_uindex
    on permission_code (code);

` + `

create table record
(
    urid       integer
        constraint record_pk
            primary key,
    uuid       integer,
    wuid       integer,
    time_start integer,
    time_end   integer default 0,
    coin       integer default 0,
    tag        text
);

` + `

create table "transaction"
(
    utid      integer
        constraint transaction_pk
            primary key,
    uuid      integer,
	wuid      integer,
    unix_time integer,
    rmb       integer default 0,
    coin      integer default 0
);

` + `

create table user
(
    uuid       integer
        constraint user_pk
            primary key,
    username   text,
    nickname   text,
    password   text,
    permission integer default 0,
    plant      integer default 0,
    rmb_in     integer default 0,
    rmb_out    integer default 0,
    coin       integer default 0,
    deleted    integer default 0
);

create unique index user_username_uindex
    on user (username);

` + `


create table work
(
    wuid      integer
        constraint work_pk
            primary key,
    uuid      integer not null,
    name      text,
    type      integer not null,
    unit      integer not null,
    coin      integer default 0,
    associate integer default 0,
    deleted   integer default 0,
    constraint work_pk_2
        unique (uuid, name)
);

` + `

INSERT INTO inc (name, val) VALUES ('user', 0);
INSERT INTO inc (name, val) VALUES ('work', 0);
INSERT INTO inc (name, val) VALUES ('record', 0);
INSERT INTO inc (name, val) VALUES ('transaction', 0);
INSERT INTO inc (name, val) VALUES ('permission_code', 1);

INSERT INTO permission_code (pcid, code, permission, deadline, generated_by) 
	VALUES (1, 'admin', 11747296, 4753268904, 0);
`
)

func InitDatabase() {
	if !IsFile("rdm.db") {
		log.Println("rdm.db do not exist")
		os.Exit(-1)
	}

	err := dam.Exec(sql).Error
	if err != nil {
		log.Println(err)
		os.Exit(-2)
	}

	log.Println("Finished")

	os.Exit(0)
}

func Exists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	if !Exists(path) {
		return false
	}
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}