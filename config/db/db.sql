-- Create the members table
CREATE TABLE members (
  ID_MEMBER INT AUTO_INCREMENT PRIMARY KEY,
  USERNAME VARCHAR(255) NOT NULL,
  GENDER VARCHAR(255) NOT NULL,
  SKINTYPE VARCHAR(255) NOT NULL,
  SKINCOLOR VARCHAR(255) NOT NULL
);

-- Create the products table
CREATE TABLE products (
  ID_PRODUCT INT AUTO_INCREMENT PRIMARY KEY,
  PRODUCT_NAME VARCHAR(255) NOT NULL,
  PRICE DECIMAL(10, 2) NOT NULL
);

-- Create the review_products table
CREATE TABLE review_products (
  ID_REVIEW INT AUTO_INCREMENT PRIMARY KEY,
  ID_MEMBER INT NOT NULL,
  ID_PRODUCT INT NOT NULL,
  DESC_REVIEW TEXT,
  FOREIGN KEY (ID_MEMBER) REFERENCES members (ID_MEMBER) ON DELETE CASCADE,
  FOREIGN KEY (ID_PRODUCT) REFERENCES products (ID_PRODUCT) ON DELETE CASCADE
);

-- Create the like_review table
CREATE TABLE like_reviews (
  ID_LIKE INT AUTO_INCREMENT PRIMARY KEY,
  ID_REVIEW INT NOT NULL,
  ID_MEMBER INT NOT NULL,
  FOREIGN KEY (ID_REVIEW) REFERENCES review_products (ID_REVIEW) ON DELETE CASCADE,
  FOREIGN KEY (ID_MEMBER) REFERENCES members (ID_MEMBER) ON DELETE CASCADE
);

-- Insert dummy data into the members table
INSERT INTO members (USERNAME, GENDER, SKINTYPE, SKINCOLOR)
VALUES
  ('User1', 'Male', 'Oily', 'Fair'),
  ('User2', 'Female', 'Combination', 'Medium'),
  ('User3', 'Male', 'Dry', 'Dark'),
  ('User4', 'Female', 'Normal', 'Fair'),
  ('User5', 'Male', 'Oily', 'Medium'),
  ('User6', 'Female', 'Combination', 'Fair'),
  ('User7', 'Male', 'Dry', 'Dark'),
  ('User8', 'Female', 'Oily', 'Medium'),
  ('User9', 'Male', 'Combination', 'Fair'),
  ('User10', 'Female', 'Normal', 'Dark');

-- Insert dummy data into the products table
INSERT INTO products (PRODUCT_NAME, PRICE)
VALUES
  ('Product1', 9.99),
  ('Product2', 19.99),
  ('Product3', 14.99),
  ('Product4', 24.99),
  ('Product5', 29.99);

-- Insert dummy data into the review_products table
INSERT INTO review_products (ID_MEMBER, ID_PRODUCT, DESC_REVIEW)
VALUES
  (1, 1, 'Great product! Highly recommended.'),
  (2, 1, 'Average product. Could be better.'),
  (3, 2, 'Excellent quality and value.'),
  (4, 2, 'Not satisfied with the product.'),
  (5, 3, 'Works well for my skin type.'),
  (6, 3, 'Didn''t see any noticeable results.'),
  (7, 4, 'Amazing product! Will repurchase.'),
  (8, 4, 'Didn''t work for me.'),
  (9, 5, 'Impressed with the packaging and performance.'),
  (10, 5, 'Disappointed with the product.');

-- Insert dummy data into the like_reviews table
INSERT INTO like_reviews (ID_MEMBER, ID_REVIEW)
VALUES
  (1, 1),
  (2, 1),
  (3, 2),
  (4, 3),
  (5, 3),
  (6, 4),
  (7, 5),
  (8, 6),
  (9, 6),
  (10, 7);
