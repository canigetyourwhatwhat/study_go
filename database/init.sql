-- 記事データを格納するためのテーブル
create table if not exists articles (
                                        id integer unsigned auto_increment primary key,
                                        title varchar(100) not null,
                                        contents text not null,
                                        username varchar(100) not null,
                                        nice_num integer not null,
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
(title, contents, username, nice_num, created_at) values
    ('firstPost', 'This is my first blog', 'saki', 2, now());

insert into articles
(title, contents, username, nice_num) values
    ('2nd', 'Second blog post', 'saki', 4);


-- コメントデータ 2 つ
insert into comments
(article_id, message, created_at) values
    (1, '1st comment yeah', now());

insert into comments
(article_id, message) values
    (1, 'welcome');