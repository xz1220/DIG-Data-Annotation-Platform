create database labelproject;

use labelproject;


create table if not exists image
(
    image_id        int(64) auto_increment,
    image_name      varchar(1024) null,
    image_thumb     varchar(1024) null,
    user_confirm_id int(64)       null,
    task_id         int(64)       null,
    width           int           null,
    height          int           null,
    constraint image_id_index
        unique (image_id)
);

alter table image
    add primary key (image_id);

create table if not exists imagedata
(
    data_id    int(64) auto_increment,
    image_id   int(64)       null,
    label_id   int(64)       null,
    user_id    int(64)       null,
    data_desc  varchar(1024) null,
    label_type int           null,
    iscrowd    int           null,
    constraint data_id_index
        unique (data_id)
);

alter table imagedata
    add primary key (data_id);

create table if not exists imagedatapoints
(
    data_id  int(64) null,
    `order`  int     null,
    x        double  null,
    y        double  null,
    image_id int(64) null,
    user_id  int(64) null
);

create table if not exists imagedatarle
(
    data_id  int(64)  null,
    image_id int(64)  null,
    user_id  int(64)  null,
    data_rle longtext null
);

create table if not exists imagelabel
(
    label_id    int(64) auto_increment,
    label_name  varchar(50) null,
    label_type  int         null,
    label_color varchar(30) null,
    constraint label_label_id_uindex
        unique (label_id)
);

alter table imagelabel
    add primary key (label_id);

create table if not exists task
(
    task_id      int(64) auto_increment,
    task_name    varchar(50)   null,
    task_desc    varchar(1024) null,
    image_number int(64)       null,
    is_created  int null,
    task_type    int           null,
    constraint task_id_index
        unique (task_id)
);

alter table task
    add primary key (task_id);

create table if not exists tasklabelinfo
(
    task_id  int(64) null,
    label_id int(64) null
);

create table if not exists taskreviewerinfo
(
    task_id     int(64) null,
    reviewer_id int(64) null
);

create table if not exists taskuserinfo
(
    task_id int(64) null,
    user_id int(64) null
);

create table if not exists test
(
    id       int(10) auto_increment
        primary key,
    test     varchar(255) null,
    testLong text         null
);

create table if not exists user
(
    user_id     int(64) auto_increment,
    username    varchar(50)            null,
    password    varchar(100)           null,
    authorities varchar(20) default '' null,
    constraint user_id_uindex
        unique (user_id)
);

alter table user
    add primary key (user_id);

insert into user (username,password,authorities) values("admin","admin","ROLE_ADMIN");



create table if not exists userfinished
(
    user_id  int(64) null,
    task_id  int(64) null,
    image_id int(64) null
);

create table if not exists video
(
    video_id        int(64) auto_increment
        primary key,
    video_name      varchar(1024) null,
    video_thumb     varchar(1024) null,
    task_id         int(64)       null,
    duration        double        null,
    user_confirm_id int(64)       null
);

create index video_id_index
    on video (video_id);

create table if not exists videodata
(
    data_id    int(64) auto_increment
        primary key,
    user_id    int(64)      null,
    label_id   int(64)      null,
    type       int          null,
    video_id   int(64)      null,
    start_time varchar(255) null,
    end_time   varchar(255) null,
    sentence   mediumtext   null
);

create index data_id_index
    on videodata (data_id);

create table if not exists videolabel
(
    label_id int(64) auto_increment
        primary key,
    question varchar(1024) null,
    type     int           null,
    selector varchar(1024) null
);

create index videolabel_label_id_index
    on videolabel (label_id);


