package graph

import "github.com/RossiniM/full-cycle-graphql/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	Categories []*model.Category
	Coursers []*model.Course
	Chapters []*model.Chapter
}
