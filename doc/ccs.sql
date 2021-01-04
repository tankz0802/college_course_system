/*==============================================================*/
/* DBMS name:      MySQL 5.0                                    */
/* Created on:     2021/1/4 14:58:54                            */
/*==============================================================*/



/*==============================================================*/
/* Table: cg_teacher                                            */
/*==============================================================*/
create table cg_teacher
(
   cg_id                integer not null  comment '',
   tid                  integer not null  comment '',
   primary key (cg_id, tid)
);

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
/* Table: course                                                */
/*==============================================================*/
create table course
(
   id                   integer not null  comment '',
   name                 varchar(24) not null  comment '',
   class_hour           int  comment '',
   credit               int  comment '',
   category             int  comment '',
   stu_num              int  comment '',
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
/* Table: course_schedule                                       */
/*==============================================================*/
create table course_schedule
(
   id                   integer not null auto_increment  comment '',
   cid                  integer  comment '',
   week                 integer  comment '',
   day                  integer  comment '',
   start                integer  comment '',
   end                  integer  comment '',
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
   sex                  char(2)  comment '',
   password             varchar(20) not null  comment '',
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
   sex                  char(2)  comment '',
   title                int  comment '',
   password             varchar(20)  comment '',
   primary key (id)
);

/*==============================================================*/
/* Table: teacher_course                                        */
/*==============================================================*/
create table teacher_course
(
   tid                  integer not null  comment '',
   cid                  integer not null  comment '',
   primary key (tid, cid)
);

alter table class add constraint FK_CLASS_REFERENCE_STUDENT foreign key (id)
      references student (id) on delete restrict on update restrict;

alter table class add constraint FK_CLASS_REFERENCE_TEACHER foreign key (tid)
      references teacher (id) on delete restrict on update restrict;

alter table course add constraint FK_COURSE_REFERENCE_COURSE_S foreign key (id)
      references course_schedule (id) on delete restrict on update restrict;

