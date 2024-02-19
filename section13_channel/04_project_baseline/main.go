package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"

	"time"
)

type Photos []struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type Image struct {
	filePath string
	img      []byte
}

func main() {
	defer func() {
		fmt.Println(" main program exit successful")
	}()
	// getJson()
	photos := Photos{}
	getJson("https://jsonplaceholder.typicode.com/photos", &photos)

	//	check and create folder
	dir := "mydownload" + time.Now().Format("15_04_05")
	if _, err := os.Stat(dir); err != nil {
		os.Mkdir(dir, 0777)
	}
	chImg := make(chan Image, len(photos))

	for _, v := range photos {
		v := v
		go func() {
			// downloader() แปลง string ไป []byte
			img, err := downloader(v.ThumbnailURL)
			if err != nil {
				log.Fatal(err)
			}

			// analyze file type
			// image Decode
			format, err := decodeImage(img)
			if err != nil {
				log.Fatal(err)
			}
			// save image to local disk
			file := fmt.Sprintf("%d.%s", v.ID, format)
			//err = saveImage(filepath.Join(dir, file), img)
			// if err != nil {
			// 	log.Println(err)
			// }

			chImg <- Image{filePath: file, img: img}

		}()
	}
	for range photos {
		v := <-chImg
		err := saveImage(v.filePath, v.img)
		if err != nil {
			log.Println(err)
		}
	}
	time.Sleep(10 * time.Millisecond)
}
func saveImage(file string, img []byte) error {

	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	defer f.Close()

	_, err = io.Copy(f, bytes.NewReader(img))
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return err
}

func decodeImage(img []byte) (string, error) {
	_, format, err := image.Decode(bytes.NewReader(img))
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	return format, err
}

func downloader(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer resp.Body.Close()

	img, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return img, err
}

func getJson(url string, strucyType interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	defer resp.Body.Close()

	switch v := strucyType.(type) {
	case *Photos:
		decode := json.NewDecoder(resp.Body)
		decode.Decode(strucyType.(*Photos))
		return nil
	default:
		return fmt.Errorf("%v", v)
	}
}
