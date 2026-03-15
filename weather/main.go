// =============================================================================
// ЗАДАЧА: Агрегатор погоды 🌤️
// =============================================================================
//
// ЛЕГЕНДА:
// Ты пишешь сервис, который запрашивает погоду из нескольких источников
// одновременно и возвращает среднюю температуру по всем ответам.
//
// Источники ненадёжные: каждый иногда возвращает ошибку.
//
// =============================================================================
// ТВОЯ ЗАДАЧА:
// =============================================================================
//
//   Реализуй две версии агрегатора с разным поведением при ошибке:
//
//   1. FetchWeatherStrict
//      - Запроси все источники ПАРАЛЛЕЛЬНО
//      - Если любой вернул ошибку — сразу отмени остальные и верни ошибку
//      - Если все успешно — верни среднюю температуру
//
//   2. FetchWeatherBestEffort
//      - Запроси все источники ПАРАЛЛЕЛЬНО
//      - Дождись ВСЕХ источников, даже если часть упала
//      - Посчитай среднюю температуру только по успешным ответам
//      - Если хотя бы один ответил успешно — верни среднее и nil
//      - Если ВСЕ источники упали — верни ошибку
//
//   Пример разницы (3 источника, 1 упал):
//
//   FetchWeatherStrict     → error  (как только упал первый)
//   FetchWeatherBestEffort → среднее по 2 успешным, nil
//
// Проверь себя:
//   go test -race -v
//
// Установи зависимость если нужно:
//   go get golang.org/x/sync/errgroup
//
// =============================================================================

package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

// WeatherSource — источник погоды
type WeatherSource struct {
	Name      string
	ErrorRate float64
}

// Fetch симулирует HTTP-запрос к внешнему API.
// Прерывается если ctx отменён.
func (s *WeatherSource) Fetch(ctx context.Context) (float64, error) {
	delay := time.Duration(rand.Intn(300)+100) * time.Millisecond

	select {
	case <-time.After(delay):
		if rand.Float64() < s.ErrorRate {
			return 0, fmt.Errorf("источник %q недоступен", s.Name)
		}
		temp := 15.0 + rand.Float64()*10
		fmt.Printf("  ✅ %s: %.1f°C\n", s.Name, temp)
		return temp, nil

	case <-ctx.Done():
		fmt.Printf("  ❌ %s: отменён\n", s.Name)
		return 0, fmt.Errorf("источник %q: запрос отменён", s.Name)
	}
}

// FetchWeatherStrict запрашивает все источники параллельно.
// При первой ошибке немедленно отменяет остальные и возвращает ошибку.
func FetchWeatherStrict(ctx context.Context, sources []*WeatherSource) (float64, error) {
	result := make([]float64, len(sources))
	eg := new(errgroup.Group)

	for i, s := range sources {
		eg.Go(func() error {
			temp, err := s.Fetch(ctx)
			if err != nil {
				return err
			}
			result[i] = temp
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return 0, err
	}
	var total float64
	for _, val := range result {
		total += val
	}

	return total / float64(len(result)), nil
}

// FetchWeatherBestEffort запрашивает все источники параллельно.
// Ждёт всех, считает среднее только по успешным ответам.
// Возвращает ошибку только если ВСЕ источники упали.
func FetchWeatherBestEffort(ctx context.Context, sources []*WeatherSource) (float64, error) {
	result := make([]float64, len(sources))
	eg := new(errgroup.Group)
	errors := make([]error, len(sources))
	for i, s := range sources {
		eg.Go(func() error {
			temp, err := s.Fetch(ctx)
			if err != nil {
				errors[i] = err
			}
			result[i] = temp
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return 0, err
	}
	var total float64
	var cnt int
	var err error
	for i := range result {
		if errors[i] != nil {
			err = errors[i]
			continue
		}
		cnt++
		total += result[i]
	}
	if cnt == 0 {
		return 0, err
	}

	return total / float64(cnt), nil
}

func main() {
	sources := []*WeatherSource{
		{
			Name:      "Yandex",
			ErrorRate: 0.2,
		},
		{
			Name:      "Gismeteo",
			ErrorRate: 0.1,
		},
		{
			Name:      "AccuWeather",
			ErrorRate: 0.19,
		},
	}
	ctx := context.Background()
	res, err := FetchWeatherStrict(ctx, sources)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fetch Weather Strict. Avg temperature ", res)
}
