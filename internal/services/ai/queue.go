package ai

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/Dokito555/mizuki/internal/models"
	"github.com/Dokito555/mizuki/internal/repositories"
	"github.com/sirupsen/logrus"
)

type EnrichmentJob struct {
	FlowID   uint
	UploadID uint
}

type AIEnrichmentQueue struct {
	jobs    chan EnrichmentJob
	sem     chan struct{}
	workers int
	flowRepo   repositories.FlowRepository
	flowAIRepo repositories.FlowAIRepository
	uploadRepo repositories.UploadRepository
	providers   []Provider
	log         *logrus.Logger
	stopOnce    sync.Once
	stopCh      chan struct{}
}

func NewAIEnrichmentQueue(
	concurrency int,
	flowRepo repositories.FlowRepository,
	flowAIRepo repositories.FlowAIRepository,
	uploadRepo repositories.UploadRepository,
	providers []Provider,
	log *logrus.Logger,
) *AIEnrichmentQueue {
	if concurrency < 1 {
		concurrency = 3
	}
	q := &AIEnrichmentQueue{
		jobs:        make(chan EnrichmentJob, 100),
		sem:         make(chan struct{}, concurrency),
		workers:     concurrency,
		flowRepo:    flowRepo,
		flowAIRepo: flowAIRepo,
		uploadRepo:  uploadRepo,
		providers:   providers,
		log:         log,
		stopCh:      make(chan struct{}),
	}
	q.start()
	return q
}

func (q *AIEnrichmentQueue) start() {
	for i := 0; i < q.workers; i++ {
		go q.worker()
	}
}

func (q *AIEnrichmentQueue) worker() {
	for {
		select {
		case <-q.stopCh:
			return
		case job := <-q.jobs:
			q.sem <- struct{}{}
			q.processJob(job)
			<-q.sem
		}
	}
}

func (q *AIEnrichmentQueue) Enqueue(job EnrichmentJob) {
	select {
	case q.jobs <- job:
	default:
		q.log.WithField("job", job).Warn("AI enrichment queue full, dropping job")
	}
}

func (q *AIEnrichmentQueue) Stop() {
	q.stopOnce.Do(func() {
		close(q.stopCh)
	})
}

func (q *AIEnrichmentQueue) processJob(job EnrichmentJob) {
	ctx := context.Background()
	log := q.log.WithFields(logrus.Fields{
		"flow_id":   job.FlowID,
		"upload_id": job.UploadID,
	})

	flow, err := q.flowRepo.FindByID(ctx, job.FlowID)
	if err != nil {
		log.Errorf("find flow: %v", err)
		return
	}

	var provider Provider
	if len(q.providers) > 0 {
		provider = q.providers[0]
	} else {
		provider = NewFallbackProvider()
	}

	related, _, _ := q.flowRepo.FindAll(ctx, models.FlowFilter{
		UploadID: job.UploadID,
		SrcIP:    flow.SrcIP,
	})

	result, err := provider.Analyze(ctx, flow, related)
	if err != nil {
		log.Errorf("AI analysis failed: %v; using fallback", err)
		result, _ = NewFallbackProvider().Analyze(ctx, flow, related)
	}

	resultJSON, _ := json.Marshal(result)

	if err := q.flowAIRepo.Save(ctx, job.FlowID, resultJSON); err != nil {
		log.Errorf("save AI analysis: %v", err)
	}
	if err := q.flowRepo.MarkAIAnalyzed(ctx, job.FlowID, provider.Name()); err != nil {
		log.Errorf("mark AI analyzed: %v", err)
	}
}
