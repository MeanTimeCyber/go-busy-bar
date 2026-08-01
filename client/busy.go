package client

import (
	"fmt"
	"os"

	"github.com/markkurossi/tabulate"
)

type BusySnapshot struct {
	// Snapshot contains the current BUSY timer snapshot and BUSY bar settings.
	Snapshot BusySnapshotPayload `json:"snapshot"`
	// SnapshotTimestampMS is the Unix timestamp in milliseconds when the snapshot was captured.
	SnapshotTimestampMS int `json:"snapshot_timestamp_ms"`
}

type BusySnapshotPayload struct {
	// Type is the snapshot type (for example NOT_STARTED, SIMPLE, INTERVAL, or INFINITE).
	Type string `json:"type"`
	// CardID is the active card UUID for running timer modes.
	CardID string `json:"card_id,omitempty"`
	// TimeLeftMS is remaining time in milliseconds for SIMPLE snapshots.
	TimeLeftMS int `json:"time_left_ms,omitempty"`
	// CurrentInterval is the current interval index for INTERVAL snapshots.
	CurrentInterval int `json:"current_interval,omitempty"`
	// CurrentIntervalTimeTotalMS is total duration of the current interval in milliseconds.
	CurrentIntervalTimeTotalMS int `json:"current_interval_time_total_ms,omitempty"`
	// CurrentIntervalTimeLeftMS is remaining duration of the current interval in milliseconds.
	CurrentIntervalTimeLeftMS int `json:"current_interval_time_left_ms,omitempty"`
	// IsPaused indicates whether the timer is currently paused.
	IsPaused bool `json:"is_paused,omitempty"`
	// IntervalSettings contains interval timer configuration for INTERVAL snapshots.
	IntervalSettings *BusyTimerIntervalSettings `json:"interval_settings,omitempty"`
	// BusyBarSettings controls BUSY bar rendering and automation behavior.
	BusyBarSettings BusyBarSettings `json:"busy_bar_settings"`
}

type BusyProfile struct {
	// SortOrder is the profile ordering index.
	SortOrder int `json:"sort_order"`
	// Title is the human-readable profile title.
	Title string `json:"title"`
	// ID is the profile UUID.
	ID string `json:"id"`
	// TimerSettings holds timer configuration for the selected timer mode.
	TimerSettings any `json:"timer_settings"`
	// BusyBarSettings contains BUSY bar display and integration settings.
	BusyBarSettings BusyBarSettings `json:"busy_bar_settings"`
	// ProfileTimestamp is the Unix timestamp in milliseconds of the profile snapshot.
	ProfileTimestamp int `json:"profile_timestamp_ms"`
}

type BusyBarSettings struct {
	// Theme is the BUSY bar theme identifier.
	Theme string `json:"theme"`
	// ShowWorkPhaseOnly controls whether only work phase is shown.
	ShowWorkPhaseOnly bool `json:"show_work_phase_only"`
	// TriggerSmartHome controls whether smart home automation is triggered.
	TriggerSmartHome bool `json:"trigger_smart_home"`
}

type BusyTimerIntervalSettings struct {
	// Type is the timer settings type.
	Type string `json:"type"`
	// IntervalWorkMS is work interval duration in milliseconds.
	IntervalWorkMS int `json:"interval_work_ms"`
	// IntervalRestMS is rest interval duration in milliseconds.
	IntervalRestMS int `json:"interval_rest_ms"`
	// IntervalWorkCyclesCount is the number of work cycles.
	IntervalWorkCyclesCount int `json:"interval_work_cycles_count"`
	// IsAutostartEnabled controls whether intervals auto-start.
	IsAutostartEnabled bool `json:"is_autostart_enabled"`
}

func (b *BusySnapshot) PrettyPrint() {
	fmt.Printf("\nBusy Snapshot\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Type")
	row.Column(b.Snapshot.Type)

	row = tab.Row()
	row.Column("Paused")
	row.Column(fmt.Sprintf("%v", b.Snapshot.IsPaused))

	row = tab.Row()
	row.Column("Timestamp (ms)")
	row.Column(fmt.Sprintf("%d", b.SnapshotTimestampMS))

	tab.Print(os.Stdout)
	fmt.Println()
}

func (b *BusyProfile) PrettyPrint() {
	fmt.Printf("\nBusy Profile\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Title")
	row.Column(b.Title)

	row = tab.Row()
	row.Column("ID")
	row.Column(b.ID)

	row = tab.Row()
	row.Column("Sort Order")
	row.Column(fmt.Sprintf("%d", b.SortOrder))

	tab.Print(os.Stdout)
	fmt.Println()
}
