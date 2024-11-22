package cron

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type CronJob struct {
	cron *cron.Cron
}

func NewCronJob() *CronJob {
	// CRON_TZ=Asia/Seoul 설정으로 한국 시간 기준으로 동작하도록 설정
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		log.Printf("Failed to load location: %v", err)
		return nil
	}

	c := cron.New(cron.WithLocation(loc))
	return &CronJob{
		cron: c,
	}
}

func (c *CronJob) Start() {
	// 테스트용: 2초마다 실행
	c.cron.AddFunc("@every 2s", func() {
		log.Println("Cron job executed at:", time.Now().Format("2006-01-02 15:04:05"))
		// 여기에 실행하고자 하는 작업을 추가하세요
	})

	// 실제 배포용: 매일 오전 11시에 실행
	// c.cron.AddFunc("0 11 * * *", func() {
	//     log.Println("Daily cron job executed at:", time.Now().Format("2006-01-02 15:04:05"))
	//     // 여기에 실행하고자 하는 작업을 추가하세요
	// })

	c.cron.Start()
	log.Println("Cron job started")
}

func (c *CronJob) Stop() {
	if c.cron != nil {
		c.cron.Stop()
		log.Println("Cron job stopped")
	}
}
