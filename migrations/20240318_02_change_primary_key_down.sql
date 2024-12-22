-- 1단계: 기존 id를 다시 PK로 복구
ALTER TABLE
    school_notifications DROP PRIMARY KEY,
ADD
    PRIMARY KEY (id);

-- 2단계: id 인덱스 제거 (PK가 되므로 불필요)
ALTER TABLE
    school_notifications DROP INDEX idx_school_notifications_id;