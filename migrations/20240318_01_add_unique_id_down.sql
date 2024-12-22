-- 1단계: unique_id 관련 제약조건 및 컬럼 제거
ALTER TABLE
    school_notifications DROP CONSTRAINT unique_id_constraint,
    DROP COLUMN unique_id;