CREATE DATABASE IF NOT EXISTS online_shop;
USE online_shop;

CREATE TABLE IF NOT EXISTS members (
    member_id VARCHAR(50) NOT NULL,
    username VARCHAR(50) NOT NULL,
    gender ENUM("Male", "Female"),
    skin_type ENUM("Normal", "Sensitive", "Dry", "Oily") NOT NULL,
    skin_color ENUM("White", "Yellow", "Black") NOT NULL,
    CONSTRAINT PK_Member PRIMARY KEY (member_id)
);

CREATE TABLE IF NOT EXISTS products (
    product_id VARCHAR(50) NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    price DECIMAL(12, 3) NOT NULL,
    CONSTRAINT PK_Product PRIMARY KEY (product_id)
);

CREATE TABLE IF NOT EXISTS reviews (
    review_id VARCHAR(50) NOT NULL,
    member_id VARCHAR(50) NOT NULL,
    product_id VARCHAR(50) NOT NULL,
    desc_review VARCHAR(500),
    CONSTRAINT PK_Review PRIMARY KEY (review_id),
    CONSTRAINT FK_Review_Member FOREIGN KEY (member_id) 
    REFERENCES members(member_id)
    ON DELETE CASCADE,
    CONSTRAINT FK_Review_Product FOREIGN KEY (product_id) 
    REFERENCES products(product_id)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS like_reviews (
    review_id VARCHAR(50) NOT NULL,
    member_id VARCHAR(50) NOT NULL,
    CONSTRAINT PK_Review PRIMARY KEY (review_id, member_id),
    CONSTRAINT FK_LikeReview_Review FOREIGN KEY (review_id) 
    REFERENCES reviews(review_id)
    ON DELETE CASCADE,
    CONSTRAINT FK_LikeReview_Member FOREIGN KEY (member_id) 
    REFERENCES members(member_id)
    ON DELETE CASCADE
);
