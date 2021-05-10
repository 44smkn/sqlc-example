CREATE TABLE chair (
    id          BIGINT  NOT NULL AUTO_INCREMENT,
    name        TEXT    NOT NULL,
    description TEXT    NOT NULL,
    thumbnail   TEXT    NOT NULL,
    price       INT     NOT NULL,
    height      INT     NOT NULL,
    width       INT     NOT NULL,
    depth       INT     NOT NULL,
    color       TEXT    NOT NULL,
    features    TEXT    NOT NULL,
    kind        TEXT    NOT NULL,
    popularity  INT     NOT NULL,
    stock       INT     NOT NULL,
    PRIMARY KEY (id)
);