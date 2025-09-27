package sequence_test

import (
	"testing"

	"github.com/Hyp9r/sequncer/domain/sequence"
)

func TestNewSequence(t *testing.T) {
	steps := []sequence.Step{
		{Subject: "s1", Content: "c1"},
	}
	seq := sequence.NewSequence("Test Seq", true, false, steps)

	if seq.Name != "Test Seq" {
		t.Errorf("expected Name 'Test Seq', got '%s'", seq.Name)
	}
	if !seq.OpenTrackingEnabled {
		t.Errorf("expected OpenTrackingEnabled true")
	}
	if seq.ClickTrackingEnabled {
		t.Errorf("expected ClickTrackingEnabled false")
	}
	if len(seq.Steps) != 1 {
		t.Errorf("expected 1 step, got %d", len(seq.Steps))
	}
}

func TestUpdateStep(t *testing.T) {
	seq := sequence.NewSequence("Test", true, true, []sequence.Step{
		{Subject: "old", Content: "old content"},
	})
	stepID := seq.Steps[0].ID

	// update both subject and content
	subject := "new subject"
	content := "new content"
	if err := seq.UpdateStep(stepID, &subject, &content); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	step := seq.Steps[0]
	if step.Subject != "new subject" || step.Content != "new content" {
		t.Fatal("step not updated correctly")
	}

	// update only subject
	subject2 := "subject only"
	if err := seq.UpdateStep(stepID, &subject2, nil); err != nil {
		t.Fatal(err)
	}
	step = seq.Steps[0]
	if step.Subject != "subject only" || step.Content != "new content" {
		t.Fatal("step partial update failed")
	}

	// update only content
	content2 := "content only"
	if err := seq.UpdateStep(stepID, nil, &content2); err != nil {
		t.Fatal(err)
	}
	step = seq.Steps[0]
	if step.Subject != "subject only" || step.Content != "content only" {
		t.Fatal("step partial update failed")
	}

	// update non-existent step
	subject3 := "x"
	content3 := "y"
	err := seq.UpdateStep("non-existent-id", &subject3, &content3)
	if err == nil {
		t.Fatal("expected error for non-existent step")
	}
}

func TestDeleteStep(t *testing.T) {
	seq := sequence.NewSequence("Test", true, true, []sequence.Step{
		{Subject: "s", Content: "c"},
	})
	stepID := seq.Steps[0].ID

	// delete existing
	if err := seq.DeleteStep(stepID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(seq.Steps) != 0 {
		t.Fatal("step not deleted")
	}

	// delete non-existent
	err := seq.DeleteStep("non-existent-id")
	if err == nil {
		t.Fatal("expected error for non-existent step")
	}
}

func TestUpdateTracking(t *testing.T) {
	seq := sequence.NewSequence("Test", false, false, nil)

	// update both
	open := true
	click := true
	seq.UpdateTracking(&open, &click)
	if !seq.OpenTrackingEnabled || !seq.ClickTrackingEnabled {
		t.Fatal("tracking not updated correctly")
	}

	// update only open to true
	open2 := true
	click2 := false
	seq.UpdateTracking(&open2, &click2)
	if !seq.OpenTrackingEnabled || seq.ClickTrackingEnabled {
		t.Fatal("partial tracking update failed")
	}

	// update only click to true
	open3 := false
	click3 := true
	seq.UpdateTracking(&open3, &click3)
	if seq.OpenTrackingEnabled || !seq.ClickTrackingEnabled {
		t.Fatal("partial tracking update failed")
	}
}
