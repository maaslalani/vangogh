package style

import (
  "bytes"
  "encoding/json"
  "io"
  "log"
  "mime/multipart"
  "net/http"
  "os"
)

const STYLE_TRANSFER_URL string = "https://api.deepai.org/api/fast-style-transfer"
const PAINTINGS_URL string = "https://www.vangoghgallery.com/img/"

var key string = os.Getenv("VANGOGH_API_KEY")
var paintings = map[string]string{
  "starry":    "starry_night_full.jpg",
  "potato":    "potato_full.jpg",
  "sunflower": "sunflower_full.jpg",
  "poppies":   "poppies_full.jpeg",
  "irises":    "irises_full.jpeg",
  "bedroom":   "bedroom_full.jpeg",
  "cafe":      "cafe_full.jpeg",
  "mulberry":  "mulberry_full.jpg",
  "blossom":   "blossom_full.jpeg",
}

func Command(content, style string) {
  contentFile, err := os.Open(content)
  defer contentFile.Close()

  if err != nil {
    log.Fatalln(err)
  }

  var requestBody bytes.Buffer

  writer := multipart.NewWriter(&requestBody)
  defer writer.Close()

  contentWriter, err := writer.CreateFormFile("content", "content.png")
  if err != nil {
    log.Fatalln(err)
  }

  _, err = io.Copy(contentWriter, contentFile)
  if err != nil {
    log.Fatalln(err)
  }

  styleWriter, err := writer.CreateFormField("style")
  if err != nil {
    log.Fatalln(err)
  }

  _, err = styleWriter.Write([]byte(PAINTINGS_URL + paintings[style]))
  if err != nil {
    log.Fatalln(err)
  }

  req, err := http.NewRequest("POST", STYLE_TRANSFER_URL, &requestBody)
  if err != nil {
    log.Fatalln(err)
  }

  req.Header.Set("api-key", key)
  req.Header.Set("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  response, err := client.Do(req)
  if err != nil {
    log.Fatalln(err)
  }

  var result map[string]interface{}

  json.NewDecoder(response.Body).Decode(&result)

  fileUrl := result["output_url"]

  if err := DownloadFile("stylized.png", fileUrl.(string)); err != nil {
    panic(err)
  }
}

func DownloadFile(filepath string, url string) error {
  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  out, err := os.Create(filepath)
  if err != nil {
    return err
  }
  defer out.Close()

  _, err = io.Copy(out, resp.Body)
  return err
}
