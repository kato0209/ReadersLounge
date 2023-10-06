INSERT INTO users DEFAULT VALUES;
INSERT INTO users DEFAULT VALUES;

-- user_details テーブルにユーザー詳細を挿入
INSERT INTO user_details (user_id, name, profile_text, profile_image) VALUES (2, 'John Doe', 'Hello, I am John!', 'https://res.cloudinary.com/dvh5ehszr/image/upload/v1689442197/media/john_doe.png');
INSERT INTO user_details (user_id, name, profile_text, profile_image) VALUES (3, 'Jane Smith', 'Hi there!', 'https://res.cloudinary.com/dvh5ehszr/image/upload/v1689442197/media/jane_smith.png');

-- user_auths テーブルに認証情報を挿入
INSERT INTO user_auths (user_id, identity_type, identifier, credential) VALUES (2, 'email', 'john@example.com', 'hashed_password_1');
INSERT INTO user_auths (user_id, identity_type, identifier, credential) VALUES (3, 'email', 'jane@example.com', 'hashed_password_2');

INSERT INTO books (ISBNcode, title, author, price, publisher, published_at, item_url, image)
VALUES
    ('978-0-13-110362-7', 'Craftsmanship', 'Robert C. Martin', '29.99', 'Prentice', '2008-08-11', 'https://www.example.com/book1', 'https://www.example.com/book1.jpg'),
    ('978-0-321-14618-7', 'Mastery', 'Andrew Hunt, David Thomas', '34.99', 'Addison-Wesley', '1999-10-30', 'https://www.example.com/book2', 'https://www.example.com/book2.jpg'),
    ('978-0-596-00797-3', 'DObject-Oriented', 'Erich Gamma', '44.99', 'Addison-Wesley', '1994-11-10', 'https://www.example.com/book3', 'https://www.example.com/book3.jpg');

-- posts テーブルに投稿を挿入
INSERT INTO posts (user_id, book_id, created_at) VALUES (1, 1, '2023-10-07');
INSERT INTO posts (user_id, book_id, created_at) VALUES (3, 2, '2023-10-07');

-- post_details テーブルに投稿詳細を挿入
INSERT INTO post_details (post_id, content, rating, image, updated_at) VALUES (3, 'This book is great!', 5, 'https://www.example.com/image1.jpg', NULL);
INSERT INTO post_details (post_id, content, rating, image, updated_at) VALUES (4, 'Highly recommend!', 4, 'https://www.example.com/image2.jpg', NULL);