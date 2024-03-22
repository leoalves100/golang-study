-- Insert informations in users table
INSERT INTO users 
(name, nick, mail, password)
VALUES
("User 01", "user_01", "user_01@gmail.com", "$2a$10$JcNHE6OI9qfNl4XoPAtEsuENylNJCk4ybZORtEjtUYHUdwGh/q9e2"),
("User 02", "user_02", "user_02@gmail.com", "$2a$10$JcNHE6OI9qfNl4XoPAtEsuENylNJCk4ybZORtEjtUYHUdwGh/q9e2"),
("User 03", "user_03", "user_03@gmail.com", "$2a$10$JcNHE6OI9qfNl4XoPAtEsuENylNJCk4ybZORtEjtUYHUdwGh/q9e2")

-- Insert informations in followers table
INSERT INTO  followers 
(user_id, followers_id)
VALUES
(1, 2),
(3, 1),
(1, 3)