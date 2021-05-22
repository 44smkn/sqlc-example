CREATE TABLE chair (
    id          BIGINT  NOT NULL AUTO_INCREMENT,
    name        TEXT    NOT NULL,
    description TEXT    NOT NULL,
    thumbnail   TEXT    NOT NULL,
    price       INT             ,
    height      INT             ,
    width       INT             ,
    depth       INT             ,
    color       TEXT            ,
    features    TEXT            ,
    kind        TEXT            ,
    popularity  INT     NOT NULL,
    stock       INT     NOT NULL,
    PRIMARY KEY (id)
);