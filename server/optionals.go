package main

import "encoding/json"

type OptionalString struct {
	v string
}

func (o OptionalString) Value() string {
	return o.v
}
func (o OptionalString) Ok() bool {
	return o.v != ""
}
func (o OptionalString) MarshalJSON() ([]byte, error) {
	if o.v == "" {
		return []byte("null"), nil
	}
	return json.Marshal(o.v)
}
func (o *OptionalString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.v = ""
		return nil
	}
	return json.Unmarshal(data, &o.v)
}

type Optional[T any] struct {
	v  T
	ok bool
}

func (o Optional[T]) Value() T {
	return o.v
}
func (o Optional[T]) Ok() bool {
	return o.ok
}
func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if !o.ok {
		return []byte("null"), nil
	}
	return json.Marshal(o.v)
}
func (o Optional[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.ok = false
		return nil
	}
	o.ok = true
	return json.Unmarshal(data, &o.v)
}

func optvalue[T any](v T) Optional[T] {
	return Optional[T]{
		v:  v,
		ok: true,
	}
}
func optnone[T any]() Optional[T] {
	return Optional[T]{}
}
