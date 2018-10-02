package entities

import "../data"

type ent_base struct {
	id    uint64
	name  string
	ent_t data.entity_types
}
