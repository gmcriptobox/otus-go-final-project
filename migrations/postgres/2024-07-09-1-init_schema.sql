create table black_list
(
    id            serial PRIMARY KEY,
    ip            varchar(15) NOT NULL,
    mask          varchar(2)  NOT NULL,
    binary_prefix varchar(32) NOT NULL,
    created_at    timestamp DEFAULT now(),
    UNIQUE (binary_prefix, mask)
);

create table white_list
(
    id            serial PRIMARY KEY,
    ip            varchar(15) NOT NULL,
    mask          varchar(2)  NOT NULL,
    binary_prefix varchar(32) NOT NULL,
    created_at    timestamp DEFAULT now(),
    UNIQUE (binary_prefix, mask)
);
