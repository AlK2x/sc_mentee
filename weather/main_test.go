package main

import (
	"context"
	"testing"
)

func TestFetchWeatherStrict(t *testing.T) {
	cases := []struct {
		sources         []*WeatherSource
		errorMsg        string
		createContextFn func() context.Context
	}{
		{
			sources:  []*WeatherSource{{Name: "100% error", ErrorRate: 1}},
			errorMsg: "источник \"100% error\" недоступен",
		},
		{
			sources:  []*WeatherSource{{Name: "0% error", ErrorRate: 0}},
			errorMsg: "",
		},
		{
			sources: []*WeatherSource{
				{Name: "1 source", ErrorRate: 0},
				{Name: "2 source", ErrorRate: 0},
				{Name: "3 source", ErrorRate: 0},
			},
			errorMsg: "",
		},
		{
			sources: []*WeatherSource{
				{Name: "1 source", ErrorRate: 0},
			},
			errorMsg: "источник \"1 source\": запрос отменён",
			createContextFn: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			},
		},
		{
			sources: []*WeatherSource{
				{Name: "1 source", ErrorRate: 0},
				{Name: "2 source", ErrorRate: 1},
				{Name: "3 source", ErrorRate: 0},
			},
			errorMsg: "источник \"2 source\" недоступен",
		},
	}

	for i, c := range cases {
		ctx := context.Background()
		if c.createContextFn != nil {
			ctx = c.createContextFn()
		}
		result, err := FetchWeatherStrict(ctx, c.sources)
		if c.errorMsg == "" && (result < 15 || result > 25) {
			t.Errorf("[%d] Temperature is out of acceptable range: %.1f", i, result)
		}

		if c.errorMsg != "" && (err == nil || err.Error() != c.errorMsg) {
			t.Errorf("[%d] Expected error %s, got %v", i, c.errorMsg, err)
		}
	}
}

func TestFetchWeatherBestEffort(t *testing.T) {
	cases := []struct {
		sources         []*WeatherSource
		errorMsg        string
		createContextFn func() context.Context
	}{
		{
			sources:  []*WeatherSource{{Name: "100% error", ErrorRate: 1}},
			errorMsg: "источник \"100% error\" недоступен",
		},
		{
			sources:  []*WeatherSource{{Name: "0% error", ErrorRate: 0}},
			errorMsg: "",
		},
		{
			sources: []*WeatherSource{
				{Name: "1 source", ErrorRate: 0},
				{Name: "2 source", ErrorRate: 0},
				{Name: "3 source", ErrorRate: 0},
			},
			errorMsg: "",
		},
		{
			sources: []*WeatherSource{
				{Name: "1 source", ErrorRate: 0},
			},
			errorMsg: "источник \"1 source\": запрос отменён",
			createContextFn: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			},
		},
		{
			sources: []*WeatherSource{
				{Name: "1 source", ErrorRate: 0},
				{Name: "2 source", ErrorRate: 1},
				{Name: "3 source", ErrorRate: 0},
			},
			errorMsg: "",
		},
	}

	for i, c := range cases {
		ctx := context.Background()
		if c.createContextFn != nil {
			ctx = c.createContextFn()
		}
		result, err := FetchWeatherBestEffort(ctx, c.sources)
		if c.errorMsg == "" && (result < 15 || result > 25) {
			t.Errorf("[%d] Temperature is out of acceptable range: %.1f", i, result)
		}

		if c.errorMsg != "" && (err == nil || err.Error() != c.errorMsg) {
			t.Errorf("[%d] Expected error %s, got %v", i, c.errorMsg, err)
		}
	}
}
