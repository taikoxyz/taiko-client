package handler

// handleNewBlockProposedEvent handles the new block proposed event.
// func (h *BlockProposedEventHandler) handleNewBlockProposedEvent(ctx context.Context, e *bindings.TaikoL1ClientBlockProposed) error {
// 	// Check whether the block has been verified.
// 	isVerified, err := h.isBlockVerified(e.BlockId)
// 	if err != nil {
// 		return fmt.Errorf("failed to check if the current L2 block is verified: %w", err)
// 	}
// 	if isVerified {
// 		log.Info("📋 Block has been verified", "blockID", e.BlockId)
// 		return nil
// 	}

// 	// Check whether the block's proof is still needed.
// 	proofStatus, err := rpc.GetBlockProofStatus(
// 		ctx,
// 		h.rpc,
// 		e.BlockId,
// 		h.proverAddress,
// 	)
// 	if err != nil {
// 		return fmt.Errorf("failed to check whether the L2 block needs a new proof: %w", err)
// 	}

// 	if proofStatus.IsSubmitted {
// 		// If there is already a proof submitted and there is no need to contest
// 		// it, we skip proving this block here.
// 		if !proofStatus.Invalid {
// 			log.Info(
// 				"A valid proof has been submitted, skip proving",
// 				"blockID", e.BlockId,
// 				"parent", proofStatus.ParentHeader.Hash(),
// 			)
// 			return nil
// 		}

// 		// If there is an invalid proof, but current prover is not in contest mode, we skip proving this block.
// 		if !p.cfg.ContesterMode {
// 			log.Info(
// 				"An invalid proof has been submitted, but current prover is not in contest mode, skip proving",
// 				"blockID", e.BlockId,
// 				"parent", proofStatus.ParentHeader.Hash(),
// 			)
// 			return nil
// 		}

// 		// The proof submitted to protocol is invalid.
// 		return p.handleInvalidProof(
// 			ctx,
// 			e.BlockId,
// 			new(big.Int).SetUint64(e.Raw.BlockNumber),
// 			proofStatus.ParentHeader.Hash(),
// 			proofStatus.CurrentTransitionState.Contester,
// 			&e.Meta,
// 			proofStatus.CurrentTransitionState.Tier,
// 		)
// 	}

// 	provingWindow, err := p.getProvingWindow(e)
// 	if err != nil {
// 		return fmt.Errorf("failed to get proving window: %w", err)
// 	}

// 	var (
// 		now                    = uint64(time.Now().Unix())
// 		provingWindowExpiresAt = e.Meta.Timestamp + uint64(provingWindow.Seconds())
// 		provingWindowExpired   = now > provingWindowExpiresAt
// 		timeToExpire           = time.Duration(provingWindowExpiresAt-now) * time.Second
// 	)
// 	if provingWindowExpired {
// 		// If the proving window is expired, we need to check if the current prover is the assigned prover
// 		// at first, if yes, we should skip proving this block, if no, then we check if the current prover
// 		// wants to prove unassigned blocks.
// 		log.Info(
// 			"Proposed block's proving window has expired",
// 			"blockID", e.BlockId,
// 			"prover", e.AssignedProver,
// 			"now", now,
// 			"expiresAt", provingWindowExpiresAt,
// 			"minTier", e.Meta.MinTier,
// 		)
// 		if e.AssignedProver == p.proverAddress {
// 			log.Warn(
// 				"Assigned prover is the current prover, but the proving window has expired, skip proving",
// 				"blockID", e.BlockId,
// 				"prover", e.AssignedProver,
// 				"expiresAt", provingWindowExpiresAt,
// 			)
// 			return nil
// 		}
// 		if !p.cfg.ProveUnassignedBlocks {
// 			log.Info(
// 				"Skip proving expired blocks",
// 				"blockID", e.BlockId,
// 				"prover", e.AssignedProver,
// 				"expiresAt", provingWindowExpiresAt,
// 			)
// 			return nil
// 		}
// 	} else {
// 		// If the proving window is not expired, we need to check if the current prover is the assigned prover,
// 		// if no and the current prover wants to prove unassigned blocks, then we should wait for its expiration.
// 		if e.AssignedProver != p.proverAddress {
// 			log.Info(
// 				"Proposed block is not provable",
// 				"blockID", e.BlockId,
// 				"prover", e.AssignedProver,
// 				"expiresAt", provingWindowExpiresAt,
// 				"timeToExpire", timeToExpire,
// 			)

// 			if p.cfg.ProveUnassignedBlocks {
// 				log.Info(
// 					"Add proposed block to wait for proof window expiration",
// 					"blockID", e.BlockId,
// 				)
// 				time.AfterFunc(
// 					// Add another 60 seconds, to ensure one more L1 block will be mined before the proof submission
// 					timeToExpire+60*time.Second,
// 					func() { p.proofWindowExpiredCh <- e },
// 				)
// 			}

// 			return nil
// 		}
// 	}

// 	tier := e.Meta.MinTier
// 	if p.IsGuardianProver() {
// 		tier = encoding.TierGuardianID
// 	}

// 	log.Info(
// 		"Proposed block is provable",
// 		"blockID", e.BlockId,
// 		"prover", e.AssignedProver,
// 		"expiresAt", provingWindowExpiresAt,
// 		"minTier", e.Meta.MinTier,
// 		"currentTier", tier,
// 	)

// 	metrics.ProverProofsAssigned.Inc(1)

// 	if proofSubmitter := p.selectSubmitter(tier); proofSubmitter != nil {
// 		return proofSubmitter.RequestProof(ctx, e)
// 	}

// 	return nil
// }
