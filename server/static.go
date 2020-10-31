package server

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type StaticPath struct {
	path   string
	prefix string
}

func GetStaticHandle(path StaticPath) Route {
	//routes := make([]Route, len(paths))
	//for index, path := range paths {
	log.Debug(path)
	route := Route{
		path.prefix, getStaticHandler(path),
	}
	return route
	//routes[index] = route
	//}
	//return routes
}

func getStaticHandler(sp StaticPath) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		target := sp.path + getTargetPath(sp.prefix, req.URL.EscapedPath())
		log.Debug(target)
		fileInfo, err := os.Stat(target)
		if err != nil {
			log.Error(err)
			res.WriteHeader(http.StatusNotFound)
			_, err = res.Write([]byte(err.Error()))
		}
		if !fileInfo.IsDir() {
			file, err := ioutil.ReadFile(target)
			if err != nil {
				log.Error(err)
				res.WriteHeader(http.StatusForbidden)
			}
			_, err = res.Write(file)
			if err != nil {
				log.Error(err)
			}
		} else {
			infos, err := ioutil.ReadDir(target)
			if err != nil {
				log.Error(err)
				res.WriteHeader(http.StatusInternalServerError)
				_, err = res.Write([]byte(err.Error()))
			}
			dirInfo := getDirInfoList(infos)
			jsonInfo, err := json.Marshal(dirInfo)
			log.Debug(string(jsonInfo))
			if err != nil {
				log.Error(err)
			}
			_, err = res.Write(jsonInfo)
			if err != nil {
				log.Error(err)
			}
		}
	}
}

func getTargetPath(prefix string, path string) string {
	return path[len(prefix):]
}

func getSysPath(base string, path string) string {
	return base + path
}

type DirInfo struct {
	Name    string    `json:"name"`
	IsDir   bool      `json:"isDir"`
	Size    int64     `json:"size"`
	ModTime time.Time `json:"modTime"`
}

func getDirInfoList(infos []os.FileInfo) []DirInfo {
	dirInfo := make([]DirInfo, len(infos))
	for index, val := range infos {
		info := DirInfo{
			Name:    val.Name(),
			IsDir:   val.IsDir(),
			Size:    val.Size(),
			ModTime: val.ModTime(),
		}
		dirInfo[index] = info
	}
	log.Debug(dirInfo)
	return dirInfo
}
