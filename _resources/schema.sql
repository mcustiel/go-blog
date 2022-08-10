PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE posts (id integer primary key asc, slug varchar(255), title varchar(255), text text);
CREATE TABLE comments (id bigint primary key asc, post_id bigint not null, text text, foreign key (post_id) references posts(id));
CREATE UNIQUE INDEX post_slug on posts(slug);
COMMIT;
