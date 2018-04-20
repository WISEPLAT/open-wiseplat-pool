package shifts

import (
	"log"

	"github.com/wiseplat/open-wiseplat-pool/storage"
	"github.com/wiseplat/open-wiseplat-pool/util"
)

type ShiftsConfig struct {
	Enabled      bool   `json:"enabled"`
	LongInterval     string `json:"longInterval"`
	ShortInterval    string `json:"shortInterval"`
	FlushInterval    string `json:"flushInterval"`
	KeepLong         string `json:"keepLong"`
	KeepShort        string `json:"keepShort"`
}

type ShiftsProcessor struct {
	config   *ShiftsConfig
	backend  *storage.RedisClient
}

func NewShiftsProcessor(cfg *ShiftsConfig, backend *storage.RedisClient) *ShiftsProcessor {
	s := &ShiftsProcessor{config: cfg, backend: backend}
	return s
}

func (u *ShiftsProcessor) Start() {
	log.Println("Starting shifts")

	// Shifts
	longIntv, shortIntv := util.MustParseDuration(u.config.LongInterval), util.MustParseDuration(u.config.ShortInterval)
	log.Printf("Set shifting intervals to (%v, %v)", longIntv, shortIntv)
	
	// Windows
	longWindow, shortWindow := util.MustParseDuration(u.config.KeepLong), util.MustParseDuration(u.config.KeepShort)
	log.Printf("Set shift windows to (%v, %v)", longWindow, shortWindow)
	
	// Flushing
	flushIntv := util.MustParseDuration(u.config.FlushInterval)
	log.Printf("Set shifts flushing interval to %v", flushIntv)
	
	processFlush := func() {
		users, err := u.backend.GetMiners()
		if err != nil {
			log.Println("Error while retrieving miners from backend:", err)
		} else {
			entries, err2 := u.backend.FlushShifts(longWindow, shortWindow, users)
			if err2 != nil {
				log.Println("Error while flushing entries:", err2)
			}
			log.Printf("Flushed %v entries", entries)
		}
	}
	
	util.Schedule(u.processLong, longIntv)
	util.Schedule(u.processShort, shortIntv)
	util.Schedule(processFlush, flushIntv)
}

func (u *ShiftsProcessor) processLong() {
	users, err := u.backend.GetMiners()
	if err != nil {
		log.Println("Error while retrieving miners from backend:", err)
		return
	}
	
	shiftsDone := 0
	for _, login := range users {
		u.backend.WriteLongShift(login)
		shiftsDone++
	}
	
	log.Printf("%v long shifts created", shiftsDone)
}

func (u *ShiftsProcessor) processShort() {
	users, err := u.backend.GetMiners()
	if err != nil {
		log.Println("Error while retrieving miners from backend:", err)
		return
	}
	
	shiftsDone := 0
	for _, login := range users {
		u.backend.WriteShortShift(login)
		shiftsDone++
	}
	
	log.Printf("%v short shifts created", shiftsDone)
}
