
-- +migrate Up
-- common_classeses
INSERT INTO common_classes (class) VALUES ('胸部');
INSERT INTO common_classes (class) VALUES ('背部');
INSERT INTO common_classes (class) VALUES ('肩部');
INSERT INTO common_classes (class) VALUES ('上腕部');
INSERT INTO common_classes (class) VALUES ('前腕部');
INSERT INTO common_classes (class) VALUES ('腹部');
INSERT INTO common_classes (class) VALUES ('大腿部');
INSERT INTO common_classes (class) VALUES ('下腿部');
INSERT INTO common_classes (class) VALUES ('その他');

-- 胸部
WITH chest AS (
    SELECT id FROM common_classes WHERE class = '胸部'
)
INSERT INTO common_parts (class, part) VALUES (chest, '');
INSERT INTO common_parts (class, part) VALUES (chest, '大胸筋');
INSERT INTO common_parts (class, part) VALUES (chest, '小胸筋');
INSERT INTO common_parts (class, part) VALUES (chest, '前鋸筋');
-- 背部
WITH back AS (
    SELECT id FROM common_classes WHERE class = '背部'
)
INSERT INTO common_parts (class, part) VALUES (back, '');
INSERT INTO common_parts (class, part) VALUES (back, '広背筋');
INSERT INTO common_parts (class, part) VALUES (back, '僧帽筋');
INSERT INTO common_parts (class, part) VALUES (back, '脊柱起立筋');
-- 肩部
WITH shoulder AS (
    SELECT id FROM common_classes WHERE class = '肩部'
)
INSERT INTO common_parts (class, part) VALUES (shoulder, '');
INSERT INTO common_parts (class, part) VALUES (shoulder, '三角筋');
INSERT INTO common_parts (class, part) VALUES (shoulder, '回旋筋腱板');
-- 上腕部
WITH brachium AS (
    SELECT id FROM common_classes WHERE class = '上腕部'
)
INSERT INTO common_parts (class, part) VALUES (brachium, '');
INSERT INTO common_parts (class, part) VALUES (brachium, '上腕二頭筋');
INSERT INTO common_parts (class, part) VALUES (brachium, '上腕三頭筋');
-- 前腕部
WITH antebrachium AS (
    SELECT id FROM common_classes WHERE class = '前腕部'
)
INSERT INTO common_parts (class, part) VALUES (antebrachium, '');
-- 腹部
WITH abdominal AS (
    SELECT id FROM common_classes WHERE class = '腹部'
)
INSERT INTO common_parts (class, part) VALUES (abdominal, '');
INSERT INTO common_parts (class, part) VALUES (abdominal, '腹直筋');
INSERT INTO common_parts (class, part) VALUES (abdominal, '腹斜筋');
INSERT INTO common_parts (class, part) VALUES (abdominal, '腹横筋');
-- 大腿部
WITH femur AS (
    SELECT id FROM common_classes WHERE class = '大腿部'
)
INSERT INTO common_parts (class, part) VALUES (femur, '');
INSERT INTO common_parts (class, part) VALUES (femur, '大腿四頭筋');
INSERT INTO common_parts (class, part) VALUES (femur, 'ハムストリングス');
-- 下腿部
WITH crus AS (
    SELECT id FROM common_classes WHERE class = '下腿部'
)
INSERT INTO common_parts (class, part) VALUES (crus, '');
INSERT INTO common_parts (class, part) VALUES (crus, '下腿三頭筋');
-- その他
WITH others AS (
    SELECT id FROM common_classes WHERE class = 'その他'
)
INSERT INTO common_parts (class, part) VALUES (others, '');

-- status
INSERT INTO status (id, state) VALUES (0, '平常');
INSERT INTO status (id, state) VALUES (1, '疲労');
INSERT INTO status (id, state) VALUES (2, '重疲労');

-- +migrate Down
TRUNCATE TABLE common_parts CASCADE;
TRUNCATE TABLE common_classes CASCADE;