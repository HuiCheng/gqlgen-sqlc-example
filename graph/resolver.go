package graph

import "idea2/pkgs/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	dbQueries *db.Queries
}

func NewResolver(dbtx db.DBTX) (r *Resolver, err error) {
	r = &Resolver{
		dbQueries: db.New(dbtx),
	}
	return
}
