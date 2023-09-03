package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"html/template"
	"time"
	"math/rand"
)

var tmpl = template.Must(template.ParseGlob("*.html"))
const MAX_UPLOAD_SIZE = 5*1024*1024
var uploadPath = "./uploads"


func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/upload", UploadFileHandler)
	fs := http.FileServer(http.Dir("./uploads"))
	http.Handle("/file/", http.StripPrefix("/file/", fs))
	fs1 := http.FileServer(http.Dir("./edits"))
	http.Handle("/edit/", http.StripPrefix("/edit/", fs1))
	log.Println("strating server")
	if err := http.ListenAndServe(":7000", nil); err != nil {
		log.Fatal(err)
	}
}


type Progress struct {
	TotalSize int64
	BytesRead int64
}

func (pr *Progress) Write(p []byte) (n int, err error) {
	n, err = len(p), nil
	pr.BytesRead += int64(n)
	pr.Print()
	return
}

func (pr *Progress) Print() {
	if pr.BytesRead == pr.TotalSize {
		fmt.Println("DONE!")
		return
	}

	fmt.Printf("File upload in progress: %d\n", pr.BytesRead)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "xhr.html")
}


func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "xhr.html")
	}
		if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
			fmt.Printf("Could not parse multipart form: %v\n", err)
			renderError(w, "CANT_PARSE_FORM", http.StatusInternalServerError)
		}
		var fileEndings string
		var fileName string

		files := r.MultipartForm.File["imgfile"]
		for _, fileHeader := range files {
			log.Println("hao")
			file, err := fileHeader.Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			defer file.Close()
			log.Println("file OK")

			fileSize := fileHeader.Size
			fmt.Printf("File size (bytes): %v\n", fileSize)

			if fileSize > MAX_UPLOAD_SIZE {
				renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			}
			fileBytes, err := io.ReadAll(file)
			if err != nil {
				renderError(w, "INVALID_FILE"+http.DetectContentType(fileBytes), http.StatusBadRequest)
			}

			// check file type, detectcontenttype only needs the first 512 bytes
			detectedFileType := http.DetectContentType(fileBytes)
			switch detectedFileType {
			case "image/png":
				fileEndings = ".png"
				break
			case "image/jpeg":
				fileEndings = ".jpg"
				break
			default:
				renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
			}
			fileName = GenerateName(12)
			//fileEndings, err := mime.ExtensionsByType(detectedFileType)

			if err != nil {
				renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
			}
			newFileName := fileName + fileEndings

			newPath := filepath.Join(uploadPath, newFileName)
			fmt.Printf("FileType: %s, File: %s\n", detectedFileType, newPath)

			// write file
			newFile, err := os.Create(newPath)
			if err != nil {
				renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			}
			defer newFile.Close() // idempotent, okay to call twice
			if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
				renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			}
		}
		var ff = FFFilters {
			fileName : fileName, 
			fileExt : fileEndings, 
			desiredExt : ".png", 
			width : r.FormValue("imgwidth"), 
			filter : "lanczos",
			angle : "45",
		}
		switch r.FormValue("task") {
			case "scale" :  {
				ffScale(ff);		
				sendImageAsHTML(w,r,"/edit/scale_"+fileName+ff.desiredExt)
			}
			case "hflip" :  {
				ffHflip(ff);		
				sendImageAsHTML(w,r,"/edit/hflip_"+fileName+ff.desiredExt)
			}
			case "vflip" :  {
				ffVflip(ff);		
				sendImageAsHTML(w,r,"/edit/vflip_"+fileName+ff.desiredExt)
			}
			case "rotateclk" :  {
				ffRotateClk(ff);		
				sendImageAsHTML(w,r,"/edit/rotateclk_"+fileName+ff.desiredExt)
			}
			case "rotateanticlk" :  {
				ffRotateAntiClk(ff);		
				sendImageAsHTML(w,r,"/edit/rotateanticlk_"+fileName+ff.desiredExt)
			}
			case "rotate180" :  {
				ffRotate180(ff);		
				sendImageAsHTML(w,r,"/edit/rotate180_"+fileName+ff.desiredExt)
			}
			case "rotateangle" :  {
				ffRotateAngle(ff);		
				sendImageAsHTML(w,r,"/edit/rotate"+ff.angle+"_"+fileName+ff.desiredExt)
			}
			case "negval" :  {
				ffNegval(ff);		
				sendImageAsHTML(w,r,"/edit/negval_"+fileName+ff.desiredExt)
			}
			case "blur" :  {
				ffBlur(ff);		
				sendImageAsHTML(w,r,"/edit/blur_"+fileName+ff.desiredExt)
			}
			case "canny" :  {
				ffCanny(ff);		
				sendImageAsHTML(w,r,"/edit/canny_"+fileName+ff.desiredExt)
			}
			case "negclr" :  {
				ffNegclr(ff);		
				sendImageAsHTML(w,r,"/edit/negclr_"+fileName+ff.desiredExt)
			}
			case "yelo" :  {
				ffYelo(ff);		
				sendImageAsHTML(w,r,"/edit/yelo_"+fileName+ff.desiredExt)
			}
			case "vintage" :  {
				ffVintage(ff);		
				sendImageAsHTML(w,r,"/edit/vintage_"+fileName+ff.desiredExt)
			}
			case "darken" :  {
				ffDarken(ff);		
				sendImageAsHTML(w,r,"/edit/darken_"+fileName+ff.desiredExt)
			}
			case "duotone" :  {
				ffDuotone(ff);		
				sendImageAsHTML(w,r,"/edit/duotone_"+fileName+ff.desiredExt)
			}
			case "mrgp" :  {
				ffMergePlanes(ff);		
				sendImageAsHTML(w,r,"/edit/mrgp_"+fileName+ff.desiredExt)
			}
			case "pixel" :  {
				ffPixelated(ff);		
				go sendImageAsHTML(w,r,"/edit/pixel_"+fileName+ff.desiredExt)
			}
		}
		os.Remove("./uploads/"+fileName+ff.fileExt)
		fmt.Println("./uploads/"+fileName+ff.fileExt)
}

type FFFilters struct {
	fileName , fileExt , desiredExt, height, width, filter, angle string
}

func sendImageAsHTML(w http.ResponseWriter, r *http.Request, a string) {
	fmt.Fprintf(w,a)
}

func sendImageAsAttachment(w http.ResponseWriter, r *http.Request) {
    buf, err := os.ReadFile("F46ZPQ0bQAACFYs.jpg")
    if err != nil {
        log.Fatal(err)
    }
    w.Header().Set("Content-Type", "image/jpeg")
    w.Header().Set("Content-Disposition", `attachment;filename="F46ZPQ0bQAACFYs.jpg"`)
    w.Write(buf)
}

func sendImageAsBytes(w http.ResponseWriter, r *http.Request, a string) {
    buf, err := os.ReadFile("./uploads/"+a)
    if err != nil {
        log.Fatal(err)
    }
    w.Header().Set("Content-Type", "image/jpeg")
    w.Write(buf)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789_ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateName(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func init() {
    rand.Seed(time.Now().UnixNano())
}


func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}
