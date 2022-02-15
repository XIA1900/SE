CREATE TABLE USER (
    ID INT NOT NULL AUTO_INCREMENT,
    USERNAME VARCHAR(20) NOT NULL UNIQUE,
    PASSWORD VARCHAR(20) NOT NULL,
    SALT VARCHAR(6) NOT NULL,
    NICKNAME VARCHAR(20) NOT NULL,
    BIRTHDAY DATE NOT NULL,
    GENDER VARCHAR(10), -- male, female, unknown
    DEPARTMENT VARCHAR(100),
    PRIMARY KEY (ID)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Role Management Table
-- Using Casbin
CREATE TABLE `casbin_rule` (
    `id` bigint unsigned AUTO_INCREMENT,
    `ptype` varchar(100),
    `v0` varchar(100),
    `v1` varchar(100),
    `v2` varchar(100),
    `v3` varchar(100),
    `v4` varchar(100),
    `v5` varchar(100),
    PRIMARY KEY (`id`),
    UNIQUE INDEX unique_index (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

