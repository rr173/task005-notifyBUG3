package notify

import (
	"testing"
	"time"
)

func TestCreateScheduleAtPointerIndependence(t *testing.T) {
	s := New()
	now := time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC)
	future := now.Add(24 * time.Hour)

	input := CreateInput{
		ID: "P1", Recipient: "u", Content: "c",
		ScheduleAt: &future,
	}
	s.Create(input, now)

	// 修改原始 input 的 ScheduleAt 指向的值
	*input.ScheduleAt = time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC)

	// store 内部不应被影响
	got, _ := s.Get("P1")
	if !got.ScheduleAt.Equal(now.Add(24 * time.Hour)) {
		t.Errorf("Create stores ScheduleAt pointer from input without copying: got %v, want %v",
			got.ScheduleAt, now.Add(24*time.Hour))
	}
}
