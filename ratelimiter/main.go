// =============================================================================
// ЗАДАЧА: Rate Limiter — Fixed Window 🚦
// =============================================================================
//
// Реализуй потокобезопасный rate limiter по алгоритму Fixed Window.
//
// Алгоритм:
//   - Время делится на фиксированные окна одинаковой длины
//   - В каждом окне ведётся счётчик запросов клиента
//   - Запрос разрешён если счётчик < лимита, иначе — отклонён
//   - Когда окно истекает — счётчик сбрасывается в 0
//
//   Пример (limit=3, window=1s):
//
//   t=0.0s  счётчик=0 → ✅ разрешён  (счётчик=1)
//   t=0.3s  счётчик=1 → ✅ разрешён  (счётчик=2)
//   t=0.7s  счётчик=2 → ✅ разрешён  (счётчик=3)
//   t=0.9s  счётчик=3 → ❌ отклонён
//   t=1.0s  новое окно → счётчик=0
//   t=1.1s  счётчик=0 → ✅ разрешён  (счётчик=1)
//
// Запусти тесты:
//   go test -race -v
//
// =============================================================================

package main

import (
	"sync"
	"time"
)

// FixedWindow — счётчик запросов для одного клиента
type FixedWindow struct {
	mu          sync.Mutex
	limit       int
	window      time.Duration
	windowStart time.Time
	count       int
}

// NewFixedWindow создаёт окно с заданным лимитом и длиной окна
func NewFixedWindow(limit int, window time.Duration) *FixedWindow {
	return &FixedWindow{
		limit:       limit,
		window:      window,
		windowStart: time.Now(),
	}
}

// Allow возвращает true если запрос разрешён и увеличивает счётчик.
// Если окно истекло — сбрасывает счётчик и начинает новое.
func (w *FixedWindow) Allow() bool {
	w.mu.Lock()
	defer w.mu.Unlock()
	if time.Since(w.windowStart) >= w.window {
		w.count = 0
		w.windowStart = time.Now()
	}

	if w.limit > w.count {
		w.count += 1
		return true
	}
	return false
}

// RateLimiter — менеджер лимитов для множества клиентов.
// Каждый клиент имеет своё независимое окно.
type RateLimiter struct {
	fixedWindows sync.Map
	limit        int
	window       time.Duration
}

// NewRateLimiter создаёт RateLimiter где все клиенты имеют одинаковые limit и window
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		fixedWindows: sync.Map{},
		limit:        limit,
		window:       window,
	}
}

// Allow проверяет разрешён ли запрос для clientID.
// Если клиент новый — создаёт для него FixedWindow.
func (rl *RateLimiter) Allow(clientID string) bool {
	fw, ok := rl.fixedWindows.Load(clientID)
	if !ok {
		fw = rl.createFixedWindow()
		rl.fixedWindows.Store(clientID, fw)
	}

	return fw.(*FixedWindow).Allow()
}

func (rl *RateLimiter) createFixedWindow() *FixedWindow {
	return NewFixedWindow(rl.limit, rl.window)
}
