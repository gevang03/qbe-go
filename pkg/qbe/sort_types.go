package qbe

import "fmt"

// TypeDefs must be defined in order without forward referencing.
// This is basically a topological sort using DFS.
func (mod *Module) sortTypes() []TypeDef {
	type status int
	const (
		unvisited = status(iota)
		visiting
		visited
	)
	statuses := make(map[TypeDef]status, len(mod.typeDefs))
	typeDefs := make([]TypeDef, 0, len(mod.typeDefs))
	stack := make([]TypeDef, 0, 1)
	for _, typeDef := range mod.typeDefs {
		if statuses[typeDef] != unvisited {
			continue
		}
		stack = append(stack, typeDef)
		for len(stack) > 0 {
			typeDef := stack[len(stack)-1]
			switch statuses[typeDef] {
			case unvisited:
				statuses[typeDef] = visiting
				for field := range typeDef.getFields() {
					switch statuses[field] {
					case unvisited:
						stack = append(stack, field)
					case visiting:
						panic(fmt.Sprintf("Cyclic type definitions detected involving %v and %v",
							typeDef.Name(), field.Name()))
					}
				}
			case visiting:
				statuses[typeDef] = visited
				typeDefs = append(typeDefs, typeDef)
				stack = stack[:len(stack)-1]
			case visited:
				stack = stack[:len(stack)-1]
			}
		}
	}
	return typeDefs
}
