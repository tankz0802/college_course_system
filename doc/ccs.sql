/*==============================================================*/
/* DBMS name:      MySQL 5.0                                    */
/* Created on:     2021/1/8 12:53:13                            */
/*==============================================================*/




/*==============================================================*/
/* Table: class                                                 */
/*==============================================================*/
create table class
(
   id                   integer not null  comment '',
   name                 varchar(20)  comment '',
   tid                  integer not null  comment '',
   primary key (id)
);

/*==============================================================*/
/* Table: class_course                                          */
/*==============================================================*/
create table class_course
(
   id                   integer not null  comment '',
   cid                  integer not null  comment '',
   primary key (id, cid)
);

/*==============================================================*/
/* Table: course                                                */
/*==============================================================*/
create table course
(
   id                   char(15) not null  comment '',
   name                 varchar(24) not null  comment '',
   duration             int  comment '',
   credit               float  comment '',
   category             varchar(20)  comment '',
   stu_num              int  comment '',
   max_num              int  comment '',
   status               int  comment '',
   primary key (id)
);

/*==============================================================*/
/* Table: course_group                                          */
/*==============================================================*/
create table course_group
(
   id                   integer not null auto_increment  comment '',
   name                 varchar(20)  comment '',
   tid                  integer  comment '',
   primary key (id)
);

/*==============================================================*/
/* Table: course_group_teacher                                  */
/*==============================================================*/
create table course_group_teacher
(
   cg_id                integer not null  comment '',
   tid                  integer not null  comment '',
   primary key (cg_id, tid)
);

/*==============================================================*/
/* Table: course_schedule                                       */
/*==============================================================*/
create table course_schedule
(
   id                   integer not null auto_increment  comment '',
   cid                  char(15)  comment '',
   week_start           integer  comment '',
   week_end             integer  comment '',
   week_day             integer  comment '',
   section_start        integer  comment '',
   section_end          integer  comment '',
   tid                  integer  comment '',
   primary key (id)
);

/*==============================================================*/
/* Table: student                                               */
/*==============================================================*/
create table student
(
   id                   integer not null  comment '',
   name                 char(20) not null  comment '',
   password             char(64) not null  comment '',
   cid                  integer not null  comment '',
   primary key (id)
);

/*==============================================================*/
/* Table: student_course                                        */
/*==============================================================*/
create table student_course
(
   sid                  integer not null  comment '',
   cid                  integer not null  comment '',
   grade                int  comment '',
   gpa                  double  comment '',
   primary key (sid, cid)
);

/*==============================================================*/
/* Table: teacher                                               */
/*==============================================================*/
create table teacher
(
   id                   integer not null  comment '',
   name                 char(20) not null  comment '',
   password             char(64)  comment '',
   title                int  comment '',
   primary key (id)
);

/*==============================================================*/
/* Table: teacher_course                                        */
/*==============================================================*/
create table teacher_course
(
   tid                  integer not null  comment '',
   cid                  char(15) not null  comment '',
   primary key (tid, cid)
);

alter table class add constraint FK_CLASS_REFERENCE_TEACHER foreign key (tid)
      references teacher (id) on delete restrict on update restrict;

alter table course_schedule add constraint FK_COURSE_S_REFERENCE_COURSE foreign key (cid)
      references course (id) on delete restrict on update restrict;

alter table course_schedule add constraint FK_COURSE_S_REFERENCE_TEACHER foreign key (tid)
      references teacher (id) on delete restrict on update restrict;

alter table student add constraint FK_STUDENT_REFERENCE_CLASS foreign key (cid)
      references class (id) on delete restrict on update restrict;

