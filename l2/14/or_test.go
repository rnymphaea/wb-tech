package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	sig := func(after time.Duration) <-chan any {
		c := make(chan any)
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	t.Run("single channel", func(t *testing.T) {
		start := time.Now()
		<-or(sig(100 * time.Millisecond))
		elapsed := time.Since(start)
		if elapsed < 100*time.Millisecond || elapsed > 150*time.Millisecond {
			t.Errorf("expected ~100ms, got %v", elapsed)
		}
	})

	t.Run("multiple channels", func(t *testing.T) {
		start := time.Now()
		<-or(
			sig(2*time.Hour),
			sig(5*time.Minute),
			sig(100*time.Millisecond),
			sig(1*time.Hour),
			sig(1*time.Minute),
		)
		elapsed := time.Since(start)
		if elapsed < 100*time.Millisecond || elapsed > 150*time.Millisecond {
			t.Errorf("expected ~100ms, got %v", elapsed)
		}
	})

	t.Run("no channels", func(t *testing.T) {
		c := or()
		if c != nil {
			t.Error("should return nil for no channels")
		}
	})

	t.Run("immediate close", func(t *testing.T) {
		c := make(chan any)
		close(c)
		start := time.Now()
		<-or(c)
		elapsed := time.Since(start)
		if elapsed > 10*time.Millisecond {
			t.Error("should return immediately")
		}
	})

	t.Run("mixed closed and open", func(t *testing.T) {
		c1 := make(chan any)
		c2 := make(chan any)
		close(c1)

		start := time.Now()
		<-or(c1, c2)
		elapsed := time.Since(start)
		if elapsed > 10*time.Millisecond {
			t.Error("should return immediately")
		}
	})

	t.Run("all closed together", func(t *testing.T) {
		c1 := make(chan any)
		c2 := make(chan any)
		c3 := make(chan any)

		go func() {
			time.Sleep(100 * time.Millisecond)
			close(c1)
			close(c2)
			close(c3)
		}()

		start := time.Now()
		<-or(c1, c2, c3)
		elapsed := time.Since(start)
		if elapsed < 100*time.Millisecond || elapsed > 150*time.Millisecond {
			t.Errorf("expected ~100ms, got %v", elapsed)
		}
	})

	t.Run("very short timeout", func(t *testing.T) {
		start := time.Now()
		<-or(sig(1 * time.Microsecond))
		elapsed := time.Since(start)
		if elapsed > 10*time.Millisecond {
			t.Error("should be very fast")
		}
	})

	t.Run("mixed timeouts", func(t *testing.T) {
		start := time.Now()
		<-or(
			sig(200*time.Millisecond),
			sig(50*time.Millisecond),
			sig(150*time.Millisecond),
		)
		elapsed := time.Since(start)
		if elapsed < 50*time.Millisecond || elapsed > 80*time.Millisecond {
			t.Errorf("expected ~50ms, got %v", elapsed)
		}
	})

	t.Run("with nil channel", func(t *testing.T) {
		start := time.Now()
		<-or(
			nil,
			sig(100*time.Millisecond),
		)
		elapsed := time.Since(start)
		if elapsed < 100*time.Millisecond || elapsed > 150*time.Millisecond {
			t.Errorf("expected ~100ms, got %v", elapsed)
		}
	})

	t.Run("many channels", func(t *testing.T) {
		var chs []<-chan any
		for i := 0; i < 100; i++ {
			chs = append(chs, sig(time.Duration(i+1)*100*time.Millisecond))
		}

		start := time.Now()
		<-or(chs...)
		elapsed := time.Since(start)
		if elapsed < 100*time.Millisecond || elapsed > 199*time.Millisecond {
			t.Errorf("expected ~100ms, got %v", elapsed)
		}
	})
}
