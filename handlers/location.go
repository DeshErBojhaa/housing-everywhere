package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/DeshErBojhaa/housing-everywhere/web"
)

// Location ...
type Location struct {
	log *log.Logger
}

// Atlas finds the location value from the coordinates.
func (l *Location) Atlas(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	lv, err := getLocation(ctx, w, r, params)
	if err != nil {
		return err
	}
	loc := struct {
		Loc float64 `json:"loc"`
	}{
		Loc: lv,
	}
	return web.Respond(ctx, w, loc, http.StatusOK)
}

// Mama finds the location value from the coordinates.
func (l *Location) Mama(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	lv, err := getLocation(ctx, w, r, params)
	if err != nil {
		return err
	}
	loc := struct {
		Location float64 `json:"location"`
	}{
		Location: lv,
	}
	return web.Respond(ctx, w, loc, http.StatusOK)
}

func getLocation(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) (float64, error) {
	var c struct {
		X   string `json:"x"`
		Y   string `json:"y"`
		Z   string `json:"z"`
		Vel string `json:"vel"`
	}
	if err := web.Decode(r, &c); err != nil {
		return -1.0, err
	}
	id, err := strconv.ParseFloat(params["sector_id"], 64)
	if err != nil {
		return -1.0, fmt.Errorf("invalid sector id %v", params["sector_id"])
	}

	ans := 0.0

	v := reflect.ValueOf(c)
	for i := 0; i < v.NumField(); i++ {
		s := v.Field(i).String()
		val, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return -1.0, fmt.Errorf("invalid sector id %v", s)
		}
		ans += val * id
	}
	return ans, nil
}
