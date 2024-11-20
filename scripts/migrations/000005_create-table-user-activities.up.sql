CREATE TABLE IF NOT EXISTS user_activities (
    id INT PRIMARY KEY AUTO_INCREMENT,
    post_id INT NOT NULL,
    user_id BIGINT NOT NULL,
    is_liked BOOLEAN NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by LONGTEXT NOT NULL,
    updated_by LONGTEXT NOT NULL,
    CONSTRAINT fk_post_id_user_activities FOREIGN KEY (post_id) REFERENCES posts(id),
    CONSTRAINT fk_user_id_user_activities FOREIGN KEY (user_id) REFERENCES users(id)
);