package wavutils

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/antyiwei/goutils/fileutils"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

/*
		Save wave file
		input param is write data
	    name param is file url
*/
func Save(input []byte, name string) {

	exist, _ := fileutils.PathExists(name)
	if exist {
		saveWave(name, input)
	} else {
		createNewWave(name, input)
	}
}

type wavworker struct {
	numChannels    int
	sampleRate     int
	sourceBitDepth int
	audioFormat    int
	e              *wav.Encoder
}

var (
	numChannels    = 1
	sampleRate     = 8000
	sourceBitDepth = 8
	audioFormat    = 1
)

func genWavWork(out io.WriteSeeker, numChannels, sampleRate, sourceBitDepth, audioFormat int) *wavworker {
	return &wavworker{
		numChannels:    numChannels,
		sampleRate:     sampleRate,
		sourceBitDepth: sourceBitDepth,
		audioFormat:    audioFormat,
		e:              wav.NewEncoder(out, sampleRate, int(sourceBitDepth), numChannels, int(audioFormat)),
	}
}

func (w *wavworker) write(buf []byte) error {

	inBuf := audio.IntBuffer{
		Format: &audio.Format{
			NumChannels: numChannels,
			SampleRate:  sampleRate,
		},
		SourceBitDepth: sourceBitDepth,
	}

	inBuf.Data = make([]int, len(buf))
	for i := 0; i < len(buf); i++ {
		inBuf.Data[i] = int(buf[i])
	}

	w.e.WriteFrame(inBuf)

	return nil
}

func (w *wavworker) close() {
	w.e.Close()
}

func createNewWave(url string, input []byte) {

	if len(input) == 0 {
		log.Println("input is nil")
		return
	}

	out, err := os.Create(url)
	if err != nil {
		log.Println(fmt.Sprintf("couldn't create output file %s - %v", url, err))
		return
	}
	defer out.Close()

	w := genWavWork(out, numChannels, sampleRate, sourceBitDepth, audioFormat)
	defer w.close()

	err = w.write(input)
	if err != nil {
		log.Println(fmt.Sprintf("couldn't write output file %s - %v", url, err))
		return
	}

}

func saveWave(url string, input []byte) {
	if len(input) == 0 {
		log.Println("input is nil")
		return
	}

	fout, err := os.OpenFile(url, os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer fout.Close()

	w := genWavWork(fout, numChannels, sampleRate, sourceBitDepth, audioFormat)
	defer w.close()

	err = w.write(input)
	if err != nil {
		log.Println(fmt.Sprintf("couldn't write output file %s - %v", url, err))
		return
	}

}
