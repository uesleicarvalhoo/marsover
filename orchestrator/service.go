package orchestrator

import (
	"context"
	"fmt"
	"sync"

	"github.com/uesleicarvalhoo/marsrover/internal/logger"
	"github.com/uesleicarvalhoo/marsrover/plateau"
	"github.com/uesleicarvalhoo/marsrover/rover"
)

type MissionUseCase interface {
	Execute(ctx context.Context, params MissionParams) ([]rover.Params, error)
}

type service struct {
	concurrency int
}

func NewMissionService(concurrency int) MissionUseCase {
	if concurrency < 0 {
		concurrency = 1
	}

	return &service{
		concurrency: concurrency,
	}
}

func (s *service) Execute(ctx context.Context, params MissionParams) ([]rover.Params, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	plat, err := plateau.New(params.PlateauLimits)
	if err != nil {
		logger.WithContext(ctx).ErrorF("error while creating plateau", logger.Fields{
			"params": params,
			"error":  err,
		})
	}

	var (
		sem    = make(chan struct{}, s.concurrency)
		wg     = sync.WaitGroup{}
		errCh  = make(chan error, len(params.RoverInstructions))
		result = make([]rover.Params, len(params.RoverInstructions))
	)

	for i, instruction := range params.RoverInstructions {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			sem <- struct{}{}
			wg.Add(1)

			go func(idx int, instruction RoverInstructions) {
				defer wg.Done()

				r, err := rover.New(plat, instruction.Params)
				if err != nil {
					logger.WithContext(ctx).ErrorF("error while creating rover", logger.Fields{
						"params": instruction.Params,
						"error":  err,
					})
					errCh <- fmt.Errorf("error while creating rover: %w", err)
					return
				}
				cor, dir, err := r.ExecuteCommands(ctx, instruction.Commands)
				if err != nil {
					logger.WithContext(ctx).ErrorF("error while running rover commands", logger.Fields{
						"commands": instruction.Commands,
						"error":    err,
					})
					errCh <- fmt.Errorf("error while running rover commands: %w", err)
				}

				result[idx] = rover.Params{
					Name:        r.Name,
					Coordinates: cor,
					Direction:   dir,
				}
			}(i, instruction)
		}
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
