-- universities
INSERT INTO universities (id, name) VALUES 
(1, '名城大学'),
(2, '名古屋大学'),
(3, 'ケンブリッジ大学');

-- faculties
INSERT INTO faculties (id, name, university_id) VALUES 
(1, '情報工学部', 1),
(2, '工学部', 1),
(3, '法学部', 2);

-- departments
INSERT INTO departments (id, name, faculty_id) VALUES 
(1, '情報工学科', 1),
(2, '化学科', 1),
(3, '機械工学科', 2);

-- users
INSERT INTO users (id, email, password, name, university_id, faculty_id, department_id, grade) VALUES 
(1, 'user1@example.com', 'password123', '山田太郎', 1, 1, 1, 2),
(2, 'user2@example.com', 'password456', '佐藤花子', 2, 3, NULL, 3);

-- plans
INSERT INTO plans (id, title, content, user_id, created_at, updated_at) VALUES 
(1, '情報学勉強計画', '情報学の基礎から応用まで', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, '法学習得プラン', '憲法と民法を中心に', 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- courses
INSERT INTO courses (id, name, content, plan_id, created_at, updated_at) VALUES 
(1, '力学入門', '運動方程式を学ぶ', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, '憲法概論', '日本国憲法の基礎', 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- posts
INSERT INTO posts (id, content, published, author_id, plan_id) VALUES 
(1, '今日は運動方程式について学びました', true, 1, 1),
(2, '憲法の基本原理についてまとめました', true, 2, 2);

-- favorite_plans
INSERT INTO favorite_plans (id, user_id, plan_id) VALUES 
(1, 2, 1),
(2, 1, 2);