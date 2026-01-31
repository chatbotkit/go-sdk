// Package params provides utilities for converting API parameter types to URL query values.
package params

import (
	"fmt"
	"net/url"
)

// ListParams is an interface that all list parameter types can implement.
// This allows for type-safe conversion of list parameters to URL query values.
type ListParams interface {
	// ToQuery converts the parameters to URL query values.
	ToQuery() url.Values
}

// BaseListParams contains the common fields for all list parameter types.
// This struct can be embedded in generated types that need query parameter conversion.
type BaseListParams struct {
	Cursor *string           `json:"cursor,omitempty"`
	Meta   map[string]string `json:"meta,omitempty"`
	Take   *int64            `json:"take,omitempty"`
}

// ToQuery converts BaseListParams to URL query values.
func (p *BaseListParams) ToQuery() url.Values {
	query := url.Values{}
	if p == nil {
		return query
	}
	if p.Cursor != nil {
		query.Set("cursor", *p.Cursor)
	}
	if p.Take != nil {
		query.Set("take", fmt.Sprintf("%d", *p.Take))
	}
	for k, v := range p.Meta {
		query.Set("meta["+k+"]", v)
	}
	return query
}

// SetCursor sets the cursor parameter.
func SetCursor(q url.Values, cursor *string) {
	if cursor != nil {
		q.Set("cursor", *cursor)
	}
}

// SetOrder sets the order parameter from any string-based order type.
func SetOrder[T ~string](q url.Values, order *T) {
	if order != nil {
		q.Set("order", string(*order))
	}
}

// SetTake sets the take parameter.
func SetTake(q url.Values, take *int64) {
	if take != nil {
		q.Set("take", fmt.Sprintf("%d", *take))
	}
}

// SetMeta sets the meta parameters.
func SetMeta(q url.Values, meta map[string]string) {
	for k, v := range meta {
		q.Set("meta["+k+"]", v)
	}
}

// BuildListQuery builds URL query values from common list parameters.
// It takes pointers to the common list fields and returns the query values.
// This function is generic over the Order type to support type-safe order parameters.
func BuildListQuery[T ~string](cursor *string, order *T, take *int64, meta map[string]string) url.Values {
	query := url.Values{}
	SetCursor(query, cursor)
	SetOrder(query, order)
	SetTake(query, take)
	SetMeta(query, meta)
	return query
}
