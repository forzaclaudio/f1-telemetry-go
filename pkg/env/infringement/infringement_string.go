// Code generated by "stringer -type=Infringement -linecomment"; DO NOT EDIT.

package infringement

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BlockingBySlowDriving-0]
	_ = x[BlockingByWrongWayDriving-1]
	_ = x[ReversingOffTheStartLine-2]
	_ = x[BigCollision-3]
	_ = x[SmallCollision-4]
	_ = x[CollisionFailedToHandBackPositionSingle-5]
	_ = x[CollisionFailedToHandBackPositionMultiple-6]
	_ = x[CornerCuttingGainedTime-7]
	_ = x[CornerCuttingOvertakeSingle-8]
	_ = x[CornerCuttingOvertakeMultiple-9]
	_ = x[CrossedPitExitLane-10]
	_ = x[IgnoringBlueFlags-11]
	_ = x[IgnoringYellowFlags-12]
	_ = x[IgnoringDriveThrough-13]
	_ = x[TooManyDriveThroughs-14]
	_ = x[DriveThroughReminderServeWithinNLaps-15]
	_ = x[DriveThroughReminderServeThisLap-16]
	_ = x[PitLaneSpeeding-17]
	_ = x[ParkedForTooLong-18]
	_ = x[IgnoringTyreRegulations-19]
	_ = x[TooManyPenalties-20]
	_ = x[MultipleWarnings-21]
	_ = x[ApproachingDisqualification-22]
	_ = x[TyreRegulationsSelectSingle-23]
	_ = x[TyreRegulationsSelectMultiple-24]
	_ = x[LapInvalidatedCornerCutting-25]
	_ = x[LapInvalidatedRunningWide-26]
	_ = x[CornerCuttingRanWideGainedTimeMinor-27]
	_ = x[CornerCuttingRanWideGainedTimeSignificant-28]
	_ = x[CornerCuttingRanWideGainedTimeExtreme-29]
	_ = x[LapInvalidatedWallRiding-30]
	_ = x[LapInvalidatedFlashbackUsed-31]
	_ = x[LapInvalidatedResetToTrack-32]
	_ = x[BlockingThePitLane-33]
	_ = x[JumpStart-34]
	_ = x[SafetyCarToCarCollision-35]
	_ = x[SafetyCarIllegalOvertake-36]
	_ = x[SafetyCarExceedingAllowedPace-37]
	_ = x[VirtualSafetyCarExceedingAllowedPace-38]
	_ = x[FormationLapBelowAllowedSpeed-39]
	_ = x[FormationLapParking-40]
	_ = x[RetiredMechanicalFailure-41]
	_ = x[RetiredTerminallyDamaged-42]
	_ = x[SafetyCarFallingTooFarBack-43]
	_ = x[BlackFlagTimer-44]
	_ = x[UnServedStopGoPenalty-45]
	_ = x[UnServedDriveThroughPenalty-46]
	_ = x[EngineComponentChange-47]
	_ = x[GearboxChange-48]
	_ = x[ParcFermeChange-49]
	_ = x[LeagueGridPenalty-50]
	_ = x[RetryPenalty-51]
	_ = x[IllegalTimeGain-52]
	_ = x[MandatoryPitStop-53]
	_ = x[AttributeAssigned-54]
}

const _Infringement_name = "Blocking by slow drivingBlocking by wrong way drivingReversing off the start lineBig CollisionSmall CollisionCollision failed to hand back position singleCollision failed to hand back position multipleCorner cutting gained timeCorner cutting overtake singleCorner cutting overtake multipleCrossed pit exit laneIgnoring blue flagsIgnoring yellow flagsIgnoring drive throughToo many drive throughsDrive through reminder serve within n lapsDrive through reminder serve this lapPit lane speedingParked for too longIgnoring tyre regulationsToo many penaltiesMultiple warningsApproaching disqualificationTyre regulations select singleTyre regulations select multipleLap invalidated corner cuttingLap invalidated running wideCorner cutting ran wide gained time minorCorner cutting ran wide gained time significantCorner cutting ran wide gained time extremeLap invalidated wall ridingLap invalidated flashback usedLap invalidated reset to trackBlocking the pitlaneJump startSafety car to car collisionSafety car illegal overtakeSafety car exceeding allowed paceVirtual safety car exceeding allowed paceFormation lap below allowed speedFormation lap parkingRetired mechanical failureRetired terminally damagedSafety car falling too far backBlack flag timerUnserved stop go penaltyUnserved drive through penaltyEngine component changeGearbox changeParc Fermé changeLeague grid penaltyRetry penaltyIllegal time gainMandatory pitstopAttribute assigned"

var _Infringement_index = [...]uint16{0, 24, 53, 81, 94, 109, 154, 201, 227, 257, 289, 310, 329, 350, 372, 395, 437, 474, 491, 510, 535, 553, 570, 598, 628, 660, 690, 718, 759, 806, 849, 876, 906, 936, 956, 966, 993, 1020, 1053, 1094, 1127, 1148, 1174, 1200, 1231, 1247, 1271, 1301, 1324, 1338, 1356, 1375, 1388, 1405, 1422, 1440}

func (i Infringement) String() string {
	if i >= Infringement(len(_Infringement_index)-1) {
		return "Infringement(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Infringement_name[_Infringement_index[i]:_Infringement_index[i+1]]
}