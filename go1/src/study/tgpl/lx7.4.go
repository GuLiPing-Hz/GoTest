package main

import (
	"fmt"
	"golang.org/x/net/html"
	"html/template"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

type NewReader struct {
}

func (r *NewReader) Read(p []byte) (n int, err error) {
	root, err := html.Parse(strings.NewReader(string(p)))
	if err != nil {
		return 0, err
	}

	for node := root.FirstChild; node != nil; node = node.NextSibling {
		fmt.Printf("node =%v\n", node)
	}
	return len(p), nil
}

const Html = `
<html>
<head>标题</head>
<body>
<div><p>内容</p></div>
</body>
</html>
`

type LReader struct {
	r     io.Reader
	cnt   int64
	limit int64
}

func (r *LReader) Read(p []byte) (int, error) {
	remain := r.limit - r.cnt
	if int64(len(p)) > remain {
		n, err := r.r.Read(p[:remain])
		if err != nil {
			return n, err
		}
		return int(remain), io.EOF
	}

	return r.r.Read(p)
}

//练习题7.5
func LimitReader(r io.Reader, n int64) io.Reader {
	lr := &LReader{r: r, limit: n}
	return lr
}

//练习题7.8
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type SortTracks struct {
	T       []*Track
	keys    [5]string
	counter map[string]int
}

// Len is the number of elements in the collection.
func (st *SortTracks) Len() int {
	return len(st.T)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (st *SortTracks) Less(i, j int) bool {
	for k := range st.keys {
		switch st.keys[k] {
		case "Title":
			if st.T[i].Title == st.T[j].Title {
				continue
			} else {
				return st.T[i].Title < st.T[j].Title
			}
		default:
			continue
		case "Artist":
			if st.T[i].Artist == st.T[j].Artist {
				continue
			} else {
				return st.T[i].Artist < st.T[j].Artist
			}
		case "Album":
			if st.T[i].Album == st.T[j].Album {
				continue
			} else {
				return st.T[i].Album < st.T[j].Album
			}
		case "Year":
			if st.T[i].Year == st.T[j].Year {
				continue
			} else {
				return st.T[i].Year < st.T[j].Year
			}
		case "Length":
			if st.T[i].Length == st.T[j].Length {
				continue
			} else {
				return st.T[i].Length < st.T[j].Length
			}
		}
	}

	return st.T[i].Title < st.T[j].Title
}

// Swap swaps the elements with indexes i and j.
func (st *SortTracks) Swap(i, j int) {
	st.T[i], st.T[j] = st.T[j], st.T[i]
}

var sTracks = &SortTracks{
	tracks,
	[5]string{},
	make(map[string]int),
}

func tracksHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	key := r.Form.Get("key")
	if key != "" {
		for i := 4; i > 0; i-- {
			if sTracks.keys[i-1] == key {
				sTracks.keys[i-1] = ""
				continue
			}
			sTracks.keys[i] = sTracks.keys[i-1]
		}
		sTracks.keys[0] = key
	}

	sort.Sort(sTracks)
	t, err := template.New("list").Parse(HTMLTemplate)
	if err != nil {
		fmt.Fprintf(w, "err=%s", err.Error())
		return
	}
	t.Execute(w, sTracks)
}

//注意HTML模板中的变量必须是可以导出的变量。即大写。
const HTMLTemplate = `
<h1>list</h1> 
<table> 
<tr style='text-align: left'> 
<td><a href='/tracks?key=Title'>Title</a></td> 
<td><a href='/tracks?key=Artist'>Artist</a></td> 
<td><a href='/tracks?key=Album'>Album</a></td> 
<td><a href='/tracks?key=Year'>Year</a></td> 
<td><a href='/tracks?key=Length'>Length</a></td> </tr> 
{{range .T}} 
<tr><td>{{.Title}}</td> <td>{{.Artist}}</td> <td>{{.Album}}</td> <td>{{.Year}}</td> <td>{{.Length}}</td></tr> 
{{end}}
</table>
`

func main() {
	reader := NewReader{}
	reader.Read([]byte(Html))

	http.HandleFunc("/tracks", tracksHandler)
	http.ListenAndServe(":8000", nil)
}
