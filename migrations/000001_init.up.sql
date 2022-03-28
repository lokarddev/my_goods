create table goods
(
    id          serial
        primary key,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    title       text,
    description text
);

create table dishes
(
    id          serial
        primary key,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    title       text,
    description text
);

create table lists
(
    id          serial
        primary key,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    title       text,
    description text
);

create table dish_goods
(
    dish_id  int not null
        constraint fk_dish_goods_dish
            references dishes,
    goods_id int not null
        constraint fk_dish_goods_goods
            references goods,
    primary key (dish_id, goods_id)
);

create table list_dishes
(
    list_id int not null
        constraint fk_list_dishes_list
            references lists,
    dish_id int not null
        constraint fk_list_dishes_dish
            references dishes,
    primary key (list_id, dish_id)
);
