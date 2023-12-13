create table t_account
(
    account  varchar(50)   not null,
    password varchar(50)   not null,
    user_id  varchar(50)   not null,
    state    int default 0 not null,
    primary key (account)
);

