-- 1단계: PK 변경
ALTER TABLE
    school_notifications DROP PRIMARY KEY,
ADD
    PRIMARY KEY (unique_id);

-- 2단계: 기존 id 컬럼에 인덱스 추가
ALTER TABLE
    school_notifications
ADD
    INDEX idx_school_notifications_id (id);