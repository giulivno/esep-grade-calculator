package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeHighScores(t *testing.T) {
	// Weighted = 0.50*100 + 0.35*95 + 0.15*91 = 97.9 → "A"
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

// New tests to cover remaining branches 

func TestGetGradeC(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a", 70, Assignment)
	gc.AddGrade("e", 70, Exam)
	gc.AddGrade("s", 70, Essay)
	if got := gc.GetFinalGrade(); got != "C" {
		t.Fatalf("expected C, got %s", got)
	}
}

func TestGetGradeD(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a", 60, Assignment)
	gc.AddGrade("e", 60, Exam)
	gc.AddGrade("s", 60, Essay)
	if got := gc.GetFinalGrade(); got != "D" {
		t.Fatalf("expected D, got %s", got)
	}
}

func TestGetGradeF_LowScores(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a", 10, Assignment)
	gc.AddGrade("e", 10, Exam)
	gc.AddGrade("s", 10, Essay)
	if got := gc.GetFinalGrade(); got != "F" {
		t.Fatalf("expected F, got %s", got)
	}
}

func TestEmptyCategoriesHandled(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a1", 100, Assignment)
	if got := gc.GetFinalGrade(); got != "F" {
		t.Fatalf("expected F with empty exams/essays, got %s", got)
	}
}

func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Fatalf("Assignment.String mismatch")
	}
	if Exam.String() != "exam" {
		t.Fatalf("Exam.String mismatch")
	}
	if Essay.String() != "essay" {
		t.Fatalf("Essay.String mismatch")
	}
}
