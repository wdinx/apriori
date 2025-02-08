package service

import (
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type CronjobServiceImpl struct {
	AprioriService AprioriService
	Scheduler      gocron.Scheduler
}

func NewCronJobServiceImpl(aprioriService AprioriService, scheduler gocron.Scheduler) CronjobService {
	return &CronjobServiceImpl{
		AprioriService: aprioriService,
		Scheduler:      scheduler,
	}
}

// Inisiasi Cronjob
func (c *CronjobServiceImpl) InitCronJob() {
	job, err := c.Scheduler.NewJob(
		gocron.DurationJob(1*time.Hour),
		gocron.NewTask(c.UpdateRecommendationItem),
	)
	if err != nil {
		log.Println("Error while creating new job", err.Error())
	}
	log.Println("Cronjob Init, Job ID:", job.ID())
}

// Melakukan update rekomendasi item setiap jam 00.00
func (c *CronjobServiceImpl) UpdateRecommendationItem() {
	log.Println("Cronjob Update Recommendation Item Running")
	// Process Apriori
	midnight := time.Now().Truncate(24 * time.Hour)
	oneAM := midnight.Add(1 * time.Hour)

	if time.Now().Before(oneAM) && time.Now().After(midnight) {
		err := c.AprioriService.CreateRecommendationItem()

		if err != nil {
			log.Println("Error while processing apriori", err.Error())
		}
	}
}
