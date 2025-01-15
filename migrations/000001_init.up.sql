CREATE TABLE IF NOT EXISTS users
(
    id           UUID PRIMARY KEY,
    username     VARCHAR(100) NOT NULL UNIQUE,
    password     VARCHAR(100) NOT NULL,
    access_level INT CHECK (access_level IN (1, 2)) -- 1: Moderator, 2: Client
);

CREATE TABLE IF NOT EXISTS house
(
    id                UUID PRIMARY KEY,
    street            VARCHAR(100) NOT NULL,
    builder           VARCHAR(100) NOT NULL,
    year_construction DATE         NOT NULL,
    created_at        DATE         NOT NULL DEFAULT CURRENT_DATE,
    update_at         DATE         NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE IF NOT EXISTS flat
(
    id                UUID PRIMARY KEY,
    house_id          UUID NOT NULL REFERENCES house (id),
    price             INT  NOT NULL,
    count_rooms       INT  NOT NULL,
    apartment_number  INT  NOT NULL,
    moderation_status INT CHECK (moderation_status IN (1, 2, 3, 4)) -- 1: Created, 2: Approved, 3:Declined, 4: On moderation
);
