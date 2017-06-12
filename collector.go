package ask

type Collector []Doer

func (c Collector) Do() error {
	for _, d := range c {
		if err := d.Do(); err != nil {
			return err
		}
	}
	return nil
}

func (c *Collector) Push(d Doer) {
	*c = append(*c, d)
}
