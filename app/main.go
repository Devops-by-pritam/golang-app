package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Pritam Mane</h1>
	<p>🔭 I’m currently working in Tata Elxsi Ltd.</p>
	<p>🌱 Learning, Building projects on DevOps| Cloud - <a href="https://github.com/Devops-by-pritam">GitHub</a></p>
	<p>Twitter: <a href="https://x.com/priitam03">@priitam03</a></p>
	<p>Email: pritammane7666@gmail.com</p>
	<p>LinkedIn: <a href="https://www.linkedin.com/in/pritam-mane03/">Pritam Mane</a></p>
	<p>YouTube: <a href="https://www.youtube.com/@pritam_mane">Channel</a></p>
	<p>Medium: <a href="https://medium.com/@priitam47">Articles</a></p>
	<p>Portfolio: <a href="https://prritam.github.io/">prritam.github.io</a></p>`)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting at port 3000...")
	http.ListenAndServe(":3000", nil)
}
