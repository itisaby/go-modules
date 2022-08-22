package main
import()
type Video struct {
	Id string
	Title string
	Description string
	Image string
	Url string
}

func getVideos()(Videos []Video){
	video1 := Video{
		Id: "1",
		Title: "Video 1",
		Description: "Video 1 Description",
		Image: "https://via.placeholder.com/150",
		Url: "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
		}
	video2 := Video{
		Id: "2",
		Title: "Video 2",
		Description: "Video 2 Description",
		Image: "https://via.placeholder.com/150",
		Url: "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
		}
	return []Video{video1, video2}

}