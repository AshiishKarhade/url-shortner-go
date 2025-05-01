// Package snowflake provides utilities and components for the application.
// This file implements the Snowflake algorithm for generating unique IDs
// in a distributed system. The algorithm ensures IDs are time-ordered
// and unique across multiple machines.
package snowflake

import (
	"fmt"
	"sync"
	"time"
)

const (
	// epoch is the custom starting point in milliseconds (e.g., January 1, 2021).
	epoch = int64(1609459200000)

	// machineIdBits defines the number of bits allocated for the machine ID.
	machineIdBits = uint(10)

	// sequenceBits defines the number of bits allocated for the sequence number.
	sequenceBits = uint(12)

	// machineIdShift is the bit shift for the machine ID in the final ID.
	machineIdShift = sequenceBits

	// timestampShift is the bit shift for the timestamp in the final ID.
	timestampShift = sequenceBits + machineIdBits

	// maxMachineID is the maximum value for the machine ID.
	maxMachineID = -1 ^ (-1 << machineIdBits)

	// maxSequence is the maximum value for the sequence number.
	maxSequence = -1 ^ (-1 << sequenceBits)
)

// Snowflake is a struct that implements the Snowflake algorithm for generating unique IDs.
// It ensures thread-safe ID generation using a mutex.
type Snowflake struct {
	mu            sync.Mutex // Mutex to ensure thread-safe access.
	machineID     int64      // ID of the machine generating the IDs.
	sequence      int64      // Sequence number for IDs within the same millisecond.
	lastTimestamp int64      // Tracks the last timestamp used to ensure uniqueness.
}

func NewSnowflake(machineID int64) (*Snowflake, error) {
	if machineID < 0 || machineID > maxMachineID {
		return nil, fmt.Errorf("machine ID must be between 0 and %d", maxMachineID)
	}
	return &Snowflake{
		machineID: machineID,
	}, nil
}

func (s *Snowflake) NextID() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixNano() / 1e6 // Convert to milliseconds
	//sequence := int64(0)

	if now == s.lastTimestamp {
		// If the current timestamp is the same as the last, increment the sequence.
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			// If the sequence overflows, wait for the next millisecond.
			for now <= s.lastTimestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		// Reset sequence if we are in a new millisecond.
		s.sequence = 0
	}
	s.lastTimestamp = now
	// Generate the unique ID by combining the timestamp, machine ID, and sequence.
	id := ((now - epoch) << timestampShift) | (s.machineID << machineIdShift) | s.sequence
	return id
}
