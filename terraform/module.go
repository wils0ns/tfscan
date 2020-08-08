package terraform

// Module represents state top level modules
type Module struct {
	Resources    []*Resource `json:"resources"`
	Address      string      `json:"address"`
	ChildModules []*Module   `json:"child_modules"`
}

// ModuleVisitor is the interface that wraps the Visit method
// The Visit Method takes the module to be visited and the its parent module
type ModuleVisitor interface {
	Visit(module, parent *Module)
}

// VisitModules runs Visitor.Visit on the module and all its child modules
func (m *Module) VisitModules(visitor ModuleVisitor, parent *Module) {
	visitor.Visit(m, parent)

	for _, childMod := range m.ChildModules {
		childMod.VisitModules(visitor, m)
	}
}
