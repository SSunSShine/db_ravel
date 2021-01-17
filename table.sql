use travel;

create table customer
(
    cust_name varchar(50)   not null
        primary key,
    password  varchar(50)   not null,
    type      int default 0 null
);

create table city
(
    city_name varchar(20) not null
        primary key
);

create table car
(
    car_num   varchar(50)             not null
        primary key,
    price     varchar(50) default '0' not null,
    city_name varchar(50)             not null
);

create table flight
(
    flight_num varchar(50)   not null
        primary key,
    price      int default 0 null,
    num_seats  int default 0 null,
    from_city  varchar(50)   not null,
    ariv_city  varchar(50)   not null
);

create table hotel
(
    hotel_name varchar(50)   not null
        primary key,
    price      int default 0 null,
    num_rooms  int default 0 null,
    city_name  varchar(50)   not null
);

create table reservation
(
    resv_key  varchar(50) not null
        primary key,
    cust_name varchar(50) not null,
    type      int         not null,
    res_date  date        null
);


