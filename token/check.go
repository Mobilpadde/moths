package token

// Check is checking if everything
// is working
//
// Meaning no errors is generated
// from getting codes
func (g *Generator) Check() error {
	if _, err := g.getToken(); err != nil {
		return err
	}

	return nil
}
