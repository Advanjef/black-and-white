CREATE TABLE player(
    player_id INT(11) AUTO_INCREMENT NOT NULL, 
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY (player_id));


CREATE TABLE opening(
    opening_id INT(11) AUTO_INCREMENT NOT NULL, 
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY (opening_id));

INSERT INTO opening (opening_id, name) VALUES 
(null, "なし"),
(null, "居飛車棒銀"),
(null, "三間飛車"),
(null, "四間飛車"),
(null, "向い飛車"),
(null, "中飛車");

CREATE TABLE castling(
    castling_id INT(11) AUTO_INCREMENT NOT NULL, 
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY (castling_id));

INSERT INTO castling (castling_id, name) VALUES 
(null, "なし"),
(null, "美濃囲い"),
(null, "銀冠"),
(null, "穴熊"),
(null, "金無双"),
(null, "矢倉囲い"),
(null, "舟囲い"),
(null, "雁木囲い"),
(null, "早囲い");