package audio

import (
	. "ScamfariHelper/modules/logModule"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"os"
	"time"
)

func StreamAudio(file string) {
	f, err := os.Open(file)
	if err != nil {
		Logger.Printf("Error in StreamAudio: missing file %v", file)
		return
	}
	streamer, format, err := wav.Decode(f)
	if err != nil {
		Logger.Printf("Error in StreamAudio: can not decode wav file %v", file)
		return
		//log.Fatal(err)
	}
	defer streamer.Close()
	err2 := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		Logger.Println("Error in StreamAudio speakers.Init. Error: ", err2)
	}
	speaker.Play(streamer)
	// Wait
	select {}
}
