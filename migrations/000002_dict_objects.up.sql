create table measure
(
    id          serial
        primary key,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    code       varchar(100),
    value       varchar(100)
);

alter table goods add column measure_id int references measure;
alter table dish_goods add column amount int;
create table list_goods
(
    list_id int not null
        constraint fk_list_dishes_list
            references lists,
    goods_id int not null
        constraint fk_dish_goods_goods
            references goods,
    amount int,
    primary key (goods_id, list_id)
);
