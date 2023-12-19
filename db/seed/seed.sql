INSERT INTO users DEFAULT VALUES;
INSERT INTO users DEFAULT VALUES;

-- user_details テーブルにユーザー詳細を挿入
INSERT INTO user_details (user_id, name, profile_text) VALUES (1, 'John Doe', 'Hello, I am John!');
INSERT INTO user_details (user_id, name, profile_text) VALUES (2, 'Jane Smith', 'Hi there!');

-- user_auths テーブルに認証情報を挿入
INSERT INTO user_auths (user_id, identity_type, identifier, credential) VALUES (1, 'email', 'john@example.com', 'hashed_password_1');
INSERT INTO user_auths (user_id, identity_type, identifier, credential) VALUES (2, 'email', 'jane@example.com', 'hashed_password_2');

INSERT INTO books (ISBNcode, title, author, price, publisher, published_at, item_url, image)
VALUES
    ('978-0-13-110362-7', 'Craftsmanship', 'Robert C. Martin', '29.99', 'Prentice', '2008-08-11', 'https://www.example.com/book1', 'https://www.example.com/book1.jpg'),
    ('978-0-321-14618-7', 'Mastery', 'Andrew Hunt, David Thomas', '34.99', 'Addison-Wesley', '1999-10-30', 'https://www.example.com/book2', 'https://www.example.com/book2.jpg');

-- posts テーブルに投稿を挿入
INSERT INTO posts (user_id, book_id) VALUES (1, 1);
INSERT INTO posts (user_id, book_id) VALUES (2, 2);

-- post_details テーブルに投稿詳細を挿入
INSERT INTO post_details (post_id, content, rating, image) VALUES (1, 'This book is great!', 5, 'https://www.example.com/image1.jpg');
INSERT INTO post_details (post_id, content, rating, image) VALUES (2, 'Highly recommend!', 4, 'https://www.example.com/image2.jpg');

-- chat_rooms テーブルにチャットルームを挿入
INSERT INTO chat_rooms DEFAULT VALUES;

-- entries テーブルにチャットルーム参加者を挿入
INSERT INTO entries (chat_room_id, user_id) VALUES (1, 1);

-- chat_messages テーブルにチャットメッセージを挿入
INSERT INTO chat_messages (chat_room_id, user_id, content) VALUES (1, 1, 'Hello!');
