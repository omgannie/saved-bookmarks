DROP TABLE IF EXISTS bookmark;
CREATE TABLE bookmark (
    id INT AUTO_INCREMENT NOT NULL,
    link_text VARCHAR(128) NOT NULL,
    href VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

INSERT INTO bookmark
(link_text, href, description)
VALUES
('Learn go', 'https://go.dev', 'Build simple, secure, scalable systems with Go'),
('Modern JavaScript', 'https://javascript.info', 'How its done now!'),
('devhints', 'https://devhints.io', 'Ricos Cheatsheets');
    