package ask

type Collector []Doer

func (c Collector) Do() error {
	for _, d := range c {
		switch err := d.Do(); err {
		case nil:
			// continue
		case Skip:
			// continue
		default:
			return err
		}
	}
	return nil
}

func (c *Collector) Push(d Doer) {
	*c = append(*c, d)
}
