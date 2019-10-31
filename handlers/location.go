package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
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

	x, err := strconv.ParseFloat(c.X, 64)
	if err != nil {
		return -1.0, fmt.Errorf("invalid sector id %v", c.X)
	}

	y, err := strconv.ParseFloat(c.Y, 64)
	if err != nil {
		return -1.0, fmt.Errorf("invalid sector id %v", c.Y)
	}

	z, err := strconv.ParseFloat(c.Z, 64)
	if err != nil {
		return -1.0, fmt.Errorf("invalid sector id %v", c.Z)
	}

	vel, err := strconv.ParseFloat(c.Vel, 64)
	if err != nil {
		return -1.0, fmt.Errorf("invalid sector id %v", c.Vel)
	}
	return x*id + y*id + z*id + vel*id, nil
}
