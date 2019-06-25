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
(null, "居飛車"),
(null, "振り飛車"),
(null, "三間飛車"),
(null, "四間飛車"),
(null, "向い飛車"),
(null, "中飛車"),
(null, "石田流"),
(null, "早石田"),
(null, "ゴキゲン中飛車"),
(null, "棒銀"),
(null, "右玉"),
(null, "柚飛車"),
(null, "右四間飛車"),
(null, "矢倉３七銀"),
(null, "雀刺し"),
(null, "相掛かり"),
(null, "角換わり"),
(null, "早繰り銀"),
(null, "横歩取り");

CREATE TABLE castling(
    castling_id INT(11) AUTO_INCREMENT NOT NULL, 
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY (castling_id));

INSERT INTO castling (castling_id, name) VALUES 
(null, "なし"),
(null, "美濃囲い"),
(null, "高美濃"),
(null, "銀冠"),
(null, "ダイヤモンド美濃"),
(null, "片美濃"),
(null, "穴熊"),
(null, "金無双"),
(null, "矢倉囲い"),
(null, "銀矢倉"),
(null, "矢倉穴熊"),
(null, "左美濃囲い"),
(null, "天守閣美濃"),
(null, "舟囲い"),
(null, "雁木囲い"),
(null, "ポナンザ囲い"),
(null, "早囲い");