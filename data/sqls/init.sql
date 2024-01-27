USE share_database;

CREATE TABLE user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    createdAt DATETIME,
    upcatedAt DATETIME,
    deletedAt DATETIME
);

CREATE TABLE authorization (
    loginId VARCHAR(100) PRIMARY KEY,
    userId INT NOT NULL,
    password VARCHAR(20) NOT NULL,
    createdAt DATETIME,
    updatedAt DATETIME,
    FOREIGN KEY fk_auth_userId(userId) REFERENCES user(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE summary (
    id INT AUTO_INCREMENT PRIMARY KEY,
    userId INT NOT NULL,
    title VARCHAR(50) NOT NULL,
    markdown TEXT NOT NULL,
    FOREIGN KEY fk_summary_userId(userId) REFERENCES user(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE paper_authors (
    id INT AUTO_INCREMENT PRIMARY KEY,
    paperId INT NOT NULL,
    userId INT NOT NULL,
    FOREIGN KEY fk_pa_paperId(paperId) REFERENCES summary(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY fk_pa_userId(userId) REFERENCES user(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE author (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);
