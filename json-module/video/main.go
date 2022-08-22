package main
import (
	"fmt"
"io/ioutil"
"encoding/json"
)

type Video struct {
	Id string
	Title string
	Description string
	Image string
	Url string
}

func getVideos()(Videos []Video){
	fileBytes, err := ioutil.ReadFile("./videos.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &Videos)

	if err != nil {
		panic(err)
	}

	return Videos  
}

func saveVideos(Videos []Video){
	file, err := json.Marshal(Videos)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./videos-updated.json", file, 0644)

	if err != nil {
		panic(err)
	}
} 

func main(){
	videos := getVideos()
	for i, _ := range videos {
		videos[i].Description = "This is a description" + videos[i].Description
	}

	fmt.Println(videos) 

	saveVideos(videos)
}