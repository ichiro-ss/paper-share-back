USE share_database;

CREATE TABLE user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    createdAt DATETIME,
    updatedAt DATETIME,
    deletedAt DATETIME
);

CREATE TABLE authorizations (
    loginId VARCHAR(100) PRIMARY KEY,
    userId INT NOT NULL,
    password VARCHAR(20) NOT NULL,
    createdAt DATETIME,
    updatedAt DATETIME,
    FOREIGN KEY fk_auth_userId(userId) REFERENCES user(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE summaries (
    id INT AUTO_INCREMENT PRIMARY KEY,
    userId INT NOT NULL,
    title VARCHAR(50) NOT NULL,
    markdown TEXT NOT NULL,
    FOREIGN KEY fk_summary_userId(userId) REFERENCES user(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE authors (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE summary_authors (
    id INT AUTO_INCREMENT PRIMARY KEY,
    summaryId INT NOT NULL,
    authorId INT NOT NULL,
    FOREIGN KEY fk_pa_summaryId(summaryId) REFERENCES summaries(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY fk_pa_authorId(authorId) REFERENCES authors(id) ON UPDATE CASCADE ON DELETE CASCADE
);
