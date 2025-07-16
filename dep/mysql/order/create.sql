CREATE TABLE orders (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    delivery_latitude DOUBLE NOT NULL,
    delivery_longitude DOUBLE NOT NULL,
    client_id BIGINT UNSIGNED NOT NULL,
    drop_off_start BIGINT UNSIGNED NOT NULL,
    drop_off_end BIGINT UNSIGNED NOT NULL,
    region_code VARCHAR(50) NOT NULL,
    route_id BIGINT UNSIGNED NOT NULL
);