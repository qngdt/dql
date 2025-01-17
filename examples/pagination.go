package examples

import (
	"fmt"

	"dql/dql"
)

func Pagination() {
	genreBlock := dql.NewAttribute("genre").
		WithDirectives("(orderasc: name@en)", "(first: 3)").
		WithAttributes(
			dql.NewAttribute("name@en"),
		)

	directorFilmBlock := dql.NewAttribute("director.film").
		WithDirectives("(first: -2)").
		WithAttributes(
			dql.NewAttribute("name@en"),
			dql.NewAttribute("initial_release_date"),
			genreBlock,
		)

	queryBlock := dql.NewQueryBlock("me", `allofterms(name@en, "Steven Spielberg")`).
		WithAttributes(
			directorFilmBlock,
		)

	query := dql.NewQuery("", queryBlock)

	fmt.Println(query.PrettyPrint())
}
