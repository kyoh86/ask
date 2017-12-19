package ask

// Collector will collect values to ask.
type Collector []Doer

// Do will get all values from input.
func (c Collector) Do() error {
	for _, d := range c {
		switch err := d.Do(); err {
		case nil:
			// continue
		case ErrSkip:
			// continue
		default:
			return err
		}
	}
	return nil
}

// Push new value to ask.
func (c *Collector) Push(d Doer) {
	*c = append(*c, d)
}
