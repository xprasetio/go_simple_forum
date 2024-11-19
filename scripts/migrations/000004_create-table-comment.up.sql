CREATE TABLE IF NOT EXISTS comments (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    post_id INT NOT NULL,
    comment_content LONGTEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    created_by VARCHAR(255) NOT NULL,
    updated_by VARCHAR(255) NOT NULL,
    CONSTRAINT fk_post_id_comments FOREIGN KEY (post_id) REFERENCES posts(id),
    CONSTRAINT fk_user_id_comments FOREIGN KEY (user_id) REFERENCES users(id)

);