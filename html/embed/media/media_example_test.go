package media_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/embed/media"
	"github.com/protolambda/chord/html/text"
)

func ExampleVideo() {
	core.Dump(media.Video(media.Src("movie.mp4"), media.Controls()))
	// Output: <video src="movie.mp4" controls></video>
}

func ExampleVideo_sources() {
	core.Dump(media.Video(media.Controls(),
		media.Source(core.Attribute("src", "movie.webm"), core.Attribute("type", "video/webm")),
		media.Source(core.Attribute("src", "movie.mp4"), core.Attribute("type", "video/mp4")),
	))
	// Output: <video controls><source src="movie.webm" type="video/webm"/><source src="movie.mp4" type="video/mp4"/></video>
}

func ExampleAudio() {
	core.Dump(media.Audio(media.Src("song.mp3"), media.Controls()))
	// Output: <audio src="song.mp3" controls></audio>
}

func ExampleTrack() {
	core.Dump(media.Video(media.Src("movie.mp4"),
		media.Track(core.Attribute("src", "subs.vtt"), core.Attribute("kind", "subtitles"), core.Attribute("srclang", "en")),
	))
	// Output: <video src="movie.mp4"><track src="subs.vtt" kind="subtitles" srclang="en"/></video>
}

func ExampleVideo_fallback() {
	core.Dump(media.Video(media.Src("movie.mp4"), text.Text("Video not supported")))
	// Output: <video src="movie.mp4">Video not supported</video>
}
