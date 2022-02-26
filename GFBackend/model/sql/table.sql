CREATE TABLE User (
    ID INT NOT NULL AUTO_INCREMENT,
    Username VARCHAR(20) NOT NULL UNIQUE,
    Password VARCHAR(32) NOT NULL,
    Salt VARCHAR(6) NOT NULL,
    Nickname VARCHAR(20),
    Birthday DATE,
    Gender VARCHAR(10), -- male, female, unknown
    Department VARCHAR(100),
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

CREATE TABLE Follow (
    Followee VARCHAR(20) NOT NULL,
    Follower VARCHAR(20) NOT NULL,
    Create_Day DATE,
    PRIMARY KEY (Followee, Follower)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE COMMUNITY (
    ID INT NOT NULL AUTO_INCREMENT,
    Creator VARCHAR(20) NOT NULL,
    Name VARCHAR(20) NOT NULL UNIQUE,
    Description VARCHAR(500),
    Num_Member INT NOT NULL DEFAULT '1',
    Create_Time DATETIME NOT NULL,
    PRIMARY KEY (ID)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- User Private Space Info
CREATE TABLE Space (
    ID INT NOT NULL AUTO_INCREMENT,
    Username VARCHAR(20) NOT NULL,
    Capacity FLOAT(6,2) DEFAULT '10.00',      -- MB
    Remaining FLOAT(6,2) DEFAULT '10.00',      -- MB
    PRIMARY KEY (ID)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;