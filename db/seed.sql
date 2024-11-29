-- universities
INSERT INTO universities (id, name) VALUES 
(1, '名城大学'),
(2, '名古屋大学'),
(3, 'ケンブリッジ大学');
-- faculties
INSERT INTO faculties (id, name) VALUES 
(1, '工学部'),
(2, '法学部'),
(3, '理学部');
-- departments
INSERT INTO departments (id, name) VALUES 
(1, '情報学科'),
(2, '機械工学科'),
(3, '法律学科'),
(4, '物理学科');
-- users
INSERT INTO users (id, email, password, name, university_id, faculty_id, department_id, grade) VALUES 
(1, 'user1@example.com', 'password123', '山田太郎', 1, 1, 1, 2), 
(2, 'user2@example.com', 'password456', '佐藤花子', 2, 2, 3, 3);
-- plans
INSERT INTO plans (id, title, content, user_id, created_at, updated_at) VALUES 
(1, '情報学勉強計画', '情報学の基礎から応用まで', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, '法学習得プラン', '憲法と民法を中心に', 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
-- courses
INSERT INTO courses (id, name, content, plan_id, created_at, updated_at) VALUES 
(1, '力学入門', '運動方程式を学ぶ', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, '憲法概論', '日本国憲法の基礎', 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
-- posts
INSERT INTO posts (id, content, plan_id, created_at, author_id) VALUES 
(1, '力学入門の勉強法を教えてください', 1, CURRENT_TIMESTAMP, 1),
(2, '憲法概論の勉強法を教えてください', 2, CURRENT_TIMESTAMP, 2);
-- favorite_plans
INSERT INTO favorite_plans (id, user_id, plan_id) VALUES
(1, 2, 1), 
(2, 1, 2);