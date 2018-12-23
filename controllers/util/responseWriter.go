package util

import (
	"compress/gzip"
	"net/http"
	"strings"
)

//CloseableResponseWriter : Custom gzip
type CloseableResponseWriter interface {
	http.ResponseWriter
	Close()
}

type gzipResponseWriter struct {
	http.ResponseWriter
	*gzip.Writer
}

//Write : Custom gzip
func (gzipRespWtr gzipResponseWriter) Write(data []byte) (int, error) {
	return gzipRespWtr.Writer.Write(data)
}

func (gzipRespWtr gzipResponseWriter) Close() {
	gzipRespWtr.Writer.Close()
}

func (gzipRespWtr gzipResponseWriter) Header() http.Header {
	return gzipRespWtr.ResponseWriter.Header()
}

type closeableResponseWriter struct {
	http.ResponseWriter
}

func (gzipRespWtr closeableResponseWriter) Close() {
}

//GetResponseWriter : Custom Response Writer for compressing/gzip content
func GetResponseWriter(w http.ResponseWriter, req *http.Request) CloseableResponseWriter {
	if strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		gRW := gzipResponseWriter{
			ResponseWriter: w,
			Writer:         gzip.NewWriter(w),
		}
		return gRW
	}
	return closeableResponseWriter{ResponseWriter: w}
}
