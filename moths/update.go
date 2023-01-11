package moths

import (
	"log"
	"time"
)

// IDEA: Use this to update time
func (m *Moths) update() {
	update := time.NewTicker(m.interval).C

	for {
		// m.time = m.time.Add(time.Since(m.time))
		<-update
		log.Println("UPDATE", m.time)
	}
}
