package static

import (
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/dll02/webgo/framework/gin"
)

const INDEX = "index.html"

type ServeFileSystem interface {
	http.FileSystem
	Exists(prefix string, path string) bool
}

type localFileSystem struct {
	http.FileSystem
	root    string
	indexes bool
}

func LocalFile(root string, indexes bool) *localFileSystem {
	return &localFileSystem{
		FileSystem: gin.Dir(root, indexes),
		root:       root,
		indexes:    indexes,
	}
}

func (l *localFileSystem) Exists(prefix string, filepath string) bool {
	// 使用 TrimPrefix 剥离前缀，以获取相对于根目录的文件路径
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		// 拼接文件的完整路径
		name := path.Join(l.root, p)

		// 获取文件信息
		stats, err := os.Stat(name)
		if err != nil {
			return false // 如果出错，表示文件不存在
		}

		// 检查是否是目录
		if stats.IsDir() {
			if !l.indexes {
				// 如果目录中不允许索引文件（indexes 为 false）
				// 则检查是否存在索引文件（INDEX 为索引文件名）
				index := path.Join(name, INDEX)
				_, err := os.Stat(index)
				if err != nil {
					return false // 如果索引文件不存在，表示目录不存在
				}
			}
		}
		return true // 文件或目录存在
	}
	return false // 文件路径不匹配前缀，返回不存在
}

func ServeRoot(urlPrefix, root string) gin.HandlerFunc {
	return Serve(urlPrefix, LocalFile(root, false))
}

// Static returns a middleware handler that serves static files in the given directory.
func Serve(urlPrefix string, fs ServeFileSystem) gin.HandlerFunc {
	fileserver := http.FileServer(fs)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}
	return func(c *gin.Context) {
		// urlPrefix: 指定的静态文件夹 Path:请求的文件名
		if fs.Exists(urlPrefix, c.Request.URL.Path) {
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}
