# 취Go Backend Server

대학생 대상 교육 및 편의 기능 지원 서비스 [취Go] 서버 레포지토리

---

## 주요 기능

- 사용자 인증 및 관리
  - 회원가입/로그인
  - JWT 기반 인증
- 이력서 관리
  - 이력서 작성 및 조회
  - 원티드, 프로그래머스, 링크드인 형식 변환 지원
- 학교/학과 공지사항 알림
  - 웹 크롤링을 통한 공지사항 수집 배치 작업
  - 텔레그램 봇을 통한 키워드 기반 알림 발송
- 교육 콘텐츠 제공
  - YouTube 플레이리스트 기반 강의 관리
  - 강의 목록 및 상세 정보 제공

## 실행 방법

1. 환경 설정

env 파일 생성

```bash
touch .env
```

환경변수 설정

```
DB_HOST
DB_USERNAME
DB_PASSWORD
DB_NAME
JWT_SECRET
YOUTUBE_API_KEY
REDIS_ADDR
SCHOOL_NOTICE_URL
DEPT_NOTICE_URL_1
DEPT_NOTICE_URL_2
DEPT_NOTICE_URL_3
```

2. 서버 실행
   로컬 실행

```bash
go run main.go
```

Docker 실행

```bash
docker build -t chee-go-backend .
docker run -p 8080:8080 chee-go-backend
```

## 프로젝트 구조

```
.
├── internal/                  # 내부 패키지
│   ├── config/               # 환경설정 및 초기화
│   ├── domain/
│   │   ├── entity/          # 도메인 모델
│   │   ├── repository/      # 레포지토리 인터페이스
│   │   └── service/         # 서비스 인터페이스
│   ├── http/
│   │   ├── dto/            # 요청/응답 데이터 구조체
│   │   └── handler/        # HTTP 핸들러
│   ├── infrastructure/
│   │   ├── crawler/        # 웹 크롤링
│   │   ├── cron/          # 크론 작업
│   │   ├── telegram/      # 텔레그램 봇
│   │   └── youtube/       # YouTube API
│   ├── repository/         # 레포지토리 구현체
│   └── service/           # 서비스 구현체
├── test/
│   ├── e2e/              # E2E 테스트
│   ├── mock/             # 테스트용 Mock
│   └── util/             # 테스트 유틸리티
└── deploy.sh             # 배포 스크립트
```

## 기술스택

- 언어 및 프레임워크
  - Go 1.23
  - Gin Web Framework
- 데이터베이스
  - MySQL
  - Redis (알림 상태 관리)
- 인프라
  - Docker
  - GitHub Actions (CI/CD)
  - AWS EC2, ECR
- 외부 서비스
  - YouTube Data API
  - Telegram Bot API
- 테스트
  - Go testing
  - Testify
