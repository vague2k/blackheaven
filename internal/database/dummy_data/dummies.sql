-- FIXME: GITIGNORE THIS FILE LATER PLEASE
CREATE TABLE inquiries (
id INTEGER PRIMARY KEY AUTOINCREMENT,
topic TEXT NOT NULL,
email TEXT NOT NULL,
name TEXT NOT NULL,
order_number TEXT NOT NULL,
subject TEXT NOT NULL,
content TEXT NOT NULL,
created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
) ;

INSERT INTO inquiries (topic,
email,
content,
name,
subject,
order_number) VALUES
('order',
'user1@example.com',
'Inquiry content 1',
'User 1',
'Subject 1',
'ORD1001'),
('support', 'user2@example.com', 'Inquiry content 2', '', 'Subject 2', ''),
('feedback', 'user3@example.com', 'Inquiry content 3', 'User 3', '', ''),
('order', 'user4@example.com', 'Inquiry content 4', 'User 4', '', 'ORD1002'),
('product', 'user5@example.com', 'Inquiry content 5', '', '', ''),
('support',
'user6@example.com',
'Inquiry content 6',
'User 6',
'Subject 6',
''),
('order',
'user7@example.com',
'Inquiry content 7',
'',
'Subject 7',
'ORD1003'),
('feedback', 'user8@example.com', 'Inquiry content 8', '', '', ''),
('product', 'user9@example.com', 'Inquiry content 9', 'User 9', '', ''),
('order',
'user10@example.com',
'Inquiry content 10',
'User 10',
'Subject 10',
'ORD1004'),
('support', 'user11@example.com', 'Inquiry content 11', '', '', ''),
('feedback',
'user12@example.com',
'Inquiry content 12',
'User 12',
'Subject 12',
''),
('order',
'user13@example.com',
'Inquiry content 13',
'User 13',
'',
'ORD1005'),
('product',
'user14@example.com',
'Inquiry content 14',
'',
'Subject 14',
''),
('support', 'user15@example.com', 'Inquiry content 15', 'User 15', '', ''),
('order', 'user16@example.com', 'Inquiry content 16', '', '', 'ORD1006'),
('feedback',
'user17@example.com',
'Inquiry content 17',
'User 17',
'Subject 17',
''),
('product', 'user18@example.com', 'Inquiry content 18', '', '', ''),
('order',
'user19@example.com',
'Inquiry content 19',
'User 19',
'Subject 19',
'ORD1007'),
('support', 'user20@example.com', 'Inquiry content 20', '', '', ''),
('order', 'user21@example.com', 'Inquiry content 21', '', '', 'ORD1008'),
('product', 'user22@example.com', 'Inquiry content 22', 'User 22', '', ''),
('support',
'user23@example.com',
'Inquiry content 23',
'',
'Subject 23',
''),
('order',
'user24@example.com',
'Inquiry content 24',
'User 24',
'',
'ORD1009'),
('feedback', 'user25@example.com', 'Inquiry content 25', '', '', ''),
('support', 'user26@example.com', 'Inquiry content 26', 'User 26', '', ''),
('order',
'user27@example.com',
'Inquiry content 27',
'',
'Subject 27',
'ORD1010'),
('product', 'user28@example.com', 'Inquiry content 28', '', '', ''),
('feedback',
'user29@example.com',
'Inquiry content 29',
'User 29',
'Subject 29',
''),
('order',
'user30@example.com',
'Inquiry content 30',
'User 30',
'',
'ORD1011'),
('support', 'user31@example.com', 'Inquiry content 31', '', '', ''),
('product', 'user32@example.com', 'Inquiry content 32', 'User 32', '', ''),
('order', 'user33@example.com', 'Inquiry content 33', '', '', 'ORD1012'),
('feedback',
'user34@example.com',
'Inquiry content 34',
'User 34',
'Subject 34',
''),
('support', 'user35@example.com', 'Inquiry content 35', '', '', ''),
('order',
'user36@example.com',
'Inquiry content 36',
'User 36',
'',
'ORD1013'),
('product',
'user37@example.com',
'Inquiry content 37',
'',
'Subject 37',
''),
('feedback', 'user38@example.com', 'Inquiry content 38', 'User 38', '', ''),
('order', 'user39@example.com', 'Inquiry content 39', '', '', 'ORD1014'),
('support',
'user40@example.com',
'Inquiry content 40',
'User 40',
'Subject 40',
''),
('order', 'user41@example.com', 'Inquiry content 41', '', '', 'ORD1015'),
('product', 'user42@example.com', 'Inquiry content 42', 'User 42', '', ''),
('support', 'user43@example.com', 'Inquiry content 43', '', '', ''),
('order',
'user44@example.com',
'Inquiry content 44',
'User 44',
'Subject 44',
'ORD1016'),
('feedback', 'user45@example.com', 'Inquiry content 45', '', '', ''),
('order', 'user46@example.com', 'Inquiry content 46', '', '', 'ORD1017'),
('support', 'user47@example.com', 'Inquiry content 47', 'User 47', '', ''),
('product', 'user48@example.com', 'Inquiry content 48', '', '', ''),
('order',
'user49@example.com',
'Inquiry content 49',
'User 49',
'Subject 49',
'ORD1018'),
('feedback', 'user50@example.com', 'Inquiry content 50', '', '', ''),
('order',
'user51@example.com',
'Inquiry content 51',
'User 51',
'',
'ORD1019'),
('support', 'user52@example.com', 'Inquiry content 52', '', '', ''),
('order', 'user53@example.com', 'Inquiry content 53', '', '', 'ORD1020'),
('product',
'user54@example.com',
'Inquiry content 54',
'User 54',
'Subject 54',
''),
('feedback', 'user55@example.com', 'Inquiry content 55', '', '', ''),
('order',
'user56@example.com',
'Inquiry content 56',
'User 56',
'',
'ORD1021'),
('support',
'user57@example.com',
'Inquiry content 57',
'',
'Subject 57',
''),
('product', 'user58@example.com', 'Inquiry content 58', 'User 58', '', ''),
('order', 'user59@example.com', 'Inquiry content 59', '', '', 'ORD1022'),
('feedback',
'user60@example.com',
'Inquiry content 60',
'User 60',
'Subject 60',
''),
('support', 'user61@example.com', 'Inquiry content 61', '', '', ''),
('order',
'user62@example.com',
'Inquiry content 62',
'User 62',
'',
'ORD1023'),
('product', 'user63@example.com', 'Inquiry content 63', '', '', ''),
('feedback', 'user64@example.com', 'Inquiry content 64', 'User 64', '', ''),
('order',
'user65@example.com',
'Inquiry content 65',
'',
'Subject 65',
'ORD1024'),
('support', 'user66@example.com', 'Inquiry content 66', 'User 66', '', ''),
('order', 'user67@example.com', 'Inquiry content 67', '', '', 'ORD1025'),
('product',
'user68@example.com',
'Inquiry content 68',
'User 68',
'Subject 68',
''),
('feedback', 'user69@example.com', 'Inquiry content 69', '', '', ''),
('order',
'user70@example.com',
'Inquiry content 70',
'User 70',
'',
'ORD1026'),
('support', 'user71@example.com', 'Inquiry content 71', '', '', ''),
('product', 'user72@example.com', 'Inquiry content 72', 'User 72', '', ''),
('order',
'user73@example.com',
'Inquiry content 73',
'',
'Subject 73',
'ORD1027'),
('feedback', 'user74@example.com', 'Inquiry content 74', 'User 74', '', ''),
('support', 'user75@example.com', 'Inquiry content 75', '', '', ''),
('order',
'user76@example.com',
'Inquiry content 76',
'User 76',
'Subject 76',
'ORD1028'),
('product', 'user77@example.com', 'Inquiry content 77', '', '', ''),
('order', 'user78@example.com', 'Inquiry content 78', '', '', 'ORD1029'),
('support', 'user79@example.com', 'Inquiry content 79', 'User 79', '', ''),
('feedback', 'user80@example.com', 'Inquiry content 80', '', '', ''),
('order',
'user81@example.com',
'Inquiry content 81',
'User 81',
'',
'ORD1030'),
('product',
'user82@example.com',
'Inquiry content 82',
'',
'Subject 82',
''),
('support', 'user83@example.com', 'Inquiry content 83', 'User 83', '', ''),
('order', 'user84@example.com', 'Inquiry content 84', '', '', 'ORD1031'),
('feedback',
'user85@example.com',
'Inquiry content 85',
'User 85',
'Subject 85',
''),
('product', 'user86@example.com', 'Inquiry content 86', '', '', ''),
('order',
'user87@example.com',
'Inquiry content 87',
'User 87',
'',
'ORD1032'),
('support', 'user88@example.com', 'Inquiry content 88', '', '', ''),
('order',
'user89@example.com',
'Inquiry content 89',
'',
'Subject 89',
'ORD1033'),
('product', 'user90@example.com', 'Inquiry content 90', 'User 90', '', ''),
('feedback', 'user91@example.com', 'Inquiry content 91', '', '', ''),
('order',
'user92@example.com',
'Inquiry content 92',
'User 92',
'Subject 92',
'ORD1034'),
('support', 'user93@example.com', 'Inquiry content 93', '', '', ''),
('order', 'user94@example.com', 'Inquiry content 94', '', '', 'ORD1035'),
('product',
'user95@example.com',
'Inquiry content 95',
'User 95',
'Subject 95',
''),
('feedback', 'user96@example.com', 'Inquiry content 96', '', '', ''),
('order',
'user97@example.com',
'Inquiry content 97',
'User 97',
'',
'ORD1036'),
('support', 'user98@example.com', 'Inquiry content 98', '', '', ''),
('product', 'user99@example.com', 'Inquiry content 99', 'User 99', '', ''),
('order',
'user100@example.com',
'Inquiry content 100',
'',
'Subject 100',
'ORD1037') ;
