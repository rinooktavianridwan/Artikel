create table artikel(
    id int auto_increment primary key,
    title varchar(200),
    content text,
    category varchar(100),
    created_at timestamp,
    update_at timestamp,
    status varchar(100)
)