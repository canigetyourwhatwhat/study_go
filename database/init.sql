-- 記事データを格納するためのテーブル
create table if not exists articles (
                                        id integer unsigned auto_increment primary key,
                                        title varchar(100) not null,
                                        contents text not null,
                                        username varchar(100) not null,
                                        nice_num integer not null default 0,
                                        created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                        updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- コメントデータを格納するためのテーブル
create table if not exists comments (
                                        id integer unsigned auto_increment primary key,
                                        article_id integer unsigned not null,
                                        message text not null,
                                        created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                        foreign key (article_id) references articles(id)
);



-- 記事データ2つ
insert into articles
(title, contents, username, nice_num, created_at, updated_at) values
    ('firstPost', 'hello', 'john', 2, '2000-01-01 00:00:00', '2000-01-01 00:00:00');

insert into articles
(title, contents, username, nice_num, created_at, updated_at) values
    ('2nd', 'welcome', 'bob', 4, '2000-01-01 00:00:00', '2000-01-01 00:00:00');

-- コメントデータ 5 つ
insert into comments
(article_id, message, created_at) values
    (1, 'comment1', now());

insert into comments
(article_id, message) values
    (1, 'comment2');

insert into comments
(article_id, message) values
    (1, 'comment3');

insert into comments
(article_id, message) values
    (2, 'comment4');

insert into comments
(article_id, message) values
    (1, 'comment5');