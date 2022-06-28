create table order_table
(
    order_number int auto_increment,
    seller_id    int         not null,
    buyer_id     int         not null,
    prod_id      int         not null,
    time_bought  datetime    not null,
    order_status varchar(30) not null,
    constraint order_table_order_number_uindex
        unique (order_number)
);

alter table order_table
    add primary key (order_number);

create table user_table
(
    uid      int auto_increment
        primary key,
    username varchar(30)  not null,
    password varchar(256) not null,
    email    varchar(256) not null,
    coin     int(10)      not null
);

create table add_table
(
    loc_id   int          not null
        primary key,
    uid      int          not null,
    address  varchar(256) not null,
    city     varchar(100) null,
    postcode varchar(20)  not null,
    country  varchar(60)  null,
    constraint address_table_user_table_uid_fk
        foreign key (uid) references user_table (uid)
);

create table prod_table
(
    prod_id          int auto_increment
        primary key,
    prod_name        varchar(30)   not null,
    details          varchar(1000) null,
    start_time       datetime      not null,
    end_time         datetime      not null,
    initial_price    double        not null,
    discounted_price double        not null,
    stock            int           not null,
    num_sold         int           not null,
    uid              int           not null,
    constraint prod_table_user_table_uid_fk
        foreign key (uid) references user_table (uid)
);


