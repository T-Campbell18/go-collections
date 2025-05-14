package gocollections

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack[int]()

	// Test IsEmpty on new stack
	if !s.IsEmpty() {
		t.Error("expected stack to be empty")
	}

	// Test Size on new stack
	if s.Size() != 0 {
		t.Errorf("expected size 0, got %d", s.Size())
	}

	// Test Pop on empty stack
	if _, err := s.Pop(); err == nil {
		t.Error("expected error when popping from empty stack")
	}

	// Test Peek on empty stack
	if _, err := s.Peek(); err == nil {
		t.Error("expected error when peeking from empty stack")
	}

	// Test Push
	s.Push(10)
	s.Push(20)
	s.Push(30)

	if s.IsEmpty() {
		t.Error("expected stack to not be empty after pushes")
	}

	if s.Size() != 3 {
		t.Errorf("expected size 3, got %d", s.Size())
	}

	// Test Peek
	if val, err := s.Peek(); err != nil || val != 30 {
		t.Errorf("expected peek to return 30, got %v, err: %v", val, err)
	}

	// Test Pop
	if val, err := s.Pop(); err != nil || val != 30 {
		t.Errorf("expected pop to return 30, got %v, err: %v", val, err)
	}
	if val, err := s.Pop(); err != nil || val != 20 {
		t.Errorf("expected pop to return 20, got %v, err: %v", val, err)
	}
	if val, err := s.Pop(); err != nil || val != 10 {
		t.Errorf("expected pop to return 10, got %v, err: %v", val, err)
	}

	// Stack should be empty again
	if !s.IsEmpty() {
		t.Error("expected stack to be empty after popping all elements")
	}

	// Test Clear
	s.Push(1)
	s.Push(2)
	s.Clear()
	if !s.IsEmpty() || s.Size() != 0 {
		t.Error("expected stack to be empty after clear")
	}
}
