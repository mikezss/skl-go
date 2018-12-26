package controllers

import (
	"archive/zip"
	"bufio"
	"encoding/json"

	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	_ "os/exec"
	_ "path"
	"path/filepath"
	"regexp"
	"github.com/mikezss/skl-go/models"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
)

// Operations about Users
type COMMONController struct {
	//beego.Controller
	BASEController
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]

func (ctl *COMMONController) GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//获得文件的扩展名
func (ctl *COMMONController) GetExt(filename string) string {
	fileextarr := strings.Split(filename, ".")
	fileext := fileextarr[1]
	return fileext
}
func (ctl *COMMONController) GetFilename(filename string) string {
	fileextarr := strings.Split(filename, ".")
	filename1 := fileextarr[0]
	return filename1
}
func (ctl *COMMONController) Uploadfile() {
	var status, result, rename string
	rename = "true"
	err := ctl.Ctx.Request.ParseMultipartForm(32 << 20)
	if err != nil {
		status = "false"
		result = err.Error()
		ctl.Data["json"] = map[string]string{"status": status, "result": result}
		ctl.ServeJSON()
		return
	}
	LOCAL_FILE_DIR := ctl.Ctx.Request.FormValue("filepath")
	err = os.MkdirAll(LOCAL_FILE_DIR, 0777)
	if err != nil {
		status = "false"
		result = err.Error()
		ctl.Data["json"] = map[string]string{"status": status, "result": result}
		ctl.ServeJSON()
		return
	}
	rename1 := ctl.Ctx.Request.FormValue("renamefilename")
	if rename1 != "" {
		rename = rename1
	}
	log.Println(LOCAL_FILE_DIR)
	//file, handler, err := ctl.Ctx.Request.FormFile("filelist")
	//LOCAL_FILE_DIR := ctl.GetString("filepath")
	//	f, h, err := ctl.GetFile("fileList")
	//	defer f.Close()
	//	if err != nil {
	//		status = "false"
	//		result = err.Error()
	//	} else {
	//		var Url string
	//		ext := ctl.Ext(h.Filename)
	//		fileExt := strings.TrimLeft(ext, ".")
	//		fileSaveName := fmt.Sprintf("%s_%d%s", fileExt, time.Now().Unix(), ext)
	//		filePath := fmt.Sprintf("%s/%s", LOCAL_FILE_DIR, fileSaveName)

	//		err = ctl.SaveToFile("fileList", filePath) // 保存位置在 static/upload,没有文件夹要先创建
	//		if err != nil {
	//			status = "false"
	//			result = err.Error()
	//		}
	//		Url = "/" + filePath
	//		status = "ok"
	//		result = Url
	//	}
	// GetFiles return multi-upload files
	files, err := ctl.GetFiles("filelist")
	if err != nil {
		status = "false"
		result = err.Error()
		ctl.Data["json"] = map[string]string{"status": status, "result": result}
		ctl.ServeJSON()
		return
	}
	for i, _ := range files {
		status = "ok"
		if rename == "true" {
			filename := ctl.GetFilename(files[i].Filename)
			fileExt := ctl.GetExt(files[i].Filename)
			fileSaveName := fmt.Sprintf("%s_%d%s%s", filename, time.Now().Unix(), ".", fileExt)
			result = LOCAL_FILE_DIR + fileSaveName
		} else {
			result = LOCAL_FILE_DIR + files[i].Filename
		}
		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			status = "false"
			result = err.Error()
			break
		}
		//create destination file making sure the path is writeable.
		dst, err := os.Create(result)
		defer dst.Close()
		if err != nil {
			status = "false"
			result = err.Error()
			break
		}
		//copy the uploaded file to the destination file
		_, err = io.Copy(dst, file)
		if err != nil {
			status = "false"
			result = err.Error()
			break
		}
	}

	ctl.Data["json"] = map[string]string{"status": status, "result": result}
	ctl.ServeJSON()
}

//文件按行读入数组
func (this *COMMONController) GetFileContentAsStringLines(filepath string) ([]string, error) {

	result := []string{}
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	s := string(b)
	for _, lineStr := range strings.Split(s, "\n") {
		lineStr = strings.TrimSpace(lineStr)
		//		if lineStr == "" {
		//			continue
		//		}
		result = append(result, lineStr)
	}
	fmt.Println(result)
	fmt.Println(len(result))

	return result, nil
}

//文件中是否包含此字符串
func (this *COMMONController) IsFileIncludestring(filePath string, includedstr string) bool {

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	s := string(b)
	return strings.Contains(s, includedstr)
}

//文件转换为字符串
func (this *COMMONController) Readfile2string(filePath string, charset string) (s string, err1 error) {

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	switch charset {
	case "GBK":
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(b)
		s = string(decodeBytes)
	case "TGBK":
		decodeBytes, _ := traditionalchinese.Big5.NewDecoder().Bytes(b)
		s = string(decodeBytes)
	default:
		s = string(b)

	}

	return s, nil
}

func (ctl *COMMONController) Getfilelistbyfiid() {
	ob := models.FIFLOW{}
	flist := make([]models.FILELIST, 0)
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	fmt.Println(ob.Fiid)
	uri := "/static/files/" + strconv.Itoa(ob.Fiid)
	dirname := ctl.GetCurrentDirectory() + uri
	fmt.Println("DIRNAME:" + dirname)
	fif, err1 := ioutil.ReadDir(dirname)
	if err1 != nil {
		fmt.Println(err1)
		ctl.Data["json"] = map[string]string{"status": err1.Error()}
		ctl.ServeJSON()
		return
	}
	for index, fif1 := range fif {
		flist = append(flist, models.FILELIST{Uid: index, Name: fif1.Name(), Size: fif1.Size(), Url: uri + "/" + fif1.Name()})
	}

	ctl.Data["json"] = flist
	ctl.ServeJSON()

}
func (ctl *COMMONController) Replacefilecontent(filepath string, fromstr string, tostr string) error {
	filecontent, err := ctl.Readfile2string(filepath, "utf8")
	if err != nil {
		fmt.Println(err)
		return err
	}
	reg := regexp.MustCompile(fromstr)
	resultstr := reg.ReplaceAllString(filecontent, tostr)
	data := []byte(resultstr)
	err = ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//文件转换为字符串
func (this *COMMONController) Readfilefirstline(filePath string, charset string) (s string, err1 error) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)

	line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(line)

	return line, nil
}
func (this *COMMONController) IsFDexists(filePath string) bool {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return false
	}
	return true
}
func (this *COMMONController) DeCompressZip(zipfilename string, unzipdir string) (err error, unzipfilenames []string) {

	unzipfilenames = make([]string, 0)
	err = os.MkdirAll(unzipdir, 0777) //创建一个目录
	if err != nil {
		return err, nil
	}

	cf, err := zip.OpenReader(zipfilename) //读取zip文件
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	defer cf.Close()
	for _, file := range cf.File {
		rc, err := file.Open()
		if err != nil {
			fmt.Println(err)
			return err, nil
		}
		unzipfilepath := ""
		if strings.LastIndex(unzipdir, "/") == 0 {
			unzipfilepath = unzipdir + file.Name
		} else {
			unzipfilepath = unzipdir + "/" + file.Name
		}
		//unzipdir.LastIndexOf("is", 1)
		f, err := os.Create(unzipfilepath)
		if err != nil {
			fmt.Println(err)
			return err, nil
		}
		defer f.Close()
		_, err = io.Copy(f, rc)
		if err != nil {
			fmt.Println(err)
			return err, nil
		}
		unzipfilenames = append(unzipfilenames, unzipfilepath)

	}
	return nil, unzipfilenames

}
