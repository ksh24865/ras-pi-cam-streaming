package handler

import (
	"github.com/gin-gonic/gin"
)


func ShowHomePage(c *gin.Context) {

	
	render(c, gin.H{
		"title": "Home Page"}, "index.html")
}
func ShowStreamPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Stream Page"}, "stream.html")
}
func ShowVideoLogPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Stream Page"}, "video_log.html")
}
