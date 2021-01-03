
-- +migrate Up
INSERT INTO parts (class, detail) VALUES ('胸部', '大胸筋');
INSERT INTO parts (class, detail) VALUES ('胸部', '小胸筋');
INSERT INTO parts (class, detail) VALUES ('胸部', '前鋸筋');
INSERT INTO parts (class, detail) VALUES ('背部', '広背筋');
INSERT INTO parts (class, detail) VALUES ('背部', '僧帽筋');
INSERT INTO parts (class, detail) VALUES ('背部', '脊柱起立筋');
INSERT INTO parts (class, detail) VALUES ('肩部', '三角筋');
INSERT INTO parts (class, detail) VALUES ('肩部', '回旋筋腱板');
INSERT INTO parts (class, detail) VALUES ('上腕部', '上腕二頭筋');
INSERT INTO parts (class, detail) VALUES ('上腕部', '上腕筋');
INSERT INTO parts (class, detail) VALUES ('上腕部', '上腕三頭筋');
INSERT INTO parts (class, detail) VALUES ('前腕部', '前腕筋群');
INSERT INTO parts (class, detail) VALUES ('腹部', '腹直筋');
INSERT INTO parts (class, detail) VALUES ('腹部', '腹斜筋');
INSERT INTO parts (class, detail) VALUES ('腹部', '腹横筋');
INSERT INTO parts (class, detail) VALUES ('大腿部', '大腿四頭筋');
INSERT INTO parts (class, detail) VALUES ('大腿部', 'ハムストリングス');
INSERT INTO parts (class, detail) VALUES ('下腿部', '下腿三頭筋');
INSERT INTO parts (class, detail) VALUES ('その他', 'その他');

-- +migrate Down
TRUNCATE TABLE parts CASCADE;