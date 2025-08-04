package main
import(
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)
type DeleteRequest struct {
    Filename string `json:"filename" binding:"required"`
}

func main(){
	r := gin.Default()
	r.POST("/create_page", Create_page)
	r.POST("/detete_page", Delete_page)
	r.GET("/all-pages", All_pages)
	r.Run(":8080")
}
func Create_page(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "Файл не найден"})
		return
	}
	contenttype := file.Header.Get("Content-Type")
	if contenttype != "text/html" {
		c.JSON(400, gin.H{"error": "Только HTML-файлы разрешены"})
		return
	}
	c.SaveUploadedFile(file, "/home/wake_up/MY_PROJ/wiki_file" + file.Filename)
	c.JSON(200, gin.H{"message": "Страница успешно создана", "filename": file.Filename})
}

func Delete_page(c *gin.Context) {
	var req DeleteRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": "Некорректные данные"})
        return
    }	
	
	if err := os.Remove("/home/wake_up/MY_PROJ/wiki_file/" + req.Filename); err != nil {
		c.JSON(500, gin.H{"error": "Не удалось удалить страницу"})
		return
	}
	c.JSON(200, gin.H{"message": "Страница успешно удалена"})
} 
func All_pages(c *gin.Context) {
	file, err := os.ReadDir("/home/wake_up/MY_PROJ/wiki_file")
	if err != nil{
		c.JSON(500, gin.H{"error": "Не удалось получить список страниц"})
	}
	var pages []string
	for _, f := range file {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".html"){
		pages = append(pages, f.Name())
		}
	}
	c.JSON(200, gin.H{"pages": pages})
}