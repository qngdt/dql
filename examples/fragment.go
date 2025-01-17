package examples

import (
	"fmt"

	"dql/dql"
)

func Fragment() {
	fragmentB := dql.NewFragment("TestFragB").
		WithAttributes(
			dql.NewAttribute("country"),
		)

	fragmentA := dql.NewFragment("TestFrag").
		WithAttributes(
			dql.NewAttribute("initial_release_date"),
			dql.NewAttribute("...TestFragB"),
		)

	queryBlock := dql.NewQueryBlock("debug", "uid(1)").
		WithAttributes(
			dql.NewAttribute("name@en"),
			dql.NewAttribute("...TestFrag"),
		)

	query := dql.NewQuery("query", queryBlock).
		WithFragments(fragmentA, fragmentB)

	fmt.Println(query.PrettyPrint())
}
