package main

func (p *Piles) aEstTriee() bool {
	for i := 0; i < len(p.a)-1; i++ {
		if p.a[i] > p.a[i+1] {
			return false
		}
	}
	return true
}
