package crawler

import (
	"chee-go-backend/internal/domain/entity"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	SchoolNoticePrefix = "hiu-gen"     // 학교 공지사항 prefix
	DeptNoticePrefix1  = "hiu-dep-gen" // 학과 일반 공지사항 prefix
	DeptNoticePrefix2  = "hiu-dep-sch" // 학과 장학 공지사항 prefix
	DeptNoticePrefix3  = "hiu-dep-job" // 학과 취업 공지사항 prefix
)

// 게시판 타입 정의
type BoardType struct {
	URL    string
	Prefix string
}

type Crawler interface {
	// 학교 공지사항 크롤링
	FetchSchoolNotices() ([]entity.SchoolNotification, error)
	// 학과 공지사항 크롤링 (3개 게시판)
	FetchDepartmentNotices() ([]entity.SchoolNotification, error)
}

type crawler struct {
	SchoolNoticeURL string      // 학교 공지사항 URL
	DeptNoticeURLs  []BoardType // 학과 공지사항 URL들과 prefix
}

func NewCrawler(schoolURL string, deptURLs []string) Crawler {
	// 각 게시판별 prefix 설정
	deptBoards := []BoardType{
		{URL: deptURLs[0], Prefix: DeptNoticePrefix1}, // 일반 공지
		{URL: deptURLs[1], Prefix: DeptNoticePrefix2}, // 장학 공지
		{URL: deptURLs[2], Prefix: DeptNoticePrefix3}, // 취업 공지
	}

	return &crawler{
		SchoolNoticeURL: schoolURL,
		DeptNoticeURLs:  deptBoards,
	}
}

// 학교 공지사항 크롤링
func (c *crawler) FetchSchoolNotices() ([]entity.SchoolNotification, error) {
	resp, err := http.Get(c.SchoolNoticeURL)
	if err != nil {
		return nil, fmt.Errorf("HTTP 요청 실패: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("잘못된 상태 코드: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("페이지 파싱 실패: %v", err)
	}

	var notices []entity.SchoolNotification
	today := time.Now().In(time.FixedZone("KST", 9*60*60)).Format("2006.01.02")

	log.Printf("학교 공지사항 크롤링 시작 - URL: %s, 오늘 날짜: %s", c.SchoolNoticeURL, today)

	// 기본 URL 추출 (쿼리스트링 제외)
	baseURL := c.SchoolNoticeURL
	if idx := strings.Index(baseURL, "?"); idx != -1 {
		baseURL = baseURL[:idx]
	}

	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		// 날짜는 td:nth-child(4)에 있음
		dateStr := strings.TrimSpace(s.Find("td:nth-child(4)").Text())
		log.Printf("발견된 날짜: '%s'", dateStr)

		// 날짜 형식이 YYYY.MM.DD 형식임
		if dateStr != today {
			return
		}

		// 공지사항 ID는 td:nth-child(1)에 있음
		rawID := strings.TrimSpace(s.Find("td:nth-child(1)").Text())
		id := fmt.Sprintf("%s-%s", SchoolNoticePrefix, rawID)

		// 제목과 URL은 td:nth-child(2) 내부의 a 태그에 있음
		titleCell := s.Find("td:nth-child(2)")
		title := strings.TrimSpace(titleCell.Find(".b-title").Text())
		detailURL, exists := titleCell.Find("a").Attr("href")
		if !exists {
			return
		}

		// URL이 상대 경로인 경우 절대 경로로 변환
		fullURL := detailURL
		if !strings.HasPrefix(detailURL, "http") {
			if strings.HasPrefix(detailURL, "?") {
				// 쿼리스트링만 있는 경우 기존 URL에 추가
				fullURL = baseURL + detailURL
			} else if strings.HasPrefix(detailURL, "/") {
				// 절대 경로인 경우
				fullURL = fmt.Sprintf("https://www.hongik.ac.kr%s", detailURL)
			} else {
				// 상대 경로인 경우
				fullURL = fmt.Sprintf("https://www.hongik.ac.kr/%s", detailURL)
			}
		}

		content, err := c.fetchNoticeContent(fullURL)
		if err != nil {
			log.Printf("상세 내용 크롤링 실패: %v", err)
			content = "" // 실패 시 빈 문자열로 설정
		}

		notice := entity.SchoolNotification{
			ID:      id,
			Title:   title,
			Url:     fullURL,
			Date:    parseDate(dateStr),
			Content: content,
		}

		log.Printf("발견된 공지사항: %+v", notice)
		notices = append(notices, notice)
	})

	log.Printf("학교 공지사항 총 %d개 발견", len(notices))
	return notices, nil
}

// 날짜 문자열을 time.Time으로 파싱하는 헬퍼 함수
func parseDate(dateStr string) time.Time {
	t, err := time.Parse("2006.01.02", dateStr)
	if err != nil {
		log.Printf("날짜 파싱 실패: %v", err)
		return time.Time{}
	}
	return t
}

// 학과 공지사항 크롤링 (3개 게시판 통합)
func (c *crawler) FetchDepartmentNotices() ([]entity.SchoolNotification, error) {
	var allNotices []entity.SchoolNotification

	for _, board := range c.DeptNoticeURLs {
		notices, err := c.fetchDeptNoticesFromURL(board.URL, board.Prefix)
		if err != nil {
			log.Printf("게시판 크롤링 실패 - URL: %s, 에러: %v", board.URL, err)
			continue
		}
		allNotices = append(allNotices, notices...)
	}

	return allNotices, nil
}

// 개별 학과 게시판 크롤링
func (c *crawler) fetchDeptNoticesFromURL(url string, prefix string) ([]entity.SchoolNotification, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP 요청 실패: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("잘못된 상태 코드: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("페이지 파싱 실패: %v", err)
	}

	var notices []entity.SchoolNotification
	today := time.Now().In(time.FixedZone("KST", 9*60*60)).Format("2006.01.02")

	log.Printf("학과 공지사항 크롤링 시작 - URL: %s, 오늘 날짜: %s", url, today)

	// 기본 URL 추출 (쿼리스트링 제외)
	baseURL := url
	if idx := strings.Index(url, "?"); idx != -1 {
		baseURL = url[:idx]
	}

	doc.Find("div.bn-list-common table tbody tr").Each(func(i int, s *goquery.Selection) {
		// 날짜는 td:nth-child(4)에 있음
		dateStr := strings.TrimSpace(s.Find("td:nth-child(4)").Text())
		log.Printf("발견된 날짜: '%s'", dateStr)

		// 날짜 형식이 YYYY.MM.DD 형식임
		if dateStr != today {
			return
		}

		// 공지사항 ID는 td:nth-child(1)에 있음
		rawID := strings.TrimSpace(s.Find("td:nth-child(1)").Text())
		id := fmt.Sprintf("%s-%s", prefix, rawID)

		// 제목과 URL은 td:nth-child(2) 내부의 a 태그에 있음
		titleCell := s.Find("td:nth-child(2)")
		title := strings.TrimSpace(titleCell.Find(".b-title").Text())
		detailURL, exists := titleCell.Find("a").Attr("href")
		if !exists {
			return
		}

		// URL이 상대 경로인 경우 절대 경로로 변환
		fullURL := detailURL
		if !strings.HasPrefix(detailURL, "http") {
			if strings.HasPrefix(detailURL, "?") {
				// 쿼리스트링만 있는 경우 기존 URL에 추가
				fullURL = baseURL + detailURL
			} else if strings.HasPrefix(detailURL, "/") {
				// 절대 경로인 경우
				fullURL = fmt.Sprintf("https://wwwce.hongik.ac.kr%s", detailURL)
			} else {
				// 상대 경로인 경우
				fullURL = fmt.Sprintf("https://wwwce.hongik.ac.kr/%s", detailURL)
			}
		}

		content, err := c.fetchNoticeContent(fullURL)
		if err != nil {
			log.Printf("상세 내용 크롤링 실패: %v", err)
			content = "" // 실패 시 빈 문자열로 설정
		}

		notice := entity.SchoolNotification{
			ID:      id,
			Title:   title,
			Url:     fullURL,
			Date:    parseDate(dateStr),
			Content: content,
		}

		log.Printf("발견된 공지사항: %+v", notice)
		notices = append(notices, notice)
	})

	log.Printf("학과 공지사항 총 %d개 발견 (URL: %s)", len(notices), url)
	return notices, nil
}

// 공지사항 상세 내용 크롤링
func (c *crawler) fetchNoticeContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("HTTP 요청 실패: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("잘못된 상태 코드: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("페이지 파싱 실패: %v", err)
	}

	// 상세 내용 찾기
	var content string

	// 학교/학과 공지사항 형식 (.fr-view)
	contentBox := doc.Find(".fr-view")
	if contentBox.Length() > 0 {
		var texts []string
		contentBox.Find("p").Each(func(i int, s *goquery.Selection) {
			// 이미지 태그는 건너뛰기
			if s.Find("img").Length() > 0 {
				return
			}

			// 텍스트 추출 및 전처리
			text := strings.TrimSpace(s.Text())
			if text != "" {
				texts = append(texts, text)
			}
		})
		content = strings.Join(texts, "\n\n")
	}

	// 다른 가능한 클래스들도 시도
	if content == "" {
		contentBox = doc.Find(".b-content-box")
		if contentBox.Length() > 0 {
			content = contentBox.Text()
		}
	}

	if content == "" {
		contentBox = doc.Find(".board-view-content")
		if contentBox.Length() > 0 {
			content = contentBox.Text()
		}
	}

	// 상세 내용을 찾지 못한 경우에만 디버깅 로그 출력
	if content == "" {
		log.Printf("상세 내용을 찾을 수 없음 - URL: %s", url)
		// HTML 구조 출력하여 디버깅
		html, _ := doc.Html()
		log.Printf("페이지 HTML 구조: %s", html)
		return "", fmt.Errorf("상세 내용을 찾을 수 없음")
	}

	// 내용 전처리
	content = strings.TrimSpace(content)
	content = strings.ReplaceAll(content, "\n\n\n", "\n\n")
	content = strings.ReplaceAll(content, "\t", "")

	return content, nil
}
