// =============================================================================
// ЗАДАЧА: Map-Reduce сортировка событий
// =============================================================================
//
// ЛЕГЕНДА:
// К тебе приходит большой неотсортированный поток событий от разных сервисов.
// Нужно обработать их параллельно и отдать единый отсортированный стрим.
//
// Ты разбиваешь входные данные на чанки, каждый чанк обрабатывает
// отдельный воркер (Map), а затем результаты всех воркеров
// сливаются в один упорядоченный канал (Reduce).
//
//   Входные события (неотсортированы):
//   [5, 2, 8, 1, 9, 3, 7, 4, 6, 0]
//           │
//           ▼
//      Разбивка на чанки
//           │
//    ┌──────┼──────┐
//    ▼      ▼      ▼
// Worker1 Worker2 Worker3   <- каждый сортирует свой чанк
//  [1,5,8] [2,3,9] [0,4,6,7]   и отдаёт отсортированный канал
//    │      │      │
//    └──────┼──────┘
//           ▼
//         MergeN             <- сливаем N каналов в один упорядоченный
//           │
//           ▼
//   out: 0→1→2→3→4→5→6→7→8→9
//
// =============================================================================
// ТВОЯ ЗАДАЧА:
// =============================================================================
//
//   Реализуй три функции:
//
//   1. MapWorker(events []Event) <-chan Event
//      - Получает чанк событий
//      - Сортирует их по Timestamp
//      - Возвращает канал из которого можно читать отсортированные события
//
//   2. MergeN(channels []<-chan Event) <-chan Event
//      - Получает N отсортированных каналов
//      - Возвращает один канал где все события идут по возрастанию Timestamp
//      - Когда все входные каналы закрыты — закрывает выходной канал
//
//   3. MapReduce(events []Event, numWorkers int) <-chan Event
//      - Разбивает events на numWorkers чанков
//      - Запускает каждый чанк через MapWorker параллельно
//      - Передаёт все каналы воркеров в MergeN
//      - Возвращает итоговый отсортированный канал
//
//   Код не должен содержать data race:
//      go test -race -v
//
// =============================================================================

package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

// Event — событие от внешнего сервиса
type Event struct {
	Timestamp int
	Payload   string
}

// MapWorker получает чанк событий, сортирует их и возвращает канал.
func MapWorker(events []Event) <-chan Event {
	out := make(chan Event)
	go func() {
		defer close(out)
		slices.SortFunc(events, func(lt, rt Event) int {
			if lt.Timestamp < rt.Timestamp {
				return -1
			} else {
				return 1
			}
		})

		for _, e := range events {
			out <- e
		}
	}()

	return out
}

// MergeN сливает N отсортированных каналов в один упорядоченный.
func MergeN(channels [](<-chan Event)) <-chan Event {
	result := make(chan Event)
	values := make([]Event, len(channels))

	for i, ch := range channels {
		v, ok := <-ch
		if !ok {
			values[i] = Event{Timestamp: -1}
			continue
		}

		values[i] = v
	}

	go func(values []Event) {
		defer close(result)
		for {
			// find min
			minEvent := Event{Timestamp: -1}
			minIdx := -1
			for i := 0; i < len(values); i++ {
				if values[i].Timestamp == -1 {
					continue
				}
				if minEvent.Timestamp == -1 || minEvent.Timestamp > values[i].Timestamp {
					minEvent = values[i]
					minIdx = i
				}
			}
			if minIdx == -1 {
				return
			}

			v, ok := <-channels[minIdx]
			if !ok {
				values[minIdx] = Event{Timestamp: -1}
			} else {
				values[minIdx] = v
			}

			result <- minEvent
		}
	}(values)

	return result
}

// MapReduce — точка входа.
// Разбивает события на чанки, запускает воркеры и сливает результаты.
func MapReduce(events []Event, numWorkers int) <-chan Event {
	if len(events) == 0 {
		ch := make(chan Event)
		close(ch)
		return ch
	}

	if numWorkers > len(events) {
		numWorkers = len(events)
	}
	workersCh := make([]<-chan Event, numWorkers)

	chankSize := len(events) / numWorkers
	mod := len(events) % numWorkers
	lt := 0
	rt := mod + chankSize
	for i := 0; i < numWorkers; i++ {
		workersCh[i] = MapWorker(events[lt:rt])
		lt, rt = rt, rt+chankSize
	}

	return MergeN(workersCh)
}

// =============================================================================
// Не меняй код ниже
// =============================================================================

func generateEvents(n int) []Event {
	events := make([]Event, n)
	for i := range events {
		events[i] = Event{
			Timestamp: rand.Intn(n * 10),
			Payload:   fmt.Sprintf("event_%d", i),
		}
	}
	return events
}

func main() {
	rand.Seed(time.Now().UnixNano())

	events := generateEvents(20)

	fmt.Println("Входные события (неотсортированы):")
	for _, e := range events {
		fmt.Printf("  ts=%d\n", e.Timestamp)
	}

	fmt.Println("\nЗапускаем MapReduce с 4 воркерами...")
	out := MapReduce(events, 4)

	fmt.Println("\nРезультат:")
	prev := -1
	ok := true
	for e := range out {
		fmt.Printf("  ts=%d  %s\n", e.Timestamp, e.Payload)
		if e.Timestamp < prev {
			fmt.Println("  ❌ НАРУШЕН ПОРЯДОК!")
			ok = false
		}
		prev = e.Timestamp
	}

	if ok {
		fmt.Println("\n✅ Все события в правильном порядке!")
	}
}
