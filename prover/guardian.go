package prover

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

// gurdianProverHeartbeatLoop keeps sending heartbeats to the guardian prover health check server
// on an interval.
func (p *Prover) gurdianProverHeartbeatLoop(ctx context.Context) {
	// Only guardian provers need to send heartbeat.
	if !p.IsGuardianProver() {
		return
	}

	ticker := time.NewTicker(heartbeatInterval)
	p.wg.Add(1)

	defer func() {
		ticker.Stop()
		p.wg.Done()
	}()

	for {
		select {
		case <-p.ctx.Done():
			return
		case <-ticker.C:
			latestL1Block, err := p.rpc.L1.BlockNumber(ctx)
			if err != nil {
				log.Error("Failed to get L1 head", err)
				continue
			}

			latestL2Block, err := p.rpc.L2.BlockNumber(ctx)
			if err != nil {
				log.Error("Failed to get L2 head", err)
				continue
			}

			if err := p.guardianProverHeartbeater.SendHeartbeat(
				ctx,
				latestL1Block,
				latestL2Block,
			); err != nil {
				log.Error("Failed to send guardian prover heartbeat", "error", err)
			}
		}
	}
}
